package configs

import (
	"crypto/tls"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// this global variable can be access anywhere after initialize with function InitConfigWithState
var AppConfig = &AppConfigs{}

type AppConfigs struct {
	HTTPTransport *http.Transport
	Validator     *validator.Validate
	HttpTransport struct {
		MaxIdleCons           int           `mapstructure:"max_idle_cons"`
		MaxIdleConsPerHost    int           `mapstructure:"max_idle_cons_per_host"`
		IdleConTimeout        time.Duration `mapstructure:"idle_con_timeout"`
		TLSHandshakeTimeout   time.Duration `mapstructure:"tls_handshake_timeout"`
		ResponseHeaderTimeout time.Duration `mapstructure:"response_header_timeout"`
		ExpectContinueTimeout time.Duration `mapstructure:"expect_continue_timeout"`
	} `mapstructure:"http_transport"`
	Server struct {
		GracefulShutdownTime time.Duration `mapstructure:"graceful_shutdown_time"`
		HTTPServer           struct {
			ReadTimeout  time.Duration `mapstructure:"read_timeout"`
			WriteTimeout time.Duration `mapstructure:"write_timeout"`
			IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
		} `mapstructure:"http_server"`
	} `mapstructure:"server"`
}

func InitAppConfig(configPath string) error {

	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("app_config")

	if err := v.ReadInConfig(); err != nil {
		log.Errorln("read config file error:", err)
		return err
	}

	if err := bindingAppConfig(v, AppConfig); err != nil {
		log.Errorln("binding config error:", err)
		return err
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Infoln("config file changed:", e.Name)
		if err := bindingAppConfig(v, AppConfig); err != nil {
			log.Errorln("binding error:", err)
		}
		log.Printf("config: %+v", AppConfig)
	})

	return nil
}

func bindingAppConfig(vp *viper.Viper, appConf *AppConfigs) error {
	if err := vp.Unmarshal(&appConf); err != nil {
		return err
	}
	tp := http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:          appConf.HttpTransport.MaxIdleCons,
		MaxIdleConnsPerHost:   appConf.HttpTransport.MaxIdleConsPerHost,
		IdleConnTimeout:       appConf.HttpTransport.IdleConTimeout,
		TLSHandshakeTimeout:   appConf.HttpTransport.TLSHandshakeTimeout,
		ResponseHeaderTimeout: appConf.HttpTransport.ResponseHeaderTimeout,
		ExpectContinueTimeout: appConf.HttpTransport.ExpectContinueTimeout,
	}
	appConf.HTTPTransport = &tp

	appConf.Validator = validator.New(&validator.Config{TagName: "validate"})
	return nil
}

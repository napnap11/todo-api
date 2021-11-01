package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/napnap11/todo-api/internal/pkg/configs"
)

// Server struct for use by graceful shutdown
type Server struct {
	Server *http.Server
}

type ServerHandlerFunc func(*Server, http.ResponseWriter, *http.Request)

func (s *Server) H(f ServerHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(s, w, r)
	}
}

func NewServer(port string, routes *gin.Engine) *Server {
	return &Server{
		Server: &http.Server{
			Addr:         fmt.Sprintf(":%s", port),
			Handler:      routes,
			ReadTimeout:  configs.AppConfig.Server.HTTPServer.ReadTimeout,
			WriteTimeout: configs.AppConfig.Server.HTTPServer.WriteTimeout,
			IdleTimeout:  configs.AppConfig.Server.HTTPServer.IdleTimeout,
		},
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Server.Handler.ServeHTTP(w, r)
}

func (s *Server) ListenAndServe() {
	s.Server.ListenAndServe()
}

func (s *Server) ListenAndServeWithGracefulShutdown(t time.Duration) {
	trigger := make(chan os.Signal, 1)
	signal.Notify(trigger, os.Interrupt)

	go s.Server.ListenAndServe()
	<-trigger

	sync, _ := context.WithTimeout(context.Background(), t)
	s.Server.Shutdown(sync)
}

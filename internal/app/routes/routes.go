package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const (
	get  = "GET"
	post = "POST"
)

type Route struct {
	desc    string
	path    string
	method  string
	handler gin.HandlerFunc
}

func NewRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// ===================== START ROUTES (add routes here) ========================
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes := []Route{

		{
			desc:    "health check",
			method:  get,
			path:    "/status",
			handler: status,
		},
	}

	v1 := []Route{
	}

	log.Printf("==========PATH==========\n")
	for _, route := range routes {
		r.Handle(route.method, route.path, route.handler)
	}
	rv1 := r.Group("/v1")
	for _, e := range v1 {
		rv1.Handle(e.method, e.path, e.handler)
	}
	log.Printf("========================\n")
	return r
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "todo OK",
	})
}

package routes

import (
	"github.com/napnap11/todo-api/internal/pkg/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	createHandler "github.com/napnap11/todo-api/internal/app/services/create/handler"
	createService "github.com/napnap11/todo-api/internal/app/services/create/service"
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
	repo := repository.NewRepository()
	create := createHandler.NewHandler(createService.NewService(repo))

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
		{
			desc:    "create",
			method:  post,
			path:    "/create",
			handler: create.Create,
		},
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

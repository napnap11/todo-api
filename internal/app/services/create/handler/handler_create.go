package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/napnap11/todo-api/internal/app/services/create/service"
)

type Handler struct {
	s service.Service
}

func NewHandler(s service.Service) Handler {
	return Handler{s: s}
}

func (h Handler) Create(ctx *gin.Context) {
	return
}

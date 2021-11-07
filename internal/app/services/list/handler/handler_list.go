package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/napnap11/todo-api/internal/app/services/list/dto"
	"github.com/napnap11/todo-api/internal/app/services/list/service"
)

type Handler struct {
	s service.Service
}

func NewHandler(s service.Service) Handler {
	return Handler{s: s}
}

func (h Handler) List(ctx *gin.Context) {
	var req dto.ListRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		log.Errorf("[List] decode error: %s",err)
		ctx.JSON(http.StatusOK, dto.ListResponse{
			ErrorCode: "10",
			ErrorDesc: "invalid request",
		})
		return
	}
	defer ctx.Request.Body.Close()

	resp := h.s.List(req)
	ctx.JSON(http.StatusOK, resp)
}

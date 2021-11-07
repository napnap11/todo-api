package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/napnap11/todo-api/internal/app/services/update/dto"
	"github.com/napnap11/todo-api/internal/app/services/update/service"
)

type Handler struct {
	s service.Service
}

func NewHandler(s service.Service) Handler {
	return Handler{s: s}
}

func (h Handler) Update(ctx *gin.Context) {
	var req dto.UpdateRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		log.Errorf("[Update] decode error: %s",err)
		ctx.JSON(http.StatusOK, dto.UpdateResponse{
			ErrorCode: "10",
			ErrorDesc: "invalid request",
		})
		return
	}
	defer ctx.Request.Body.Close()

	resp := h.s.Update(req)
	ctx.JSON(http.StatusOK, resp)
}

package service

import (
	"errors"
	"github.com/napnap11/todo-api/internal/app/services/create/dto"
	"github.com/napnap11/todo-api/internal/app/services/create/repo"
	"github.com/napnap11/todo-api/internal/pkg/configs"
	"github.com/napnap11/todo-api/internal/pkg/models"
	log "github.com/sirupsen/logrus"
	"time"
)

type Service interface {
	Create(req dto.CreateRequest) dto.CreateResponse
}

type service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) Service {
	return service{repo: repo}
}

func (s service) Create(req dto.CreateRequest) dto.CreateResponse {
	if err := s.ValidateRequest(req); err != nil {
		return dto.CreateResponse{
			ErrorCode: "10",
			ErrorDesc: "invalid request",
		}
	}

	if err := s.repo.CheckDuplicateID(req.ID); err != nil {
		return dto.CreateResponse{
			ErrorCode: "10",
			ErrorDesc: "duplicate id",
		}
	}

	todo := models.Todo{
		ID:     req.ID,
		Title:  req.Title,
		Date:   req.Date,
		Status: req.Status,
		Image:  req.Image,
	}

	if err := s.repo.CreateTodo(todo); err != nil {
		return dto.CreateResponse{
			ErrorCode: "10",
			ErrorDesc: "cannot create todo",
		}
	}

	return dto.CreateResponse{
		ErrorCode: "00",
		ErrorDesc: "success",
	}
}

func (s service) ValidateRequest(req dto.CreateRequest) error {
	if err := configs.AppConfig.Validator.Struct(req); err != nil {
		log.Errorf("[Create] validate error: %s", err)
		return err
	}
	_, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		log.Errorf("[Create] validate date error: %s", err)
		return err
	}

	if req.Status != "IN_PROGRESS" && req.Status != "COMPLETE" {
		log.Errorf("[Create] validate invalid")
		return errors.New("invalid status")
	}
	return nil
}

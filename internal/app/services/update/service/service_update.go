package service

import (
	"errors"
	"github.com/napnap11/todo-api/internal/app/services/update/dto"
	"github.com/napnap11/todo-api/internal/app/services/update/repo"
	"github.com/napnap11/todo-api/internal/pkg/configs"
	"github.com/napnap11/todo-api/internal/pkg/models"
	"github.com/napnap11/todo-api/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"time"
)

type Service interface {
	Update(req dto.UpdateRequest) dto.UpdateResponse
}

type service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) Service {
	return service{repo: repo}
}

func (s service) Update(req dto.UpdateRequest) dto.UpdateResponse {
	if err := s.ValidateRequest(req); err != nil {
		return dto.UpdateResponse{
			ErrorCode: "10",
			ErrorDesc: "invalid request",
		}
	}
	todo, err := s.repo.GetTodoById(req.ID)
	if err != nil {
		if err == repository.ErrorNotFound {
			return dto.UpdateResponse{
				ErrorCode: "10",
				ErrorDesc: "not found",
			}
		}
		return dto.UpdateResponse{
			ErrorCode: "10",
			ErrorDesc: "internal error",
		}
	}

	todo = models.Todo{
		ID:     req.ID,
		Title:  req.Title,
		Date:   req.Date,
		Status: req.Status,
		Image:  req.Image,
	}

	if err := s.repo.UpdateTodo(todo); err != nil {
		return dto.UpdateResponse{
			ErrorCode: "10",
			ErrorDesc: "cannot update todo",
		}
	}

	return dto.UpdateResponse{
		ErrorCode: "00",
		ErrorDesc: "success",
	}
}

func (s service) ValidateRequest(req dto.UpdateRequest) error {
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

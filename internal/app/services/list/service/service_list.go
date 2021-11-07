package service

import (
	"errors"
	"github.com/napnap11/todo-api/internal/app/services/list/dto"
	"github.com/napnap11/todo-api/internal/app/services/list/repo"
	"github.com/napnap11/todo-api/internal/pkg/configs"
	"github.com/napnap11/todo-api/internal/pkg/models"
	log "github.com/sirupsen/logrus"
	"sort"
	"strings"
)

type Service interface {
	List(req dto.ListRequest) dto.ListResponse
}

type service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) Service {
	return service{repo: repo}
}

func (s service) List(req dto.ListRequest) dto.ListResponse {
	if err := s.ValidateRequest(req); err != nil {
		return dto.ListResponse{
			ErrorCode: "10",
			ErrorDesc: "invalid request",
		}
	}
	todos, err := s.repo.GetTodos()
	if err != nil {
		return dto.ListResponse{
			ErrorCode: "10",
			ErrorDesc: "internal error",
		}
	}

	todos, err = s.FilterSearch(req, todos)
	if err != nil {
		return dto.ListResponse{
			ErrorCode: "10",
			ErrorDesc: "internal error",
		}
	}

	todos, err = s.Sort(req, todos)
	if err != nil {
		return dto.ListResponse{
			ErrorCode: "10",
			ErrorDesc: "internal error",
		}
	}

	return dto.ListResponse{
		Todos:     todos,
		ErrorCode: "00",
		ErrorDesc: "success",
	}
}

func (s service) ValidateRequest(req dto.ListRequest) error {
	if err := configs.AppConfig.Validator.Struct(req); err != nil {
		log.Errorf("[List] validate error: %s", err)
		return err
	}

	if req.SortBy != "title" && req.SortBy != "date" && req.SortBy != "status" && req.SortBy != "" {
		log.Errorf("[List] invalid sort_by")
		return errors.New("invalid sort_by")
	}

	if req.SortType != "desc" && req.SortType != "asc" && req.SortType != "" {
		log.Errorf("[List] invalid sort_type")
		return errors.New("invalid sort_type")
	}

	if req.SortBy != "" && req.SortType == "" {
		log.Errorf("[List] invalid sort_type")
		return errors.New("invalid sort_type")
	}

	return nil
}

func (s service) FilterSearch(req dto.ListRequest, todos []models.Todo) ([]models.Todo, error) {
	filtered := make([]models.Todo, 0)
	if req.Title == "" && req.Description == "" {
		return todos, nil
	}
	if req.Title != "" {
		for _, todo := range todos {
			if strings.Contains(todo.Title, req.Title) {
				filtered = append(filtered, todo)
			}
		}
	}
	if req.Description != "" {
		for _, todo := range todos {
			if strings.Contains(todo.Description, req.Description) {
				filtered = append(filtered, todo)
			}
		}
	}
	return filtered, nil
}

func (s service) Sort(req dto.ListRequest, todos []models.Todo) ([]models.Todo, error) {
	sort.SliceStable(todos, func(i, j int) bool {
		switch req.SortBy {
		case "title":
			switch req.SortType {
			case "desc":
				return todos[i].Title < todos[j].Title
			case "asc":
				return todos[i].Title > todos[j].Title
			}
		case "date":
			switch req.SortType {
			case "desc":
				return todos[i].Date < todos[j].Date
			case "asc":
				return todos[i].Date > todos[j].Date
			}
		case "status":
			switch req.SortType {
			case "desc":
				return todos[i].Status < todos[j].Status
			case "asc":
				return todos[i].Status > todos[j].Status
			}
		}
		return false
	})
	return todos, nil
}

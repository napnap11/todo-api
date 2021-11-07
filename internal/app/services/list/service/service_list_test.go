package service

import (
	"github.com/napnap11/todo-api/internal/app/services/list/dto"
	"github.com/napnap11/todo-api/internal/pkg/configs"
	"github.com/napnap11/todo-api/internal/pkg/models"

	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v8"

	"testing"
)
type mockRepo struct {
}

func(r mockRepo) GetTodos() ([]models.Todo, error) {
	return []models.Todo{}, nil
}

func TestServiceCreateSuccessCase(t *testing.T) {
	configs.AppConfig.Validator = validator.New(&validator.Config{TagName: "validate"})
	req := dto.ListRequest{
		SortBy:      "title",
		SortType:    "desc",
		Title:       "",
		Description: "",
	}
	repo := mockRepo{}
	s := NewService(repo)
	resp := s.List(req)

	expected := dto.ListResponse{
		Todos:     []models.Todo{},
		ErrorCode: "00",
		ErrorDesc: "success",
	}

	assert.Equal(t, expected.ErrorCode, resp.ErrorCode)
	assert.Equal(t, expected.ErrorDesc, resp.ErrorDesc)
	assert.Equal(t, expected.Todos, resp.Todos)
}

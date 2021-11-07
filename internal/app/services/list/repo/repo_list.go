package repo

import "github.com/napnap11/todo-api/internal/pkg/models"

type Repository interface {
	GetTodos() ([]models.Todo, error)
}

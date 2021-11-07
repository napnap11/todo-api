package repo

import "github.com/napnap11/todo-api/internal/pkg/models"

type Repository interface {
	CheckDuplicateID(id string) error
	CreateTodo(todo models.Todo) error
}

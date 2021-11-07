package repo

import "github.com/napnap11/todo-api/internal/pkg/models"

type Repository interface {
	GetTodoById(id string) (models.Todo, error)
	UpdateTodo(todo models.Todo) error
}

package repository

import (
	"errors"

	"github.com/napnap11/todo-api/internal/pkg/models"
)

type Repository struct {
	Jobs chan Job
}

func NewRepository() *Repository {
	InitJob()
	return &Repository{}
}

func (r Repository) GetTodos() ([]models.Todo, error) {
	job := NewReadTodosJob()
	r.Jobs <- job
	if err := <-job.ExitChan(); err != nil {
		return nil, err
	}

	todos := <-job.todos
	return todos, nil
}

func (r Repository) WriteTodos(todos []models.Todo) error {
	job := NewWriteTodosJob(todos)
	r.Jobs <- job
	return <-job.ExitChan()
}

func (r Repository) CheckDuplicateID(id string) error {
	todos, err := r.GetTodos()
	if err != nil {
		return err
	}
	for _, todo := range todos {
		if todo.ID == id {
			return errors.New("duplication id")
		}
	}
	return nil
}
func (r Repository) CreateTodo(newTodo models.Todo) error {
	todos, err := r.GetTodos()
	if err != nil {
		return err
	}
	todos = append(todos, newTodo)
	return r.WriteTodos(todos)
}

func (r Repository) GetTodosWithSortAndSearch(sortBy, sortType, title, description string) ([]models.Todo, error) {
	todos, err := r.GetTodos()
	if err != nil {
		return []models.Todo{}, err
	}
	return todos, nil
}

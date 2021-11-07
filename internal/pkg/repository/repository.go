package repository

import (
	"errors"

	"github.com/napnap11/todo-api/internal/pkg/models"
)

var ErrorNotFound = errors.New("not found")

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

	todosSlice := make([]models.Todo, 0)
	for _, todo := range todos {
		todosSlice = append(todosSlice, todo)
	}
	return todosSlice, nil
}

func (r Repository) GetTodoById(id string) (models.Todo, error) {
	job := NewReadTodosJob()
	r.Jobs <- job
	if err := <-job.ExitChan(); err != nil {
		return models.Todo{}, err
	}

	todos := <-job.todos
	todo, isValid := todos[id]
	if isValid {
		return todo, nil
	}
	return models.Todo{}, ErrorNotFound
}

func (r Repository) WriteTodos(todos []models.Todo) error {
	job := NewWriteTodosJob(todos)
	r.Jobs <- job
	return <-job.ExitChan()
}

func (r Repository) CheckDuplicateID(id string) error {
	_, err := r.GetTodoById(id)
	if err != nil && err != ErrorNotFound {
		return err
	}
	if err == ErrorNotFound {
		return nil
	}
	return errors.New("duplication id")
}
func (r Repository) CreateTodo(newTodo models.Todo) error {
	todos, err := r.GetTodos()
	if err != nil {
		return err
	}
	todos = append(todos, newTodo)
	return r.WriteTodos(todos)
}

func (r Repository) UpdateTodo(todo models.Todo) error {
	job := NewReadTodosJob()
	r.Jobs <- job
	if err := <-job.ExitChan(); err != nil {
		return err
	}
	todos := <-job.todos

	todos[todo.ID] = todo
	todosSlice := make([]models.Todo, 0)
	for _, todo := range todos {
		todosSlice = append(todosSlice, todo)
	}

	return r.WriteTodos(todosSlice)
}

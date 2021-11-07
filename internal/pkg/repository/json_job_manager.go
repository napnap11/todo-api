package repository

import (
	"encoding/json"
	"github.com/napnap11/todo-api/internal/pkg/models"
	"io/ioutil"
)

type Job interface {
	ExitChan() chan error
	Modify(data []models.Todo) ([]models.Todo, error)
}

func ProcessJobs(jobs chan Job, db string) {
	for {
		j := <-jobs
		var data []models.Todo
		content, err := ioutil.ReadFile(db)
		if err == nil {
			if err = json.Unmarshal(content, &data); err == nil {
				dataModified, err := j.Modify(data)
				if err == nil && dataModified != nil {
					b, err := json.Marshal(dataModified)
					if err == nil {
						err = ioutil.WriteFile(db, b, 0644)
					}
				}
			}
		}

		j.ExitChan() <- err
	}
}

type ReadTodosJob struct {
	todos    chan []models.Todo
	exitChan chan error
}

func NewReadTodosJob() *ReadTodosJob {
	return &ReadTodosJob{
		todos:    make(chan []models.Todo, 1),
		exitChan: make(chan error, 1),
	}
}
func (j ReadTodosJob) ExitChan() chan error {
	return j.exitChan
}
func (j ReadTodosJob) Modify(todos []models.Todo) ([]models.Todo, error) {
	j.todos <- todos

	return nil, nil
}

type WriteTodosJob struct {
	newTodos []models.Todo
	exitChan chan error
}

func NewWriteTodosJob(newTodos []models.Todo) *WriteTodosJob {
	return &WriteTodosJob{
		newTodos: newTodos,
		exitChan: make(chan error, 1),
	}
}
func (j WriteTodosJob) ExitChan() chan error {
	return j.exitChan
}
func (j WriteTodosJob) Modify(todos []models.Todo) ([]models.Todo, error) {
	return j.newTodos, nil
}
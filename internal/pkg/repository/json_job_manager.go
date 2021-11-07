package repository

import (
	"encoding/json"
	"github.com/napnap11/todo-api/internal/pkg/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type Job interface {
	ExitChan() chan error
	Modify(data map[string]models.Todo) (map[string]models.Todo, error)
}

func InitJob() chan Job {
	db := "./db.json"
	jobs := make(chan Job)
	log.Infof("init job")
	go ProcessJobs(jobs, db)
	return jobs
}

func ProcessJobs(jobs chan Job, db string) {
	for {
		log.Infof("wait job")
		j := <-jobs
		log.Infof("get a job")
		data := make(map[string]models.Todo, 0)
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
	todos    chan map[string]models.Todo
	exitChan chan error
}

func NewReadTodosJob() *ReadTodosJob {
	return &ReadTodosJob{
		todos:    make(chan map[string]models.Todo, 1),
		exitChan: make(chan error, 1),
	}
}
func (j ReadTodosJob) ExitChan() chan error {
	return j.exitChan
}
func (j ReadTodosJob) Modify(todos map[string]models.Todo) (map[string]models.Todo, error) {
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
func (j WriteTodosJob) Modify(todos map[string]models.Todo) (map[string]models.Todo, error) {
	data := make(map[string]models.Todo, 0)
	for _, newTodo := range j.newTodos {
		data[newTodo.ID] = newTodo
	}
	return data, nil
}

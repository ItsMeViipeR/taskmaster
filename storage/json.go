package storage

import (
	"encoding/json"
	"os"

	"github.com/ItsMeViipeR/taskmaster/task"
)

type JSON struct {
	filePath string
	tasks    []task.Task
}

func NewJSON(path string) *JSON {
	return &JSON{
		filePath: path,
		tasks:    []task.Task{},
	}
}

func (j *JSON) save() error {
	f, err := os.Create(j.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(j.tasks)
}

func (j *JSON) Load() error {
	f, err := os.Open(j.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(&j.tasks)
}

func (j *JSON) AddTask(t *task.Task) error {
	j.tasks = append(j.tasks, *t)
	return j.save()
}

func (j *JSON) Tasks() []task.Task {
	return j.tasks
}

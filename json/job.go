package json

import (
	"encoding/json"
	"github.com/Tiofx/jobShop/base"
)

type Job base.Job

func (j Job) String() string {
	bytes, err := json.Marshal(j.toTasks())
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

type tasks []Task

func (j Job) toTasks() tasks {
	tasks := make(tasks, len(j))
	for i, task := range j {
		tasks[i] = Task(task)
	}

	return tasks
}

package json

import (
	"encoding/json"
	"github.com/Tiofx/jobShop/base"
)

type Task base.Task

func NewTask(bytes []byte) Task {
	var res Task
	json.Unmarshal(bytes, &res)

	return res
}

func (t Task) String() string {
	bytes, err := json.Marshal(base.Task(t))
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
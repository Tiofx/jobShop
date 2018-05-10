package json

import "github.com/Tiofx/jobShop/base"

type TaskString string

func (ts TaskString) ToTaskOrPanic() base.Task {
	return base.Task(NewTask([]byte(ts)))
}
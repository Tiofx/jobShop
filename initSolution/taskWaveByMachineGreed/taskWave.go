package taskWaveByMachineGreed

import (
	. "github.com/Tiofx/jobShop/base"
	"fmt"
	"github.com/Tiofx/jobShop/state"
)

type TaskInfo struct {
	Job int
	*Task
}

type TaskInfoSet []TaskInfo

type TaskWave []TaskInfoSet

func (tw TaskWave) Add(job int, task *Task) {
	currentTasks := tw[task.Machine]
	tw[task.Machine] = append(currentTasks, TaskInfo{job, task})
}

func (tw TaskWave) GetBiggest() []TaskInfo {
	var (
		biggest        int
		indexOfBiggest int
	)

	for machine, arr := range tw {
		if currentLen := len(arr); currentLen > biggest {
			indexOfBiggest = machine
			biggest = currentLen
		}
	}

	return tw[indexOfBiggest]
}

func (ti TaskInfo) String() string {
	if ti.Task == nil {
		return fmt.Sprintf("{%v {no Task}}", ti.Job)
	}

	return fmt.Sprintf("{%v %v}", ti.Job, *ti.Task)
}

type StateBy state.State

func (s StateBy) NextTaskWave() TaskWave {
	tw := make(TaskWave, s.Jobs.MachineNumber())

	for i := range s.Jobs {
		s := state.State(s)
		if task, ok := s.NextTaskOf(i); ok {
			tw.Add(i, task)
		}
	}

	return tw
}

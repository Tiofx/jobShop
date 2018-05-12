package taskWaveByMachineGreed

import (
	. "github.com/Tiofx/jobShop/base"
	"fmt"
	"github.com/Tiofx/jobShop/state"
)

type TaskInfo struct {
	Job uint64
	*Task
}

type TaskInfoSet []TaskInfo

type TaskWave []TaskInfoSet

func (tw TaskWave) Add(job uint64, task *Task) {
	currentTasks := tw[task.Machine]
	tw[task.Machine] = append(currentTasks, TaskInfo{job, task})
}

func (tw TaskWave) GetBiggest() []TaskInfo {
	var (
		biggest        uint64
		indexOfBiggest uint64
	)

	for machine, arr := range tw {
		if currentLen := uint64(len(arr)); currentLen > biggest {
			indexOfBiggest = uint64(machine)
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
		if task, ok := s.NextTaskOf(uint64(i)); ok {
			tw.Add(uint64(i), task)
		}
	}

	return tw
}

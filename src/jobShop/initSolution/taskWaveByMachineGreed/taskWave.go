package taskWaveByMachineGreed

import (
	. "jobShop/base"
	"fmt"
)

type TaskInfo struct {
	Job int
	*Task
}

type TaskInfoSet []TaskInfo

type TaskWave map[Machine]TaskInfoSet

func (tw TaskWave) Add(job int, task *Task) {
	currentTasks := tw[task.Machine]
	tw[task.Machine] = append(currentTasks, TaskInfo{job, task})
}

func (tw TaskWave) GetBiggest() []TaskInfo {
	var (
		biggest        int
		indexOfBiggest Machine
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
package taskWaveByMachineGreed

import (
	"github.com/Tiofx/jobShop/util"
	. "github.com/Tiofx/jobShop/state"
)

const OptimalPermutationLimit = 3

type Resolver struct {
	MaxTasksOnWave uint64
	State
}

func (r *Resolver) IsBetterThan(second Resolver) bool {
	return r.State.IsBetterThan(second.State)
}

func (r *Resolver) Copy() Resolver {
	return Resolver{MaxTasksOnWave: r.MaxTasksOnWave, State: r.State.Copy()}
}

func (r *Resolver) ExecuteByTaskInfo(info TaskInfo) {
	taskPosition := r.Executed[info.Job]
	r.Execute(info.Job, taskPosition)
}

func (r Resolver) NextTaskWave() TaskWave {
	tw := make(TaskWave, r.Jobs.MachineNumber())

	for i := range r.Jobs {
		if task, ok := r.NextTaskOf(uint64(i)); ok {
			tw.Add(uint64(i), task)
		}
	}

	return tw
}

func (r *Resolver) Next() Resolver {
	tasks := r.NextTaskWave().GetBiggest()
	if uint64(len(tasks)) > r.MaxTasksOnWave {
		tasks = tasks[:r.MaxTasksOnWave-1]
	}
	nextSolution := r.GreedChoice(tasks)
	nextSolution.Parent = &r.State

	return nextSolution
}

func (r Resolver) GreedChoice(tasks TaskInfoSet) Resolver {
	var best *Resolver

	c := util.Combination(uint64(len(tasks)))
	for tasksOrder, isChanOpen := <-c; isChanOpen; tasksOrder, isChanOpen = <-c {
		newState := r.Copy()

		for _, index := range tasksOrder {
			currentTask := tasks[index]
			newState.ExecuteByTaskInfo(currentTask)
		}

		if best == nil || newState.IsBetterThan(*best) {
			best = &newState
		}
	}

	return *best
}

func (r Resolver) FindSolution() State {
	var currentState Resolver

	for currentState = r; !currentState.IsFinish(); currentState = currentState.Next() {

	}

	return currentState.State
}

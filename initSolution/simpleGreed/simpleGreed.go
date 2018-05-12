package simpleGreed

import (
	"github.com/Tiofx/jobShop/state"
)

type job uint64
type task uint64

type TaskWave []*task

func (tw TaskWave) addTask(job job, task task) {
	tw[job] = &task
}

type Resolver struct {
	state.State
	Parent *Resolver
}

func (r *Resolver) IsBetterThan(second Resolver) bool {
	return r.State.IsBetterThan(second.State)
}

func (r *Resolver) ExecuteBy(job uint64, task task) {
	r.Execute(job, uint64(task))
}

func (r *Resolver) Copy() Resolver {
	return Resolver{State: r.State.Copy()}
}

func (r *Resolver) nextTaskWave() TaskWave {
	tasks := make(TaskWave, len(r.Jobs))

	for i := uint64(0); i < uint64(len(r.Jobs)); i++ {
		if taskIndex, ok := r.NextTaskIndexOf(i); ok {
			tasks.addTask(job(i), task(taskIndex))
		}
	}

	return tasks
}

func (r *Resolver) Next() Resolver {
	tasks := r.nextTaskWave()
	choice := r.GreedChoice(tasks)
	choice.Parent = r

	return choice
}

func (r *Resolver) GreedChoice(tasks TaskWave) Resolver {
	var best *Resolver
	copy := r.Copy()

	for job, task := range tasks {
		if task == nil {
			continue
		}

		task := *task
		r.ExecuteBy(uint64(job), task)

		if best == nil {
			best = &copy
			best.ExecuteBy(uint64(job), task)
		} else if r.IsBetterThan(*best) {
			best.Undo()
			best.ExecuteBy(uint64(job), task)
		}

		r.Undo()

	}

	return *best
}

func (r Resolver) FindSolution() state.State {
	var currentState Resolver

	for currentState = r; !currentState.IsFinish(); currentState = currentState.Next() {

	}

	return currentState.State
}

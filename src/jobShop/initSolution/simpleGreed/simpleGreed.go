package simpleGreed

import (
	"jobShop/state"
)

type job int
type task int

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

func (r *Resolver) ExecuteBy(job int, task task) {
	r.Execute(job, int(task))
}

func (r *Resolver) Copy() Resolver {
	return Resolver{State: r.State.Copy()}
}

func (r *Resolver) nextTaskWave() TaskWave {
	tasks := make(TaskWave, len(r.Jobs))

	for i := 0; i < len(r.Jobs); i++ {
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
		r.ExecuteBy(job, task)

		if best == nil {
			best = &copy
			best.ExecuteBy(job, task)
		} else if r.IsBetterThan(*best) {
			best.Undo()
			best.ExecuteBy(job, task)
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

package simpleGreed

import (
	"jobShop/state"
)

type job int
type task int

type TaskWave map[job]task

func (tw TaskWave) addTask(job job, task task) {
	tw[job] = task
}

type Resolver struct {
	state.State
	Parent *Resolver
}

func (r *Resolver) IsBetterThan(second Resolver) bool {
	return r.State.IsBetterThan(second.State)
}

func (resolver *Resolver) ExecuteBy(job job, task task) {
	resolver.Execute(int(job), int(task))
}

func (s *Resolver) Copy() Resolver {
	return Resolver{State: s.State.Copy()}
}

func (r *Resolver) nextTaskWave() (tasks TaskWave) {
	tasks = make(TaskWave)

	for i := 0; i < len(r.Jobs); i++ {
		if taskIndex, ok := r.NextTaskIndexOf(i); ok {
			tasks.addTask(job(i), task(taskIndex))
		}
	}

	return
}

func (r *Resolver) Next() Resolver {
	tasks := r.nextTaskWave()
	choice := r.GreedChoice(tasks)
	choice.Parent = r

	return choice
}

func (resolver *Resolver) GreedChoice(tasks TaskWave) Resolver {
	var best *Resolver
	copy := resolver.Copy()

	for job, task := range tasks {

		resolver.ExecuteBy(job, task)

		if best == nil {
			best = &copy
			best.ExecuteBy(job, task)
		} else if resolver.IsBetterThan(*best) {
			best.Undo()
			best.ExecuteBy(job, task)
		}

		resolver.Undo()

	}

	return *best
}

func (s Resolver) FindSolution() state.State {
	var currentState Resolver

	for currentState = s; !currentState.IsFinish(); currentState = currentState.Next() {

	}

	return currentState.State
}

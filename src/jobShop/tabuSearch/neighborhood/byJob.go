package neighborhood

import (
	"jobShop/base"
	"jobShop/state"
	"jobShop/tabuSearch/graph_state"
)

type byJob struct {
	JobState *state.State
	Graph    *graph_state.DisjunctiveGraph
}

type taskPosition int

type job int
type criticalTasks map[base.Machine][]taskPosition

func (r byJob) taskPositionFor(critical job) criticalTasks {
	tasks := make(criticalTasks)

	for machine, jobOrder := range *r.Graph {
		for index, job := range jobOrder {
			if int(job) == int(critical) {
				machine := base.Machine(machine)
				tasks[machine] = append(tasks[machine], taskPosition(index))
			}
		}
	}

	return tasks
}

func (r byJob) generateFor(job job, tasks criticalTasks) (iterator <-chan Move) {
	ch := make(chan Move)

	go func(consumer chan<- Move) {
		defer close(consumer)

		for machine, tasksOfMachine := range tasks {
			for _, task := range tasksOfMachine {
				for taskIndex := range (*r.Graph)[machine] {
					if taskIndex != int(task) {
						consumer <- Move{Machine: int(machine), I: taskIndex, J: int(task)}
					}
				}
			}
		}
	}(ch)

	return ch
}
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

func (r byJob) generateFor(tasks criticalTasks) []Move {
	var res []Move

	for machine := 0; machine < r.JobState.Jobs.MachineNumber(); machine++ {
		machine := base.Machine(machine)
		tasksOfMachine := tasks[machine]

		for _, task := range tasksOfMachine {
			for taskIndex := range (*r.Graph)[machine] {
				if taskIndex != int(task) && !r.theSameJob(machine, taskIndex, task) {
					res = append(res, Move{Machine: int(machine), I: taskIndex, J: int(task)})
				}
			}
		}
	}

	return res
}
func (r byJob) theSameJob(machine base.Machine, taskIndex int, task taskPosition) bool {
	return (*r.Graph)[machine][taskIndex] == (*r.Graph)[machine][task]
}

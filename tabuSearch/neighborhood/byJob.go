package neighborhood

import (
	"jobShop/base"
	"jobShop/state"
	"jobShop/tabuSearch/graph_state"
	"jobShop/util"
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
				if r.canMove(taskIndex, int(task), machine) {
					res = append(res, Move{Machine: int(machine), I: taskIndex, J: int(task)})
				}
			}
		}
	}

	return res
}
func (r byJob) canMove(taskIndex int, task int, machine base.Machine) bool {
	return taskIndex != task && !r.theSameJob(machine, taskIndex, task) && !r.willLeadToImpossibleTaskOrder(machine, taskIndex, task)
}
func (r byJob) theSameJob(machine base.Machine, taskIndex int, task int) bool {
	return (*r.Graph)[machine][taskIndex] == (*r.Graph)[machine][task]
}
func (r byJob) willLeadToImpossibleTaskOrder(machine base.Machine, i int, j int) bool {
	jobI, jobJ := (*r.Graph)[machine][i], (*r.Graph)[machine][j]
	i1, i2 := util.Min(i, j), util.Max(i, j)
	for _, job := range (*r.Graph)[machine][i1+1:i2] {
		if job == jobI || job == jobJ {
			return true
		}
	}

	return false
}

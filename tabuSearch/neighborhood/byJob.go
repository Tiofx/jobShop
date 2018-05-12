package neighborhood

import (
	"github.com/Tiofx/jobShop/base"
	"github.com/Tiofx/jobShop/state"
	"github.com/Tiofx/jobShop/tabuSearch/graph_state"
	"github.com/Tiofx/jobShop/util"
)

type byJob struct {
	JobState *state.State
	Graph    *graph_state.DisjunctiveGraph
}

type taskPosition uint64

type job uint64
type criticalTasks map[base.Machine][]taskPosition

func (r byJob) taskPositionFor(critical job) criticalTasks {
	tasks := make(criticalTasks)

	for machine, jobOrder := range *r.Graph {
		for index, job := range jobOrder {
			if uint64(job) == uint64(critical) {
				machine := base.Machine(machine)
				tasks[machine] = append(tasks[machine], taskPosition(index))
			}
		}
	}

	return tasks
}

func (r byJob) generateFor(tasks criticalTasks) []Move {
	var res []Move

	for machine := uint64(0); machine < r.JobState.Jobs.MachineNumber(); machine++ {
		machine := base.Machine(machine)
		tasksOfMachine := tasks[machine]

		for _, task := range tasksOfMachine {
			for taskIndex := range (*r.Graph)[machine] {
				if r.canMove(uint64(taskIndex), uint64(task), machine) {
					res = append(res, Move{Machine: uint64(machine), I: uint64(taskIndex), J: uint64(task)})
				}
			}
		}
	}

	return res
}
func (r byJob) canMove(taskIndex uint64, task uint64, machine base.Machine) bool {
	return taskIndex != task && !r.theSameJob(machine, taskIndex, task) && !r.willLeadToImpossibleTaskOrder(machine, taskIndex, task)
}
func (r byJob) theSameJob(machine base.Machine, taskIndex uint64, task uint64) bool {
	return (*r.Graph)[machine][taskIndex] == (*r.Graph)[machine][task]
}
func (r byJob) willLeadToImpossibleTaskOrder(machine base.Machine, i uint64, j uint64) bool {
	jobI, jobJ := (*r.Graph)[machine][i], (*r.Graph)[machine][j]
	i1, i2 := util.Min(i, j), util.Max(i, j)
	for _, job := range (*r.Graph)[machine][i1+1:i2] {
		if job == jobI || job == jobJ {
			return true
		}
	}

	return false
}

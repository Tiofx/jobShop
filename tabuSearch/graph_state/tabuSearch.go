package graph_state

import (
	"github.com/Tiofx/jobShop/state"
	"github.com/Tiofx/jobShop/base"
	"fmt"
)

type job int
type jobSequence []job
type DisjunctiveGraph []jobSequence

type task struct {
	job  job
	task int
}
type CriticalPath []task

type State struct {
	Jobs     base.Jobs
	Executed []int
	DisjunctiveGraph
}

func NewState(jobs base.Jobs, graph DisjunctiveGraph) State {
	return State{
		Jobs:             jobs,
		Executed:         make([]int, jobs.MachineNumber()),
		DisjunctiveGraph: graph,
	}
}

func (s *State) ExecuteNextOf(machine base.Machine) {
	s.Executed[machine]++
}

func (s *State) NextOf(machine base.Machine) int {
	return int(s.DisjunctiveGraph[machine][s.Executed[machine]])
}

//func (s *State) Makespan() int {
//	return s.To().Makespan()
//}
func (graph DisjunctiveGraph) AddTo(machine base.Machine, nextJob int) {
	graph[machine] = append(graph[machine], job(nextJob))
}

func From(state state.State) (graph DisjunctiveGraph) {
	graph = make(DisjunctiveGraph, state.Jobs.MachineNumber())
	currentTaskNumber := make([]int, len(state.Jobs))

	for _, job := range state.JobOrder {
		task := state.Jobs[job][currentTaskNumber[job]]
		graph.AddTo(task.Machine, job)
		currentTaskNumber[job]++
	}

	return
}

func (s *State) To(jobState *state.State) (success bool) {
	jobNumber := len(jobState.Jobs)

	for !jobState.IsFinish() {
		hasNoExecutedTaskDueIteration := true

		for job := 0; job < jobNumber; job++ {
			if task, ok := jobState.NextTaskOf(job); ok {
				if s.NextOf(task.Machine) == job {
					taskIndex, _ := jobState.NextTaskIndexOf(job)
					jobState.Execute(job, taskIndex)

					s.ExecuteNextOf(task.Machine)

					hasNoExecutedTaskDueIteration = false
				}
			}
		}

		if hasNoExecutedTaskDueIteration {
			return false
		}
	}

	return true
}

func (g DisjunctiveGraph) String() string {
	var res string

	for i, js := range g {
		res += fmt.Sprintf("Machine %v: %v\n", i, js)
	}

	return res
}

package tabuSearch

import (
	"jobShop/state"
	"jobShop/base"
	"github.com/getlantern/deepcopy"
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

func neighbours() {}

//func move()       {}

func (s State) ExecuteNextOf(machine base.Machine) {
	s.Executed[machine]++
}

func (s State) NextOf(machine base.Machine) int {
	return int(s.DisjunctiveGraph[machine][s.Executed[machine]])
}

func (s State) Makespan() int {
	return s.To().Makespan()
}
func (graph DisjunctiveGraph) AddTo(machine base.Machine, nextJob int) {
	graph[machine] = append(graph[machine], job(nextJob))
}

func From(state state.State) (graph DisjunctiveGraph) {
	graph = make(DisjunctiveGraph, len(state.Jobs.ToMachines()))
	currentTaskNumber := make([]int, len(state.Jobs))

	for _, job := range state.JobOrder {
		task := state.Jobs[job][currentTaskNumber[job]]
		graph.AddTo(task.Machine, job)
		currentTaskNumber[job]++
	}

	return
}

func (s State) To() (jobState state.State) {
	jobState = state.NewState(s.Jobs)
	jobNumber := len(jobState.Jobs)

	for !jobState.IsFinish() {
		isNotExecutedTaskDueIteration := true

		for job := 0; job < jobNumber; job++ {
			if task, ok := jobState.NextTaskOf(job); ok {
				if s.NextOf(task.Machine) == job {
					taskIndex, _ := jobState.NextTaskIndexOf(job)
					jobState.Execute(job, taskIndex)

					s.ExecuteNextOf(task.Machine)

					isNotExecutedTaskDueIteration = false
				}
			}
		}

		if isNotExecutedTaskDueIteration {
			panic("endless loop")
		}
	}

	return
}

//Deprected
func (graph DisjunctiveGraph) To(jobs base.Jobs) (jobState *state.State, exist bool) {
	stateOfGraph := State{
		Executed:         make([]int, len(graph)),
		DisjunctiveGraph: graph,
	}
	newState := state.NewState(jobs)
	jobState = &newState
	jobNumber := len(jobState.Jobs)

	for !jobState.IsFinish() {
		isNotExecutedTaskDueIteration := true

		for job := 0; job < jobNumber; job++ {
			if task, ok := jobState.NextTaskOf(job); ok {
				if stateOfGraph.NextOf(task.Machine) == job {
					taskIndex, _ := jobState.NextTaskIndexOf(job)
					jobState.Execute(job, taskIndex)

					stateOfGraph.ExecuteNextOf(task.Machine)

					isNotExecutedTaskDueIteration = false
				}
			}
		}

		if isNotExecutedTaskDueIteration {
			return nil, false
		}
	}

	return jobState, true
}

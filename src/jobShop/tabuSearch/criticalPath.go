package tabuSearch

import (
	"jobShop/base"
	"jobShop/state"
)

type taskPosition int

type criticalTasks map[base.Machine][]taskPosition

func criticalJob(state state.State) job {
	var (
		criticalJob       int
		timeOfCriticalJob int
	)

	for job, time := range state.JobTimeWave {
		if time > timeOfCriticalJob {
			criticalJob = job
			timeOfCriticalJob = time
		}
	}

	return job(criticalJob)
}

func (graph DisjunctiveGraph) criticalTaskPositionFor(criticalJob job) criticalTasks {
	tasks := make(criticalTasks)

	for machine, jobOrder := range graph {
		for index, job := range jobOrder {
			if job == criticalJob {
				machine := base.Machine(machine)
				tasks[machine] = append(tasks[machine], taskPosition(index))
			}
		}
	}

	return tasks
}

func (s *Solver) worstJobs() (iterator <-chan move) {
	ch := make(chan move)

	criticalJob := criticalJob(s.CurrentSolution.jobState)
	tasks := s.CurrentSolution.graph.criticalTaskPositionFor(criticalJob)
	graph := s.CurrentSolution.graph

	go func(consumer chan<- move) {
		defer close(consumer)

		for machine, tasksOfMachine := range tasks {
			for _, task := range tasksOfMachine {
				for taskIndex, _ := range graph[machine] {
					if taskIndex != int(task) {
						consumer <- move{machine: int(machine), i: taskIndex, j: int(task)}
					}
				}
			}
		}
	}(ch)

	return ch
}

package tabuSearch

import (
	"jobShop/base"
	"jobShop/state"
	"math"
	"github.com/getlantern/deepcopy"
)

type Solver struct {
	jobs            base.Jobs
	CurrentSolution *neighbour
	tabuList        tabuList

	bestLocal    *neighbour
	bestSolution *neighbour
}

func (s *Solver) BestLocalMakespan() int {
	if s.bestLocal == nil {
		return math.MaxInt64
	}

	return s.bestLocal.Makespan()
}

func (s *Solver) BestMakespan() int {
	if s.bestSolution == nil {
		return math.MaxInt64
	}

	return s.bestSolution.Makespan()
}

func NewSolver(state state.State) Solver {
	initialSolution := From(state)
	best, _ := newNeighbour(initialSolution, state.Jobs)

	var bestCopy neighbour
	best.copyIn(&bestCopy)

	return Solver{
		jobs:            state.Jobs,
		CurrentSolution: best,
		bestSolution:    &bestCopy,
	}
}

func (s *Solver) setUpBestNeighbour() (isBestSolutionChanged bool) {
	criticalJob := criticalJob(s.CurrentSolution.jobState)
	tasks := s.CurrentSolution.graph.criticalTaskPositionFor(criticalJob)
	graph := s.CurrentSolution.graph

	s.bestLocal = nil
	for machine, tasksOfMachine := range tasks {
		for _, task := range tasksOfMachine {
			for taskIndex, _ := range graph[machine] {
				if taskIndex != int(task) {
					//todo: refactoring for iterating by moves
					move := move{machine: int(machine), i: taskIndex, j: int(task)}
					s.CurrentSolution.apply(move)

					//TODO: optimize, make checking in tabu before update graph
					success := s.CurrentSolution.updateByGraph()
					//if success && !s.tabuList.contain(*s.CurrentSolution) && s.CurrentSolution.Makespan() < s.BestMakespan() {
					//	s.CurrentSolution.copyIn(s.bestSolution)
					if success && !s.tabuList.contain(*s.CurrentSolution) &&
						(s.bestLocal == nil || s.CurrentSolution.Makespan() < s.bestLocal.Makespan()) {

						if s.bestLocal == nil {
							s.bestLocal = &neighbour{}
						}
						s.CurrentSolution.copyIn(s.bestLocal)
						//isBestSolutionChanged = true
					}

					s.CurrentSolution.redo(move)
				}
			}
		}
	}

	if s.bestLocal.Makespan() < s.BestMakespan() {
		isBestSolutionChanged = true
		s.bestSolution = s.bestLocal
		s.bestLocal = nil
	}

	return
}

func (s *Solver) Next() {
	var current neighbour
	s.CurrentSolution.copyIn(&current)

	isBestSolutionChanged := s.setUpBestNeighbour()

	if !isBestSolutionChanged {
		s.tabuList = append(s.tabuList, current)

		if s.bestLocal != nil {
			s.CurrentSolution = s.bestLocal
			s.bestLocal = nil

		} else {
			s.CurrentSolution = &s.tabuList[0]
			s.tabuList = s.tabuList[1:]
		}

	} else {
		s.tabuList = append(s.tabuList, current)
		s.bestSolution.copyIn(s.CurrentSolution)
	}

	//if isBestSolutionChanged {
	//	s.tabuList = append(s.tabuList, current)
	//	s.bestSolution.copyIn(s.CurrentSolution)
	//
	//} else {
	//	s.CurrentSolution = &s.tabuList[0]
	//	s.tabuList = s.tabuList[1:]
	//}
}

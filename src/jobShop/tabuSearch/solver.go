package tabuSearch

import (
	"jobShop/base"
	"jobShop/state"
	"math"
)

type Solver struct {
	jobs            base.Jobs
	CurrentSolution *neighbour
	tabuList        TabuList

	bestLocal    *neighbour
	bestSolution *neighbour
}

func (s *Solver) GetBest() state.State {
	return s.bestSolution.jobState
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
	best := neighbour{
		jobState: state,
		graph:    initialSolution,
	}

	var bestCopy neighbour
	best.copyIn(&bestCopy)

	return Solver{
		jobs:            state.Jobs,
		CurrentSolution: &best,
		bestSolution:    &bestCopy,
		tabuList:        tabuList{},
	}
}

func (s *Solver) setUpBestNeighbour() (bestMove move) {
	//criticalJob := job(rand.Int() % len(s.jobs))
	s.bestLocal = nil
	for move := range s.worstJobs() {
		s.CurrentSolution.apply(move)

		//TODO: optimize, make checking in tabu before update graph
		//TODO: store impossible move to prevent useless graph update
		success := s.CurrentSolution.updateByGraph()

		if success {
			if s.CurrentSolution.Makespan() < s.bestSolution.Makespan() {
				s.tabuList.Forget(move)
			}

			if s.bestLocal == nil {
				s.bestLocal = &neighbour{}
				s.CurrentSolution.copyIn(s.bestLocal)

			} else if s.CurrentSolution.Makespan() < s.bestLocal.Makespan() &&
				!s.tabuList.Contain(move) {

				bestMove = move
				s.CurrentSolution.copyIn(s.bestLocal)
			}
		}

		s.CurrentSolution.redo(move)
	}

	return
}

func (s *Solver) Next() {
	//var current neighbour
	//s.CurrentSolution.copyIn(&current)

	bestMove := s.setUpBestNeighbour()

	isBestSolutionChanged := false
	if s.bestLocal != nil && s.bestLocal.Makespan() < s.BestMakespan() {
		s.bestSolution = s.bestLocal
		s.bestLocal = nil
		isBestSolutionChanged = true
	}

	if !isBestSolutionChanged {
		s.tabuList.Add(bestMove)

		if s.bestLocal != nil {
			s.CurrentSolution = s.bestLocal
			s.bestLocal = nil

		} else {
			newMove := s.tabuList.ForgetOldest()

			s.CurrentSolution.apply(newMove)
			s.CurrentSolution.updateByGraph()
			if s.CurrentSolution.Makespan() < s.BestMakespan() {
				s.CurrentSolution.copyIn(s.bestSolution)
			}
		}

	} else {
		s.tabuList.Add(bestMove)
		s.bestSolution.copyIn(s.CurrentSolution)
	}
}

package tabuSearch

import (
	"jobShop/base"
	"jobShop/state"
	"math"
	. "jobShop/tabuSearch/neighborhood"
	"jobShop/tabuSearch/neighborhood"
	"jobShop/tabuSearch/graph_state"
)

type Solver struct {
	jobs            base.Jobs
	CurrentSolution *Neighbour
	tabuList        TabuList

	bestLocal    *Neighbour
	bestSolution *Neighbour
}

func (s *Solver) GetBest() *Neighbour {
	return s.bestSolution
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
	initialSolution := graph_state.From(state)
	best := Neighbour{
		JobState: state,
		Graph:    initialSolution,
	}

	var bestCopy Neighbour
	best.CopyIn(&bestCopy)

	return Solver{
		jobs:            state.Jobs,
		CurrentSolution: &best,
		bestSolution:    &bestCopy,
		tabuList:        tabuList{},
	}
}

func (s *Solver) setUpBestNeighbour() (bestMove Move) {
	s.bestLocal = nil
	iterator := neighborhood.NewByCriticalPath(&s.CurrentSolution.JobState, &s.CurrentSolution.Graph)
	for move := range iterator.Generator() {
		s.CurrentSolution.Apply(move)

		//TODO: optimize, make checking in tabu before update graph
		//TODO: store impossible Move to prevent useless graph update
		success := s.CurrentSolution.UpdateByGraph()

		if success {
			if s.CurrentSolution.Makespan() < s.bestSolution.Makespan() {
				s.tabuList.Forget(move)
			}

			if s.bestLocal == nil {
				s.bestLocal = &Neighbour{}
				s.CurrentSolution.CopyIn(s.bestLocal)

			} else if s.CurrentSolution.Makespan() < s.bestLocal.Makespan() &&
				!s.tabuList.Contain(move) {

				bestMove = move
				s.CurrentSolution.CopyIn(s.bestLocal)
			}
		}

		s.CurrentSolution.Redo(move)
	}

	return
}

func (s *Solver) Next() {
	//var current Neighbour
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

			s.CurrentSolution.Apply(newMove)
			s.CurrentSolution.UpdateByGraph()
			if s.CurrentSolution.Makespan() < s.BestMakespan() {
				s.CurrentSolution.CopyIn(s.bestSolution)
			}
		}

	} else {
		s.tabuList.Add(bestMove)
		s.bestSolution.CopyIn(s.CurrentSolution)
	}
}

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
		tabuList:        &tabuList{},
	}
}

func (s *Solver) setUpBestNeighbour() (bestMove Move) {
	s.bestLocal = nil
	iterator := neighborhood.NewByCriticalPath(&s.CurrentSolution.JobState, &s.CurrentSolution.Graph)

	//fmt.Println("best: ", s.BestMakespan())
	//fmt.Println(s.CurrentSolution.Makespan(), " ", s.CurrentSolution.Graph)
	//fmt.Println("current: ",s.CurrentSolution.Makespan())

	for move := range iterator.Generator() {
		s.CurrentSolution.Apply(move)

		//TODO: optimize, make checking in tabu before update graph
		//TODO: store impossible Move to prevent useless graph update
		success := s.CurrentSolution.UpdateByGraph()

		if success {
			if s.CurrentSolution.Makespan() < s.bestSolution.Makespan() {
				s.tabuList.Forget(move)
				//s.bestLocal = &Neighbour{}
				//s.CurrentSolution.CopyIn(s.bestLocal)
				//bestMove = move
			}

			//fmt.Println(s.CurrentSolution.Makespan(), "-move: ", move)

			if s.bestLocal == nil {
				s.bestLocal = &Neighbour{}
				s.CurrentSolution.CopyIn(s.bestLocal)
				bestMove = move
				//fmt.Println(s.CurrentSolution.Makespan(), " first move: ", move)

			} else if s.CurrentSolution.Makespan() < s.bestLocal.Makespan() &&
				!s.tabuList.Contain(move) {
				//fmt.Println(s.CurrentSolution.Makespan(), " move: ", move)

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
	//empty := neighborhood.Move{}
	//if bestMove == empty {
	//fmt.Println("empty")
	//fmt.Println(s.bestLocal.Makespan())
	//}

	isBestSolutionChanged := false
	if s.bestLocal != nil && s.bestLocal.Makespan() < s.BestMakespan() {
		s.bestSolution = s.bestLocal
		s.bestLocal = nil
		isBestSolutionChanged = true
	}

	if !isBestSolutionChanged {

		if s.bestLocal != nil {
			s.tabuList.Add(bestMove)
			s.CurrentSolution = s.bestLocal
			s.bestLocal = nil

		} else {
			//fmt.Println("-- tabu: ", s.tabuList)
			newMove := s.tabuList.ForgetOldest()
			//fmt.Println(newMove)
			//fmt.Println("-- tabu: ", s.tabuList)

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

	//fmt.Println("tabu: ", s.tabuList)
}

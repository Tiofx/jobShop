package tabuSearch

import (
	"jobShop/base"
	"jobShop/state"
	"math"
	"sort"
)

type Solver struct {
	jobs base.Jobs
	//CurrentGraph    DisjunctiveGraph
	CurrentSolution *neighbour
	//neighboursLisp neighboursSet
	tabuList tabuList

	bestSolution *neighbour
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

	return Solver{
		jobs:            state.Jobs,
		CurrentSolution: best,
		bestSolution:    best,
	}
}

func (s *Solver) Next() {
	//currentNeihgbhour, _ := newNeighbour(s.CurrentGraph, s.jobs)
	//currentState, exist := s.CurrentGraph.To(s.jobs)
	//if !exist {
	//	log.Fatal("unaapropriate graph:\n", s)
	//}
	criticalJob := criticalJob(s.CurrentSolution.jobState)
	criticals := s.CurrentSolution.graph.criticalTaskPositionFor(criticalJob)
	neighboursGraphs := s.CurrentSolution.graph.neighbours(criticals)

	neighbours := createNeighboursWithoutTabu(neighboursGraphs, s.jobs, s.tabuList)
	sort.Sort(neighbours)

	var neighbour neighbour

	if len(neighbours) != 0 {
		neighbour = neighbours[0]
	} else {
		neighbour = s.tabuList[0]
		s.tabuList = s.tabuList[1:]
	}

	//fmt.Println("tabu: ", len(s.tabuList))
	//fmt.Println("\nneighbours: ", len(neighbours))
	//fmt.Println("selected: ", neighbour.jobState.JobOrder)

	//if !exist {
	//	log.Fatal("no next neibhour:\n", s)
	//}
	s.tabuList = append(s.tabuList, neighbour)

	if neighbour.Makespan() < s.BestMakespan() {
		s.bestSolution = &neighbour
	}

	s.CurrentSolution = &neighbour

}

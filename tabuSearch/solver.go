package tabuSearch

import (
	"jobShop/base"
	"jobShop/state"
	"math"
	. "jobShop/tabuSearch/neighborhood"
	"jobShop/tabuSearch/graph_state"
	"jobShop/initSolution/taskWaveByMachineGreed"
)

type Solver struct {
	jobs            base.Jobs
	CurrentSolution *Neighbour
	tabuList        TabuList

	bestLocal    *Neighbour
	bestSolution *Neighbour

	maxIteration            int
	iterationWithoutChanges int
	maxWithoutImprovement   int
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

func Solve(jobs base.Jobs, maxIterationNumber, maxWithoutImprovement, memoryCapacity int) state.State {
	initState := taskWaveByMachineGreed.Resolver{
		MaxTasksOnWave: taskWaveByMachineGreed.OptimalPermutationLimit,
		State:          state.New(jobs),
	}
	solution := initState.FindSolution()
	solver := NewSolver(solution, memoryCapacity, maxIterationNumber, maxWithoutImprovement)

	return solver.FindSolution()
}

func NewSolver(state state.State, memoryCapacity, maxIteration, maxWithoutImprovement int) Solver {
	initialSolution := graph_state.From(state)
	best := Neighbour{
		JobState: state,
		Graph:    initialSolution,
	}

	var bestCopy Neighbour
	best.CopyIn(&bestCopy)

	return Solver{
		jobs:                  state.Jobs,
		CurrentSolution:       &best,
		bestSolution:          &bestCopy,
		tabuList:              &tabuList{memoryCapacity: memoryCapacity},
		maxIteration:          maxIteration,
		maxWithoutImprovement: maxWithoutImprovement,
	}
}

func (s *Solver) setUpBestNeighbour() (bestMove Move) {
	s.bestLocal = nil
	iterator := NewByCriticalPath(&s.CurrentSolution.JobState, &s.CurrentSolution.Graph)

	bestLocalMakespan := math.MaxInt64

	for _, move := range iterator.Generate() {
		s.CurrentSolution.Apply(move)

		//TODO: optimize, make checking in tabu before update graph
		//TODO: store impossible Move to prevent useless graph update
		success := s.CurrentSolution.UpdateByGraph()

		if success {
			if s.CurrentSolution.Makespan() < s.bestSolution.Makespan() {
				s.tabuList.Forget(move)
			}

			if s.tabuList.Contain(move) {
				s.CurrentSolution.Redo(move)
				continue
			}

			if s.CurrentSolution.Makespan() < bestLocalMakespan {
				bestMove = move
				bestLocalMakespan = s.CurrentSolution.Makespan()
			}

		}

		s.CurrentSolution.Redo(move)
	}

	if bestLocalMakespan != math.MaxInt64 {
		s.CurrentSolution.Apply(bestMove)
		s.CurrentSolution.UpdateByGraph()

		s.bestLocal = s.CurrentSolution
	}

	return
}

func (s *Solver) Next() {
	bestMove := s.setUpBestNeighbour()

	isBestSolutionChanged := false
	if s.bestLocal != nil && s.bestLocal.Makespan() < s.BestMakespan() {
		s.bestLocal.CopyIn(s.bestSolution)
		s.bestLocal = nil
		isBestSolutionChanged = true
	}

	if !isBestSolutionChanged {
		s.iterationWithoutChanges++

		if s.bestLocal != nil {
			s.tabuList.Add(bestMove)
			s.CurrentSolution = s.bestLocal
			s.bestLocal = nil

		} else {
			for {
				newMove := s.tabuList.ForgetOldest()
				s.CurrentSolution.Apply(newMove)
				success := s.CurrentSolution.UpdateByGraph()
				if success {
					break
				}
				s.CurrentSolution.Redo(newMove)
			}
			if s.CurrentSolution.Makespan() < s.BestMakespan() {
				s.CurrentSolution.CopyIn(s.bestSolution)
			}
		}

	} else {
		s.iterationWithoutChanges = 0
		s.tabuList.Add(bestMove)
		s.bestSolution.CopyIn(s.CurrentSolution)
	}
}

func (s *Solver) FindSolution() state.State {
	min := s.BestMakespan()

	for i := 0; i < s.maxIteration && s.iterationWithoutChanges < s.maxWithoutImprovement; i++ {
		s.Next()

		if s.BestMakespan() < min {
			min = s.BestMakespan()
			//fmt.Println(i, " : ", min)
		}
	}
	return s.GetBest().JobState
}
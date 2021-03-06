package stats

import (
	"time"
	"github.com/Tiofx/jobShop/initSolution"
	"github.com/Tiofx/jobShop/state"
	"github.com/Tiofx/jobShop/base"
	"github.com/Tiofx/jobShop/initSolution/taskWaveByMachineGreed"
	"github.com/Tiofx/jobShop/tabuSearch"
	"fmt"
	"sync"
)

type SolverWithStatistic struct {
	wg              sync.WaitGroup
	improvementChan chan Snapshot
	startTime       time.Time
	stat            Snapshot

	initSolver   initSolution.Solver
	solverGetter func(state.State) *tabuSearch.Solver

	*tabuSearch.Solver
}

func NewSolverWithStatistic(jobs base.Jobs, memoryCapacity, maxIteration, maxWithoutImprovement uint64) SolverWithStatistic {
	initSolver :=
		taskWaveByMachineGreed.Resolver{
			MaxTasksOnWave: taskWaveByMachineGreed.OptimalPermutationLimit,
			State:          state.New(jobs),
		}

	return SolverWithStatistic{
		improvementChan: make(chan Snapshot, maxIteration),
		initSolver:      initSolver,
		solverGetter: func(state state.State) *tabuSearch.Solver {
			solver := tabuSearch.NewSolver(state, memoryCapacity, maxIteration, maxWithoutImprovement)
			return &solver
		},
		Solver: nil,
	}
}

func (s *SolverWithStatistic) GetImprovementStatsChannel() <-chan Snapshot {
	return s.improvementChan
}

func (s *SolverWithStatistic) Next() {
	if s.stat.Makespan != s.BestMakespan() {
		s.stat.Makespan, s.stat.TimeFromStart = s.BestMakespan(), time.Since(s.startTime)
		s.improvementChan <- s.stat
	}

	s.Solver.Next()
	s.stat.IterationNumber++
}

func (s *SolverWithStatistic) FindSolution() state.State {
	defer close(s.improvementChan)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
			s.Solver.Log()
			time.Sleep(100 * time.Millisecond)
		}
	}()
	s.startTime = time.Now()

	s.setUpSolver()
	for i := uint64(0); s.IsNeedOneMoreIteration(i); i++ {
		s.Next()
	}

	return s.GetBest().JobState
}

func (s *SolverWithStatistic) WaitForProcessingStats() {
	s.wg.Add(1)
	s.wg.Wait()
}

func (s *SolverWithStatistic) FinishProcessingStats() {
	s.wg.Done()
}

func (s *SolverWithStatistic) setUpSolver() {
	init := s.initSolver.FindSolution()
	s.Solver = s.solverGetter(init)
}

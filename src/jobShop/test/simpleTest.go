package test

import (
	"jobShop/base"
	"jobShop/state"
	"jobShop/initSolution/simpleGreed"
	"jobShop/tabuSearch"
	"fmt"
)

func SimpleTest() {
	jobs := base.Jobs{
		base.Job{
			base.Task{0, 3},
			base.Task{1, 2},
			base.Task{2, 2},
		},

		base.Job{
			base.Task{Machine: 0, Time: 2},
			base.Task{Machine: 2, Time: 1},
			base.Task{Machine: 1, Time: 4},
		},

		base.Job{
			base.Task{1, 4},
			base.Task{2, 3},
		},
	}

	newState := state.NewState(jobs)
	solution := simpleGreed.Resolver{State: newState}.FindSolution()
	resultState := solution.(simpleGreed.Resolver).State

	solver := tabuSearch.NewSolver(resultState)
	for i := 0; i < 100; i++ {
		fmt.Println(solver.BestMakespan(), " ", solver.BestLocalMakespan())
		solver.Next()
	}
}

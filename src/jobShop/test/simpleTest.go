package test

import (
	"jobShop/state"
	"jobShop/initSolution/simpleGreed"
	"jobShop/tabuSearch"
	"fmt"
)

func SimpleTest() {
	jobs := simpleTestCase().Jobs

	newState := state.NewState(jobs)
	solution := simpleGreed.Resolver{State: newState}.FindSolution()
	resultState := solution

	solver := tabuSearch.NewSolver(resultState)
	for i := 0; i < 100; i++ {
		fmt.Println(solver.BestMakespan(), " ", solver.BestLocalMakespan())
		solver.Next()
	}
}

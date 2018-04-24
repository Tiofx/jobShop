package test

import (
	"jobShop/base"
	"jobShop/state"
	"jobShop/initSolution/taskWaveByMachineGreed"
	"jobShop/tabuSearch"
	"fmt"
	"os"
	"time"
	"log"
)

func TabuSearch(permutationLimit, maxIterationNumber, maxWithoutImprovement, memoryCapacity int, tests ...testCase) Results{
	res := test(func(jobs base.Jobs) state.State {
		initState := taskWaveByMachineGreed.Resolver{MaxTasksOnWave: permutationLimit, State: state.New(jobs)}
		solution := initState.FindSolution()
		solver := tabuSearch.NewSolver(solution, memoryCapacity, maxIterationNumber, maxWithoutImprovement)

		return solver.FindSolution()
	}, tests...)

	return res
}

func saveResult(name string, res Results, description string) {
	return
	file, err := os.Create(fmt.Sprintf("./testResults/%v_%v.txt", name, time.Now().Format(time.RFC3339)))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(description)
	file.WriteString(res.String())
}

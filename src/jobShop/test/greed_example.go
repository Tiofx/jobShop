package test

import (
	"jobShop/base"
	"jobShop/state"
	"jobShop/initSolution/simpleGreed"
	"fmt"
	"jobShop/initSolution/taskWaveByMachineGreed"
	"jobShop/tabuSearch"
	"time"
	"os"
	"log"
)

func ExampleResolver_FindSolution() {
	testCase := allTestsCases()

	res := test(func(jobs base.Jobs) state.State {
		initState := simpleGreed.Resolver{State: state.New(jobs)}
		return initState.FindSolution()
	}, testCase...)

	fmt.Println(res)
	saveResult("greed_1", res, fmt.Sprintf("solution by greed\n", ))
}

func ExampleResolver_FindSolution2() {
	const max = 4
	testCase := allTestsCases()

	res := test(func(jobs base.Jobs) state.State {
		initState := taskWaveByMachineGreed.Resolver{MaxTasksOnWave: max, State: state.New(jobs)}
		return initState.FindSolution()
	}, testCase...)

	fmt.Println(res)
	saveResult("greed_2", res, fmt.Sprintf("solution by task wave with max %v\n", max))
}

func ExampleTabuSearch() {
	const (
		iterationNumber       = 5000
		max                   = 3
		maxWithoutImprovement = 256
		memoryCapacity        = 55
	)

	testCase := []testCase{testCaseNumber(5)}
	//testCase := allTestsCases()[:5]

	res := test(func(jobs base.Jobs) state.State {
		initState := taskWaveByMachineGreed.Resolver{MaxTasksOnWave: max, State: state.New(jobs)}
		//initState := simpleGreed.Resolver{State: state.NewState(jobs)}
		solution := initState.FindSolution()
		solver := tabuSearch.NewSolver(solution, memoryCapacity, iterationNumber, maxWithoutImprovement)

		return solver.FindSolution()
	}, testCase...)

	fmt.Println(res)
	saveResult("tabu_search", res,
		fmt.Sprintf("after rid of copy in loop\ntabu search: %v iterations, max %v without changes iteration and after tabu reset\nfirst solution by task wave with max %v\n",
			iterationNumber, maxWithoutImprovement, max))
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

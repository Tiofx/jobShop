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
		initState := simpleGreed.Resolver{State: state.NewState(jobs)}
		return initState.FindSolution()
	}, testCase...)

	fmt.Println(res)
	saveResult("greed_1", res, fmt.Sprintf("solution by greed\n", ))
}

func ExampleResolver_FindSolution2() {
	const max = 4
	testCase := allTestsCases()

	res := test(func(jobs base.Jobs) state.State {
		initState := taskWaveByMachineGreed.Resolver{MaxTasksOnWave: max, State: state.NewState(jobs)}
		return initState.FindSolution()
	}, testCase...)

	fmt.Println(res)
	saveResult("greed_2", res, fmt.Sprintf("solution by task wave with max %v\n", max))
}

func ExampleTabuSearch() {
	const iterationNumber = 100
	const max = 3
	testCase := allTestsCases()

	res := test(func(jobs base.Jobs) state.State {
		initState := taskWaveByMachineGreed.Resolver{MaxTasksOnWave: max, State: state.NewState(jobs)}
		//initState := simpleGreed.Resolver{State: state.NewState(jobs)}
		solution := initState.FindSolution()
		solver := tabuSearch.NewSolver(solution)

		//min := solver.BestMakespan()
		//fmt.Println("==========================")
		for i := 0; i < iterationNumber; i++ {
			solver.Next()

			//if solver.BestMakespan() < min {
			//	min = solver.BestMakespan()
			//	fmt.Println(i, " : ", min)
			//}
		}
		return solver.GetBest().JobState
	}, testCase...)

	fmt.Println(res)
	saveResult("tabu_search", res,
		fmt.Sprintf("tabu search: %v iterations\nfirst solution by task wave with max %v\n",
			iterationNumber, max))
}

func saveResult(name string, res Results, description string) {
	file, err := os.Create(fmt.Sprintf("./testResults/%v_%v.txt", name, time.Now().Format(time.RFC3339)))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(description)
	file.WriteString(res.String())
}

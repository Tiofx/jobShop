package test

import (
	"jobShop/base"
	"jobShop/state"
	"jobShop/initSolution/simpleGreed"
	"fmt"
	"jobShop/initSolution/taskWaveByMachineGreed"
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
	TabuSearch(taskWaveByMachineGreed.OptimalPermutationLimit, 5000, 256, 55, testCaseNumber(5))
}

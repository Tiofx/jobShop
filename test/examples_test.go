package test

import (
	"github.com/Tiofx/jobShop/base"
	"github.com/Tiofx/jobShop/state"
	"github.com/Tiofx/jobShop/initSolution/simpleGreed"
	"fmt"
	"github.com/Tiofx/jobShop/initSolution/taskWaveByMachineGreed"
	"testing"
)

func ExampleResolver_FindSolution() {
	testCase := AllTestsCases()

	res := test(func(jobs base.Jobs) state.State {
		initState := simpleGreed.Resolver{State: state.New(jobs)}
		return initState.FindSolution()
	}, testCase...)

	fmt.Println(res)
	saveResult("greed_1", res, fmt.Sprintf("solution by greed\n", ))
}

func ExampleResolver_FindSolution2() {
	const max = 4
	testCase := AllTestsCases()

	res := test(func(jobs base.Jobs) state.State {
		initState := taskWaveByMachineGreed.Resolver{MaxTasksOnWave: max, State: state.New(jobs)}
		return initState.FindSolution()
	}, testCase...)

	fmt.Println(res)
	saveResult("greed_2", res, fmt.Sprintf("solution by task wave with max %v\n", max))
}

func Test_TabuSearch(t *testing.T) {
	//res := TabuSearch(taskWaveByMachineGreed.OptimalPermutationLimit, 5000, 256, 55, TestCaseNumber(5))
	//fmt.Println(res)
	fmt.Println(TestCaseNumber(6))
}

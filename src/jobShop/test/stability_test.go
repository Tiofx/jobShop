package test

import (
	"testing"
	"jobShop/base"
	"jobShop/state"
	"jobShop/initSolution/simpleGreed"
	"reflect"
	"fmt"
	"jobShop/initSolution/taskWaveByMachineGreed"
	"jobShop/tabuSearch"
)

func init() {
	testDir = "/Users/andrej/GoglandProjects/job-shop/testinstances"
}

func TestStabilityOfSimpleGreed(t *testing.T) {
	testCase := someRandomTests()
	solver := func(jobs base.Jobs) state.State {
		initState := simpleGreed.Resolver{State: state.NewState(jobs)}
		return initState.FindSolution()
	}

	testStability(t, testCase, solver)
}

func TestStabilityOfSecondGreedAlgorithm(t *testing.T) {
	testCase := someRandomTests()
	solver := func(jobs base.Jobs) state.State {
		initState := taskWaveByMachineGreed.Resolver{MaxTasksOnWave: 3, State: state.NewState(jobs)}
		return initState.FindSolution()
	}

	testStability(t, testCase, solver)
}

func TestStabilityOfTabuSearch(t *testing.T) {
	testCase := someRandomTests()
	solver := func(jobs base.Jobs) state.State {
		initState := simpleGreed.Resolver{State: state.NewState(jobs)}
		solution := initState.FindSolution()
		solver := tabuSearch.NewSolver(solution)

		for i := 0; i < 100; i++ {
			solver.Next()
		}
		return solver.GetBest().JobState
	}

	testStability(t, testCase, solver)
}

func testStability(t *testing.T, tests []testCase, solver func(jobs base.Jobs) state.State) {
	res1 := test(solver, tests...)
	res2 := test(solver, tests...)

	for i := 0; i < len(tests); i++ {
		r1, r2 := res1[i].Result, res2[i].Result

		if !reflect.DeepEqual(r1, r2) {
			fmt.Println(res1[i])
			fmt.Println(res2[i])
			t.Fail()
		}
	}
}

func someRandomTests() []testCase { return []testCase{randomTest(), randomTest(), randomTest()} }

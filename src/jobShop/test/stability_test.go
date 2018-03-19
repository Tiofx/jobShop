package test

import (
	"testing"
	"jobShop/base"
	"jobShop/state"
	"jobShop/initSolution/simpleGreed"
	"reflect"
	"fmt"
	"math/rand"
	"jobShop/initSolution/taskWaveByMachineGreed"
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

func someRandomTests() []testCase {
	t1, t2, t3 := rand.Intn(testsNumber()), rand.Intn(testsNumber()), rand.Intn(testsNumber())
	all := allTestsCases()
	return []testCase{all[t1], all[t2], all[t3]}
}

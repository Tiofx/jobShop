package test

import (
	"jobShop/base"
	"jobShop/state"
	"jobShop/initSolution/simpleGreed"
	"fmt"
)

func ExampleResolver_FindSolution() {
	testCase := allTestsCases()

	res := test(func(jobs base.Jobs) state.State {
		initState := simpleGreed.Resolver{State: state.NewState(jobs)}
		return initState.FindSolution()
	}, testCase...)

	fmt.Println(res)
}

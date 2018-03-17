package test

import (
	"jobShop/base"
	"jobShop/state"
	"fmt"
	"time"
)

type Result struct {
	Problem ProblemDescription
	Result  state.State
	Elapsed time.Duration
}

type ProblemDescription struct{ testCase }

type ResultComparator interface {
	HowMuchWorse() float32
}

func (r Result) String() (res string) {
	res += fmt.Sprintf("%v", r.Problem)
	res += fmt.Sprintf("expected makespan: %v\n", r.Problem.Optimum)
	res += fmt.Sprintf("actual makespan:   %v\n", r.Result.Makespan())
	res += fmt.Sprintf("how much worse:    %.2f%%\n", r.HowMuchWorse()*100)
	res += fmt.Sprintf("elapsed time is    %v\n", r.Elapsed)

	return
}

func (r Result) HowMuchWorse() float32 {
	return float32(r.Result.Makespan()-r.Problem.Optimum) / float32(r.Problem.Optimum)
}

func test(testCase testCase, solver func(jobs base.Jobs) state.State) Result {
	start := time.Now()

	result := solver(testCase.Jobs)

	duration := time.Since(start)

	return Result{
		Problem: ProblemDescription{testCase},
		Result:  result,
		Elapsed: duration,
	}
}

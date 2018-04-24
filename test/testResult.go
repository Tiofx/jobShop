package test

import (
	"github.com/Tiofx/jobShop/base"
	"github.com/Tiofx/jobShop/state"
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
	res += fmt.Sprintf("------------ Result --------------\n")
	res += fmt.Sprintf("expected makespan: %v\n", r.Problem.Optimum)
	res += fmt.Sprintf("actual makespan:   %v\n", r.Result.Makespan())
	res += fmt.Sprintf("how much worse:    %.2f%%\n", r.HowMuchWorse()*100)
	res += fmt.Sprintf("elapsed time is    %v\n", r.Elapsed)
	res += fmt.Sprintf("----------------------------------\n")

	return
}

func (r Result) HowMuchWorse() float64 {
	return float64(r.Result.Makespan()-r.Problem.Optimum) / float64(r.Problem.Optimum)
}

func test(solver func(jobs base.Jobs) state.State, tests ...testCase) Results {
	results := make(Results, 0, len(tests))

	for _, testCase := range tests {
		start := time.Now()
		s := solver(testCase.Jobs)
		duration := time.Since(start)

		r := Result{
			Problem: ProblemDescription{testCase},
			Result:  s,
			Elapsed: duration,
		}

		results = append(results, r)
	}

	return results
}

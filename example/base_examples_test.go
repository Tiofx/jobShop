package example_test

import (
	"fmt"
	"github.com/Tiofx/jobShop/base"
	"github.com/Tiofx/jobShop/tabuSearch"
	"github.com/Tiofx/jobShop/tabuSearch/graph_state"
	"github.com/Tiofx/jobShop/initSolution/taskWaveByMachineGreed"
	"github.com/Tiofx/jobShop/test"
)

func Example_initJobs() {
	jobs := base.Jobs{
		{
			{0, 3},
			{1, 2},
			{2, 2},
		},

		{
			{Machine: 0, Time: 2},
			{Machine: 2, Time: 1},
			{Machine: 1, Time: 4},
		},

		{
			{1, 4},
			{2, 3},
		},
	}

	fmt.Println(jobs)

	//Output:
	//job 0: [{0 3} {1 2} {2 2}]
	//job 1: [{0 2} {2 1} {1 4}]
	//job 2: [{1 4} {2 3}]
}

func Example_timeOfSchedule() {
	solution := tabuSearch.Solve(test.ComplexTestCase().Jobs, 100, 16, 10)
	fmt.Printf("Time to finish schedule: %v\n", solution.Makespan())

	// Output:
	// Time to finish schedule: 850
}

func Example_solution1() {
	solution := tabuSearch.Solve(test.ComplexTestCase().Jobs, 100, 16, 10)
	fmt.Println("Order of job in which schedule has to be accomplish:")
	fmt.Println(solution.JobOrder)

	// Output:
	// Order of job in which schedule has to be accomplish:
	// [0 2 3 4 5 7 10 12 13 14 16 0 1 2 3 4 5 9 11 12 17 18 0 1 2 4 5 6 8 9 12 14 15 16 17 18 19 0 1 2 3 4 6 7 10 12 13 14 15 17 0 1 3 4 5 6 7 11 13 15 17 0 2 3 5 6 7 10 13 16 17 18 2 3 4 5 11 15 17 18 1 2 3 4 9 10 13 14 15 17 18 19 1 2 4 6 8 10 11 14 15 19 4 6 8 12 13 14 16 17 18 19 0 3 4 6 7 8 10 12 13 19 0 1 3 5 6 7 13 16 19 1 2 3 5 6 7 13 14 15 18 19 1 2 4 5 7 9 10 12 13 14 17 0 1 7 9 10 11 12 13 14 15 16 0 1 4 6 8 9 10 12 13 17 18 0 1 2 6 7 9 12 13 14 15 16 17 19 0 1 3 5 6 7 11 12 13 14 15 16 18 19 2 3 4 8 10 12 13 14 16 18 0 2 3 4 5 7 8 9 10 11 14 15 18 0 2 3 4 5 7 8 9 10 14 15 16 0 1 5 6 8 11 12 14 15 16 17 19 1 3 5 7 8 9 11 12 17 18 19 2 5 7 8 10 11 15 16 18 19 8 9 10 11 16 17 18 19 8 9 10 11 12 16 17 6 8 9 11 16 18 19 6 8 9 11 15 19 11 9]
}

// Sequence of job's task on each machine.
// Number of task is as if all jobs would be combined.
func Example_solution2() {
	solution := tabuSearch.Solve(test.ComplexTestCase().Jobs, 100, 16, 10)
	graph := graph_state.From(solution)
	fmt.Println(graph)

	// Output:
	//Machine 0: [2 11 17 0 3 15 14 13 12 1 4 9 19 8 5 7 10 6 16 18]
	//Machine 1: [4 9 3 13 5 17 1 10 19 2 0 14 16 7 18 12 15 6 11 8]
	//Machine 2: [0 12 17 2 16 6 1 14 4 13 7 5 3 18 15 11 8 10 9 19]
	//Machine 3: [16 0 1 6 14 2 4 12 7 8 19 5 10 17 18 9 15 13 3 11]
	//Machine 4: [0 7 3 15 13 6 10 11 14 8 5 2 4 16 12 1 19 9 17 18]
	//Machine 5: [7 12 18 16 10 17 4 11 9 15 13 3 19 5 1 14 0 6 8 2]
	//Machine 6: [2 1 3 4 6 7 0 5 13 12 16 14 8 18 10 15 19 17 11 9]
	//Machine 7: [2 3 4 15 5 18 10 0 16 13 1 9 12 17 14 19 7 11 8 6]
	//Machine 8: [5 13 12 6 4 19 8 17 0 10 7 15 1 14 16 11 2 3 9 18]
	//Machine 9: [3 14 0 15 19 7 4 6 2 18 13 10 8 11 5 1 17 12 16 9]
	//Machine 10: [10 17 18 2 4 1 7 13 3 5 6 15 14 0 16 8 11 9 12 19]
	//Machine 11: [4 5 1 11 17 8 14 18 3 19 13 0 12 6 2 16 7 10 15 9]
	//Machine 12: [12 9 17 1 3 10 13 2 14 6 19 16 4 7 18 0 5 8 11 15]
	//Machine 13: [4 2 6 8 18 1 11 12 7 15 13 19 14 0 9 10 5 3 17 16]
	//Machine 14: [5 14 15 13 18 6 7 9 10 17 0 1 3 12 4 8 11 19 2 16]
}

func Example_testCase() {
	res := test.TabuSearch(taskWaveByMachineGreed.OptimalPermutationLimit, 1000, 32, 55, test.TestCaseNumber(6))[0]
	fmt.Println(res)

	//Example of output:
	//----- Short test description -----
	//file name:		ft10.txt
	//job number:  		10
	//avg. task number:	10
	//total task number: 	100
	//machine number:		10
	//optimum:     		930
	//----------------------------------
	//------------ Result --------------
	//expected makespan: 930
	//actual makespan:   1088
	//how much worse:    16.99%
	//elapsed time is    190.809ms
	//----------------------------------
}

func Example_tests() {
	res := test.TabuSearch(taskWaveByMachineGreed.OptimalPermutationLimit, 100, 16, 32, test.AllTestsCases()[:5]...)
	fmt.Println(res)

	// Example of output:
	//== Total test number: 5
	//==============================
	//========== Test №1 ===========
	//==============================
	//----- Short test description -----
	//file name:			abz5.txt
	//job number:  		10
	//avg. task number:	10
	//total task number: 	100
	//machine number:		10
	//optimum:     		1234
	//----------------------------------
	//------------ Result --------------
	//expected makespan: 1234
	//actual makespan:   1290
	//how much worse:    4.54%
	//elapsed time is    130.586ms
	//----------------------------------
	//==============================
	//========== Test №2 ===========
	//==============================
	//----- Short test description -----
	//file name:			abz6.txt
	//job number:  		10
	//avg. task number:	10
	//total task number: 	100
	//machine number:		10
	//optimum:     		943
	//----------------------------------
	//------------ Result --------------
	//expected makespan: 943
	//actual makespan:   1144
	//how much worse:    21.31%
	//elapsed time is    95.874ms
	//----------------------------------
	//==============================
	//========== Test №3 ===========
	//==============================
	//----- Short test description -----
	//file name:			abz7.txt
	//job number:  		20
	//avg. task number:	15
	//total task number: 	300
	//machine number:		15
	//optimum:     		651
	//----------------------------------
	//------------ Result --------------
	//expected makespan: 651
	//actual makespan:   846
	//how much worse:    29.95%
	//elapsed time is    1.516808s
	//----------------------------------
	//==============================
	//========== Test №4 ===========
	//==============================
	//----- Short test description -----
	//file name:			abz8.txt
	//job number:  		20
	//avg. task number:	15
	//total task number: 	300
	//machine number:		15
	//optimum:     		627
	//----------------------------------
	//------------ Result --------------
	//expected makespan: 627
	//actual makespan:   939
	//how much worse:    49.76%
	//elapsed time is    1.408411s
	//----------------------------------
	//==============================
	//========== Test №5 ===========
	//==============================
	//----- Short test description -----
	//file name:			abz9.txt
	//job number:  		20
	//avg. task number:	15
	//total task number: 	300
	//machine number:		15
	//optimum:     		650
	//----------------------------------
	//------------ Result --------------
	//expected makespan: 650
	//actual makespan:   896
	//how much worse:    37.85%
	//elapsed time is    1.576006s
	//----------------------------------
	//==============================
	//======== Statistics ==========
	//==============================
	//-- Result error statistics:
	//avg.:  28.68%
	//best:  4.54%
	//worst: 49.76%
	//---------------------------
	//-- Time statistics:
	//total: 4.727685s
	//mean:  945.537ms
	//min:   95.874ms
	//max:   1.576006s
	//---------------------------
}

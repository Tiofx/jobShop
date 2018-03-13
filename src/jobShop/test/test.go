package test

import (
	"io/ioutil"
	"log"
	"os"
	"jobShop/parser"
	"jobShop/state"
	"jobShop/tabuSearch"
	"jobShop/initSolution/simpleGreed"
)

type TestResult struct {
	Expected, Actual int
}

const testDir = "./testinstances"

func getAllFiles() []os.FileInfo {
	infos, err := ioutil.ReadDir(testDir)
	if err != nil {
		log.Fatal(err)
	}

	return infos
}

func TestAll() {
	for _, fileinfo := range getAllFiles() {
		//for _, fileinfo := range getAllFiles()[0:1] {
		//for _, fileinfo := range getAllFiles()[2:3] {
		//fmt.Println(fileinfo.Name())
		testCase := parser.NewTestCase(testDir + "/" + fileinfo.Name())

		//state := taskWaveByMachineGreed.Resolver{State: state.NewState(testCase.Jobs)}
		initState := simpleGreed.Resolver{State: state.NewState(testCase.Jobs)}

		solution := initState.FindSolution()

		//fmt.Println(testCase.Optimum)
		//result := solution.Makespan()
		//fmt.Println(result)
		newState := solution.(simpleGreed.Resolver).State
		//newState := solution.(taskWaveByMachineGreed.Resolver).State

		//fmt.Println("first: ", solution.Makespan())
		solver := tabuSearch.NewSolver(newState)
		for i := 0; i < 200; i++ {
			solver.Next()
			//fmt.Println(solver.BestMakespan(), " ", solver.BestLocalMakespan())
		}
		//fmt.Println(solver.BestMakespan())

		//fmt.Println(1.0 - float32(testCase.Optimum)/float32(result))
		//fmt.Println(float32(result) / float32(testCase.Optimum))
		//fmt.Println("------")
	}
}

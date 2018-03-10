package test

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"
	"jobShop/parser"
	"jobShop/state"
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
	//for _, fileinfo := range getAllFiles() {
	for _, fileinfo := range getAllFiles()[2:3] {
		fmt.Println(fileinfo.Name())
		testCase := parser.NewTestCase(testDir + "/" + fileinfo.Name())

		//state := taskWaveByMachineGreed.Resolver{state.NewState(testCase.Jobs)}
		state := simpleGreed.Resolver{State: state.NewState(testCase.Jobs)}

		solution := state.FindSolution()

		fmt.Println(testCase.Optimum)
		result := solution.Makespan()
		fmt.Println(result)
		fmt.Println(1.0 - float32(testCase.Optimum)/float32(result))
		fmt.Println(float32(result) / float32(testCase.Optimum))
		fmt.Println("------")
	}
}

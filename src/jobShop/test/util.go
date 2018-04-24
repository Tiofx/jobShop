package test

import (
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"jobShop/base"
	"math/rand"
)

const testDir = "./testinstances"

func getAllFiles() []os.FileInfo {
	infos, err := ioutil.ReadDir(testDir)
	if err != nil {
		log.Fatal(err)
	}

	return infos
}

func fullPathTo(filename string) string   { return testDir + "/" + filename }
func testCaseOf(filename string) testCase { return newTestCase(filename) }
func testsNumber() int                    { return len(getAllFiles()) }

func testCaseNumber(index int) testCase {
	files := getAllFiles()
	lastIndex := len(files) - 1
	if index < 0 || index > lastIndex {
		log.Fatal(fmt.Sprintf("no test case with index %v. Maximum index is %v", index, lastIndex))
	}

	return testCaseOf(files[index].Name())
}

func simpleTestCase() testCase {
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

	return testCase{
		Filename:    "no file",
		JobsNumber:  3,
		TasksNumber: 3,
		Optimum:     11,
		Jobs:        jobs,
	}
}

func generalTestCase() testCase { return testCaseNumber(0) }
func complexTestCase() testCase { return testCaseNumber(2) }
func randomTest() testCase      { return testCaseNumber(rand.Intn(testsNumber())) }

func allTestsCases() []testCase {
	var res []testCase
	for _, fileInfo := range getAllFiles() {
		res = append(res, testCaseOf(fileInfo.Name()))
	}
	return res
}

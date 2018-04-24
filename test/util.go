package test

import (
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"github.com/Tiofx/jobShop/base"
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

func TestCaseNumber(index int) testCase {
	files := getAllFiles()
	lastIndex := len(files) - 1
	if index < 0 || index > lastIndex {
		log.Fatal(fmt.Sprintf("no test case with index %v. Maximum index is %v", index, lastIndex))
	}

	return testCaseOf(files[index].Name())
}

func From(jobs base.Jobs, optimum int) testCase {
	return testCase{
		Filename:    "no file",
		JobsNumber:  len(jobs),
		TasksNumber: -1,
		Optimum:     optimum,
		Jobs:        jobs,
	}

}

func SimpleTestCase() testCase {
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

	return From(jobs, 11)
}

func GeneralTestCase() testCase { return TestCaseNumber(0) }
func ComplexTestCase() testCase { return TestCaseNumber(2) }
func RandomTest() testCase      { return TestCaseNumber(rand.Intn(testsNumber())) }

func AllTestsCases() []testCase {
	var res []testCase
	for _, fileInfo := range getAllFiles() {
		res = append(res, testCaseOf(fileInfo.Name()))
	}
	return res
}

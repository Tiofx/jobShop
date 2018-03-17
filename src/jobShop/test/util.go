package test

import (
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"jobShop/base"
)

const testDir = "./testinstances"

func getAllFiles() []os.FileInfo {
	infos, err := ioutil.ReadDir(testDir)
	if err != nil {
		log.Fatal(err)
	}

	return infos
}

func testCaseOf(filename string) TestCase {
	return NewTestCase(testDir + "/" + filename)
}

func testsOf(filenames []string) []TestCase {
	res := make([]TestCase, len(filenames), 0)
	for _, filename := range filenames {
		res = append(res, testCaseOf(filename))
	}

	return res
}

func testCaseNumber(index int) TestCase {
	files := getAllFiles()
	lastIndex := len(files) - 1
	if index < 0 || index > lastIndex {
		log.Fatal(fmt.Sprintf("no test case with index %v. Maximum index is %v", index, lastIndex))
	}

	return testCaseOf(files[index].Name())
}

func simpleTestCase() TestCase {
	jobs := base.Jobs{
		base.Job{
			base.Task{0, 3},
			base.Task{1, 2},
			base.Task{2, 2},
		},

		base.Job{
			base.Task{Machine: 0, Time: 2},
			base.Task{Machine: 2, Time: 1},
			base.Task{Machine: 1, Time: 4},
		},

		base.Job{
			base.Task{1, 4},
			base.Task{2, 3},
		},
	}

	return TestCase{
		JobsNumber:  3,
		TasksNumber: 3,
		Optimum:     11,
		Jobs:        jobs,
	}
}

func generalTestCase() TestCase {
	return testCaseNumber(0)
}

func complexTestCase() TestCase {
	return testCaseNumber(2)
}

func allTestsCases() (res []TestCase) {
	for _, fileInfo := range getAllFiles() {
		res = append(res, testCaseOf(fileInfo.Name()))
	}
	return res
}

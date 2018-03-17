package test

import (
	"jobShop/base"
	"fmt"
)

type testCase struct {
	Filename                         string
	JobsNumber, TasksNumber, Optimum int

	base.Jobs
}

func newTestCase(filename string) testCase {
	scanner := newTestParser(fullPathTo(filename))
	jobsNumber, taskNumbers, optimum, jobs := scanner.parseAllData()

	return testCase{
		Filename:    filename,
		JobsNumber:  jobsNumber,
		TasksNumber: taskNumbers,
		Optimum:     optimum,
		Jobs:        jobs,
	}
}

func (ts *testCase) String() (res string) {
	res += fmt.Sprintf("----- Short test description -----")
	res += fmt.Sprintf("file name: %v", ts.Filename)
	res += fmt.Sprintf("job number: %v", ts.JobsNumber)
	res += fmt.Sprintf("task number on each job: %v", ts.TasksNumber)
	res += fmt.Sprintf("optimum: %v", ts.Optimum)
	res += fmt.Sprintf("---------------------")
	return
}

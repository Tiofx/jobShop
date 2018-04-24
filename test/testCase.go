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

func (ts testCase) String() (res string) {
	res += fmt.Sprintf("----- Short test description -----\n")
	res += fmt.Sprintf("file name:			%v\n", ts.Filename)
	res += fmt.Sprintf("job number:  		%v\n", ts.JobsNumber)
	res += fmt.Sprintf("avg. task number:	%v\n", float64(ts.TotalTaskNumber()/ts.JobsNumber))
	res += fmt.Sprintf("total task number: 	%v\n", ts.TotalTaskNumber())
	res += fmt.Sprintf("machine number:		%v\n", ts.MachineNumber())
	res += fmt.Sprintf("optimum:     		%v\n", ts.Optimum)
	res += fmt.Sprintf("----------------------------------\n")
	return
}

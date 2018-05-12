package base

import "fmt"

func (jobs Jobs) String() (res string) {
	for i, job := range jobs {
		res += fmt.Sprintf("job %v: %v\n", i, job)
	}

	return
}

func (executed NumberOfExecutedTasks) String() (res string) {
	//jobsToString([]uint64(executed), res)
	for job, number := range executed {
		res += fmt.Sprintf("job %v: %v\n", job, number)
	}

	return
}

func (left JobsTotalTime) String() (res string) {
	for job, time := range left {
		res += fmt.Sprintf("job %v: %v\n", job, time)
	}

	return
}

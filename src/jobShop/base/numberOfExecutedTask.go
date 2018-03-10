package base

type NumberOfExecutedTasks []int

func (noet NumberOfExecutedTasks) IsExecutedAllOf(jobNumber int, job Job) bool {
	return noet[jobNumber] == len(job)
}

func (noet NumberOfExecutedTasks) IsExecutedAll(jobs Jobs) (res bool) {
	res = true

	for i, job := range jobs {
		res = res && noet.IsExecutedAllOf(i, job)
	}

	return
}

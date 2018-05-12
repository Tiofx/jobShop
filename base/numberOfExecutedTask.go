package base

type NumberOfExecutedTasks []uint64

func (noet NumberOfExecutedTasks) IsExecutedAllOf(jobNumber uint64, job Job) bool {
	return noet[jobNumber] == uint64(len(job))
}

func (noet NumberOfExecutedTasks) IsExecutedAll(jobs Jobs) (res bool) {
	res = true

	for i, job := range jobs {
		res = res && noet.IsExecutedAllOf(uint64(i), job)
	}

	return
}

package base

type Job []Task

func (job *Job) TotalTime() (total uint64) {
	for _, task := range *job {
		total += task.Time
	}

	return
}

func (job *Job) TaskNumber() uint64 {
	return uint64(len(*job))
}

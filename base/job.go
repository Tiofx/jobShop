package base

type Job []Task

func (job *Job) TotalTime() (total int) {
	for _, task := range *job {
		total += task.Time
	}

	return
}

func (job *Job) TaskNumber() int {
	return len(*job)
}

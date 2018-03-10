package base

type Job []Task

func (job Job) TotalTime() (total int) {
	for _, task := range job {
		_, time := task.Deconstruct()
		total += time
	}

	return
}

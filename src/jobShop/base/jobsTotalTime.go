package base

type JobsTotalTime []int

func NewJobsTotalTime(jobs Jobs) JobsTotalTime {
	res := make(JobsTotalTime, len(jobs))

	for i, job := range jobs {
		res[i] = job.TotalTime()
	}

	return res
}

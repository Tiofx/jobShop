package base

type JobsTotalTime []uint64

func NewJobsTotalTime(jobs Jobs) JobsTotalTime {
	res := make(JobsTotalTime, len(jobs))
	res.SetUpBy(jobs)

	return res
}

func (jtt JobsTotalTime) SetUpBy(jobs Jobs) {
	for i, job := range jobs {
		jtt[i] = job.TotalTime()
	}
}

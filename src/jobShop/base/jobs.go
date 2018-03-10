package base

type Jobs []Job

func (jobs Jobs) EstimateTime() (min, max int) {
	var total, maxJobTime int

	for _, job := range jobs {
		jobTime := job.TotalTime()
		total += jobTime

		if jobTime > maxJobTime {
			maxJobTime = jobTime
		}
	}

	min = maxJobTime
	max = total
	return
}

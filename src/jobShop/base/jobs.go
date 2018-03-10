package base

type Jobs []Job

func (jobs Jobs) ToMachines() (ms Machines) {
	ms = make(Machines)

	for _, job := range jobs {
		for _, task := range job {
			ms.AddMachine(task.Machine)
		}
	}

	return
}

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

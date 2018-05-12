package base

type Jobs []Job

func (jobs *Jobs) MachineNumber() uint64 {
	var maxMachineNumber Machine

	for _, job := range *jobs {
		for _, task := range job {
			if task.Machine > maxMachineNumber {
				maxMachineNumber = task.Machine
			}
		}
	}

	return uint64(maxMachineNumber) + 1
}

func (jobs *Jobs) TotalTaskNumber() (number uint64) {
	for _, job := range *jobs {
		number += job.TaskNumber()
	}

	return
}

func (jobs *Jobs) TaskNumber(beforeJob uint64) (number uint64) {
	if beforeJob < 0 {
		return 0
	}

	for jobIndex, job := range *jobs {
		if beforeJob == uint64(jobIndex) {
			break
		}
		number += job.TaskNumber()
	}

	return
}

func (jobs Jobs) EstimateTime() (min, max uint64) {
	var total, maxJobTime uint64

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

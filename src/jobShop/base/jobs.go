package base

type Jobs []Job

func (jobs *Jobs) MachineNumber() int {
	var maxMachineNumber Machine

	for _, job := range *jobs {
		for _, task := range job {
			if task.Machine > maxMachineNumber {
				maxMachineNumber = task.Machine
			}
		}
	}

	return int(maxMachineNumber) + 1
}

func (jobs *Jobs) TotalTaskNumber() (number int) {
	for _, job := range *jobs {
		number += job.TaskNumber()
	}

	return
}

func (jobs *Jobs) TaskNumber(beforeJob int) (number int) {
	if beforeJob < 0 {
		return 0
	}

	for jobIndex, job := range *jobs {
		if beforeJob == jobIndex {
			break
		}
		number += job.TaskNumber()
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

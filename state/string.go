package state

import "fmt"

func (times JobTimeWave) String() (res string) {
	for job, time := range times {
		res += fmt.Sprintf("job %v: %v\n", job, time)
	}

	return
}

func (times MachineTimeWave) String() (res string) {
	for machine, time := range times {
		res += fmt.Sprintf("machine %v: %v\n", machine, time)
	}

	return
}

func (state State) String() (res string) {
	res += "==== start of state ====\n"
	res += fmt.Sprintf("-- Jobs:\n%v\n", state.Jobs)
	res += fmt.Sprintf("-- Number of executed:\n%v\n", state.Executed)
	res += fmt.Sprintf("-- Jobs time wave:\n%v\n", state.JobTimeWave)
	res += fmt.Sprintf("-- Machine time wave:\n%v\n", state.MachineTimeWave)
	res += fmt.Sprintf("-- Left total time on each job:\n%v\n", state.LeftTotalTime)
	res += "-------------------------\n"
	min, max := state.EstimateTime()
	res += fmt.Sprintf("-- Estimate of total time: [min = %v; max = %v]\n", min, max)
	res += "==== end of current state ====\n"

	return
}

package state

import (
	"jobShop/util"
	. "jobShop/base"
)

type State struct {
	Jobs Jobs

	Executed        NumberOfExecutedTasks
	JobTimeWave     JobTimeWave
	MachineTimeWave MachineTimeWave

	LeftTotalTime JobsTotalTime

	Parent *State
}

func NewState(jobs Jobs) State {
	jobsNumber := len(jobs)
	machineNumber := len(jobs.ToMachines())

	return State{
		Jobs:            jobs,
		Executed:        make(NumberOfExecutedTasks, jobsNumber),
		JobTimeWave:     make(JobTimeWave, jobsNumber),
		MachineTimeWave: make(MachineTimeWave, machineNumber),
		LeftTotalTime:   NewJobsTotalTime(jobs),
		Parent:          nil,
	}
}

func (s State) startTimeFor(job, task int) (startTime int) {
	machine := s.Jobs[job][task].Machine
	startTime = util.Max(s.JobTimeWave[job], s.MachineTimeWave[machine])
	return
}

func (s State) endOf(job, task int) int {
	return s.startTimeFor(job, task) + s.Jobs[job][task].Time
}

func (s State) updateTimeWave(job, task int) {
	machine, time := s.Jobs[job][task].Deconstruct()
	endOfCurrentTask := s.endOf(job, task)

	s.JobTimeWave[job] = endOfCurrentTask
	s.MachineTimeWave[machine] = endOfCurrentTask
	s.LeftTotalTime[job] -= time
}

func (s State) Execute(job, task int) {
	s.updateTimeWave(job, task)
	s.Executed[job]++
}

func (s State) NextTaskIndexOf(job int) (int, bool) {
	nextTaskIndex := s.Executed[job]
	if nextTaskIndex >= len(s.Jobs[job]) {
		return 0, false
	}

	return nextTaskIndex, true
}

func (s State) NextTaskOf(job int) (*Task, bool) {
	if nextTaskIndex, ok := s.NextTaskIndexOf(job); ok {
		return &s.Jobs[job][nextTaskIndex], true

	} else {
		return nil, false
	}
}

func (s State) EstimateTime() (min, max int) {
	var maxEstimateOfJobEnd, totalLeftTime, maxTimeWave int

	for job, currentTime := range s.JobTimeWave {
		leftTime := s.LeftTotalTime[job]
		totalLeftTime += leftTime

		if maxTimeWave < currentTime {
			maxTimeWave = currentTime
		}

		if minPossibleTimeOfEnd := currentTime + leftTime; maxEstimateOfJobEnd < minPossibleTimeOfEnd {
			maxEstimateOfJobEnd = minPossibleTimeOfEnd
		}
	}

	min = maxEstimateOfJobEnd
	max = maxTimeWave + totalLeftTime
	return
}

func (s State) IsFinish() bool {
	return s.Executed.IsExecutedAll(s.Jobs)
}

func (s State) Makespan() int {
	min, max := s.EstimateTime()
	if min != max {
		panic("scheduling not finished yet")
	}

	return min
}

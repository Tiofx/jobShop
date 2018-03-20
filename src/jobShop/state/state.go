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

	Parent   *State
	redoData redoData
	JobOrder []int
	makespan int
}

type redoData struct {
	machine                                                          Machine
	job, task, jobTimeWave, machineTimeWave, leftTotalTime, executed int
}

func (rd *redoData) deconstruct() (job, task int, machine Machine, jobTimeWave, machineTimeWave, leftTotalTime, executed int) {
	return rd.job, rd.task, rd.machine, rd.jobTimeWave, rd.machineTimeWave, rd.leftTotalTime, rd.executed
}

func (s *State) saveForRedo(job int, task int) {
	machine := s.Jobs[job][task].Machine

	s.redoData.job = job
	s.redoData.task = task
	s.redoData.machine = machine

	s.redoData.jobTimeWave = s.JobTimeWave[job]
	s.redoData.machineTimeWave = s.MachineTimeWave[machine]
	s.redoData.leftTotalTime = s.LeftTotalTime[job]
	s.redoData.executed = s.Executed[job]
}

func NewState(jobs Jobs) State {
	jobsNumber := len(jobs)

	return State{
		Jobs:            jobs,
		Executed:        make(NumberOfExecutedTasks, jobsNumber),
		JobTimeWave:     make(JobTimeWave, jobsNumber),
		MachineTimeWave: make(MachineTimeWave, jobs.MachineNumber()),
		LeftTotalTime:   NewJobsTotalTime(jobs),
		Parent:          nil,
		JobOrder:        make([]int, 0, jobs.TotalTaskNumber()),
	}
}

func (s *State) Reset() {
	util.FillIntsWith(s.Executed, 0)
	util.FillIntsWith(s.JobTimeWave, 0)
	util.FillIntsWith(s.MachineTimeWave, 0)
	s.LeftTotalTime.SetUpBy(s.Jobs)

	util.FillIntsWith(s.JobOrder, 0)
	s.JobOrder = s.JobOrder[0:0]
	s.redoData = redoData{}
	s.makespan = 0
	s.Parent = nil
}

func (s *State) startTimeFor(job, task int) (startTime int) {
	machine := s.Jobs[job][task].Machine
	startTime = util.Max(s.JobTimeWave[job], s.MachineTimeWave[machine])
	return
}

func (s *State) endOf(job, task int) int {
	return s.startTimeFor(job, task) + s.Jobs[job][task].Time
}

func (s *State) updateTimeWave(job, task int) {
	machine, time := s.Jobs[job][task].Deconstruct()
	endOfCurrentTask := s.endOf(job, task)

	s.JobTimeWave[job] = endOfCurrentTask
	s.MachineTimeWave[machine] = endOfCurrentTask
	s.LeftTotalTime[job] -= time
}

func (s *State) Execute(job, task int) {
	s.saveForRedo(job, task)

	s.updateTimeWave(job, task)
	s.Executed[job]++

	s.JobOrder = append(s.JobOrder, job)
}

func (s *State) Undo() {
	job, _, machine,
	jobTimeWave, machineTimeWave, leftTotalTime, executed := s.redoData.deconstruct()

	s.JobTimeWave[job] = jobTimeWave
	s.MachineTimeWave[machine] = machineTimeWave
	s.LeftTotalTime[job] = leftTotalTime
	s.Executed[job] = executed

	s.JobOrder = s.JobOrder[:len(s.JobOrder)-1]
}

func (s *State) NextTaskIndexOf(job int) (int, bool) {
	nextTaskIndex := s.Executed[job]
	if nextTaskIndex >= len(s.Jobs[job]) {
		return 0, false
	}

	return nextTaskIndex, true
}

func (s *State) NextTaskOf(job int) (*Task, bool) {
	if nextTaskIndex, ok := s.NextTaskIndexOf(job); ok {
		return &s.Jobs[job][nextTaskIndex], true

	} else {
		return nil, false
	}
}

func (s *State) EstimateTime() (min, max int) {
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

func (s *State) IsFinish() bool {
	return s.Executed.IsExecutedAll(s.Jobs)
}
func (s *State) Makespan() int {
	//if s.makespan == 0 {
	//_, max := util.MinMax(s.JobTimeWave)
	max := util.MaxOf(s.JobTimeWave)
	s.makespan = max
	//}

	//return s.makespan
	return max
}

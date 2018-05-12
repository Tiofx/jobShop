package state

import (
	"github.com/Tiofx/jobShop/util"
	. "github.com/Tiofx/jobShop/base"
)

type State struct {
	Jobs Jobs

	Executed        NumberOfExecutedTasks
	JobTimeWave     JobTimeWave
	MachineTimeWave MachineTimeWave

	LeftTotalTime JobsTotalTime

	Parent   *State
	redoData redoData
	JobOrder []uint64
	makespan uint64
}

type redoData struct {
	machine                                                          Machine
	job, task, jobTimeWave, machineTimeWave, leftTotalTime, executed uint64
}

func (rd *redoData) deconstruct() (job, task uint64, machine Machine, jobTimeWave, machineTimeWave, leftTotalTime, executed uint64) {
	return rd.job, rd.task, rd.machine, rd.jobTimeWave, rd.machineTimeWave, rd.leftTotalTime, rd.executed
}

func (s *State) saveForRedo(job uint64, task uint64) {
	machine := s.Jobs[job][task].Machine

	s.redoData.job = job
	s.redoData.task = task
	s.redoData.machine = machine

	s.redoData.jobTimeWave = s.JobTimeWave[job]

	s.redoData.machineTimeWave = s.MachineTimeWave[machine]
	s.redoData.leftTotalTime = s.LeftTotalTime[job]
	s.redoData.executed = s.Executed[job]

	//s.redoData = redoData{
	//	job:     job,
	//	task:    task,
	//	machine: machine,
	//
	//	jobTimeWave:     s.JobTimeWave[job],
	//	machineTimeWave: s.MachineTimeWave[machine],
	//	leftTotalTime:   s.LeftTotalTime[job],
	//	executed:        s.Executed[job],
	//}
}

func New(jobs Jobs) State {
	jobsNumber := len(jobs)

	return State{
		Jobs:            jobs,
		Executed:        make(NumberOfExecutedTasks, jobsNumber),
		JobTimeWave:     make(JobTimeWave, jobsNumber),
		MachineTimeWave: make(MachineTimeWave, jobs.MachineNumber()),
		LeftTotalTime:   NewJobsTotalTime(jobs),
		Parent:          nil,
		JobOrder:        make([]uint64, 0, jobs.TotalTaskNumber()),
	}
}

func (s *State) Reset() {
	util.FillUintsWith(s.Executed, 0)
	util.FillUintsWith(s.JobTimeWave, 0)
	util.FillUintsWith(s.MachineTimeWave, 0)
	s.LeftTotalTime.SetUpBy(s.Jobs)

	util.FillUintsWith(s.JobOrder, 0)
	s.JobOrder = s.JobOrder[0:0]
	s.redoData = redoData{}
	s.makespan = 0
	s.Parent = nil
}

func (s *State) startTimeFor(job, task uint64) (startTime uint64) {
	machine := s.Jobs[job][task].Machine
	startTime = util.Max(s.JobTimeWave[job], s.MachineTimeWave[machine])
	return
}

func (s *State) endOf(job, task uint64) uint64 {
	return s.startTimeFor(job, task) + s.Jobs[job][task].Time
}

func (s *State) updateTimeWave(job, task uint64) {
	machine, time := s.Jobs[job][task].Deconstruct()
	endOfCurrentTask := s.endOf(job, task)

	s.JobTimeWave[job] = endOfCurrentTask
	s.MachineTimeWave[machine] = endOfCurrentTask
	s.LeftTotalTime[job] -= time
}

func (s *State) Execute(job, task uint64) {
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

func (s *State) NextTaskIndexOf(job uint64) (uint64, bool) {
	nextTaskIndex := s.Executed[job]
	if nextTaskIndex >= uint64(len(s.Jobs[job])) {
		return 0, false
	}

	return nextTaskIndex, true
}

func (s *State) NextTaskOf(job uint64) (*Task, bool) {
	if nextTaskIndex, ok := s.NextTaskIndexOf(job); ok {
		return &s.Jobs[job][nextTaskIndex], true

	} else {
		return nil, false
	}
}

func (s *State) EstimateTime() (min, max uint64) {
	var maxEstimateOfJobEnd, totalLeftTime, maxTimeWave uint64

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
func (s *State) Makespan() uint64 {
	return util.MaxOf(s.JobTimeWave)
	//if s.makespan == 0 {
	//	max := util.MaxOf(s.JobTimeWave)
	//	s.makespan = max
	//}
	//
	//return s.makespan
}

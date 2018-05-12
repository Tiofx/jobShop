package state

import (
	"github.com/getlantern/deepcopy"
)

func (s *State) CopyIn(res *State) {
	var parentRef *State
	parentRef, s.Parent = s.Parent, nil

	//res = NewState(s.Jobs)
	//
	//copy(res.Executed, s.Executed)
	//copy(res.JobTimeWave, s.JobTimeWave)
	//
	//copy(res.MachineTimeWave, s.MachineTimeWave)
	//
	//copy(res.LeftTotalTime, s.LeftTotalTime)
	//
	//res.redoData = s.redoData
	//res.JobOrder = make([]uint64, len(s.JobOrder))
	//copy(res.JobOrder, s.JobOrder)

	deepcopy.Copy(res, s)

	res.makespan = s.makespan
	res.redoData = s.redoData

	s.Parent, res.Parent = parentRef, parentRef
}

func (s *State) Copy() (res State) {
	s.CopyIn(&res)
	return
}

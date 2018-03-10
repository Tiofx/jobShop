package taskWaveByMachineGreed

import (
	"jobShop/util"
	. "jobShop/state"
	"jobShop/base"
)

type Resolver struct {
	State
}

func (r *Resolver) IsBetterThan(second Resolver) bool {
	return r.State.IsBetterThan(second.State)
}

func (s *Resolver) Copy() Resolver {
	return Resolver{s.State.Copy()}
}

func (s Resolver) ExecuteByTaskInfo(info TaskInfo) {
	taskPosition := s.Executed[info.Job]
	s.Execute(info.Job, taskPosition)
}

func (s Resolver) NextTaskWave() (tw TaskWave) {
	tw = make(TaskWave)

	for i := range s.Jobs {
		if task, ok := s.NextTaskOf(i); ok {
			tw.Add(i, task)
		}
	}

	return
}

func (s *Resolver) Next() Resolver {
	tasks := s.NextTaskWave().GetBiggest()
	nextSolution := s.GreedChoice(tasks)
	nextSolution.Parent = &s.State

	return nextSolution
}

func (s Resolver) GreedChoice(tasks TaskInfoSet) Resolver {
	var best *Resolver

	c := util.Combination(len(tasks))
	for tasksOrder, isChanOpen := <-c; isChanOpen; tasksOrder, isChanOpen = <-c {
		newState := s.Copy()

		for _, index := range tasksOrder {
			currentTask := tasks[index]
			newState.ExecuteByTaskInfo(currentTask)
		}

		if best == nil || newState.IsBetterThan(*best) {
			best = &newState
		}
	}

	return *best
}

//func (s Resolver) FindSolution() State {
//	var (
//		currentState Resolver
//		//prev         *State
//	)
//
//	for currentState = s; !currentState.IsFinish(); currentState = currentState.Next() {
//		//fmt.Println(currentState)
//		//currentState.Parent = prev
//	}
//
//	//fmt.Println(currentState)
//
//	return currentState
//}

func (s Resolver) FindSolution() base.Scheduler {
	var currentState Resolver

	for currentState = s; !currentState.IsFinish(); currentState = currentState.Next() {

	}

	return currentState
}

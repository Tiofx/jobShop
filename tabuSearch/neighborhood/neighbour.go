package neighborhood

import (
	"github.com/Tiofx/jobShop/state"
	"github.com/getlantern/deepcopy"
	"github.com/Tiofx/jobShop/util"
	. "github.com/Tiofx/jobShop/tabuSearch/graph_state"
)

type Neighbour struct {
	JobState state.State
	Graph    DisjunctiveGraph
	cycle    uint64

	graphState *State
}

func (n *Neighbour) UpdateByGraph() (success bool) {
	n.JobState.Reset()

	if n.graphState == nil {
		newState := NewState(n.JobState.Jobs, n.Graph)
		n.graphState = &newState
	} else {
		n.graphState.DisjunctiveGraph = n.Graph
		n.graphState.Jobs = n.JobState.Jobs
		util.FillUintsWith(n.graphState.Executed, 0)
	}

	success = n.graphState.To(&n.JobState)

	return
}

type Move struct{ Machine, I, J uint64 }

func (m *Move) Deconstruct() (machine, i, j uint64) {
	return m.Machine, m.I, m.J
}

func (n *Neighbour) Apply(move Move) {
	machine, i, j := move.Deconstruct()
	n.Graph[machine][i], n.Graph[machine][j] = n.Graph[machine][j], n.Graph[machine][i]
}

func (n *Neighbour) Copy() (res Neighbour) {
	deepcopy.Copy(&res, n)
	return
}

func (n *Neighbour) CopyIn(in *Neighbour) {
	n.JobState.CopyIn(&in.JobState)
	deepcopy.Copy(&in.Graph, &n.Graph)
}

func (n *Neighbour) Redo(move Move) {
	n.Apply(move)
}

func (n *Neighbour) Makespan() uint64 {
	return n.JobState.Makespan()
}

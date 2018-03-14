package tabuSearch

import (
	"jobShop/state"
	"github.com/getlantern/deepcopy"
	"jobShop/util"
)

type neighbour struct {
	jobState state.State
	graph    DisjunctiveGraph
	cycle    int

	graphState *State
}

func (n *neighbour) updateByGraph() (success bool) {
	n.jobState.Reset()

	if n.graphState == nil {
		newState := NewState(n.jobState.Jobs, n.graph)
		n.graphState = &newState
	} else {
		n.graphState.DisjunctiveGraph = n.graph
		n.graphState.Jobs = n.jobState.Jobs
		util.FillIntsWith(n.graphState.Executed, 0)
	}

	success = n.graphState.To(&n.jobState)

	return
}

type move struct{ machine, i, j int }

func (m *move) deconstruct() (machine, i, j int) {
	return m.machine, m.i, m.j
}

func (n *neighbour) apply(move move) {
	machine, i, j := move.deconstruct()
	n.graph[machine][i], n.graph[machine][j] = n.graph[machine][j], n.graph[machine][i]
}

func (n *neighbour) copy() (res neighbour) {
	deepcopy.Copy(&res, n)
	return
}

func (n *neighbour) copyIn(in *neighbour) {
	n.jobState.CopyIn(&in.jobState)
	deepcopy.Copy(&in.graph, &n.graph)
}

func (n *neighbour) redo(move move) {
	n.apply(move)
}

func (n *neighbour) Makespan() int {
	return n.jobState.Makespan()
}

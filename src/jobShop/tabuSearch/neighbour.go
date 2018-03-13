package tabuSearch

import (
	"jobShop/state"
	"jobShop/base"
	"github.com/getlantern/deepcopy"
	"jobShop/util"
)

type neighbour struct {
	jobState state.State
	graph    DisjunctiveGraph
	cycle    int
}

type neighboursSet []neighbour

type tabuList neighboursSet

func (n *neighbour) updateByGraph() (success bool) {
	jobState, exist := n.graph.To(n.jobState.Jobs)

	if !exist {
		return false
	}

	n.jobState = *jobState
	return true
}

func (n *neighbour) updateByState() {

}

func newNeighbour(graph DisjunctiveGraph, jobs base.Jobs) (newNeighbour *neighbour, exist bool) {
	if state, exist := graph.To(jobs); exist {
		return &neighbour{
			jobState: *state,
			graph:    graph,
		}, true
	}

	return nil, false
}

func (list tabuList) indexOf(neighbour neighbour) int {
	for index, tabuNeighbour := range list {

		//if tabuNeighbour.Makespan() == neighbour.Makespan() && reflect.DeepEqual(tabuNeighbour.jobState.JobOrder, neighbour.jobState.JobOrder) {
		if tabuNeighbour.Makespan() == neighbour.Makespan() && util.CompareIntSlices(tabuNeighbour.jobState.JobOrder, neighbour.jobState.JobOrder) {
			return index
		}
	}

	return -1
}

func (list tabuList) contain(neighbour neighbour) bool {
	return list.indexOf(neighbour) != -1
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
	//deepcopy.Copy(&in, &n)
	deepcopy.Copy(&in.jobState, &n.jobState)
	deepcopy.Copy(&in.graph, &n.graph)
}

func (n *neighbour) redo(move move) {
	n.apply(move)
}

func (ns neighboursSet) Len() int {
	return len(ns)
}

func (ns neighboursSet) Less(i, j int) bool {
	return ns[i].Makespan() < ns[j].Makespan()
}

func (ns neighboursSet) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func (n neighbour) Makespan() int {
	return n.jobState.Makespan()
}

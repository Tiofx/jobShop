package tabuSearch

import (
	"jobShop/state"
	"jobShop/base"
	"sort"
	"reflect"
)

type neighbour struct {
	jobState state.State
	graph    DisjunctiveGraph
	cycle    int
}

type neighboursSet []neighbour

type tabuList neighboursSet

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

		if tabuNeighbour.Makespan() == neighbour.Makespan() && reflect.DeepEqual(tabuNeighbour.jobState.JobOrder, neighbour.jobState.JobOrder) {
			return index
		}
	}

	return -1
	//return sort.Search(len(list), func(i int) bool {
	//	return reflect.DeepEqual(list[i].graph, neighbour.graph)
	//})
}

func (list tabuList) contain(neighbour neighbour) bool {
	return list.indexOf(neighbour) != -1
}

func (ns neighboursSet) getNextElement(list tabuList) (n *neighbour, exist bool) {
	for _, neighbour := range ns {
		if len(list) == 0 || !list.contain(neighbour) {
			return &neighbour, true
		}
	}

	//if len(list) > 0 {
	//	return &list[0], true
	//}

	return nil, false
}

func createNeighboursWithoutTabu(graphs []DisjunctiveGraph, jobs base.Jobs, tabu tabuList) (res neighboursSet) {
	neighbours := createExistingNeighbours(graphs, jobs)

	for _, neighbour := range neighbours {
		if tabu.contain(neighbour) {
			continue
		}
		res = append(res, neighbour)
	}

	return
}

func createExistingNeighbours(graphs []DisjunctiveGraph, jobs base.Jobs) (res neighboursSet) {
	for _, graph := range graphs {
		if newNeighbour, exist := newNeighbour(graph, jobs); exist {
			res = append(res, *newNeighbour)
		}
	}

	return
}

func (ns *neighboursSet) descendingOrder() {
	sort.Sort(sort.Reverse(ns))
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

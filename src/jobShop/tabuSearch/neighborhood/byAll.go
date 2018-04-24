package neighborhood

import (
	"jobShop/state"
	"jobShop/tabuSearch/graph_state"
)

type ByAll struct{ byJob }

func NewByAll(jobState *state.State, graph *graph_state.DisjunctiveGraph) ByAll {
	return ByAll{
		byJob: byJob{
			JobState: jobState,
			Graph:    graph,
		},
	}
}

func (r ByAll) Generate() []Move {
	var res []Move

	for machine, jobList := range *r.Graph {
		for i, jobI := range jobList {
			for _, jobJ := range jobList[i+1:] {
				res = append(res, Move{Machine: int(machine), I: int(jobI), J: int(jobJ)})
			}
		}
	}

	return res
}

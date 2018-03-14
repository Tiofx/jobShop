package neighborhood

import (
	"math/rand"
	"jobShop/state"
	"jobShop/tabuSearch/graph_state"
)

type ByRandomJob struct{ byJob }

func NewByRandomJob(jobState *state.State, graph *graph_state.DisjunctiveGraph) ByRandomJob {
	return ByRandomJob{
		byJob: byJob{
			JobState: jobState,
			Graph:    graph,
		},
	}
}

func (r byJob) Generator() (iterator <-chan Move) {
	randomJob := job(rand.Int() % len(r.JobState.Jobs))
	tasks := r.taskPositionFor(randomJob)

	return r.generateFor(randomJob, tasks)
}

package neighborhood

import (
	"github.com/Tiofx/jobShop/state"
	"github.com/Tiofx/jobShop/tabuSearch/graph_state"
)

type ByLongestJob struct{ byJob }

func NewByLongestJob(jobState *state.State, graph *graph_state.DisjunctiveGraph) ByLongestJob {
	return ByLongestJob{
		byJob: byJob{
			JobState: jobState,
			Graph:    graph,
		},
	}
}

func (r ByLongestJob) criticalJob() job {
	var (
		criticalJob       uint64
		timeOfCriticalJob uint64
	)

	for job, time := range r.JobState.JobTimeWave {
		if time > timeOfCriticalJob {
			criticalJob = uint64(job)
			timeOfCriticalJob = time
		}
	}

	return job(criticalJob)
}

func (r ByLongestJob) Generate() []Move {
	criticalJob := r.criticalJob()
	tasks := r.taskPositionFor(criticalJob)

	return r.generateFor(tasks)
}

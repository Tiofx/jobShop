package json

import (
	"github.com/Tiofx/jobShop/base"
	"encoding/json"
)

type Jobs base.Jobs

func (js Jobs) String() string {
	bytes, err := json.Marshal(base.Jobs(js))
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

type jobs []Job

func (js Jobs) toJobs() jobs {
	jobs := make(jobs, len(js))
	for i, job := range js{
		jobs[i] = Job(job)
	}

	return jobs
}

package initSolution

import "jobShop/base"

type Solver interface {
	FindSolution() base.Scheduler
}

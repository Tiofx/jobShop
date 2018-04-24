package initSolution

import (
	"github.com/Tiofx/jobShop/base"
	"github.com/Tiofx/jobShop/state"
)

type Solver interface {
	FindSolution() state.State
}

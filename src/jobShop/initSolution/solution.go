package initSolution

import (
	"jobShop/base"
	"jobShop/state"
)

type Solver interface {
	FindSolution() state.State
}

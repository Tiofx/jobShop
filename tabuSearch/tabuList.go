package tabuSearch

import (
	. "github.com/Tiofx/jobShop/tabuSearch/neighborhood"
	"math"
)

type TabuList interface {
	Contain(move Move) bool
	Add(move Move)
	Forget(move Move) bool
	ForgetOldest() Move
}

type tabuList struct {
	memoryCapacity uint64
	list           []Move
}

func (tabu *tabuList) IndexOf(move Move) (index uint64, exist bool) {
	for i, tabuMove := range tabu.list {
		if move == tabuMove {
			return uint64(i), true
		}
	}

	return math.MaxUint64, false
}

func (tabu *tabuList) Contain(move Move) bool {
	_, exist := tabu.IndexOf(move)
	if !exist {
		_, exist := tabu.IndexOf(Move{move.Machine, move.J, move.I})
		return exist
	}
	return exist
}

func (tabu *tabuList) Add(move Move) {
	tabu.list = append(tabu.list, move)

	if uint64(len(tabu.list)) > tabu.memoryCapacity {
		tabu.ForgetOldest()
	}
}

func (tabu *tabuList) Forget(move Move) (success bool) {
	removeIndex, exist := tabu.IndexOf(move)
	if !exist {
		return false
	}

	tabu.list = append(tabu.list[:removeIndex], tabu.list[removeIndex+1:]...)

	return true
}

func (tabu *tabuList) ForgetOldest() Move {
	oldest := tabu.list[0]
	tabu.list = tabu.list[1:]

	return oldest
}

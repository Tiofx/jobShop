package tabuSearch

import (
	. "jobShop/tabuSearch/neighborhood"
)

type TabuList interface {
	Contain(move Move) bool
	Add(move Move)
	Forget(move Move) bool
	ForgetOldest() Move
}

type tabuList struct {
	list []Move
}

func (tabu *tabuList) IndexOf(move Move) (index int, exist bool) {
	for i, tabuMove := range tabu.list {
		if move == tabuMove {
			return i, true
		}
	}

	return -1, false
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

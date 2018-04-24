package tabuSearch

import (
	. "github.com/Tiofx/jobShop/tabuSearch/neighborhood"
	"reflect"
)

type TabuListByState interface {
	Contain(state *Neighbour) bool
	Add(move moveWithMakespan)
	Forget(move moveWithMakespan) bool
	ForgetOldest() Move
}

type moveWithMakespan struct {
	makespan int
	Move
}

type tabuListByState struct {
	FirstState *Neighbour
	list       []moveWithMakespan
}

func (tabu *tabuListByState) IndexOf(move moveWithMakespan) (index int, exist bool) {
	for i, tabuMove := range tabu.list {
		if move == tabuMove {
			return i, true
		}
	}

	return -1, false
}

func (tabu *tabuListByState) Contain(state *Neighbour) bool {
	stateMakespan := state.Makespan()

	for i, t := range tabu.list {
		tabu.FirstState.Apply(t.Move)
		if stateMakespan == t.makespan {
			//fmt.Println("stateMakespan: ", stateMakespan)
			//fmt.Println("compare: ")
			//fmt.Println(tabu.FirstState.Graph)
			//fmt.Println(state.Graph)
			if reflect.DeepEqual(tabu.FirstState.Graph, state.Graph) {
				for j := i; j >= 0; j-- {
					tabu.FirstState.Redo(tabu.list[j].Move)
				}

				//fmt.Println("there is a duplicate")
				return true
			}
		}
	}

	for i := len(tabu.list) - 1; i >= 0; i-- {
		tabu.FirstState.Redo(tabu.list[i].Move)
	}

	return false
}

func (tabu *tabuListByState) Add(move moveWithMakespan) {
	tabu.list = append(tabu.list, move)
}

func (tabu *tabuListByState) Forget(move moveWithMakespan) (success bool) {
	removeIndex, exist := tabu.IndexOf(move)
	if !exist {
		return false
	}

	tabu.list = append(tabu.list[:removeIndex], tabu.list[removeIndex+1:]...)

	return true
}

func (tabu *tabuListByState) ForgetOldest() Move {
	oldest := tabu.list[0]
	tabu.list = tabu.list[1:]

	return oldest.Move
}

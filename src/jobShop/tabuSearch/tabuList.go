package tabuSearch

type TabuList interface {
	Contain(move move) bool
	Add(move move)
	Forget(move move) bool
	ForgetOldest() move
}

type tabuList struct {
	list []move
}

func (tabu tabuList) IndexOf(move move) (index int, exist bool) {
	for i, tabuMove := range tabu.list {
		if move == tabuMove {
			return i, true
		}
	}

	return -1, false
}

func (tabu tabuList) Contain(move move) bool {
	_, exist := tabu.IndexOf(move)
	return exist
}

func (tabu tabuList) Add(move move) {
	tabu.list = append(tabu.list, move)
}

func (tabu tabuList) Forget(move move) (success bool) {
	removeIndex, exist := tabu.IndexOf(move)
	if !exist {
		return false
	}

	tabu.list = append(tabu.list[:removeIndex], tabu.list[removeIndex+1:]...)

	return true
}

func (tabu tabuList) ForgetOldest() move {
	oldest := tabu.list[0]
	tabu.list = tabu.list[:1]

	return oldest
}
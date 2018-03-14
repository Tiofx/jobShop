package tabuSearch

type TabuList interface {
	Contain(move move) bool
	Add(move move)
}

type tabuList struct {
	list []move
}

func (tabu *tabuList) Contain(move move) bool {
	for _, tabuMove := range tabu.list {
		if move == tabuMove {
			return true
		}
	}

	return false
}

func (tabu *tabuList) Add(move move) {
	tabu.list = append(tabu.list, move)
}

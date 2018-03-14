package tabuSearch

type TabuList interface {
	Contain(move move) bool
	Add(move move)
}

type tabuList neighboursSet

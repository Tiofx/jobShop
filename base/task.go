package base

type Task struct {
	Machine Machine `json:"machine"`
	Time    uint64     `json:"time"`
}

func (t Task) Deconstruct() (Machine, uint64) {
	return t.Machine, t.Time
}



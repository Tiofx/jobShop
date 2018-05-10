package base

type Task struct {
	Machine Machine `json:"machine"`
	Time    int     `json:"time"`
}

func (t Task) Deconstruct() (Machine, int) {
	return t.Machine, t.Time
}



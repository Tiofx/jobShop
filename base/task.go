package base

type Task struct {
	Machine Machine
	Time    int
}

func (t Task) Deconstruct() (Machine, int) {
	return t.Machine, t.Time
}

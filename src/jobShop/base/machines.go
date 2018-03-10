package base

type Machines map[Machine]bool

func (ms Machines) AddMachine(machine Machine) {
	ms[machine] = true
}

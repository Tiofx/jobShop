package state

func (s State) Compare(second State) int8 {
	min, max := s.EstimateTime()
	min2, max2 := second.EstimateTime()

	switch {

	case min < min2:
		fallthrough

	case min == min2 && max < max2:
		return 1

	case min == min2 && max == max2:
		return 0

	default:
		return -1
	}
}

func (s State) IsBetterThan(second State) bool {
	return s.Compare(second) == 1
}

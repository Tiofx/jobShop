package state

import "github.com/getlantern/deepcopy"

func (s State) Copy() (res State) {
	var parentRef *State
	parentRef, s.Parent = s.Parent, nil

	deepcopy.Copy(&res, &s)

	s.Parent, res.Parent = parentRef, parentRef
	return
}

package stats

import (
	"time"
	"fmt"
)

type Snapshot struct {
	IterationNumber uint64
	Makespan        uint64
	TimeFromStart   time.Duration
}

func (s Snapshot) String() string {
	return fmt.Sprintf("On iteration number [%v] scheduler's makespan was improved to [%v]. From the start passed [%v]", s.IterationNumber, s.Makespan, s.TimeFromStart)
}
package stats

import (
	"time"
	"fmt"
)

type Snapshot struct {
	IterationNumber int
	Makespan        int
	TimeFromStart   time.Duration
}

func (s Snapshot) String() string {
	return fmt.Sprintf("On iteration number [%v] scheduler's makespan was improved to [%v]. From the start passed [%v]", s.IterationNumber, s.Makespan, s.TimeFromStart)
}
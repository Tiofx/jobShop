package stats

import (
	"time"
	"fmt"
)

type Snapshot struct {
	IterationNumber uint64        `json:"iterationNumber"`
	Makespan        uint64        `json:"makespan"`
	TimeFromStart   time.Duration `json:"timeFromStart"`
}

func (s Snapshot) String() string {
	return fmt.Sprintf("On iteration number [%v] scheduler's makespan was improved to [%v]. From the start passed [%v]", s.IterationNumber, s.Makespan, s.TimeFromStart)
}

package checker

import (
	"fmt"
	"time"
)

type CheckResult struct {
	Target       Target
	IsUp         bool
	StatusCode   int
	ResponseTime time.Duration
	Error        error
}

func (r CheckResult) String() string {
	return fmt.Sprintf(
		"%s | %v | %d | %v | %v",
		r.Target.Name,
		r.IsUp,
		r.StatusCode,
		r.ResponseTime,
		r.Error,
	)
}

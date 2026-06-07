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

func (result *CheckResult) Print() {
	fmt.Println("---------------")
	fmt.Printf(
		"%s | %v | %d | %v | %v\n",
		result.Target.Name,
		result.IsUp,
		result.StatusCode,
		result.ResponseTime,
		result.Error,
	)
}

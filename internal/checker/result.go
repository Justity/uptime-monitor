package checker

import "time"

type Result struct {
	IsUp         bool
	StatusCode   int
	ResponseTime time.Duration
	Error        error
}

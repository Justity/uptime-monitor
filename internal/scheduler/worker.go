package scheduler

import (
	"sync"

	"github.com/Justity/uptime-monitor/internal/checker"
)

func worker(
	jobs <-chan checker.Target,
	results chan<- checker.CheckResult,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for target := range jobs {
		result := checker.CheckTarget(target)

		results <- result
	}
}

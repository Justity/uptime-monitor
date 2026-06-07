package scheduler

import (
	"fmt"
	"sync"
	"time"

	"github.com/Justity/uptime-monitor/internal/checker"
)

func (s *Scheduler) Run() {

	for {

		jobs := make(chan checker.Target)

		results := make(chan checker.CheckResult)

		var wg sync.WaitGroup

		for i := 0; i < s.WorkerCount; i++ {

			wg.Add(1)

			go worker(
				jobs,
				results,
				&wg,
			)
		}

		go func() {

			for _, target := range s.Targets {

				jobs <- target
			}

			close(jobs)

		}()

		go func() {

			wg.Wait()

			close(results)

		}()

		for result := range results {
			fmt.Println(result)
		}

		time.Sleep(
			s.Interval,
		)
	}
}

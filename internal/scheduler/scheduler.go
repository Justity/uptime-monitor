package scheduler

import (
	"fmt"
	"sync"
	"time"

	"github.com/Justity/uptime-monitor/internal/checker"
)

type Scheduler struct {
	Targets  []checker.Target
	Interval time.Duration
}

func (s *Scheduler) Run() {

	for {

		var wg sync.WaitGroup

		fmt.Println("---------------")
		
		for _, target := range s.Targets {

			wg.Add(1)

			go func(target checker.Target) {

				defer wg.Done()

				result := checker.CheckTarget(target)
				fmt.Println(result)

			}(target)
		}

		wg.Wait()

		time.Sleep(s.Interval)
	}
}

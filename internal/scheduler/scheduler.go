package scheduler

import (
	"time"

	"github.com/Justity/uptime-monitor/internal/checker"
)

type Scheduler struct {
	Targets  []checker.Target
	Interval time.Duration
}

func (s *Scheduler) Run() {

	for {

		for _, target := range s.Targets {

			result := checker.CheckTarget(target)

			result.Print()
		}

		time.Sleep(s.Interval)
	}
}

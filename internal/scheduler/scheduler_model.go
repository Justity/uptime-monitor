package scheduler

import (
	"fmt"
	"time"

	"github.com/Justity/uptime-monitor/internal/checker"
	"github.com/Justity/uptime-monitor/internal/notifier"
	"github.com/Justity/uptime-monitor/internal/storage"
)

type Scheduler struct {
	Targets     []checker.Target
	Interval    time.Duration
	WorkerCount int
	Storage     storage.Storage
	Notifier    notifier.Notifier
}

func (s *Scheduler) ProcessResult(result checker.CheckResult) {
	fmt.Println(result)
	err := s.Storage.Save(result)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Notifier.Notify(result)
	if err != nil {
		fmt.Println(err)
	}
}

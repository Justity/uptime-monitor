package scheduler

import (
	"time"

	"github.com/Justity/uptime-monitor/internal/checker"
)

type Scheduler struct {
	Targets     []checker.Target
	Interval    time.Duration
	WorkerCount int
}

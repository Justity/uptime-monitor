package main

import (
	"time"

	"github.com/Justity/uptime-monitor/internal/config"
	"github.com/Justity/uptime-monitor/internal/scheduler"
)

func main() {
	targets, err := config.LoadTargets(
		"configs/targets.json",
	)

	if err != nil {
		panic(err)
	}

	s := scheduler.Scheduler{
		Targets:  targets,
		Interval: 30 * time.Second,
	}

	s.Run()
}

package main

import (
	"time"

	"github.com/Justity/uptime-monitor/internal/config"
	"github.com/Justity/uptime-monitor/internal/notifier"
	"github.com/Justity/uptime-monitor/internal/scheduler"
	"github.com/Justity/uptime-monitor/internal/storage"
)

func main() {
	memoryStorage := &storage.MemoryStorage{}
	notifierChannel := &notifier.ConsoleNotifier{}

	targets, err := config.LoadTargets(
		"configs/targets.json",
	)

	if err != nil {
		panic(err)
	}

	s := scheduler.Scheduler{
		Targets:     targets,
		Interval:    30 * time.Second,
		WorkerCount: 3,
		Storage:     memoryStorage,
		Notifier:    notifierChannel,
	}

	s.Run()
}

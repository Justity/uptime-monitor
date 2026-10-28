package storage

import "github.com/Justity/uptime-monitor/internal/checker"

type Storage interface {
	Save(
		result checker.CheckResult,
	) error
}

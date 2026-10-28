package notifier

import "github.com/Justity/uptime-monitor/internal/checker"

type Notifier interface {
	Notify(
		result checker.CheckResult,
	) error
}

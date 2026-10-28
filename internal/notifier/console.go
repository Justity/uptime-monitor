package notifier

import (
	"fmt"

	"github.com/Justity/uptime-monitor/internal/checker"
)

type ConsoleNotifier struct {
}

func (c *ConsoleNotifier) Notify(
	result checker.CheckResult,
) error {

	fmt.Println("notify", result)

	return nil
}

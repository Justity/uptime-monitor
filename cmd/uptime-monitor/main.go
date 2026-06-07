package main

import (
	"github.com/Justity/uptime-monitor/internal/checker"
	"github.com/Justity/uptime-monitor/internal/config"
)

func main() {
	targets, err := config.LoadTargets(
		"configs/targets.json",
	)

	if err != nil {
		panic(err)
	}

	for _, target := range targets {
		result := checker.CheckTarget(target)
		result.Print()
	}
}

package main

import (
	"fmt"

	"github.com/Justity/uptime-monitor/internal/checker"
)

func main() {

	targets := []checker.Target{
		{
			Name: "Google",
			URL:  "https://google.com",
		},
		{
			Name: "GitHub",
			URL:  "https://github.com",
		},
		{
			Name: "Go",
			URL:  "https://go.dev",
		},
	}

	for _, target := range targets {
		result := checker.CheckTarget(target)

		fmt.Println("---------------")
		fmt.Printf(
			"%s | %v | %d | %v | %v\n",
			target.Name,
			result.IsUp,
			result.StatusCode,
			result.ResponseTime,
			result.Error,
		)
	}
}

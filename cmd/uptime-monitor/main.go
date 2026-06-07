package main

import (
	"fmt"

	"github.com/Justity/uptime-monitor/internal/checker"
)

func main() {

	target := checker.Target{
		Name: "Google",
		URL:  "https://google.com",
	}

	result := checker.CheckTarget(
		target,
	)

	fmt.Println("---------------")
	fmt.Println(target.Name)
	fmt.Println("Available:", result.IsUp)
	fmt.Println("Status:", result.StatusCode)
	fmt.Println("Response:", result.ResponseTime)
}

package main

import (
	"fmt"

	"github.com/Justity/uptime-monitor/internal/checker"
)

func main() {
	result := checker.Check(
		"https://google.com",
	)

	fmt.Println("Available:", result.IsUp)
	fmt.Println("Status:", result.StatusCode)
	fmt.Println("Response:", result.ResponseTime)
}

package checker

import (
	"net/http"
	"time"
)

type Result struct {
	IsUp         bool
	StatusCode   int
	ResponseTime time.Duration
}

func Check(url string) Result {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		return Result{
			IsUp: false,
		}
	}

	defer resp.Body.Close()

	return Result{
		IsUp:         resp.StatusCode >= 200 && resp.StatusCode < 400,
		StatusCode:   resp.StatusCode,
		ResponseTime: time.Since(start),
	}
}

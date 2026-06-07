package checker

import (
	"net/http"
	"time"
)

func CheckTarget(
	target Target,
) Result {

	start := time.Now()

	resp, err := http.Get(target.URL)

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

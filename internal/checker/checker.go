package checker

import (
	"net/http"
	"time"
)

func CheckTarget(
	target Target,
) CheckResult {

	start := time.Now()

	resp, err := http.Get(target.URL)

	if err != nil {
		return CheckResult{
			Target: target,
			IsUp:   false,
			Error:  err,
		}
	}

	defer resp.Body.Close()

	return CheckResult{
		Target:       target,
		IsUp:         true,
		StatusCode:   resp.StatusCode,
		ResponseTime: time.Since(start),
	}
}

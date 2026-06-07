package config

import (
	"encoding/json"
	"os"

	"github.com/Justity/uptime-monitor/internal/checker"
)

func LoadTargets(
	path string,
) ([]checker.Target, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var targets []checker.Target

	err = json.Unmarshal(
		data,
		&targets,
	)

	if err != nil {
		return nil, err
	}

	return targets, nil
}

package storage

import "github.com/Justity/uptime-monitor/internal/checker"

type MemoryStorage struct {
	Results []checker.CheckResult
}

func (m *MemoryStorage) Save(
	result checker.CheckResult,
) error {

	m.Results = append(
		m.Results,
		result,
	)

	return nil
}

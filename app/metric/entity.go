package metric

import "time"

// Metric struct which contains
// data metric of a game level
type Metric struct {
	LevelName string    `json:"level_name"`
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func (m *Metric) IsInRange(start, end time.Time) bool {
	if m.Timestamp.Before(start) || m.Timestamp.Equal(end) || m.Timestamp.After(end) {
		return false
	}
	return true
}

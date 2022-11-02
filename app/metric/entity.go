package metric

import "time"

// Metric struct which contains
// data metric of a game level
type Metric struct {
	LevelName string    `json:"level_name"`
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

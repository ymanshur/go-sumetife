package util

import "time"

// ParseToTimeUTC return time object with UTC timezone
func ParseToTimeUTC(inputTime string) (time.Time, error) {
	time, err := time.Parse(time.RFC3339, inputTime)
	if err != nil {
		return time, err
	}
	return time.UTC(), nil
}

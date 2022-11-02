package metric

import (
	"testing"
	"time"
)

func TestMetricIsInRange(t *testing.T) {
	testCases := []struct {
		desc      string
		metric    Metric
		startTime time.Time
		endTime   time.Time
		want      bool
	}{
		{
			desc: "input parameter start and time are in local timezone",
			metric: Metric{
				LevelName: "level1",
				Value:     126,
				Timestamp: time.Date(2022, time.January, 1, 0, 23, 0, 0, time.UTC),
			},
			startTime: time.Date(2022, time.January, 1, 7, 20, 0, 0, time.Local),
			endTime:   time.Date(2022, time.January, 2, 0, 0, 0, 0, time.Local),
			want:      true,
		},
		{
			desc: "input parameter start is in UTC",
			metric: Metric{
				LevelName: "lobby_screen",
				Value:     73,
				Timestamp: time.Date(2022, time.January, 1, 2, 50, 0, 0, time.UTC),
			},
			startTime: time.Date(2022, time.January, 1, 7, 20, 0, 0, time.UTC),
			endTime:   time.Date(2022, time.January, 2, 0, 0, 0, 0, time.Local),
			want:      false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.metric.IsInRange(tC.startTime, tC.endTime)
			if got != tC.want {
				t.Fatalf("got %v; want %v", got, tC.want)
			}
		})
	}
}

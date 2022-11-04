package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetricResultFormatter_Success(t *testing.T) {
	// Setup
	mockResult := map[string]int{
		"lobby_screen": 161,
	}
	metricResult := MetricResultFormatter(mockResult)

	// Assertions
	assert.Equal(t, 1, len(metricResult))
	assert.Equal(t, "lobby_screen", metricResult[0].LevelName)
}

package metric

// MetricResult struct which contains
// total value of a game level
type MetricResult struct {
	LevelName  string `json:"level_name" yaml:"level_name"`
	TotalValue int    `json:"total_value" yaml:"total_value"`
}

// MetricResultFormatter return array of MetricResult
func MetricResultFormatter(result map[string]int) []MetricResult {
	metricResult := []MetricResult{}

	for key, value := range result {
		metricResult = append(metricResult, MetricResult{
			LevelName:  key,
			TotalValue: value,
		})
	}

	return metricResult
}

package metric

// MetricResultFormatter struct which contains
// total value of a game level
type MetricResultFormatter struct {
	LevelName  string `json:"level_name" yaml:"level_name"`
	TotalValue int    `json:"total_value" yaml:"total_value"`
}

// FormatResult return array of MetricResultFormatter
func FormatMetricResult(result map[string]int) []MetricResultFormatter {
	metricResult := []MetricResultFormatter{}

	for key, value := range result {
		metricResult = append(metricResult, MetricResultFormatter{
			LevelName:  key,
			TotalValue: value,
		})
	}

	return metricResult
}

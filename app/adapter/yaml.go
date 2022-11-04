package adapter

import (
	"sumetife/metric"

	"gopkg.in/yaml.v3"
)

// YAMLFileEncoder return the yaml file content ([]byte) of metric result
func YAMLFileEncoder(metricResult []metric.MetricResult) ([]byte, error) {
	fileContent, err := yaml.Marshal(metricResult)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}

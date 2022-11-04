package adapter

import (
	"os"
	"sumetife/metric"
)

type MetricFileDecoder interface {
	Decode(file *os.File, metrics *[]metric.Metric) error
}

type MetricFileEncoder interface {
	Encode(metricResult []metric.MetricResult) ([]byte, error)
}

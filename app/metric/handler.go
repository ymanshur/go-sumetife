package metric

import (
	"fmt"
	_ "log"
	"os"
)

var OpenFile = os.Open
var WriteFile = os.WriteFile

// MetricFileDecoder decode the file and save it
// to the defined metrics
type MetricFileDecoder func(file *os.File, metrics *[]Metric) error

// MetricFileEncoder unmarshal (encode)
// the metric result value to file content
type MetricFileEncoder func(metricResult []MetricResult) ([]byte, error)

type MetricHandler interface {
	// GetMetricsDataFromFile return array of mapped metric data
	GetMetricsDataFromFile(fileName string) ([]Metric, error)
	// WriteMetricResultToFile generate a file which consist all metric result
	WriteMetricResultToFile(fileName string, metricResult []MetricResult) error
}

type metricHandler struct {
	fileEncoder MetricFileEncoder
	fileDecoder MetricFileDecoder
}

func (h *metricHandler) GetMetricsDataFromFile(fileName string) ([]Metric, error) {
	// initialize our metrics array
	var metrics []Metric

	// open a metric file
	file, err := OpenFile(fileName)
	// if os.Open returns an error then handle it
	if err != nil {
		return metrics, err
	}
	// log.Println("Successfully opened", fileName)

	// defer the closing of our file so that we can parse it later on
	defer file.Close()

	// unmarshal (decoder) the file to metrics
	if err := h.fileDecoder(file, &metrics); err != nil {
		return nil, err
	}

	return metrics, nil
}

func (h *metricHandler) WriteMetricResultToFile(fileName string, metricResult []MetricResult) error {
	// marshal (encode) the result to []byte with formatter function
	fileContent, err := h.fileEncoder(metricResult)
	if err != nil {
		return err
	}

	// print result to the console
	fmt.Println(string(fileContent))

	// write the file content which contains our metric result into a file
	if err := WriteFile(fileName, fileContent, 0644); err != nil {
		return err
	}

	// log.Println("Successfully generated " + fileName)

	return nil
}

func NewMetricHandler(fileDecoder MetricFileDecoder, fileEncoder MetricFileEncoder) MetricHandler {
	return &metricHandler{
		fileDecoder: fileDecoder,
		fileEncoder: fileEncoder,
	}
}

package metric

import (
	"fmt"
	"log"
	"os"
)

var OpenFile = os.Open
var WriteFile = os.WriteFile

// MetricFileDecoder update the defined metrics
// by each metric value in the file
type MetricFileDecoder func(file *os.File, metrics *[]Metric) error

// MetricEncoder unmarshal the (any) value
type MetricEncoder func(v any) ([]byte, error)

type MetricHandler interface {
	// GetMetricsDataFromFile return array of mapped metric data
	GetMetricsDataFromFile(fileName string) ([]Metric, error)
	// WriteMetricResultToFile generate a file which consist metric result
	WriteMetricResultToFile(fileName string, result map[string]int) error
}

type metricHandler struct {
	encoder     MetricEncoder
	fileDecoder MetricFileDecoder
}

func (h metricHandler) GetMetricsDataFromFile(fileName string) ([]Metric, error) {
	// initialize our metrics array
	var metrics []Metric

	// Open a metric file
	file, err := OpenFile(fileName)
	// if os.Open returns an error then handle it
	if err != nil {
		return metrics, err
	}
	// log.Println("Successfully opened", fileName)

	// defer the closing of our file so that we can parse it later on
	defer file.Close()

	if err := h.fileDecoder(file, &metrics); err != nil {
		return nil, err
	}

	return metrics, nil
}

func (h metricHandler) WriteMetricResultToFile(fileName string, result map[string]int) error {
	metricResult := MetricResultFormatter(result)

	// print result to the console
	for _, v := range metricResult {
		fmt.Printf("Level name: %s, total value: %d\n", v.LevelName, v.TotalValue)
	}

	// marshal (encode) the result with formatter function
	fileContent, err := h.encoder(metricResult)
	if err != nil {
		return err
	}

	// white the file content which contains our result into a file
	if err := WriteFile(fileName, fileContent, 0644); err != nil {
		return err
	}

	log.Println("Successfully generated " + fileName)

	return nil
}

func NewMetricHandler(decoder MetricFileDecoder, encoder MetricEncoder) MetricHandler {
	return &metricHandler{
		fileDecoder: decoder,
		encoder:     encoder,
	}
}

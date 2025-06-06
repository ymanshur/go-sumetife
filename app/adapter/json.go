package adapter

import (
	"encoding/json"
	"io"
	"os"
	"sumetife/metric"
)

// JSONFileDecoder updates defined metrics data from a json file
func JSONFileDecoder(file *os.File, metrics *[]metric.Metric) error {
	// read our opened file as a byte array
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// unmarshal (decode) our byteArray which contains our
	// json file content into 'metrics' which we defined above
	json.Unmarshal(byteValue, &metrics)

	return nil
}

// JSONFileEncoder return the json file content ([]byte) of metric result
func JSONFileEncoder(metricResult []metric.MetricResult) ([]byte, error) {
	fileContent, err := json.MarshalIndent(metricResult, "", "\t")
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}

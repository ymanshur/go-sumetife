package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Metric struct which contains
// data metric of a game level
type Metric struct {
	LevelName string    `json:"level_name"`
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

// MetricResult struct which contains
// total value of a game level
type MetricResult struct {
	LevelName  string `json:"level_name" yaml:"level_name"`
	TotalValue int    `json:"total_value" yaml:"total_value"`
}

// MetricResultFormatter return array of MetricResult
func MetricResultFormatter(result map[string]int) []MetricResult {
	metricResults := []MetricResult{}

	for key, value := range result {
		metricResult := MetricResult{
			LevelName:  key,
			TotalValue: value,
		}
		metricResults = append(metricResults, metricResult)
	}

	return metricResults
}

// OpenMetricFile return a metric file
func OpenMetricFile(filename string) (*os.File, error) {
	// Open a metric file
	file, err := os.Open(filename)

	// if os.Open returns an error then handle it
	if err != nil {
		return file, err
	}

	fmt.Println("Successfully opened", filename)

	return file, nil
}

type FileAdapter interface {
	GetMetricsData(filename string) ([]Metric, error)
}

func NewFileAdapter(fileType string) FileAdapter {
	if fileType == "csv" {
		return CSVFileAdapter{}
	}
	return JSONFileAdapter{}
}

type CSVFileAdapter struct {
}

func (a CSVFileAdapter) GetMetricsData(filename string) ([]Metric, error) {
	// initialize our metrics array
	var metrics []Metric

	file, err := OpenMetricFile(filename)
	// if os.Open returns an error then handle it
	if err != nil {
		return metrics, err
	}

	// defer the closing of our file so that we can parse it later on
	defer file.Close()

	// read our opened file as a string matrix (array or array)
	csvReader := csv.NewReader(file)
	csvLines, _ := csvReader.ReadAll()

	// map our csvLines which contains our
	// csv file content into 'metrics' which we defined above
	for _, line := range csvLines[1:] {
		matricValue, _ := strconv.Atoi(line[2])
		matricTimestamp, _ := time.Parse(time.RFC3339, line[0])
		metric := Metric{
			LevelName: line[1],
			Value:     matricValue,
			Timestamp: matricTimestamp,
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil
}

type JSONFileAdapter struct {
}

func (a JSONFileAdapter) GetMetricsData(filename string) ([]Metric, error) {
	// initialize our metrics array
	var metrics []Metric

	file, err := OpenMetricFile(filename)
	// if os.Open returns an error then handle it
	if err != nil {
		return metrics, err
	}

	// defer the closing of our file so that we can parse it later on
	defer file.Close()

	// read our opened file as a byte array
	byteValue, _ := io.ReadAll(file)

	// unmarshal (decode) our byteArray which contains our
	// json file content into 'metrics' which we defined above
	json.Unmarshal(byteValue, &metrics)

	return metrics, nil
}

func WriteMetricResultToYAMLFile(result []MetricResult) error {
	// marshal (encode) the result with formatter function
	fileContent, err := yaml.Marshal(result)
	if err != nil {
		return err
	}

	// white the file content which contains our result into a yaml file
	if err := os.WriteFile("output/output.yaml", fileContent, 0644); err != nil {
		return err
	}

	fmt.Println("Successfully generate output.yaml")

	return nil
}

func WriteMetricResultToJSONFile(result []MetricResult) error {
	// marshal (encode) the result with formatter function
	fileContent, err := json.Marshal(result)
	if err != nil {
		return err
	}

	// white the file content which contains our result into a json file
	if err := os.WriteFile("output/output.json", fileContent, 0644); err != nil {
		return err
	}

	fmt.Println("Successfully generate output.json")

	return nil
}

func main() {
	fileType := "csv"
	dirPath := "data"

	// read the data directory
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	fileAdapter := NewFileAdapter(fileType)

	result := map[string]int{}
	// iterate through every metric file and summarize those value of each level name
	for _, file := range files {
		fileExt := filepath.Ext(file.Name())
		if fileExt != ("." + fileType) {
			continue
		}

		filename := filepath.Join(dirPath, file.Name())
		metrics, err := fileAdapter.GetMetricsData(filename)
		if err != nil {
			log.Fatal(err)
		}

		// iterate through every metric data and
		// add those value into 'result' map
		for i := 0; i < len(metrics); i++ {
			result[metrics[i].LevelName] += metrics[i].Value
		}
	}

	metricResult := MetricResultFormatter(result)
	// write the file content which contains our result into a json file
	if err := WriteMetricResultToJSONFile(metricResult); err != nil {
		log.Fatal(err)
	}
}

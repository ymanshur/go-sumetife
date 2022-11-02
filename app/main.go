package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
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

func GetMetricsDataFromCSVFile(file *os.File) []Metric {
	// initialize our metrics array
	var metrics []Metric

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

	return metrics
}

func GetMetricsDataFromJSONFile(file *os.File) []Metric {
	// initialize our metrics array
	var metrics []Metric

	// read our opened file as a byte array
	byteValue, _ := io.ReadAll(file)

	// unmarshal (decode) our byteArray which contains our
	// json file content into 'metrics' which we defined above
	json.Unmarshal(byteValue, &metrics)

	return metrics
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
	// Open a metric file
	// filename := "data/csv/01-jan.csv"
	filename := "data/json/01-jan.json"
	file, err := OpenMetricFile(filename)
	// if os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}

	// defer the closing of our file so that we can parse it later on
	defer file.Close()

	// metrics := GetMetricsDataFromCSVFile(file)
	metrics := GetMetricsDataFromJSONFile(file)

	// iterate through every metric file and
	// summarize those value for each level name at the data
	result := map[string]int{}
	for i := 0; i < len(metrics); i++ {
		result[metrics[i].LevelName] += metrics[i].Value
	}

	// white the file content which contains our result into a json file
	matricResult := MetricResultFormatter(result)
	if err := WriteMetricResultToYAMLFile(matricResult); err != nil {
		log.Fatal(err)
	}
}

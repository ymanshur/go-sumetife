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
func OpenMetricFile(fileName string) (*os.File, error) {
	// Open a metric file
	file, err := os.Open(fileName)

	// if os.Open returns an error then handle it
	if err != nil {
		return file, err
	}

	fmt.Println("Successfully opened", fileName)

	return file, nil
}

type MetricFileDecoder func(file *os.File, metrics *[]Metric) error

type MetricEncoder func(v any) ([]byte, error)

type MetricHandler interface {
	GetMetricsDataFromFile(fileName string) ([]Metric, error)
	WriteMetricResultToFile(fileName string, result []MetricResult) error
}

type metricHandler struct {
	encoder     MetricEncoder
	fileDecoder MetricFileDecoder
}

func (h metricHandler) GetMetricsDataFromFile(fileName string) ([]Metric, error) {
	// initialize our metrics array
	var metrics []Metric

	file, err := OpenMetricFile(fileName)
	// if os.Open returns an error then handle it
	if err != nil {
		return metrics, err
	}

	// defer the closing of our file so that we can parse it later on
	defer file.Close()

	if err := h.fileDecoder(file, &metrics); err != nil {
		return nil, err
	}

	return metrics, nil
}

func (h metricHandler) WriteMetricResultToFile(fileName string, result []MetricResult) error {
	// marshal (encode) the result with formatter function
	fileContent, err := h.encoder(result)
	if err != nil {
		return err
	}

	// white the file content which contains our result into a file
	if err := os.WriteFile(fileName, fileContent, 0644); err != nil {
		return err
	}

	fmt.Println("Successfully generate " + fileName)

	return nil
}

func NewMetricHandler(decoder MetricFileDecoder, encoder MetricEncoder) MetricHandler {
	return &metricHandler{
		fileDecoder: decoder,
		encoder:     encoder,
	}
}

func main() {
	inputFileType := "csv"
	outputFileType := "json"
	outputFileName := "out"
	directory := "data"

	// read the data directory
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	fileDecoder := func(file *os.File, metrics *[]Metric) error {
		// read our opened file as a byte array
		byteValue, _ := io.ReadAll(file)

		// unmarshal (decode) our byteArray which contains our
		// json file content into 'metrics' which we defined above
		json.Unmarshal(byteValue, &metrics)

		return nil
	}
	if inputFileType == "csv" {
		fileDecoder = func(file *os.File, metrics *[]Metric) error {
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
				*metrics = append(*metrics, metric)
			}

			return nil
		}
	}

	encoder := json.Marshal
	if outputFileType == "yaml" {
		encoder = yaml.Marshal
	}

	fileHandler := NewMetricHandler(fileDecoder, encoder)

	result := map[string]int{}
	// iterate through every metric file and summarize those value of each level name
	for _, file := range files {
		fileExt := filepath.Ext(file.Name())
		if fileExt != ("." + inputFileType) {
			continue
		}

		fileName := filepath.Join(directory, file.Name())
		metrics, err := fileHandler.GetMetricsDataFromFile(fileName)
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
	// write the file content which contains our result into a file
	if err := fileHandler.WriteMetricResultToFile(outputFileName+"."+outputFileType, metricResult); err != nil {
		log.Fatal(err)
	}
}

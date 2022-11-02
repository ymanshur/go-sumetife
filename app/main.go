package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sumetife/metric"
	"time"

	"gopkg.in/yaml.v3"
)

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

func CSVFileDecoder(file *os.File, metrics *[]metric.Metric) error {
	// read our opened file as a string matrix (array or array)
	csvReader := csv.NewReader(file)
	csvLines, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	// map our csvLines which contains our
	// csv file content into 'metrics' which we defined above
	for _, line := range csvLines[1:] {
		matricValue, err := strconv.Atoi(line[2])
		if err != nil {
			return err
		}

		matricTimestamp, err := time.Parse(time.RFC3339, line[0])
		if err != nil {
			return err
		}

		metric := metric.Metric{
			LevelName: line[1],
			Value:     matricValue,
			Timestamp: matricTimestamp,
		}
		*metrics = append(*metrics, metric)
	}

	return nil
}

func main() {
	inputFileType := "json"
	outputFileType := "yaml"
	outputFileName := "out"
	directory := "data"
	startTime := "2022-01-01T07:20:00+07:00"
	endTime := "2022-01-02T00:00:00+07:00"

	// read the data directory
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	fileDecoder := JSONFileDecoder
	if inputFileType == "csv" {
		fileDecoder = CSVFileDecoder
	}

	encoder := json.Marshal
	if outputFileType == "yaml" {
		encoder = yaml.Marshal
	}

	fileHandler := metric.NewMetricHandler(fileDecoder, encoder)

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

		startTime, err := time.Parse(time.RFC3339, startTime)
		if err != nil {
			log.Fatal(err)
		}
		endTime, err := time.Parse(time.RFC3339, endTime)
		if err != nil {
			log.Fatal(err)
		}
		startTimeInUTC := startTime.UTC()
		endTimeInUTC := endTime.UTC()

		// iterate through every metric data and
		// add those value into 'result' map
		for _, metric := range metrics {
			if metric.Timestamp.Before(startTimeInUTC) || metric.Timestamp.Equal(endTimeInUTC) || metric.Timestamp.After(endTimeInUTC) {
				continue
			}
			result[metric.LevelName] += metric.Value
		}
	}

	// write the file content which contains our result into a file
	if err := fileHandler.WriteMetricResultToFile(outputFileName+"."+outputFileType, result); err != nil {
		log.Fatal(err)
	}
}

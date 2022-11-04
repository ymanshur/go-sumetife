package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sumetife/adapter"
	"sumetife/metric"
	"time"
)

func main() {
	// define arguments variables
	var inputDirPath string
	var inputFileType string
	var inputStartTime string
	var inputEndTime string
	var outputFileType string
	var outputFileName string

	flag.StringVar(
		&inputDirPath,
		"d",
		"",
		"The directory path, the directory contains single type of file, it can be csv or json",
	)
	flag.StringVar(
		&inputDirPath,
		"directory",
		"",
		"The directory path, the directory contains single type of file, it can be csv or json",
	)
	flag.StringVar(
		&inputFileType,
		"t",
		"",
		"The type of the input files, supported format: json and csv",
	)
	flag.StringVar(
		&inputFileType,
		"type",
		"",
		"The type of the input files, supported format: json and csv",
	)
	flag.StringVar(
		&inputStartTime,
		"startTime",
		"",
		"The starting time to scan the data in the format of rfc3339, inclusive",
	)
	flag.StringVar(
		&inputEndTime,
		"endTime",
		"",
		"The ending time to scan the data in the format of rfc3339, exclusive",
	)
	flag.StringVar(
		&outputFileType,
		"outputFileType",
		"json",
		"The output type of the summary, supported value: json and yaml",
	)
	flag.StringVar(
		&outputFileName,
		"outputFileName",
		"out",
		"The output filename of summary",
	)
	flag.Parse()

	// validate arguments
	if inputDirPath == "" {
		panic("argument error: '-d' or '--directory' flag-value is required")
	}
	if inputFileType == "" {
		panic("argument error: '-t' or '--type' flag-value is required")
	}
	if inputFileType != "json" && inputFileType != "csv" {
		panic("argument error: only the json or csv file types were accepted as input files")
	}
	if inputStartTime == "" {
		panic("argument error: '--startTime' flag-value is required")
	}
	if inputEndTime == "" {
		panic("argument error: '--endTime' flag-value is required")
	}

	// define time range of metrics data
	startTime, err := time.Parse(time.RFC3339, inputStartTime)
	if err != nil {
		panic(err)
	}
	endTime, err := time.Parse(time.RFC3339, inputEndTime)
	if err != nil {
		panic(err)
	}

	// read the data directory
	dirEntries, err := os.ReadDir(inputDirPath)
	if err != nil {
		panic(err)
	}

	// define file decoder
	fileDecoder := adapter.JSONFileDecoder
	if inputFileType == "csv" {
		fileDecoder = adapter.CSVFileDecoder
	}

	// define file encoder
	fileEncoder := adapter.JSONFileEncoder
	if outputFileType == "yaml" {
		fileEncoder = adapter.YAMLFileEncoder
	}

	// initialize handler
	handler := metric.NewMetricHandler(fileDecoder, fileEncoder)

	// define final result
	result := map[string]int{}
	// iterate through every metric file and summarize those value of each level name
	for _, entry := range dirEntries {
		// validate the entry, only pass the file
		if entry.IsDir() {
			continue
		}
		// validate related file type
		fileExt := filepath.Ext(entry.Name())
		if fileExt != ("." + inputFileType) {
			continue
		}

		// get metrics data
		fileName := filepath.Join(inputDirPath, entry.Name())
		metrics, err := handler.GetMetricsDataFromFile(fileName)
		if err != nil {
			log.Println(err)
			continue
		}

		// iterate through every metric data and
		// sum those value into 'result' map
		for _, metric := range metrics {
			// restrict metric data which timestamp is not in the range time
			if !metric.IsInRange(startTime, endTime) {
				continue
			}
			result[metric.LevelName] += metric.Value
		}
	}

	// formatting the result map
	metricResult := metric.MetricResultFormatter(result)

	// write the file content which contains our result into a file
	if err := handler.WriteMetricResultToFile(outputFileName+"."+outputFileType, metricResult); err != nil {
		panic(err)
	}
}

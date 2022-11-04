package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"sumetife/adapter"
	"sumetife/metric"
	"sumetife/util"

	"gopkg.in/yaml.v3"
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
		log.Fatal("argument error: '-d' or '--directory' flag-value is required")
	}
	if inputFileType == "" {
		log.Fatal("argument error: '-t' or '--type' flag-value is required")
	}
	if inputFileType != "json" && inputFileType != "csv" {
		log.Fatal("argument error: only the json or csv file types were accepted as input files")
	}
	if inputStartTime == "" {
		log.Fatal("argument error: '--startTime' flag-value is required")
	}
	if inputEndTime == "" {
		log.Fatal("argument error: '--endTime' flag-value is required")
	}

	// read the data directory
	dirEntries, err := os.ReadDir(inputDirPath)
	if err != nil {
		log.Fatal(err)
	}

	// define file decoder
	fileDecoder := adapter.JSONFileDecoder
	if inputFileType == "csv" {
		fileDecoder = adapter.CSVFileDecoder
	}

	// define encoder
	encoder := json.Marshal
	if outputFileType == "yaml" {
		encoder = yaml.Marshal
	}

	// initialize handler
	handler := metric.NewMetricHandler(fileDecoder, encoder)

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
			log.Fatal(err)
		}

		startTime, err := util.ParseToTimeUTC(inputStartTime)
		if err != nil {
			log.Fatal(err)
		}
		endTime, err := util.ParseToTimeUTC(inputEndTime)
		if err != nil {
			log.Fatal(err)
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
		log.Fatal(err)
	}
}

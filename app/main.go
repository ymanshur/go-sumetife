package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// MetricFile struct which contains
// data metric of a game level
type MetricFile struct {
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

func main() {
	// Open our file
	// file, err := os.Open("data/01-jan.json")
	file, err := os.Open("data/01-jan.csv")
	// if os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Successfully opened 01-jan.json")
	fmt.Println("Successfully opened 01-jan.csv")
	// defer the closing of our file so that we can parse it later on
	defer file.Close()

	// read our opened file as a byte array
	// byteValue, _ := ioutil.ReadAll(jsonFile)
	// read our opened file as a string matrix (array or stringarray)
	csvReader := csv.NewReader(file)
	csvLines, _ := csvReader.ReadAll()

	// for _, line := range csvLines[1:] {
	// 	fmt.Printf(
	// 		"Level: %s - %s (%T) (%s)\n",
	// 		line[1],
	// 		line[2], line[2],
	// 		line[0],
	// 	)
	// }

	// initialize our MetricFiles array
	var metricFiles []MetricFile

	// unmarshal (decode) our byteArray which contains our
	// json file content into 'metricFiles' which we defined above
	// json.Unmarshal(byteValue, &metricFiles)

	// map our csvLines which contains our
	// csv file content into 'metricFiles' which we defined above
	for _, line := range csvLines[1:] {
		matricValue, _ := strconv.Atoi(line[2])
		matricTimestamp, _ := time.Parse(time.RFC3339, line[0])
		metricFile := MetricFile{
			LevelName: line[1],
			Value:     matricValue,
			Timestamp: matricTimestamp,
		}
		metricFiles = append(metricFiles, metricFile)
	}

	// iterate through every metric file and
	// summarize those value for each level name at the data
	result := map[string]int{}
	for i := 0; i < len(metricFiles); i++ {
		result[metricFiles[i].LevelName] += metricFiles[i].Value
	}

	// marshal (encode) the result with formatter function
	// fileContent, err := json.Marshal(MetricResultFormatter(result))
	fileContent, err := yaml.Marshal(MetricResultFormatter(result))
	if err != nil {
		log.Fatal(err)
	}

	// white the file content which contains our result into a json file
	if err := ioutil.WriteFile("output.yaml", fileContent, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully generate output.yaml")
}

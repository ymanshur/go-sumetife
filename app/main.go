package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
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
	LevelName  string `json:"level_name"`
	TotalValue int    `json:"total_value"`
}

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
	// Open our jsonFile
	jsonFile, err := os.Open("data/01-jan.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully opened 01-jan.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our MetricFiles array
	var metricFiles []MetricFile
	// state the result using interfaces{} to take less time for instance mapping instead using struct
	// var result []map[string]interface{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'metricFiles' which we defined above
	json.Unmarshal(byteValue, &metricFiles)
	// we unmarshal jsonFile's into 'result', as well as handle unstructured data
	// json.Unmarshal(byteValue, &result)

	// we iterate through every metric file and
	// print out the metric value, their level name, and their timestamp
	// as just an example
	for i := 0; i < len(metricFiles); i++ {
		fmt.Printf(
			"Level: %s - %d (%s)\n",
			metricFiles[i].LevelName, metricFiles[i].Value, metricFiles[i].Timestamp,
		)
	}
	// we iterate result array, consider default json number type is float64
	// for i := 0; i < len(result); i++ {
	// 	fmt.Printf(
	// 		"Level: %s - %f (%T) (%s)\n",
	// 		result[i]["level_name"],
	// 		result[i]["value"], result[i]["value"],
	// 		result[i]["timestamp"],
	// 	)
	// }

	// we iterate through every metric file and
	// summarize those value for each level name at data
	result := map[string]int{}
	for i := 0; i < len(metricFiles); i++ {
		result[metricFiles[i].LevelName] += metricFiles[i].Value
	}

	fileContent, err := json.Marshal(MetricResultFormatter(result))
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("output.json", fileContent, 0644); err != nil {
		log.Fatal(err)
	}
}

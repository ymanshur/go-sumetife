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
	// if os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully opened 01-jan.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize our MetricFiles array
	var metricFiles []MetricFile

	// unmarshal (decode) our byteArray which contains our
	// jsonFile's content into 'metricFiles' which we defined above
	json.Unmarshal(byteValue, &metricFiles)

	// iterate through every metric file and
	// summarize those value for each level name at the data
	result := map[string]int{}
	for i := 0; i < len(metricFiles); i++ {
		result[metricFiles[i].LevelName] += metricFiles[i].Value
	}

	// marshal (encode) the result with formatter function
	fileContent, err := json.Marshal(MetricResultFormatter(result))
	if err != nil {
		log.Fatal(err)
	}

	// white the file content which contains our result into a json file
	if err := ioutil.WriteFile("output.json", fileContent, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully generate output.json")
}

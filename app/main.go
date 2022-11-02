package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// MetricFile struct which contains
// a data metrics of a level
type MetricFile struct {
	LevelName string    `json:"level_name"`
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("data/01-jan.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully opened 01-jan.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our MetricFiles array
	// var metricFiles []MetricFile
	// state the result using interfaces{} to take less time for instance mapping instead using struct
	var result []map[string]interface{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'metricFiles' which we defined above
	// json.Unmarshal(byteValue, &metricFiles)
	// we unmarshal jsonFile's into 'result', as well as handle unstructured data
	json.Unmarshal(byteValue, &result)

	// we iterate through every metric file within our metricFiles array and
	// print out the metric value, their level name, and their timestamp
	// as just an example
	// for i := 0; i < len(metricFiles); i++ {
	// 	fmt.Printf(
	// 		"Level: %s - %d (%s)\n",
	// 		metricFiles[i].LevelName, metricFiles[i].Value, metricFiles[i].Timestamp,
	// 	)
	// }
	for i := 0; i < len(result); i++ {
		fmt.Printf(
			"Level: %s - %f (%s)\n",
			result[i]["level_name"], result[i]["value"], result[i]["timestamp"],
		)
	}
}

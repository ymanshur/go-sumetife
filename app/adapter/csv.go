package adapter

import (
	"encoding/csv"
	"os"
	"strconv"
	"sumetife/metric"
	"time"
)

// CSVFileDecoder updates defined metrics data from a csv file
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

package metric

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mockFileName = "01-jan.json"
var mockResult = map[string]int{}

func MockMetricFileDecoder(file *os.File, metrics *[]Metric) error {
	metric := Metric{
		LevelName: "level1",
		Value:     126,
		Timestamp: time.Date(2022, time.January, 1, 0, 23, 0, 0, time.UTC),
	}
	*metrics = append(*metrics, metric)
	return nil
}

func MockMetricEncoder(v any) ([]byte, error) {
	buf := []byte(nil)
	return buf, nil
}

func MockOpenFile(name string) (*os.File, error) {
	file := new(os.File)
	return file, nil
}

func MockWriteFile(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func TestGetMetricsDataFromFileSuccess(t *testing.T) {
	// Setup
	OpenFile = MockOpenFile
	defer func() { OpenFile = os.Open }()

	fileDecoder := MockMetricFileDecoder
	encoder := MockMetricEncoder
	fileHandler := NewMetricHandler(fileDecoder, encoder)

	metrics, err := fileHandler.GetMetricsDataFromFile(mockFileName)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(metrics))
	}
}

func TestGetMetricsDataFromFileErrorOpenFile(t *testing.T) {
	// Setup
	OpenFile = func(name string) (*os.File, error) {
		return nil, fmt.Errorf("open %s: no such file or directory", mockFileName)
	}
	defer func() { OpenFile = os.Open }()

	fileDecoder := MockMetricFileDecoder
	encoder := MockMetricEncoder
	fileHandler := NewMetricHandler(fileDecoder, encoder)

	_, err := fileHandler.GetMetricsDataFromFile(mockFileName)

	// Assertions
	assert.Error(t, err)
}

func TestGetMetricsDataFromFileErrorFileDecoder(t *testing.T) {
	// Setup
	OpenFile = MockOpenFile
	defer func() { OpenFile = os.Open }()

	fileDecoder := func(file *os.File, metrics *[]Metric) error {
		return csv.ErrTrailingComma
	}
	encoder := MockMetricEncoder
	fileHandler := NewMetricHandler(fileDecoder, encoder)

	_, err := fileHandler.GetMetricsDataFromFile(mockFileName)

	// Assertions
	assert.Error(t, err)
}

func TestWriteMetricResultToFileSuccess(t *testing.T) {
	// Setup
	WriteFile = MockWriteFile
	defer func() { WriteFile = os.WriteFile }()

	fileDecoder := MockMetricFileDecoder
	encoder := MockMetricEncoder
	fileHandler := NewMetricHandler(fileDecoder, encoder)

	// Assertions
	assert.NoError(t, fileHandler.WriteMetricResultToFile(mockFileName, mockResult))
}

func TestWriteMetricResultToFileErrorEncoder(t *testing.T) {
	// Setup
	WriteFile = MockWriteFile
	defer func() { WriteFile = os.WriteFile }()

	fileDecoder := MockMetricFileDecoder
	encoder := func(v any) ([]byte, error) {
		buf := []byte(nil)
		return buf, fmt.Errorf("json: unsupported type: func()")
	}
	fileHandler := NewMetricHandler(fileDecoder, encoder)

	// Assertions
	assert.Error(t, fileHandler.WriteMetricResultToFile(mockFileName, mockResult))
}

func TestWriteMetricResultToFileErrorWriteFile(t *testing.T) {
	// Setup
	WriteFile = func(name string, data []byte, perm fs.FileMode) error {
		return os.ErrInvalid
	}
	defer func() { WriteFile = os.WriteFile }()

	fileDecoder := MockMetricFileDecoder
	encoder := MockMetricEncoder
	fileHandler := NewMetricHandler(fileDecoder, encoder)

	// Assertions
	assert.Error(t, fileHandler.WriteMetricResultToFile(mockFileName, mockResult))
}

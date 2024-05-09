package WinRM

import (
	"encoding/json"
	"fmt"
	"os"
)

type Metric struct {
	MetricName string `json:"metric.name"`
	Metric     string `json:"metric"`
}

func MetricExtract() {
	// Open the JSON file
	file, err := os.Open("/home/raj/Work/Go/WinRM/windows.jsonq")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode JSON
	var data map[string][]Metric
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Extract metric names
	var metrics []Metric
	for _, window := range data["windows"] {
		metrics = append(metrics, Metric{
			MetricName: window.MetricName,
			Metric:     window.Metric,
		})
	}

	// Create a new file
	outputFile, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	// Encode extracted data into JSON and write to file
	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(metrics); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Extraction successful. Output written to output.json")
}

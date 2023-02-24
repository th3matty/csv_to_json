package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the CSV file
	file, err := os.Open("results.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the header row
	header := records[0]

	// Loop through the records and print out each one as a separate JSON object
	for _, record := range records[1:] {
		m := make(map[string]interface{})
		for i, value := range record {
			// Remove single quotes from the value
			value = strings.ReplaceAll(value, "'", "")
			// Check if value is a JSON object
			var jsonMap map[string]interface{}
			if json.Unmarshal([]byte(value), &jsonMap) == nil {
				m[header[i]] = jsonMap
			} else {
				m[header[i]] = value
			}
		}

		// Convert the map to a JSON object
		jsonData, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the JSON object
		fmt.Println(string(jsonData))
	}
}

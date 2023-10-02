package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <csv-file>")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	jsonData, err := json.Marshal(records)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
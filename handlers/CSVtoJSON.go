package handlers

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
)

func CSVtoJSONHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "csv/data.csv" // Path to your CSV file in the sub-folder

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Error opening file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Error reading CSV file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(records)
	if err != nil {
		http.Error(w, "Error converting to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

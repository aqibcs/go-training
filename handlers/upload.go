package handlers

import (
	"encoding/json"
	"encoding/csv"
	"net/http"
	"io"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Max size: 10 MB
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("file") // "file" is the key for the uploaded file
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Parse the CSV file
	csvReader := csv.NewReader(file)
	var csvData [][]string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, "Error reading the CSV file", http.StatusInternalServerError)
			return
		}
		csvData = append(csvData, record)
	}

	var jsonData []map[string]string
	headers := csvData[0]
	for _, row := range csvData[1:] {
		data := make(map[string]string)
		for i, col := range row {
			data[headers[i]] = col
		}
		jsonData = append(jsonData, data)
	}

	// Convert CSV data to JSON
	responseJSON, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

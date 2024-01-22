package handlers

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UploadFileHandler handles HTTP requests for uploading a CSV file, parsing its content, and converting it to JSON.
func UploadFileHandler(c echo.Context) error {
	file, err := c.FormFile("file") // "file" is the key for the uploaded file
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error retrieving the file")
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error opening the file")
	}
	defer src.Close()

	// Parse the CSV file
	csvReader := csv.NewReader(src)
	var csvData [][]string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error reading the CSV file")
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
		return c.String(http.StatusInternalServerError, "Error creating JSON response")
	}

	// Set Content-Type header to indicate JSON response
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)
	c.Response().Write(responseJSON)

	return nil
}

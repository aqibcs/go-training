package handlers

import (
	"encoding/csv"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

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
	return c.JSON(http.StatusOK, responseJSON)
}

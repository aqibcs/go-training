package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <csv-file-path>")
		return
	}

	filePath := os.Args[1]

	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read CSV records
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	var jsonData []map[string]string
	headers := records[0] // Extract headers from the first row
	for _, row := range records[1:] {
		data := make(map[string]string)
		for i, col := range row {
			data[headers[i]] = col
		}
		jsonData = append(jsonData, data)
	}

	// Convert data to JSON format
	jsonOutput, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonOutput))

	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", HelloHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, " + name + "!"))
}

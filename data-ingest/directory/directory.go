package directory

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateDirectoryForDate(target string, date string) {
	// Creates the directory to store the page of results. The directory structure looks
	// like {year}/{month}/{day}. Each result is stored as an individual JSON file.
	dateComponents := strings.Split(date, "-")
	year, month, day := dateComponents[0], dateComponents[1], dateComponents[2]

	rootPath := filepath.Join(target, "register-files")
	createIfNotExists(rootPath)

	yearPath := filepath.Join(rootPath, year)
	createIfNotExists(yearPath)

	monthPath := filepath.Join(yearPath, month)
	createIfNotExists(monthPath)

	dayPath := filepath.Join(monthPath, day)
	createIfNotExists(dayPath)
}

func createIfNotExists(path string) {
	// Creates the target filepath if it does not already exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func WriteJSON(data interface{}, filename string) {
	// Writes a struct to a JSON file
	file, err := json.MarshalIndent(&data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(filename, file, 0644)
}

func ReadJSON(data interface{}, filename string) {
	// Reads a JSON file into a struct
	file, readErr := ioutil.ReadFile(filename)
	if readErr != nil {
		log.Fatal(readErr)
	}

	marshalErr := json.Unmarshal([]byte(file), &data)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
}

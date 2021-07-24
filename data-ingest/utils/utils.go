package utils

import (
  "log"
	"os"
  "path/filepath"
  "strings"
  "strconv"
)

func createDirectory(target string, date string, page int) {
	// Creates the directory to store the page of results. The directory structure looks
	// like {year}/{month}/{day}/{page}. Each result is stored as an individual JSON file.
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

	pagePath := filepath.Join(dayPath, strconv.Itoa(page))
	createIfNotExists(pagePath)
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

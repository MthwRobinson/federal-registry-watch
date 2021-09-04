package main

import (
	"fmt"
	"github.com/MthwRobinson/federal-registry-watch/data-ingest/directory"
	"github.com/MthwRobinson/federal-registry-watch/data-ingest/fed_registry"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	target := "/home/matt/tmp/federal-regulations"
	start, _ := time.Parse("2006-01-02", "2021-08-01")
	now := time.Now()

	client := &http.Client{}
	registerFetcher := fed_registry.RegisterFetcher{Client: client}

	for date := start; date.Before(now); date = date.AddDate(0, 0, 1) {
		dateString := date.Format("2006-01-02")
		fmt.Printf("Fetching regulations for date: %s\n", dateString)

		dateComponents := strings.Split(dateString, "-")
		year, month, day := dateComponents[0], dateComponents[1], dateComponents[2]
		dirName := filepath.Join(target, "register-files", year, month, day)

		directory.CreateDirectoryForDate(target, dateString)
		results := registerFetcher.GetDailyRegisterResults(dateString)
		for _, result := range results {
			filename := result.DocumentNumber + ".json"
			directory.WriteJSON(result, filepath.Join(dirName, filename))
		}

	}

}

package main

import (
	"encoding/json"
	"github.com/MthwRobinson/federal-registry-watch/data-ingest/directory"
	"github.com/MthwRobinson/federal-registry-watch/data-ingest/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const registerBaseURL = "https://www.federalregister.gov/api/v1/documents"

func buildRegisterURL(date string, page int) string {
	// Builds and encodes the url for fetching register results
	params := url.Values{}
	params.Add("conditions[publication_date][is]", date)
	params.Add("page", strconv.Itoa(page))
	params.Add("format", "json")
	registerURL := registerBaseURL + "?" + params.Encode()
	return registerURL
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type registerFetcher struct {
	client HttpClient
}

func (r *registerFetcher) getRegisterResults(date string, page int) models.RegisterResults {
	// Collects a list of of document and links from the Federal Register for the
	// specified date and page number
	registerURL := buildRegisterURL(date, page)

	req, err := http.NewRequest("GET", registerURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resultJSON, _ := ioutil.ReadAll(resp.Body)
	var registerResults models.RegisterResults
	json.Unmarshal([]byte(resultJSON), &registerResults)

	return registerResults
}

func (r *registerFetcher) getDailyRegisterResults(date string) []models.Result {
	// Uses pagination to fetch all of the registry results for the specified day
	var totalPages int
	var results []models.Result
	for currentPage := 1; ; currentPage++ {
		registerResults := r.getRegisterResults(date, currentPage)
		totalPages = registerResults.TotalPages
		for _, result := range registerResults.Results {
			results = append(results, result)
		}
		if currentPage >= totalPages {
			break
		}
	}
	return results
}

func main() {
	client := &http.Client{}
	r := registerFetcher{client: client}
	results := r.getDailyRegisterResults("2021-06-02")
  directory.WriteJSON(results, "/home/matt/tmp/test.json")

	// fmt.Println(registerResults.NextPageURL)
	// files, _ := ioutil.ReadDir("../../../")
	// for _, file := range files {
	// 	if file.IsDir() {
	// 		fmt.Println(file.Name())
	// 	}
	// }
	// createDirectory("/home/matt/tmp", "2021-01-01", 1)
}

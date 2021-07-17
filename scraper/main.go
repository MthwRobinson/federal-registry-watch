package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

type Result struct {
	Title                  string `json:"title"`
	Type                   string `json:"type"`
	Abstract               string `json:"abstract"`
	DocumentNumber         string `json:"document_number"`
	HTMLURL                string `json:"html_url"`
	PdfURL                 string `json:"pdf_url"`
	PublicInspectionPdfURL string `json:"public_inspection_pdf_url"`
	PublicationDate        string `json:"publication_date"`
	Agencies               []struct {
		RawName  string      `json:"raw_name"`
		Name     string      `json:"name"`
		ID       int         `json:"id"`
		URL      string      `json:"url"`
		JSONURL  string      `json:"json_url"`
		ParentID interface{} `json:"parent_id"`
		Slug     string      `json:"slug"`
	} `json:"agencies"`
	Excerpts string `json:"excerpts"`
}

type RegisterResults struct {
	Count       int      `json:"count"`
	Description string   `json:"description"`
	TotalPages  int      `json:"total_pages"`
	NextPageURL string   `json:"next_page_url"`
	Results     []Result `json:"results"`
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type regulationFetcher struct {
	client HttpClient
}

func (r *regulationFetcher) getRegulations(date string, page int) RegisterResults {
	// Collects a list of of document and links from the Federal Register for the
	// specified date and page number
	registerURL := buildRegisterURL(date, page)

	req, err := http.NewRequest("GET", registerURL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	resultJSON, _ := ioutil.ReadAll(resp.Body)
	var registerResults RegisterResults
	json.Unmarshal([]byte(resultJSON), &registerResults)

	return registerResults
}

func createDirectory(target string, date string, page int) {
	// Creates the directory to store the page of results. The directory structure looks
	// like {year}/{month}/{day}/{page}. Each result is stored as an individual JSON file.
	dateComponents := strings.Split(date, "-")
	year, _, _ := dateComponents[0], dateComponents[1], dateComponents[2]
	yearPath := filepath.Join(target, year)
	createIfNotExists(yearPath)
}

func createIfNotExists(path string) {
	// Creates the target filepath if it does not already exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Creating directory: ", path)
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

}

func main() {
	// client := &http.Client{}
	// r := regulationFetcher{client: client}
	// registerResults := r.getRegulations("2021-06-02", 2)
	// fmt.Println(registerResults.NextPageURL)
	// files, _ := ioutil.ReadDir("../../../")
	// for _, file := range files {
	// 	if file.IsDir() {
	// 		fmt.Println(file.Name())
	// 	}
	// }
	// dateComponents := strings.Split("2021-01-02", "-")
	// year, month, day := dateComponents[0], dateComponents[1], dateComponents[2]
	// date := fmt.Sprintf("%s-%s-%s", year, month, day)
	// fmt.Println(date)
	createDirectory("/home/matt/tmp", "2021-01-01", 1)
}

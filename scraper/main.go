package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	// "os"
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

func main() {
	// client := &http.Client{}
	// r := regulationFetcher{client: client}
	// registerResults := r.getRegulations("2021-06-02", 2)
	// fmt.Println(registerResults.NextPageURL)
	files, _ := ioutil.ReadDir("../../../")
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
		}
	}
}

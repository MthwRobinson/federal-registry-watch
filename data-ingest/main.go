package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	// "github.com/MthwRobinson/federal-registry-watch/data-ingest/utils"
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

type registerFetcher struct {
	client HttpClient
}

func (r *registerFetcher) getRegisterResults(date string, page int) RegisterResults {
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
	var registerResults RegisterResults
	json.Unmarshal([]byte(resultJSON), &registerResults)

	return registerResults
}

func (r *registerFetcher) getDailyRegisterResults(date string) []Result {
	// Uses pagination to fetch all of the registry results for the specified day
	var totalPages int
	var results []Result
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
	r.getDailyRegisterResults("2021-06-02")
	// fmt.Println(registerResults)

	// fmt.Println(registerResults.NextPageURL)
	// files, _ := ioutil.ReadDir("../../../")
	// for _, file := range files {
	// 	if file.IsDir() {
	// 		fmt.Println(file.Name())
	// 	}
	// }
	// createDirectory("/home/matt/tmp", "2021-01-01", 1)
}

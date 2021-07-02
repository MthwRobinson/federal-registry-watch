package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	// "os"
)

const baseURL string = "https://www.federalregister.gov/documents/search?conditions[publication_date[is]=%s&page=%s#"

type PageFetcher func(url string) (*http.Response, error)

func FetchUrl(url string, pageFetcher PageFetcher) (*http.Response, error) {
	// Makes an http call to the specified url and returns a string representation of the
	// HTML response.
	resp, err := pageFetcher(url)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

type RegisterResults struct {
	Count       int    `json:"count"`
	Description string `json:"description"`
	TotalPages  int    `json:"total_pages"`
	NextPageURL string `json:"next_page_url"`
	Results     []struct {
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
	} `json:"results"`
}

func getRegulations(date string, page int) RegisterResults {
	// Collects a list of of document and links from the Federal Register for the
	// specified date and page number
	url := "https://www.federalregister.gov/api/v1/documents?conditions%5Bpublication_date%5D%5Bis%5D=2021-06-29&format=json&page=2"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	resultJSON, _ := ioutil.ReadAll(resp.Body)

	var registerResults RegisterResults
	json.Unmarshal([]byte(resultJSON), &registerResults)
	fmt.Println(registerResults.Results[0].Agencies[0].Slug)

	return registerResults

}

func FindLinks(response *http.Response) {
	// Iterates through an HTML object and finds <a> tages with href attributes
	tokenizer := html.NewTokenizer(response.Body)

	for {
		tag := tokenizer.Next()
		switch {
		case tag == html.ErrorToken:
			return
		case tag == html.StartTagToken:
			token := tokenizer.Token()

			isAnchor := token.Data == "a"
			if isAnchor {
				fmt.Println(token)
			}
		}
	}
}

func main() {
	// url := os.Args[1]
	getRegulations("2021-06-29", 1)
	// response, _ := FetchUrl(url, http.Get)
	// FindLinks(response)
}

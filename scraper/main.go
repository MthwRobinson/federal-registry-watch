package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	//"encoding/json"
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

func getRegulations(date string, page int) []string {
	// Collects a list of of document and links from the Federal Register for the
	// specified date and page number
	var urlList []string
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
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))

	return urlList

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
	getLinks("2021-06-29", 1)
	// response, _ := FetchUrl(url, http.Get)
	// FindLinks(response)
}

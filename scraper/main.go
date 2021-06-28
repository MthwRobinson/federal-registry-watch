package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

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
	url := os.Args[1]
	response, _ := FetchUrl(url, http.Get)
	FindLinks(response)
}

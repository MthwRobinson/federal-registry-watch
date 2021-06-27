package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type PageFetcher func(url string) (*http.Response, error)

func FetchUrl(url string, pageFetcher PageFetcher) (string, error) {
	// Makes an http call to the specified url and returns a string representation of the
	// HTML response.
	resp, err := pageFetcher(url)
	if err != nil {
		return "", err
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes), nil
}

func main() {
	url := os.Args[1]
	html, _ := FetchUrl(url, http.Get)
	fmt.Println(html)
}

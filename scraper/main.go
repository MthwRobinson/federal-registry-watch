package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type PageFetcher func(url string) (*http.Response, error)

func FetchUrl(url string, pageFetcher PageFetcher) string {
	// Makes an http call to the specified url and returns a string representation of the
	// HTML response.
	resp, _ := pageFetcher(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}

func main() {
	var url, html string
	url = os.Args[1]
	html = FetchUrl(url, http.Get)
	fmt.Println(html)
}

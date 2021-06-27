package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func mockPageFetcher(url string) (*http.Response, error) {
	response := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("<h1>Hello World!</h1>")),
	}
	return &response, nil
}

func TestFetchUrl(t *testing.T) {
	assert := assert.New(t)
	url := "http://fake.website"
	html := FetchUrl(url, mockPageFetcher)
	assert.Equal(html, "<h1>Hello World!</h1>", "These should be equal")
}

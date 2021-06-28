package main

import (
	"bytes"
	"errors"
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
	response, _ := FetchUrl(url, mockPageFetcher)
	bytes, _ := ioutil.ReadAll(response.Body)
	assert.Equal(string(bytes), "<h1>Hello World!</h1>")
}

func mockPageFetcherWithError(url string) (*http.Response, error) {
	response := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("Error")),
	}
	return &response, errors.New("404: Page Not Found")
}

func TestFetchUrlHandlesError(t *testing.T) {
	assert := assert.New(t)
	url := "http://fake.website"
	_, err := FetchUrl(url, mockPageFetcherWithError)
	if assert.Error(err) {
		assert.Equal(err, errors.New("404: Page Not Found"))
	}
}

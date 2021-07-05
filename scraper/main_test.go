package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildRegisterURL(t *testing.T) {
	registerURL := buildRegisterURL("2021-06-02", 4)
	expectedURL := "https://www.federalregister.gov/api/v1/documents?conditions%5Bpublication_date%5D%5Bis%5D=2021-06-02&format=json&page=4"
	assert.Equal(t, registerURL, expectedURL)
}

//
// func mockPageFetcher(url string) (*http.Response, error) {
// 	response := http.Response{
// 		Body: ioutil.NopCloser(bytes.NewBufferString("<h1>Hello World!</h1>")),
// 	}
// 	return &response, nil
// }
//
// func TestFetchUrl(t *testing.T) {
// 	assert := assert.New(t)
// 	url := "http://fake.website"
// 	response, _ := FetchUrl(url, mockPageFetcher)
// 	bytes, _ := ioutil.ReadAll(response.Body)
// 	assert.Equal(string(bytes), "<h1>Hello World!</h1>")
// }
//
// func mockPageFetcherWithError(url string) (*http.Response, error) {
// 	response := http.Response{
// 		Body: ioutil.NopCloser(bytes.NewBufferString("Error")),
// 	}
// 	return &response, errors.New("404: Page Not Found")
// }
//
// func TestFetchUrlHandlesError(t *testing.T) {
// 	assert := assert.New(t)
// 	url := "http://fake.website"
// 	_, err := FetchUrl(url, mockPageFetcherWithError)
// 	if assert.Error(err) {
// 		assert.Equal(err, errors.New("404: Page Not Found"))
// 	}
// }

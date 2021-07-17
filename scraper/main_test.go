package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestBuildRegisterURL(t *testing.T) {
	registerURL := buildRegisterURL("2021-06-02", 4)
	expectedURL := "https://www.federalregister.gov/api/v1/documents?conditions%5Bpublication_date%5D%5Bis%5D=2021-06-02&format=json&page=4"
	assert.Equal(t, registerURL, expectedURL)
}

type MockClient struct {
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	response := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(`{"count": 25}`)),
	}
	return &response, nil
}

func TestGetRegulations(t *testing.T) {
	client := &MockClient{}
	r := regulationFetcher{client: client}
	registerResults := r.getRegulations("2021-01-01", 5)
	assert.Equal(t, registerResults.Count, 25)
}

func TestCreateIfNotExists(t *testing.T) {
	dir, _ := ioutil.TempDir("", "gotest")
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "testdir")
	assert.NoDirExists(t, path)
	createIfNotExists(path)
	assert.DirExists(t, path)
}

func TestCreateDirectory(t *testing.T) {
	dir, _ := ioutil.TempDir("", "gotest")
	defer os.RemoveAll(dir)

	createDirectory(dir, "2021-02-03", 5)
	expectedPath := filepath.Join(dir, "federal-regulations", "2021", "02", "03", "5")
	assert.DirExists(t, expectedPath)
}

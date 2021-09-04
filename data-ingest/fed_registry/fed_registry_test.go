package fed_registry

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
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
		Body: ioutil.NopCloser(bytes.NewBufferString(`{
      "total_pages": 5,
      "results": [
        {"title": "Look at this new regulation!"},
        {"title": "Here's another one!"}
        ]
    }`)),
	}
	return &response, nil
}

func TestGetRegisterResults(t *testing.T) {
	client := &MockClient{}
	r := registerFetcher{client: client}
	registerResults := r.getRegisterResults("2021-01-01", 5)
	assert.Equal(t, registerResults.TotalPages, 5)
}

func TestGetDailyResults(t *testing.T) {
	client := &MockClient{}
	r := registerFetcher{client: client}
	registerResults := r.getDailyRegisterResults("2021-01-01")
	assert.Equal(t, len(registerResults), 10)
}

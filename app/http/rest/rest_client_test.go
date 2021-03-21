package rest

import (
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// do whatever you want
			return &http.Response{
				StatusCode: http.StatusBadRequest,
			}, nil
		},
	}

	request, _ := http.NewRequest("GET", "https://www.reallycoolurl.com/bad_request", nil)
	// as this is a test, we may skip error handling

	response, _ := client.Do(request)
	if response.StatusCode != http.StatusBadRequest {
		t.Error("invalid response status code")
	}
}

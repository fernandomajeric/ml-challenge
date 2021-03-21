package rest

import (
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HttpClient
)

func init() {
	Client = &http.Client{}
}

type MockClient struct {
	DoFunc func(r *http.Request) (*http.Response,error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	// just in case you want default correct return value
	return &http.Response{}, nil
}

//// Post sends a post request to the URL with the body
//func Get(url string) (*http.Response, error) {
//	request, err := http.NewRequest(http.MethodGet, url, nil)
//	if err != nil {
//		return nil, err
//	}
//	request.Header.Add("Accept", "application/json")
//	request.Header.Add("Content-Type", "application/json")
//
//	return Client.Do(request)
//}

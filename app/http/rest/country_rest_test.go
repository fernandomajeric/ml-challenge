package rest

import (
	"bytes"
	"encoding/json"
	"github.com/fernandomajeric/ml-challenge/app/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFindShouldGetCountryData(t *testing.T) {
	// Arrange
	countryData := model.Country{
		Name:       "Canada",
		Alpha2Code: "CA",
		Alpha3Code: "CAN",
		Latlng:     []float64{60.0, -95.0},
		Currencies: []model.Currency{{Name: "Canadian dollar", Code: "CAN", Symbol: "$"}},
		Languages:  []model.Language{{Name: "English"}},
	}
	parsed, _ := json.Marshal(countryData)
	client := NewCountryRestClient(&MockClient{
		DoFunc: func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(parsed)),
			}, nil
		}})

	//Action
	country, _ := client.Find("ARG")

	//Assert
	assert.NotNil(t, country)
	assert.EqualValues(t, countryData, country)
}
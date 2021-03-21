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

func TestFindShouldGetLocalizationData(t *testing.T) {
	// Arrange
	ip:="138.82.0.256"
	geoLocalizationData := model.GeoLocalization{
		CountryCode:     "CA",
		CountryCode3:    "CAN",
		CountryCodeName: "Canada",
		CountryEmoji:    "ca",
	}
	parsed, _ := json.Marshal(geoLocalizationData)
	client := NewGeoLocalizationRestClient(&MockClient{
		DoFunc: func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(parsed)),
			}, nil
		}})

	//Action
	geolocalization, _ := client.Find(ip)

	//Assert
	assert.NotNil(t, geolocalization)
	assert.EqualValues(t, geoLocalizationData, geolocalization)
}
package rest

import (
	"bytes"
	"encoding/json"
	"github.com/fernandomajeric/ml-challenge/app/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestFindShouldGetData(t *testing.T) {
	// Arrange
	currencyData := model.CurrencyRate{
		Base:  "USD",
		Date:  time.Now().String(),
		Rates: map[string]float64{},
	}
	parsed, _ := json.Marshal(currencyData)
	client := NewCurrencyRateRestClient(&MockClient{
		DoFunc: func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(parsed)),
			}, nil
		}})

	//Action
	currencyRate, _ := client.Find("USD")

	//Assert
	assert.NotNil(t, currencyRate)
	assert.EqualValues(t, currencyData, currencyRate)
}

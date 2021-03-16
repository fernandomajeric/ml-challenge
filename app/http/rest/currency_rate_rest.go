package rest

import (
	"encoding/json"
	"fmt"
	"github.com/fernandomajeric/ml-challenge/app/model"
	"github.com/fernandomajeric/ml-challenge/config"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type CurrencyRateRestClientInterface interface {
	Find(code string) model.CurrencyRate
}

type CurrencyRateRestClient struct{}

func (CurrencyRateRestClient) Find(code string) model.CurrencyRate {
	client := &http.Client{}

	req, err := http.NewRequest("GET", config.Configuration.App.Rest.Currency.Url, nil)

	if err != nil {
		log.Error(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("access_key", config.Configuration.App.Rest.Currency.Token)
	q.Add("symbols", code)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject model.CurrencyRate
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("CurrencyCore Response as struct %+v\n", responseObject)

	return responseObject
}

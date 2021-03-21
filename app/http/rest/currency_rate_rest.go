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
	Find(code string) (model.CurrencyRate, error)
}

type CurrencyRateRestClient struct{
	Client HttpClient
}

func NewCurrencyRateRestClient(client HttpClient) *CurrencyRateRestClient {
	return &CurrencyRateRestClient{Client: client}
}

func (currencyRateRest *CurrencyRateRestClient) Find(code string) (model.CurrencyRate,error) {
	req, err := http.NewRequest(http.MethodGet, config.Configuration.App.Rest.Currency.Url, nil)

	if err != nil {
		log.Error(err.Error())
		return model.CurrencyRate{}, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("access_key", config.Configuration.App.Rest.Currency.Token)
	q.Add("symbols", code)
	req.URL.RawQuery = q.Encode()

	resp, err := currencyRateRest.Client.Do(req)

	if err != nil {
		log.Error(err.Error())
		return model.CurrencyRate{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error(err.Error())
		return model.CurrencyRate{}, err
	}
	var responseObject model.CurrencyRate
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("CurrencyCore Response as struct %+v\n", responseObject)

	return responseObject,nil
}

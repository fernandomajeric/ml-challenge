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

type CountryRestClientInterface interface {
	Find(code string) (model.Country, error)
}

type CountryRestClient struct{
	Client HttpClient
}

func NewCountryRestClient(client HttpClient) *CountryRestClient {
	return &CountryRestClient{Client: client}
}

//Find information by iso code
func (countryClient *CountryRestClient) Find(code string) (model.Country, error) {
	req, err := http.NewRequest("GET", config.Configuration.App.Rest.Country.Url+code, nil)

	if err != nil {
		log.Error(err.Error())
		return model.Country{}, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := countryClient.Client.Do(req)

	if err != nil {
		log.Error(err.Error())
		return model.Country{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return model.Country{}, err
	}
	var responseObject model.Country
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("Country Response as struct %+v\n", responseObject)

	return responseObject, nil
}

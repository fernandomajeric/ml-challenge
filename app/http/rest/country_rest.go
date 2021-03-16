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
	Find(code string) model.Country
}

type CountryRestClient struct{}

func (CountryRestClient) Find(code string) model.Country {
	client := &http.Client{}

	req, err := http.NewRequest("GET", config.Configuration.App.Rest.Country.Url+code, nil)

	if err != nil {
		log.Error(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject model.Country
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("Country Response as struct %+v\n", responseObject)

	return responseObject
}

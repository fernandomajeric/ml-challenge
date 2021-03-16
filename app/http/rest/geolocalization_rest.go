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

type GeoLocalizationRestClientInterface interface {
	Find(ip string) model.GeoLocalization
}

type GeoLocalizationRestClient struct{}

func (GeoLocalizationRestClient) Find(ip string) model.GeoLocalization {
	client := &http.Client{}

	req, err := http.NewRequest("GET", config.Configuration.App.Rest.Geolocalization.Url+"?"+ip, nil)

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
	var responseObject model.GeoLocalization
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("Geolocalization Response as struct %+v\n", responseObject)

	return responseObject

}

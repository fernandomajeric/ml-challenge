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
	Find(ip string) (model.GeoLocalization, error)
}

type GeoLocalizationRestClient struct{
	Client HttpClient
}

func NewGeoLocalizationRestClient(client HttpClient) *GeoLocalizationRestClient {
	return &GeoLocalizationRestClient{Client: client}
}

func (geolocalizationRest *GeoLocalizationRestClient) Find(ip string) (model.GeoLocalization,error) {
	req, err := http.NewRequest("GET", config.Configuration.App.Rest.Geolocalization.Url+"?"+ip, nil)

	if err != nil {
		log.Error(err.Error())
		return model.GeoLocalization{}, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := geolocalizationRest.Client.Do(req)

	if err != nil {
		log.Error(err.Error())
		return model.GeoLocalization{}, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf(err.Error())
		return model.GeoLocalization{}, err
	}

	var responseObject model.GeoLocalization
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("Geolocalization Response as struct %+v\n", responseObject)

	return responseObject,nil

}

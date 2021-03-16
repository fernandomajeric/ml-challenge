// Package service implements services with business logic
package service

import (
	"errors"
	"fmt"
	"github.com/fernandomajeric/ml-challenge/app/http/rest"
	"github.com/fernandomajeric/ml-challenge/app/model"
	"github.com/fernandomajeric/ml-challenge/app/utils"
	"strings"
	"time"
)

type TraceIpServiceInterface interface {
	GetTraceIp(ip string) (result model.TraceIpResult, err error)
}
type TraceIpService struct{}

// New : build new Service
func NewTraceIpService() TraceIpServiceInterface {
	return &TraceIpService{}
}

func (TraceIpService) GetTraceIp(ip string) (result model.TraceIpResult, err error) {
	if !utils.ValidateIpAddress(ip) {
		return model.TraceIpResult{}, errors.New("format incorrecto!")
	}

	geoLocalization := rest.GeoLocalizationRestClient{}.Find(ip)

	if !validateGeolocalization(geoLocalization) {
		return model.TraceIpResult{}, errors.New("Invalid Ip")
	}

	fmt.Printf("CountryCode: ", geoLocalization.CountryCode)
	country := rest.CountryRestClient{}.Find(geoLocalization.CountryCode)

	currencyCode := country.Currencies[0].Code
	currencyRate := rest.CurrencyRateRestClient{}.Find(currencyCode)
	currencyDescription := getCurrencyDescription(currencyCode, currencyRate)
	languageDescription := getLanguageDescription(country)
	distance := getDistance(country)

	result = model.TraceIpResult{
		Ip:       ip,
		Date:     time.Now(),
		Country:  geoLocalization.CountryCodeName,
		IsoCode:  country.Alpha3Code,
		Language: languageDescription,
		Currency: currencyDescription,
		Distance: distance,
	}

	return result, err
}

func getDistance(country model.Country) float64 {
	localCountry := rest.CountryRestClient{}.Find("AR")
	return utils.Calculate(localCountry.Latlng, country.Latlng)
}

func getLanguageDescription(country model.Country) string {
	var names []string
	for _, l := range country.Languages[:] {
		names = append(names, l.Name)
	}
	return strings.Join(names, ", ")
}

func getCurrencyDescription(code string, rate model.CurrencyRate) string {
	return fmt.Sprintf("%s (1 %s = %v %s)", code, rate.Base, rate.Rates[code], code)
}

func validateGeolocalization(localization model.GeoLocalization) bool {
	if localization.CountryCode == "" {
		return false
	}
	if (model.GeoLocalization{}) == localization {
		return false
	}

	return true
}

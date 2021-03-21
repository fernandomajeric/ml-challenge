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
type TraceIpService struct {
	GeoLocalizationRestClient rest.GeoLocalizationRestClientInterface
	CountryRestClient         rest.CountryRestClientInterface
	CurrencyRateRestClient    rest.CurrencyRateRestClientInterface
}

func NewTraceIpService(geoLocalizationRestClient rest.GeoLocalizationRestClientInterface,
	countryRestClient rest.CountryRestClientInterface,
	currencyRateRestClient rest.CurrencyRateRestClientInterface) *TraceIpService {
	return &TraceIpService{GeoLocalizationRestClient: geoLocalizationRestClient,
		CountryRestClient:      countryRestClient,
		CurrencyRateRestClient: currencyRateRestClient}
}

func (traceIp *TraceIpService) GetTraceIp(ip string) (result model.TraceIpResult, err error) {
	if !utils.ValidateIpAddress(ip) {
		return model.TraceIpResult{}, errors.New("invalid ip address!")
	}

	geoLocalization, err := traceIp.GeoLocalizationRestClient.Find(ip)

	if err != nil {
		return model.TraceIpResult{}, err
	}

	if !validateGeolocalization(geoLocalization) {
		return model.TraceIpResult{}, errors.New("invalid ip address!")
	}

	country, err := traceIp.CountryRestClient.Find(geoLocalization.CountryCode)

	if err != nil {
		return model.TraceIpResult{}, err
	}

	if len(country.Currencies) == 0 {
		return model.TraceIpResult{}, errors.New(fmt.Sprintf("Currencies for country %s",country.Name))
	}

	currencyCode := country.Currencies[0].Code
	currencyRate,err := traceIp.CurrencyRateRestClient.Find(currencyCode)

	if err != nil {
		return model.TraceIpResult{}, err
	}

	currencyDescription := getCurrencyDescription(currencyCode, currencyRate)
	languageDescription := getLanguageDescription(country)
	distance, err := traceIp.getDistance(country)

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

func (traceIp *TraceIpService) getDistance(country model.Country) (float64, error) {
	localCountry,err := traceIp.CountryRestClient.Find(utils.ArgCode)

	if err != nil {
		return 0, err
	}

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

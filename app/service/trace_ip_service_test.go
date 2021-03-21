package service

import (
	"fmt"
	"github.com/fernandomajeric/ml-challenge/app/model"
	"github.com/fernandomajeric/ml-challenge/app/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// Mock of Geolocalization
type geolocalizationRestClientMock struct {
	mock.Mock
}

// Mock of Country
type countryRestClientMock struct{
	mock.Mock
}

// Mock of Currency

type currencyRestClientMock struct{
	mock.Mock
}

func (m *countryRestClientMock) Find(code string) (model.Country, error) {
	fmt.Println("Mocked Country Find notification function")
	fmt.Printf("Value passed in: %s\n", code)
	args:= m.Called(code)
	return args.Get(0).(model.Country),nil
}

// Mocking GeolocalizationRestClient
func (m *geolocalizationRestClientMock) Find(ip string) (model.GeoLocalization, error) {
	fmt.Println("Mocked  Geolocalization find notification function")
	fmt.Printf("Value passed in: %s\n", ip)
	args:= m.Called(ip)
	return args.Get(0).(model.GeoLocalization),nil
}

func (m *currencyRestClientMock) Find(code string) (model.CurrencyRate,error) {
	fmt.Println("Mocked  Currency find notification function")
	fmt.Printf("Value passed in: %s\n", code)
	args:= m.Called(code)
	return args.Get(0).(model.CurrencyRate),nil
}

func TestGetTraceIpWithInvalidIpAddress(t *testing.T) {
	// Arrange
	ip := "138.82.0.256"
	service := TraceIpService{}
	// Action
	result, err := service.GetTraceIp(ip)
	// Assert
	if assert.Errorf(t, err, "Ip Address: %s should be valid", ip) {
		assert.EqualError(t, err, "invalid ip address!")
		assert.Empty(t, result)
	}
}

func TestGetTraceShouldFindGeoLocalization(t *testing.T) {
	// Arrange
	ip := "138.82.0.1"
	countryCode := "CA"
	currencyCode := "CAD"

	// create an instance of geolocalization
	geolocalizationRestMock := new(geolocalizationRestClientMock)
	geolocalizationRestMock.On("Find", ip).Return(model.GeoLocalization{
		CountryCode:     countryCode,
		CountryCode3:    "CAN",
		CountryCodeName: "Canada",
		CountryEmoji:    "ca",
	}, nil)

	// create an instance of country
	countryRestMock := new(countryRestClientMock)
	countryRestMock.On("Find", countryCode).Return(model.Country{
		Name:       "Canada",
		Alpha2Code: countryCode,
		Alpha3Code: "CAN",
		Latlng:     []float64{60.0,-95.0},
		Currencies: []model.Currency{{Name: "Canadian dollar", Code: currencyCode, Symbol: "$"}},
		Languages:  []model.Language{{Name: "English"}},
	}, nil)

	countryRestMock.On("Find", utils.ArgCode).Return(model.Country{
		Name:       "Argentina",
		Alpha2Code: utils.ArgCode,
		Alpha3Code: "ARG",
		Latlng:     []float64{34.0,-64.0},
		Currencies: []model.Currency{{Name: "Argentine peso", Code: "ARS", Symbol: "$"}},
		Languages:  []model.Language{{Name: "Spanish"}},
	}, nil)

	currencyRestMock := new(currencyRestClientMock)
	currencyRestMock.On("Find",currencyCode).Return(model.CurrencyRate{
		Base:  "USD",
		Date:  time.Now().String(),
		Rates: map[string]float64{},
	},nil)

	sut := NewTraceIpService(geolocalizationRestMock,countryRestMock,currencyRestMock)

	// Action
	_, _ = sut.GetTraceIp(ip)
	// Assert
	geolocalizationRestMock.AssertNumberOfCalls(t, "Find",1)
	geolocalizationRestMock.AssertCalled(t, "Find",ip)
}

func TestGetTraceShouldFindCountry(t *testing.T) {
	// Arrange
	ip := "138.82.0.1"
	countryCode := "CA"
	currencyCode := "CAD"

	// create an instance of geolocalization
	geolocalizationRestMock := new(geolocalizationRestClientMock)
	geolocalizationRestMock.On("Find", ip).Return(model.GeoLocalization{
		CountryCode:     countryCode,
		CountryCode3:    "CAN",
		CountryCodeName: "Canada",
		CountryEmoji:    "ca",
	}, nil)

	// create an instance of country
	countryRestMock := new(countryRestClientMock)
	countryRestMock.On("Find", countryCode).Return(model.Country{
		Name:       "Canada",
		Alpha2Code: countryCode,
		Alpha3Code: "CAN",
		Latlng:     []float64{60.0,-95.0},
		Currencies: []model.Currency{{Name: "Canadian dollar", Code: currencyCode, Symbol: "$"}},
		Languages:  []model.Language{{Name: "English"}},
	}, nil)

	countryRestMock.On("Find", utils.ArgCode).Return(model.Country{
		Name:       "Argentina",
		Alpha2Code: utils.ArgCode,
		Alpha3Code: "ARG",
		Latlng:     []float64{34.0,-64.0},
		Currencies: []model.Currency{{Name: "Argentine peso", Code: "ARS", Symbol: "$"}},
		Languages:  []model.Language{{Name: "Spanish"}},
	}, nil)

	currencyRestMock := new(currencyRestClientMock)
	currencyRestMock.On("Find",currencyCode).Return(model.CurrencyRate{
		Base:  "USD",
		Date:  time.Now().String(),
		Rates: map[string]float64{},
	},nil)

	sut := NewTraceIpService(geolocalizationRestMock,countryRestMock,currencyRestMock)

	// Action
	_, _ = sut.GetTraceIp(ip)
	// Assert
	countryRestMock.AssertNumberOfCalls(t, "Find",2)
	countryRestMock.AssertCalled(t, "Find",utils.ArgCode)
	countryRestMock.AssertCalled(t, "Find",countryCode)
}

func TestGetTraceShouldFindCurrency(t *testing.T) {
	// Arrange
	ip := "138.82.0.1"
	countryCode := "CA"
	currencyCode := "CAD"

	// create an instance of geolocalization
	geolocalizationRestMock := new(geolocalizationRestClientMock)
	geolocalizationRestMock.On("Find", ip).Return(model.GeoLocalization{
		CountryCode:     countryCode,
		CountryCode3:    "CAN",
		CountryCodeName: "Canada",
		CountryEmoji:    "ca",
	}, nil)

	// create an instance of country
	countryRestMock := new(countryRestClientMock)
	countryRestMock.On("Find", countryCode).Return(model.Country{
		Name:       "Canada",
		Alpha2Code: countryCode,
		Alpha3Code: "CAN",
		Latlng:     []float64{60.0,-95.0},
		Currencies: []model.Currency{{Name: "Canadian dollar", Code: currencyCode, Symbol: "$"}},
		Languages:  []model.Language{{Name: "English"}},
	}, nil)

	countryRestMock.On("Find", utils.ArgCode).Return(model.Country{
		Name:       "Argentina",
		Alpha2Code: utils.ArgCode,
		Alpha3Code: "ARG",
		Latlng:     []float64{34.0,-64.0},
		Currencies: []model.Currency{{Name: "Argentine peso", Code: "ARS", Symbol: "$"}},
		Languages:  []model.Language{{Name: "Spanish"}},
	}, nil)

	currencyRestMock := new(currencyRestClientMock)
	currencyRestMock.On("Find",currencyCode).Return(model.CurrencyRate{
		Base:  "USD",
		Date:  time.Now().String(),
		Rates: map[string]float64{},
	},nil)

	sut := NewTraceIpService(geolocalizationRestMock,countryRestMock,currencyRestMock)

	// Action
	_, _ = sut.GetTraceIp(ip)
	// Assert
	currencyRestMock.AssertNumberOfCalls(t, "Find",1)
	currencyRestMock.AssertCalled(t, "Find",currencyCode)
}

func TestGetTraceIp(t *testing.T) {
	// Arrange
	ip := "138.82.0.1"
	countryCode := "CA"
	currencyCode := "CAD"
	expectedTrace := model.TraceIpResult{
		Ip:       ip,
		Date:     time.Now(),
		Country:  "Canada",
		IsoCode:  "CAN",
		Language: "English",
		Currency: "CAD (1 USD = 0 CAD)",
		Distance: 3659,
	}

	// create an instance of geolocalization
	geolocalizationRestMock := new(geolocalizationRestClientMock)
	geolocalizationRestMock.On("Find", ip).Return(model.GeoLocalization{
		CountryCode:     countryCode,
		CountryCode3:    "CAN",
		CountryCodeName: "Canada",
		CountryEmoji:    "ca",
	}, nil)

	// create an instance of country
	countryRestMock := new(countryRestClientMock)
	countryRestMock.On("Find", countryCode).Return(model.Country{
		Name:       "Canada",
		Alpha2Code: countryCode,
		Alpha3Code: "CAN",
		Latlng:     []float64{60.0,-95.0},
		Currencies: []model.Currency{{Name: "Canadian dollar", Code: currencyCode, Symbol: "$"}},
		Languages:  []model.Language{{Name: "English"}},
	}, nil)

	countryRestMock.On("Find", utils.ArgCode).Return(model.Country{
		Name:       "Argentina",
		Alpha2Code: utils.ArgCode,
		Alpha3Code: "ARG",
		Latlng:     []float64{34.0,-64.0},
		Currencies: []model.Currency{{Name: "Argentine peso", Code: "ARS", Symbol: "$"}},
		Languages:  []model.Language{{Name: "Spanish"}},
	}, nil)

	currencyRestMock := new(currencyRestClientMock)
	currencyRestMock.On("Find",currencyCode).Return(model.CurrencyRate{
		Base:  "USD",
		Date:  time.Now().String(),
		Rates: map[string]float64{},
	},nil)

	sut := NewTraceIpService(geolocalizationRestMock,countryRestMock,currencyRestMock)

	// Action
	result, _ := sut.GetTraceIp(ip)
	// Assert
	currencyRestMock.AssertExpectations(t)
	geolocalizationRestMock.AssertExpectations(t)
	countryRestMock.AssertExpectations(t)

	assert.NotNil(t, result,"Should not be empty")
	assert.EqualValues(t, result.Currency, expectedTrace.Currency,"Should be equals")
	assert.EqualValues(t, result.Country, expectedTrace.Country,"Should be equals")
	assert.EqualValues(t, result.Ip, expectedTrace.Ip,"Should be equals")
	assert.EqualValues(t, result.IsoCode, expectedTrace.IsoCode,"Should be equals")
	assert.EqualValues(t, result.Language, expectedTrace.Language,"Should be equals")
}

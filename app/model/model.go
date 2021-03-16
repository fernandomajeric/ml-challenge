// Package model implements data models
package model

import "time"

type Currency struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Symbol string `json:"symbol"`
}
type Language struct {
	Name string `json:"name"`
}

type Country struct {
	Name           string     `json:"name"`
	TopLevelDomain []string   `json:"topLevelDomain"`
	Alpha2Code     string     `json:"alpha2Code"`
	Alpha3Code     string     `json:"alpha3Code"`
	CallingCodes   string     `json:"callingCodes"`
	AltSpellings   []string   `json:"altSpellings"`
	Region         string     `json:"region"`
	Subregion      string     `json:"subregion"`
	Population     int        `json:"population"`
	Latlng         []float64  `json:"latlng"`
	Demonym        string     `json:"demonym"`
	Area           int        `json:"area"`
	Gini           int        `json:"gini"`
	Timezones      []string   `json:"timezones"`
	NativeName     string     `json:"nativeName"`
	NumericCode    string     `json:"numericCode"`
	Currencies     []Currency `json:"currencies"`
	Languages      []Language `json:"languages"`
	Flag           string     `json:"flag"`
	Cioc           string     `json:"cioc"`
}

type CurrencyRate struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

type GeoLocalization struct {
	CountryCode     string `json:"countryCode"`
	CountryCode3    string `json:"countryCode3"`
	CountryCodeName string `json:"countryName"`
	CountryEmoji    string `json:"countryEmoji"`
}

type StatisticCore struct {
	Country  string  `json:"country"`
	Distance float64 `json:"distance"`
	HitCount int64   `json:"hitCount"`
}

type StatisticItem struct {
	CountryName string
	Distance    float64
}

type Statistic struct {
	MinScore        StatisticCore `json:"minScore"`
	MaxScore        StatisticCore `json:"maxScore"`
	AverageDistance float64       `json:"averageDistance"`
}

type TraceIpResult struct {
	Ip       string
	Date     time.Time
	Country  string
	IsoCode  string
	Language string
	Currency string
	Distance float64
}

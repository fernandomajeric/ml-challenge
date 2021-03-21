package utils

import (
	"errors"
	"math"
)

func Calculate(latLng1 []float64, latLng2 []float64) (float64, error) {
	if len(latLng1) != 2 {
	  error:= errors.New("invalid latLng1")
	  return 0, error
	}

	if len(latLng2) != 2 {
		error:= errors.New("invalid latLng2")
		return 0, error
	}

	lat1 := latLng1[0]
	lon1 := latLng1[1]
	lat2 := latLng2[0]
	lon2 := latLng2[1]

	if lat1 == lat2 && lon1 == lon2 {
		return 0,nil
	}

	theta := lon1 - lon2
	distance := math.Sin(degreesToRadians(lat1))*math.Sin(degreesToRadians(lat2)) + math.Cos(degreesToRadians(lat1))*math.Cos(degreesToRadians(lat2))*math.Cos(degreesToRadians(theta))
	distance = math.Acos(distance)
	distance = radiansToDegrees(distance)
	distance = distance * 60 * 1.1515 * 1.609344

	return math.Round(distance),nil
}

func degreesToRadians(deg float64) float64 {
	return math.Pi * deg / 180.0
}

func radiansToDegrees(rad float64) float64 {
	return rad * 180.0 / math.Pi
}

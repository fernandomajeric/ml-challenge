package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateWithEmptyLatLn1(t *testing.T) {
	//Arrange
	var latLng1 []float64
	latLng2 := []float64{40,-1}
	//Action
	_, error := Calculate(latLng1, latLng2)
	//Assert
	if assert.Errorf(t, error,"Parameter latLng1 should contain two value"){
		assert.EqualError(t, error, "invalid latLng1")
	}
}

func TestCalculateWithEmptyLatLn2(t *testing.T) {
	//Arrange
	latLng1 :=[]float64{40,-1}
	var latLng2  []float64
	//Action
	_, error := Calculate(latLng1, latLng2)
	//Assert
	if assert.Errorf(t, error,"Parameter latLng1 should contain two value"){
		assert.EqualError(t, error, "invalid latLng2")
	}
}

func TestCalculateLatLong1EqualsLatLn2(t *testing.T) {
	//Arrange
	latLng1 :=[]float64{40,-2}
	latLng2 := []float64{40,-2}
	distanceExp := 0
	//Action
	d, error := Calculate(latLng1, latLng2)
	//Assert
	assert.EqualValues(t, d,distanceExp,"Distance should be zero")
	assert.NoErrorf(t, error,"No errors found")
}


func TestCalculate(t *testing.T) {
	//Arrange
	latLng1 :=[]float64{-34,-64}
	latLng2 := []float64{40,-4}
	distanceExp := 10274
	//Action
	d, error := Calculate(latLng1, latLng2)
	//Assert
	assert.EqualValues(t, d,distanceExp,"Distance should be %q", distanceExp)
	assert.NoErrorf(t, error,"No errors found")
}

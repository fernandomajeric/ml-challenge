package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateIpAddressWithoutNumericFormat(t *testing.T) {
	// Arrange
	ip := "a.82.0.0"
	// Action
	result := ValidateIpAddress(ip)
	// Assert
	assert.Falsef(t, result,"Ip address : %s should be only numbers!", ip)
}

func TestValidateIpAddressWithLessThanFourFields(t *testing.T) {
	// Arrange
	ip := "138.82.0"
	// Action
	result := ValidateIpAddress(ip)
	// Assert
	assert.Falsef(t, result,"Ip address : %q should be formatted by 4 fields!", ip)
}

func TestValidateIpAddressWithSomeFieldNumberNegative(t *testing.T) {
	// Arrange
	ip := "138.-82.0.1"
	// Action
	result := ValidateIpAddress(ip)
	// Assert
	assert.Falsef(t, result,"Ip address : %q should be valid fields", ip)
}

func TestValidateIpAddressWithSomeFieldHigherThanMaxValueAllowed(t *testing.T) {
	// Arrange
	ip := "138.82.0.256"
	// Action
	result := ValidateIpAddress(ip)
	// Assert
	assert.Falsef(t, result,"Ip address : %q should be valid fields", ip)
}

func TestValidateIpAddressWithValidIpAddress(t *testing.T) {
	// Arrange
	ip := "138.82.0.1"
	// Action
	result := ValidateIpAddress(ip)
	// Assert
	assert.Truef(t, result,"Ip address : %q well format!", ip)
}

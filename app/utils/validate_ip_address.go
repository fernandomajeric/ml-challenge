// Package model implements utility routines for services
package utils

import (
	"strconv"
	"strings"
)

//ValidateIpAddress Validate real Ip format
func ValidateIpAddress(ipAddress string) bool {
	fields := strings.Split(ipAddress, ".")
	if len(fields) != 4 {
		return false
	}

	for _, f := range fields {
		v, err := strconv.ParseInt(f, 10, 32)
		if (v < 0 || v > 255) || err != nil {
			return false
		}
	}
	return true
}

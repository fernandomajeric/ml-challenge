// Package model implements utility routines for
package utils

import (
	"strconv"
	"strings"
)

func ValidateIpAddress(ipAddress string) bool {
	tokens := strings.Split(ipAddress, ".")
	if len(tokens) != 4 {
		return false
	}

	for _, t := range tokens {
		v, err := strconv.ParseInt(t, 10, 32)
		if err != nil && (v < 0 || v > 255) {
			return false
		}
	}
	return true
}

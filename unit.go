package main

import (
	"fmt"
	"slices"
	"strings"
)

var allowedUnits = []string{"B", "KB", "MB", "GB"}

func parseUnit(unit *string) (string, error) {
	if unit == nil || *unit == "" {
		return "", fmt.Errorf("unit cannot be empty")
	}
	if !slices.Contains(allowedUnits, strings.ToUpper(*unit)) {
		return "", fmt.Errorf("unrecognized unit: %s", *unit)
	}
	return *unit, nil
}

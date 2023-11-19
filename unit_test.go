package main

import "testing"

func TestParseUnit(t *testing.T) {
	t.Run("should return error when unit is empty", func(t *testing.T) {
		_, err := parseUnit(nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("should return error when unit is not allowed", func(t *testing.T) {
		unit := "unknown unit"
		_, err := parseUnit(&unit)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("should return unit when unit is allowed", func(t *testing.T) {
		unit := "B"
		u, err := parseUnit(&unit)
		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}
		if u != unit {
			t.Errorf("expected %s, got %s", unit, u)
		}
	})
}

package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	t.Run("Should return error when template is empty", func(t *testing.T) {
		_, err := parseTemplate(nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("Should return error when template is not allowed", func(t *testing.T) {
		template := "unknown template"
		_, err := parseTemplate(&template)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("Should return template when template is allowed", func(t *testing.T) {
		template := "default"
		u, err := parseTemplate(&template)
		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}
		if u == nil {
			t.Errorf("expected function, got nil")
		}
	})
}

func TestFormatBytes(t *testing.T) {
	t.Run("Should return 0B when bytes is 0", func(t *testing.T) {
		d := Data{Total: 0, Unit: "B"}
		b := d.FormatBytes()
		if b != "0.000 B" {
			t.Errorf("expected 0.000 B, got %s", b)
		}
	})

	t.Run("Should return 1KB when bytes is 1024", func(t *testing.T) {
		d := Data{Total: 1024, Unit: "KB"}
		b := d.FormatBytes()
		if b != "1.000 KB" {
			t.Errorf("expected 1.000 KB, got %s", b)
		}
	})
}

func TestDefaultPrint(t *testing.T) {
	var bf bytes.Buffer
	d := Data{Total: 1024, Unit: "KB", Line: "my log line"}
	b := defaultPrint(&bf, d)
	if b != 11 {
		t.Errorf("expected 11, got %d", b)
	}
	if !strings.Contains(bf.String(), "⚡ Total: 1.000 KB | Line before 011 B") {
		t.Errorf("expected '⚡ Total: 1.000 KB | Line before 011 B', got %s", bf.String())
	}
}

func TestMinimalPrint(t *testing.T) {
	var bf bytes.Buffer
	d := Data{Total: 1024, Unit: "KB", Line: "this a log message"}
	b := minimalPrint(&bf, d)
	if b != 18 {
		t.Errorf("expected 18, got %d", b)
	}

	if !strings.Contains(bf.String(), "⚡ 1.000 KB | 018 B |") {
		t.Errorf("expected '⚡ 1.000 KB | 018 B |', got %s", bf.String())
	}
}

package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/inhies/go-bytesize"
)

type Data struct {
	Total int
	Unit  string
	Line  string
}

func (d *Data) FormatBytes() string {
	b, err := bytesize.Parse(strconv.Itoa(d.Total) + " B")
	if err != nil {
		panic(err)
	}
	return b.Format("%.3f ", d.Unit, false)
}

type PrintFunc func(w io.Writer, d Data) int

var (
	allowedTemplates = map[string]PrintFunc{
		"default": defaultPrint,
		"minimal": minimalPrint,
	}
)

func parseTemplate(template *string) (PrintFunc, error) {
	if template == nil || *template == "" {
		return nil, fmt.Errorf("template cannot be empty")
	}
	if _, ok := allowedTemplates[*template]; !ok {
		return nil, fmt.Errorf("unrecognized template: %s", *template)
	}
	return allowedTemplates[*template], nil
}

func defaultPrint(w io.Writer, d Data) int {
	lineBytes := len(d.Line)
	fmt.Fprintln(w, d.Line)
	fmt.Fprintf(w, "⚡ Total: %s | Line before %03d B \n", d.FormatBytes(), lineBytes)
	return lineBytes
}

func minimalPrint(w io.Writer, d Data) int {
	lineBytes := len(d.Line)
	fmt.Fprintf(w, "⚡ %s | %03d B | %s \n", d.FormatBytes(), lineBytes, d.Line)
	return lineBytes
}

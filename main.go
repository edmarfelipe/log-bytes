package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	unitPtr    = flag.String("unit", "kb", "Data unit to print  [B, KB, MB, GB]")
	templatPrt = flag.String("template", "default", "Template to print [default, minimal]")
)

func main() {
	flag.Parse()

	unit, err := parseUnit(unitPtr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	printFunc, err := parseTemplate(templatPrt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	total := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		total += printFunc(os.Stdout, Data{
			Total: total,
			Unit:  unit,
			Line:  s.Text(),
		})
	}
}

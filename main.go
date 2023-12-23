package main

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

var (
	Version string
	Tags    string
)

func printErr(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", e)
		os.Exit(-1)
	}
}

func main() {
	cfg := parseFlags()
	data := readData(cfg.data)
	tmpl, err := template.ParseFiles(cfg.templates...)
	printErr(err)

	var output io.Writer

	if *cfg.output == "" {
		output = os.Stdout
	} else {
		output, err = os.OpenFile(*cfg.output, os.O_WRONLY|os.O_CREATE, 0600)
		printErr(err)
	}

	err = tmpl.Execute(output, data)
	printErr(err)

	if *cfg.output != "" {
		output.(*os.File).Close()
	}
}

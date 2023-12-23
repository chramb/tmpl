package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type config struct {
	templates stringsFlag
	data      stringsFlag
	output    *string
}

type stringsFlag []string

func (s *stringsFlag) String() string {
	return strings.Join(*s, " ")
}

func (s *stringsFlag) Set(val string) error {
	*s = append(*s, val)
	return nil
}

func parseFlags() *config {
	help := flag.Bool("help", false, "print this help message and exit")
	version := flag.Bool("version", false, "print version info and exit")
	output := flag.String("out", "", "path to the output file")

	var templates stringsFlag
	flag.Var(&templates, "in", "path to the template file")

	var data stringsFlag
	flag.Var(&data, "data", "path to the data file")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *version {
		tags, _ := strings.CutSuffix(Tags, " ")
		if tags != "" {
			tags += "\n"
		}
		fmt.Printf("tmpl: %s\n%s", Version, tags)
		os.Exit(0)
	}

	if len(templates) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: missing at least one -in file")
		os.Exit(-1)
	}

	if len(data) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: missing at least one -data file")
		os.Exit(-1)
	}

	return &config{
		templates: templates,
		data:      data,
		output:    output,
	}
}

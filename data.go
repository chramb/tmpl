package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

var dataReader = make(map[string]func([]byte, any) error)

func init() {
	dataReader[".json"] = json.Unmarshal
}

func readData(files stringsFlag) map[string]interface{} {
	data := make(map[string]interface{})

	for _, p := range files {
		f, err := os.ReadFile(p)
		printErr(err)

		reader, ok := dataReader[path.Ext(p)]
		if !ok {
			var keys []string
			for k := range dataReader {
				format, _ := strings.CutPrefix(k, ".")
				keys = append(keys, format)
			}
			fmt.Fprintf(os.Stderr, "ERROR: unsupported file format, only supporting: %v\n", keys)
			os.Exit(-1)
		}

		err = reader(f, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to read data")
			os.Exit(-1)
		}

	}
	return data
}

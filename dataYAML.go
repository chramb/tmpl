//go:build yaml

package main

import "gopkg.in/yaml.v3"

func init() {
	dataReader[".yml"] = yaml.Unmarshal
	dataReader[".yaml"] = dataReader[".yml"]
}

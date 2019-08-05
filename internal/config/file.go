package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// fileToSchema reads the configuration file to convert to schema
func fileToSchema(file string) (s *schema, err error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(f, &s); err != nil {
		return nil, err
	}

	return
}

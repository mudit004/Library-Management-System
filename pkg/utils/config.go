package utils

import (
	"lms/pkg/types"
	"os"

	"gopkg.in/yaml.v3"
)

func NewConfig(path string) (*types.Config, error) {
	config := &types.Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

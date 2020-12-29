package config

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

func FromFile(path string, def interface{}) (interface{}, error) {
	file, err := os.Open(path)
	switch {
	case os.IsNotExist(err):
		return def, nil
	case err != nil:
		return nil, err
	}

	defer file.Close()
	return FromReader(file, def)
}

func FromReader(reader io.Reader, def interface{}) (interface{}, error) {
	cfg := def
	_, err := toml.DecodeReader(reader, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

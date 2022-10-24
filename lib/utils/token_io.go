package utils

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type TokenConfig struct {
	AccessToken  string `toml:"accesstoken"`
	RefreshToken string `toml:"refreshtoken"`
}

const FILE_PATH = "/tmp/token.toml"

func LoadToken() (TokenConfig, error) {
	file := FILE_PATH
	_, err := os.Stat(file)

	conf := TokenConfig{}

	if err == nil {
		_, err := toml.DecodeFile(file, &conf)
		if err != nil {
			wrappedError := errors.Wrapf(err, "Faield load file: %s.", FILE_PATH)
			return conf, wrappedError
		}
	}

	return conf, nil
}

func SaveToken(token TokenConfig) error {
	f, err := os.Create(FILE_PATH)
	if err != nil {
		return errors.Wrapf(err, "Faield create file: %s.", FILE_PATH)
	}
	toml.NewEncoder(f).Encode(token)
	return nil
}

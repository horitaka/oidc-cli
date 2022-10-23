package utils

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
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
			return conf, err
		}
	}

	return conf, nil
}

func SaveToken(token TokenConfig) {
	f, err := os.Create(FILE_PATH)
	if err != nil {
		// return err
		fmt.Println(err)
		return
	}
	toml.NewEncoder(f).Encode(token)
}

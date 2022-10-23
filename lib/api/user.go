package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

type TokenConfig struct {
	AccessToken  string `toml:"accesstoken"`
	RefreshToken string `toml:"refreshtoken"`
}

func GetUserInfo() {
	token, _ := loadToken()
	fmt.Println(token.AccessToken)
	fmt.Println(token.RefreshToken)

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v1/userinfo", nil) // TODO: 定数はconstantsに移動する
	req.Header.Add("Authorization", "Bearer"+token.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}

// TODO: utilsに移動
func loadToken() (TokenConfig, error) {
	file := "/tmp/token.toml"
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

package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

// TODO: api/constantnsに移動
const AUTH_TOKEN_URL = "https://oauth2.googleapis.com/token"

// TODO: apiに移動
type PostTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
}

// TODO: utilsに移動
type TokenConfig struct {
	AccessToken  string `toml:"accesstoken"`
	RefreshToken string `toml:"refreshtoken"`
}

type CallbackUrlQueryParam struct {
	Code  string
	State string
}

func Token(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load env file: %v", err)
	}

	tokenApiParam := getParams(r.URL)

	res, err := postToken(tokenApiParam)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer res.Body.Close()

	post := convertResToJosn(res)
	saveToken(post)

	// debug
	fmt.Fprintln(w, tokenApiParam.State)
	fmt.Fprintln(w, tokenApiParam.Code)
	fmt.Fprintln(w, res)
	fmt.Fprintln(w, res.StatusCode)
	fmt.Fprintln(w, post.RefreshToken)
}

func getParams(url *url.URL) CallbackUrlQueryParam {
	query := url.Query()
	params := CallbackUrlQueryParam{
		State: query.Get("state"),
		Code:  query.Get("code"),
	}
	return params
}

// TODO: apiに移動
func postToken(param CallbackUrlQueryParam) (*http.Response, error) {
	v := url.Values{}
	v.Set("code", param.Code)
	v.Add("client_id", os.Getenv("CLIENT_ID"))
	v.Add("client_secret", os.Getenv("CLIENT_SECRET"))
	v.Add("redirect_uri", "http://localhost:8080/callback") // TODO: constantnsに移動する
	v.Add("grant_type", "authorization_code")

	resp, err := http.PostForm(AUTH_TOKEN_URL, v)
	return resp, err
}

func convertResToJosn(res *http.Response) PostTokenResponse {
	body, _ := ioutil.ReadAll(res.Body)
	var posts PostTokenResponse
	if err := json.Unmarshal(body, &posts); err != nil {
		fmt.Println(err)
	}

	return posts
}

// TODO: utils/token_ioに移動

func saveToken(res PostTokenResponse) {
	f, err := os.Create("/tmp/outhtoken.toml")
	if err != nil {
		// return err
		fmt.Println(err)
		return
	}
	conf := TokenConfig{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}
	toml.NewEncoder(f).Encode(conf)
}

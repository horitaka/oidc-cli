package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/horitaka/oidc-cli/constants"
)

type PostTokenParam struct {
	Code string
}

type PostTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
}

func PostToken(param PostTokenParam) (*http.Response, error) {
	v := url.Values{}
	v.Set("code", param.Code)
	v.Add("client_id", os.Getenv("CLIENT_ID"))
	v.Add("client_secret", os.Getenv("CLIENT_SECRET"))
	v.Add("redirect_uri", constants.CALLBACK_URL)
	v.Add("grant_type", "authorization_code")

	resp, err := http.PostForm(constants.TOKEN_URL, v)

	// TODO: jsonに変換する
	return resp, err
}

func ConvertResToJosn(res *http.Response) PostTokenResponse {
	body, _ := ioutil.ReadAll(res.Body)
	var posts PostTokenResponse
	if err := json.Unmarshal(body, &posts); err != nil {
		fmt.Println(err)
	}

	return posts
}

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/horitaka/oidc-cli/constants"
	"github.com/horitaka/oidc-cli/lib/utils"
	"github.com/pkg/errors"
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

func PostToken(param PostTokenParam) (PostTokenResponse, error) {
	v := createQueryParams(param)
	resp, err := http.PostForm(constants.TOKEN_URL, v)
	if err != nil {
		wrappedError := errors.Wrap(err, "Failed to post token.")
		return PostTokenResponse{}, wrappedError
	}
	defer resp.Body.Close()

	json, err := convertResToJosn(resp)
	return json, err
}

func createQueryParams(param PostTokenParam) url.Values {
	authClient := utils.LoadEnv()

	v := url.Values{}
	v.Set("code", param.Code)
	v.Add("client_id", authClient.ClientId)
	v.Add("client_secret", authClient.ClientSecret)
	v.Add("redirect_uri", constants.CALLBACK_URL)
	v.Add("grant_type", "authorization_code")
	return v
}

func convertResToJosn(res *http.Response) (PostTokenResponse, error) {
	body, _ := ioutil.ReadAll(res.Body)
	var result PostTokenResponse
	if err := json.Unmarshal(body, &result); err != nil {
		wrappedError := errors.Wrap(err, "Failed to convert response to json.")
		return result, wrappedError
	}

	return result, nil
}

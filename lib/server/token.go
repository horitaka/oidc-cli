package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/horitaka/oidc-cli/lib/api"
	"github.com/horitaka/oidc-cli/lib/utils"
)

type CallbackUrlQueryParam struct {
	Code  string
	State string
}

func Token(w http.ResponseWriter, r *http.Request) {
	tokenApiParam := getParams(r.URL)

	resp, err := api.PostToken(tokenApiParam)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	token := utils.TokenConfig{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}
	utils.SaveToken(token)

	fmt.Fprintln(w, "Succeeded in obtaining token. Return to terminal and use CLI to call API.")

	// debug
	// fmt.Fprintln(w, tokenApiParam.State)
	fmt.Fprintln(w, tokenApiParam.Code)
	// fmt.Fprintln(w, res)
	// fmt.Fprintln(w, res.StatusCode)
	fmt.Fprintln(w, resp.RefreshToken)

}

func getParams(url *url.URL) api.PostTokenParam {
	query := url.Query()
	params := api.PostTokenParam{
		// State: query.Get("state"),
		Code: query.Get("code"),
	}
	return params
}

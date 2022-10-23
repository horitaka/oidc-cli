package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/horitaka/oidc-cli/lib/api"
	"github.com/horitaka/oidc-cli/lib/utils"
	"github.com/joho/godotenv"
)

type CallbackUrlQueryParam struct {
	Code  string
	State string
}

func Token(w http.ResponseWriter, r *http.Request) {
	// TODO: 環境変数読み込みをメソッドにする
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load env file: %v", err)
	}

	tokenApiParam := getParams(r.URL)

	res, err := api.PostToken(tokenApiParam)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer res.Body.Close()
	// ConvertResToJosnはposttokenに移動する
	post := api.ConvertResToJosn(res)

	token := utils.TokenConfig{
		AccessToken:  post.AccessToken,
		RefreshToken: post.RefreshToken,
	}
	utils.SaveToken(token)

	// TODO: メッセージを表示する

	// debug
	// fmt.Fprintln(w, tokenApiParam.State)
	fmt.Fprintln(w, tokenApiParam.Code)
	fmt.Fprintln(w, res)
	fmt.Fprintln(w, res.StatusCode)
	fmt.Fprintln(w, post.RefreshToken)
}

func getParams(url *url.URL) api.PostTokenParam {
	query := url.Query()
	params := api.PostTokenParam{
		// State: query.Get("state"),
		Code: query.Get("code"),
	}
	return params
}

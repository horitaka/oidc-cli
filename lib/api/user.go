package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/horitaka/oidc-cli/constants"
	"github.com/horitaka/oidc-cli/lib/utils"
)

func GetUserInfo() {
	token, _ := utils.LoadToken()

	client := &http.Client{}
	req, err := http.NewRequest("GET", constants.USERINFO_URL, nil)
	req.Header.Add("Authorization", "Bearer"+token.AccessToken)
	resp, err := client.Do(req)

	// TODO: tokenの有効期限切れの場合はrefresh tokenでaccess tokenを取得し直す
	// TODO: refresh tokenの期限が切れている場合はメッセージを出して再ログインをさせる

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(resp.Status)
	fmt.Println(string(body))
}

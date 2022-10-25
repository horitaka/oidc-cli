package api

import (
	"github.com/horitaka/oidc-cli/constants"
	"github.com/horitaka/oidc-cli/lib/utils"
)

func GetUserInfo() error {
	resp := utils.Get(constants.USERINFO_URL, nil)
	utils.PrintResponse(resp)
	return resp.Error
}

package api

import (
	"fmt"

	"github.com/horitaka/oidc-cli/constants"
	"github.com/horitaka/oidc-cli/lib/utils"
)

func GetCalendars() error {
	resp := utils.Get(constants.CALENDERS_URL)
	// fmt.Println(resp.Status)
	fmt.Println(string(resp.Body))
	return resp.Error
}

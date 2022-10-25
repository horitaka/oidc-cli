package api

import (
	"fmt"

	"github.com/horitaka/oidc-cli/constants"
	"github.com/horitaka/oidc-cli/lib/utils"
)

func GetCalendars() error {
	query := map[string]string{"maxResults": "1"}
	resp := utils.Get(constants.CALENDERS_URL, query)
	// fmt.Println(resp.Status)
	fmt.Println(string(resp.Body))
	return resp.Error
}

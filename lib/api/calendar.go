package api

import (
	"github.com/horitaka/oidc-cli/constants"
	"github.com/horitaka/oidc-cli/lib/utils"
)

func GetCalendars() error {
	query := map[string]string{"maxResults": "1"}
	resp := utils.Get(constants.CALENDERS_URL, query)
	utils.PrintResponse(resp)
	return resp.Error
}

func CreateCalendar() error {
	body := map[string]string{
		"summary": "test",
	}
	resp := utils.Post(constants.CALENDER_URL, body)
	utils.PrintResponse(resp)
	return resp.Error
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/horitaka/oidc-cli/lib/api"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	GET_USERINFO    = "Get Userinfo"
	GET_CALENDARS   = "Get Calendar List"
	CREATE_CALENDAR = "Create Calendar"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run API",
	Long:  "Run API",
	RunE:  hadleRunCmd,
}

func hadleRunCmd(cmd *cobra.Command, args []string) error {
	var err error

	prompt := promptui.Select{
		Label: "Select API",
		Items: []string{GET_USERINFO, GET_CALENDARS, CREATE_CALENDAR, "get xxx"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return errors.Wrap(err, "Prompt failed.")
	}

	switch result {
	case GET_USERINFO:
		err = api.GetUserInfo()
	case GET_CALENDARS:
		err = api.GetCalendars()
	case CREATE_CALENDAR:
		err = api.CreateCalendar()
	default:
		fmt.Printf("No API matched %q\n", result)
	}

	return err
}

func init() {
	rootCmd.AddCommand(runCmd)
}

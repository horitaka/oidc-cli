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
	GET_USERINFO = "Get userinfo"
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
		Items: []string{GET_USERINFO, "get xxx"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return errors.Wrap(err, "Prompt failed.")
	}

	switch result {
	case GET_USERINFO:
		err = api.GetUserInfo()
	default:
		fmt.Printf("No API matched %q\n", result)
	}

	return err
}

func init() {
	rootCmd.AddCommand(runCmd)
}

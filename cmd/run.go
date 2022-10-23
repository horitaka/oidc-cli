/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: hadleRunCmd,
}

func hadleRunCmd(cmd *cobra.Command, args []string) error {
	prompt := promptui.Select{
		Label: "Select API",
		Items: []string{GET_USERINFO, "get xxx"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	fmt.Printf("You choose %q\n", result)

	switch result {
	case GET_USERINFO:
		api.GetUserInfo()
	default:
		fmt.Printf("No API matched %q\n", result)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

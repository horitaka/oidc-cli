/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/horitaka/oidc-cli/constants"
	"github.com/horitaka/oidc-cli/lib/server"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Google",
	Long:  "Login to Google",
	RunE:  handleloginCmd,
}

func handleloginCmd(cmd *cobra.Command, args []string) error {
	fmt.Println("Open " + constants.LOGIN_URL)

	http.HandleFunc("/login", server.Login)
	http.HandleFunc("/callback", server.Token)
	http.ListenAndServe(":"+constants.PORT, nil)

	return nil
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

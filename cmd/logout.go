/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from Google",
	Long:  "Logout from Google",
	RunE:  handleLogoutCmd,
}

func handleLogoutCmd(cmd *cobra.Command, args []string) error {
	fmt.Println("logout called")
	// TODO: ログアウト機能の実装
	return nil
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

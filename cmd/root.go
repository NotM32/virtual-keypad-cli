/*
   Copyright © 2024 Riley OConnor <m32@m32.io>
*/
package cmd

import (
	"fmt"
	"os"

	query "github.com/NotM32/keypadgo/queries"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "keypadgo",
	Short: "A cli tool to interact with DMP virtual keypad acess control and alarm systems.",
	Long: `keypadgo is a cli tool for interacting with DMP white labelled access control/building alarm
systems. Commonly resold by ADT for commercial wiegand effect card readers.

This is created by reverse engineering APIs used in the web/mobile app and uses a mix of GraphQL/rest APIs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var (
	email          = os.Getenv("DMP_EMAIL")
	password       = os.Getenv("DMP_PASSWORD")
	userCode       = os.Getenv("DMP_PASSCODE")
	panel          = os.Getenv("DMP_PANEL")
	panelSecondary = os.Getenv("DMP_PANEL_SECONDARY")
	authToken string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	gqlConfig, err := query.LoadGraphQLConfig()
	if err != nil {
		fmt.Println("Error reading GraphQLConfig: ", err)
		panic(err)
	}

	if email != "" && password != ""{
		authToken = query.LoginAuthenticate(email, password, gqlConfig.Extensions.Endpoints.Default.Url)
		fmt.Println("Authenticated via email & password")
	} else if authToken == "" {
		fmt.Println("Must specify an authtoken or signin credentials")
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.godmp.yaml)")
	rootCmd.PersistentFlags().StringVar(&email, "email", os.Getenv("DMP_EMAIL"), "Email used for authentication. Defaults to $DMP_EMAIL")
	rootCmd.PersistentFlags().StringVar(&password, "password", os.Getenv("DMP_PASSWORD"), "Password used for authentication. Defaults to $DMP_PASSWORD")
	rootCmd.MarkFlagsRequiredTogether("email", "password")
	rootCmd.PersistentFlags().StringVar(&authToken, "authtoken", authToken, "Token to authenticate requests (if email/password not provided, this is required)")
	rootCmd.MarkFlagsMutuallyExclusive("authtoken", "email")
	rootCmd.MarkFlagsMutuallyExclusive("authtoken", "password")
	if os.Getenv("DMP_EMAIL") == "" {
		rootCmd.MarkFlagsOneRequired("email", "authtoken")
	}
	rootCmd.PersistentFlags().StringVar(&userCode, "passcode", os.Getenv("DMP_PASSCODE"), "Passcode/credential for the virtual user. Defaults to $DMP_PASSCODE")
	rootCmd.PersistentFlags().StringVar(&panel, "panel", os.Getenv("DMP_PANEL"), "Panel id to use. Defaults to $DMP_PANEL")
	rootCmd.PersistentFlags().StringVar(&panelSecondary, "altpanel", os.Getenv("DMP_PANEL_SECONDARY"), "Alternative panel id to use. Defaults to $DMP_PANEL_SECONDARY")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

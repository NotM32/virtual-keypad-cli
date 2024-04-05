/*
   Copyright © 2024 Riley OConnor <m32@m32.io>
*/
package cmd

import (
	"fmt"
	"slices"

	query "github.com/NotM32/virtual-keypad-cli/queries"
	"github.com/spf13/cobra"
)

// accessCmd represents the access command
var accessCmd = &cobra.Command{
	Use:   "access",
	Short: "Trigger access function on specified doors",
	Run: func(cmd *cobra.Command, args []string) {
		doors := query.GetDoors(authToken, userCode, panel, panelSecondary)
		for _, door := range doors {
			if slices.Contains(args, door.Number) {
				fmt.Println(args)
				door.Access()
				fmt.Printf("Accessing %s", door.Name)
			}
		}
	},
}

func init() {
	doorsCmd.AddCommand(accessCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

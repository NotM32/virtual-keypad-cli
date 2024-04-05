/*
   Copyright © 2024 Riley OConnor <m32@m32.io>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	query "github.com/NotM32/virtual-keypad-cli/queries"
)

// doorsCmd represents the doors command
var doorsCmd = &cobra.Command{
	Use:   "doors",
	Short: "Manage doors",
	Long: `Manage lock/unlock and triggered access functions of doors. View their current state.`,
	Run: func(cmd *cobra.Command, args []string) {
		doors := query.GetDoors(authToken, userCode, panel, panelSecondary)
		for _, door := range doors {
			fmt.Printf("[%s] | %s\n", door.Number, door.Name)
		}
	},
}


func init() {

	rootCmd.AddCommand(doorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doorsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doorsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

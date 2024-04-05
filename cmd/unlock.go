/*
   Copyright © 2024 Riley OConnor <m32@m32.io>
*/
package cmd

import (
	"fmt"
	"slices"

	query "github.com/NotM32/keypadgo/queries"
	"github.com/spf13/cobra"
)

// unlockCmd represents the unlock command
var unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock the specified doors",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		doors := query.GetDoors(authToken, userCode, panel, panelSecondary)
		for _, door := range doors {
			if slices.Contains(args, door.Number) {
				fmt.Println(args)
				door.Unlock()
				fmt.Printf("Unlocking %s", door.Name)
			}
		}
	},
}

func init() {
	doorsCmd.AddCommand(unlockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unlockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unlockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

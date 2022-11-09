/*
Copyright © 2022 Eric Peters <ericdpeters@gmail.com>

*/
package interactive

import (
	"fmt"

	"github.com/spf13/cobra"
)

// interactiveCmd represents the interactive command
var InteractiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Begin interactive mode for manual joint definition.",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		//cmd.Help()
		fmt.Println("Sorry, interactive mode still under development")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// interactiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// interactiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

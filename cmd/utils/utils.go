/*
Copyright © 2022 Eric Peters <e_peters@alum.mit.edu>

*/
package utils

import (
	"github.com/spf13/cobra"
)

// utilsCmd represents the utils command
var UtilsPal = &cobra.Command{
	Use:   "utils",
	Short: "Palette containing various utility functions",
	Long: `Utility functions and calculators that can be used independently
	from full joint definition and margin calculations.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("utils called")
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// utilsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// utilsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

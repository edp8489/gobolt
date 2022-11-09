/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package utils

import (
	"fmt"

	"github.com/spf13/cobra"
)

// preloadCmd represents the preload command
var preloadCmd = &cobra.Command{
	Use:   "preload",
	Short: "Calculate preload for a bolted joint",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("preload called")
	},
}

func init() {
	UtilsPal.AddCommand(preloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// preloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// preloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// TODO
// create promptContent and validationFunc for the following inputs
// 1. Units (select)
// 2. Nominal torque (float prompt)
// 3. Application tolerance (float prompt)
// 4. Nut factor (select)
// 5. Application method uncertainty (select)

// TODO
// calculate nominal preload using (2) and (4)
// calculate min and max preload using (3) and (5)
// print to console

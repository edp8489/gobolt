/*
Copyright © 2022 Eric Peters <ericdpeters@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Run in batch mode for predefined input file(s)",
	Long: `Given a path to an input file, or a directory with multiple
	input files, will solve and write outputs to file(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("batch called")
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// batchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// batchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

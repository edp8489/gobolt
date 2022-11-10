/*
Copyright © 2022 Eric Peters <e_peters@alum.mit.edu>

*/
package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
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
		//cmd.Help()
		cmdSel := utilsCmdSelect()
		switch cmdSel {
		case "preload":
			fmt.Printf("Sorry, preload program still under development")
		case "area":
			areaCmd.Run(cmd, args)
		}
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

// promptui select for command palette
func utilsCmdSelect() string {
	prompt := promptui.Select{
		Label: "Select a command",
		Items: []string{"area", "preload"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		//return nil
	}

	//fmt.Printf("You chose %q\n", result)
	return result
}

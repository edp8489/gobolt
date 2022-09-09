/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/edp8489/gobolt/internal/pkg/output"
	"github.com/edp8489/gobolt/internal/pkg/types"
	"github.com/spf13/cobra"
)

func Demo() {
	fmt.Println("Welcome to gobolt!")

	// define loads for MVP program
	loads := types.Loadset{
		Pshear: 1000,
		Factors: types.Factors{
			SFult: 1.4,
			SFyld: 1.2,
			FF:    1.15,
		},
	}
	output.PrintLoads(loads)

	// declare joint parameters for MVP program
	bolt := types.Bolt{
		D:   0.250,
		Fsu: 95e3,
	}
	output.PrintBolt(bolt)
	//var Fbru2d = 120e3 // sheet ult bearing strength, [psi]
	//var tA = 0.125      // sheet thickness, [in]

	// calculate bolt shear area (shank)
	// @todo refactor into utility package
	Ashank := types.Area(bolt.D)
	fmt.Printf("Bolt shank area: %.4f [in^2] (D = %.3f [in]\n", Ashank, bolt.D)

	// calculate bolt tensile area
	// @todo

	// calculate bolt thread shear area
	// @todo

	// calculate bolt shear margin
	// @todo refactor into utility package
	Psu_b := bolt.Fsu * Ashank
	fmt.Printf("Bolt ultimate shear strength = %.0f [lbf]\n", Psu_b)
}

// demoCmd represents the demo command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Runs a demo solve",
	Long:  `Long description TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		Demo()
	},
}

func init() {
	rootCmd.AddCommand(demoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// demoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// demoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

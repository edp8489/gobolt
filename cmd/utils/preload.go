/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// preloadCmd represents the preload command
var preloadCmd = &cobra.Command{
	Use:   "preload",
	Short: "Calculate preload for a bolted joint",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("preload called")
		//nutSelect()
		//uncertaintySelect()
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
// 2. Diameter (float prompt)
// 3. Nominal torque (float prompt)
// 4. Application tolerance (float prompt)
// 5. Nut factor (select)
type SelectKVPair struct {
	Name  string
	Value float64
}

var nutFactorList = []SelectKVPair{
	{Name: "Black oxide finish", Value: 0.30},
	{Name: "Zinc plated", Value: 0.20},
	{Name: "Lubricated", Value: 0.18},
	{Name: "Cadmium plated", Value: 0.16},
	{Name: "Anti-seize compound", Value: 0.12},
}

var nutFacTemplates = &promptui.SelectTemplates{
	Label:    "{{ . }}",
	Active:   "{{ .Name }} ({{ .Value }})",
	Inactive: "{{ .Name | faint }} ({{ .Value | faint }})",
	Selected: "{{ .Name | bold }} ({{ .Value | bold }})",
}

func nutSelect() int {
	prompt := promptui.Select{
		Label:     "Select bolt condition / lubrication factor",
		Items:     nutFactorList,
		Templates: nutFacTemplates,
	}

	resInd, resStr, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		//return nil
	}

	fmt.Printf("You chose %q\n", resStr)
	return resInd
}

// 6. Application method uncertainty (select)
var uncertaintyList = []SelectKVPair{
	{Name: "Torque wrench on unlubricated bolts", Value: 0.35},
	{Name: "Torque wrench on lubricated bolts", Value: 0.25},
	{Name: "Bolt elongation measurement", Value: 0.10},
}

var uncertaintyTemplates = &promptui.SelectTemplates{
	Label:    "{{ . }}",
	Active:   "{{ .Name }} ({{ .Value }})",
	Inactive: "{{ .Name | faint }} ({{ .Value | faint }})",
	Selected: "{{ .Name | bold }} ({{ .Value | bold }})",
}

func uncertaintySelect() int {
	prompt := promptui.Select{
		Label:     "Torque application method",
		Items:     uncertaintyList,
		Templates: uncertaintyTemplates,
	}

	resInd, resStr, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		//return nil
	}

	fmt.Printf("You chose %q\n", resStr)
	return resInd
}

// TODO
// calculate nominal preload using (3) and (5)
// calculate min and max preload using (4) and (6)
// print to console

/*
Copyright © 2022 Eric Peters <e_peters@alum.mit.edu>

*/
package utils

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// areaCmd represents the area command
var areaCmd = &cobra.Command{
	Use:   "area",
	Short: "Calculate the circular area for a given diameter",
	Long:  `Calculate the circular area for a given diameter`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("area called")
		createPrompt()
	},
}

func init() {
	UtilsPal.AddCommand(areaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// areaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// areaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// define a type to facilitate prompt creation
// TODO refactor into global helper package
type promptContent struct {
	errorMsg string
	label    string
}

// define a custom validation function type to facilitate prompt creation
// TODO refactor into global helper package
type validateFunc func(string, string) error

// style template for prompts
// TODO refactor into global helper package
var inputTemplates = &promptui.PromptTemplates{
	Prompt:  "{{ . }} ",
	Valid:   "{{ . | green }} ",
	Invalid: "{{ . | red }} ",
	Success: "{{ . | bold }} ",
}

// validation function for diameter prompt
func diamValidate(input string, errorMsg string) error {
	in, err := strconv.ParseFloat(input, 64)
	if err != nil || in < 1e-9 {
		return errors.New(errorMsg)
	}
	return nil
}

// validation function for unit prompt
func unitValidate(input string, errorMsg string) error {
	if len(input) <= 0 {
		return errors.New(errorMsg)
	}
	return nil
}

// Create a promptui input that returns a float value
// TODO refactor into a global helper package
func promptGetFloatInput(pc promptContent, vf validateFunc) float64 {
	validate := func(input string) error { return vf(input, pc.errorMsg) }

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: inputTemplates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	//fmt.Printf("Input: %v\n", result)
	fResult, _ := strconv.ParseFloat(result, 64)

	return fResult
}

// Create a promptui input that returns a string value
// TODO refactor into a global helper package
func promptGetStrInput(pc promptContent, vf validateFunc) string {
	validate := func(input string) error { return vf(input, pc.errorMsg) }

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: inputTemplates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	//fmt.Printf("Input: %v\n", result)

	return result
}

func createPrompt() {
	areaPromptContent := promptContent{
		"Please provide a number > 0",
		"Enter the diameter for the area calculation:",
	}

	unitPromptContent := promptContent{
		"Please define your unit system",
		"Units:",
	}

	diam := promptGetFloatInput(areaPromptContent, diamValidate)
	unit := promptGetStrInput(unitPromptContent, unitValidate)

	area := calcArea(diam)

	fmt.Printf("Diameter %f [%s], Area = %f [%s**2]", diam, unit, area, unit)
}

// TODO refactor to define with other utility functions
func calcArea(d float64) float64 {
	A := math.Pi * math.Pow(d, 2) / 4
	return A
}

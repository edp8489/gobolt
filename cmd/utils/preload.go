/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package utils

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// preloadCmd represents the preload command
var preloadCmd = &cobra.Command{
	Use:   "preload",
	Short: "Calculate preload for a bolted joint",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("preload called")

		// run prompts using go-survey
		err := survey.Ask(prompt_qs, &prompt_ans, survey.WithValidator(survey.Required))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//fmt.Println(preload_answers)
		//fmt.Printf("%s coating: K = %.2f\n", preload_answers.NutFactor, nutFactorMap[preload_answers.NutFactor])
		//fmt.Printf("%s, uncertainty = %.2f\n", preload_answers.Uncertainty, uncertaintyMap[preload_answers.Uncertainty])

		pre_calc := CalcPreload(prompt_ans.Torque,
			prompt_ans.Diameter,
			prompt_ans.Tolerance,
			nutFactorMap[prompt_ans.NutFactor],
			uncertaintyMap[prompt_ans.Uncertainty],
			prompt_ans.Units)

		fmt.Printf("Min preload: %f [%v]\n", pre_calc.Min, pre_calc.Units)
		fmt.Printf("Nominal preload: %f [%v]\n", pre_calc.Nominal, pre_calc.Units)
		fmt.Printf("Max preload: %f [%v]\n", pre_calc.Max, pre_calc.Units)
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

// create promptContent and validationFunc for the following inputs
// 1. Units (select)
// 2. Diameter (float prompt)
// 3. Nominal torque (float prompt)
// 4. Application tolerance (float prompt)
// 5. Nut factor (select)
// 6. Application method uncertainty (select)

// prompt definition using go-survey
var prompt_qs = []*survey.Question{
	{
		Name: "units",
		Prompt: &survey.Select{
			Message: "Select unit system",
			Options: []string{
				"imperial (in, lbf, psi)",
				"metric (mm, N, Pa)",
			},
			//Default: "",
		},
	},
	{
		Name:   "diameter",
		Prompt: &survey.Input{Message: "Input bolt diameter, D [in or mm]"},
	},
	{
		Name:   "nom_torque",
		Prompt: &survey.Input{Message: "Input nominal torque, T [in-lbf or Nm]"},
	},
	{
		Name:   "torque_tol",
		Prompt: &survey.Input{Message: "Input torque tolerance, +/- t [in-lbf or Nm]"},
	},
	{
		Name: "nut_factor",
		Prompt: &survey.Select{
			Message: "Select thread condition / friction factor, K",
			Options: []string{
				"Black oxide finish",
				"Zinc plated",
				"Lubricated",
				"Cadmium plated",
				"Anti-seize compound",
			},
			//Default: "",
		},
	},
	{
		Name: "uncertainty_factor",
		Prompt: &survey.Select{
			Message: "Select torque measurement method",
			Options: []string{
				"Torque wrench on unlubricated bolts",
				"Torque wrench on lubricated bolts",
				"Bolt elongation measurement",
			},
			//Default: "",
		},
	},
}

// define struct to hold answers
var prompt_ans = struct {
	Units       string  // survey will match the question and field names
	Diameter    float64 // if the types don't match, survey will convert it
	Torque      float64 `survey:"nom_torque"`
	Tolerance   float64 `survey:"torque_tol"`
	NutFactor   string  `survey:"nut_factor"`
	Uncertainty string  `survey:"uncertainty_factor"`
}{}

// create map of select key-value pairs
var nutFactorMap = map[string]float64{
	"Black oxide finish":  0.30,
	"Zinc plated":         0.20,
	"Lubricated":          0.18,
	"Cadmium plated":      0.16,
	"Anti-seize compound": 0.12,
}

var uncertaintyMap = map[string]float64{
	"Torque wrench on unlubricated bolts": 0.35,
	"Torque wrench on lubricated bolts":   0.25,
	"Bolt elongation measurement":         0.10,
}

// calculate nominal preload using (3) and (5)
// calculate min and max preload using (4) and (6)
// print to console

type preload struct {
	Nominal float64
	Min     float64
	Max     float64
	Units   string
}

func CalcPreload(tq float64, dia float64, tqtol float64, k float64, u float64, units string) preload {
	// calculate torque tolerance factors
	cmin := (tq - tqtol) / tq
	cmax := (tq + tqtol) / tq

	units = strings.Split(units, " ")[0]

	var tq_units string

	switch units {
	case "imperial":
		tq_units = "in-lbf"
	case "metric":
		dia = dia / 1000
		tq_units = "Nm"
	}

	// calculate nominal preload
	var p0 float64 = tq / (k * dia)

	// calculate minimum/maximum preload
	p0min := cmin * (1 - u) * p0
	p0max := cmax * (1 + u) * p0

	return preload{
		Nominal: p0,
		Min:     p0min,
		Max:     p0max,
		Units:   tq_units,
	}
}

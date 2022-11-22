/*
Copyright © 2022 Eric Peters <ericdpeters@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/edp8489/gobolt/cmd/utils"

	"github.com/spf13/cobra"
)

// var cfgFile string
var version = "0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gobolt",
	Version: version,
	Short:   "A bolted joint strength calculator",
	Long: `
goBolt is a bolted joint strength calculator.

This application can be run several ways:

1. Batch mode to solve predefined input file(s)
2. Interactive mode that walks the user through
   all aspects of property definition.
3. Individual calculations via the standalone utils palette
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()

		var prompt_a = struct {
			Command string
		}{}

		var prompt_q = []*survey.Question{
			{
				Name: "command",
				Prompt: &survey.Select{
					Message: "Select a command",
					Options: []string{
						"utils",
						"demo",
					},
					//Default: "",
				},
			},
		}

		err := survey.Ask(prompt_q, &prompt_a, survey.WithValidator(survey.Required))
		if err != nil {
			fmt.Println(err.Error())
		}

		cmdSel := prompt_a.Command

		switch cmdSel {
		case "demo":
			fmt.Printf("Sorry, demo program still under development")
		case "utils":
			utils.UtilsPal.Help()
			utils.UtilsPal.Run(cmd, args)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	//rootCmd.AddCommand(batch.BatchCmd)
	//rootCmd.AddCommand(interactive.InteractiveCmd)
	rootCmd.AddCommand(utils.UtilsPal)
}

func init() {
	//cobra.OnInitialize(initConfig)
	addSubcommandPalettes()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gobolt.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
/*
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gobolt" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".gobolt")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
*/

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"jinx/src/utils"
	jinxtypes "jinx/types"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "jinx",
	Short:   "An application for managing jamandbees/jinkies containers",
	Long:    `Start a jinx with jinx serve.`,
	Version: "0.0.3",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//Sets up the global configuration struct by first setting the runtime to our known sane defaults,
//then checking to see whether the user supplies a config file option on the command line. If so,
//overwrite the global configuration struct with the config file contents.
//Also checks for the default jinx.yml config file. If it exists in the same directory as the calling process,
//it is used.
func SetupGlobalConfig() jinxtypes.JinxData {
	configFilePath := "jinx.yml"
	configFileDefault := ""
	configFileOption := configFileDefault

	JinxRuntime := jinxtypes.JinxData{PullImages: true, ContainerName: "jinkies"}

	serveCmd.PersistentFlags().StringVarP(&configFileOption, "jinxconfig", "j", configFileOption, "Path to config file for jinx global options")

	//Check if the option flag was passed in and read its value if need be to override the default.
	configOptionPassedIn := configFileOption != configFileDefault

	if configOptionPassedIn {
		configFilePath = configFileOption
	}

	if _, err := os.Open(configFilePath); err == nil {
		utils.HydrateFromConfig(configFilePath, &JinxRuntime)
	}

	if configOptionPassedIn {
		//The end user asked us to verify the existence of a file by giving us the explicit expectation
		//that this filename exists. The file not existing is an error.
		fmt.Fprintf(os.Stderr, "%v does not exist", configFilePath)
	} else {
		//The jinx.yaml file does not exist, but we only checked as a convenience for the user.
		//We can safely proceed with the default Jinx Runtime.
	}

	return JinxRuntime
}

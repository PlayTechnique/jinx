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
	Use:   "jinx",
	Short: "An application for managing jamandbees/jinkies containers",
	Long: `jinx is a workflow tool for building and managing jenkins based build systems.
In the same way that ruby on rails promotes a strong opinion for how to create, test and serve websites, so jinx has
friendly hints about how to create, test and serve Jenkins based containers.

It also provides tooling to make it easier to figure out how to build and work with Jenkins and plugins in their 
programmatic interface! Check out the Plugins subcommand, for example, which generates a file of dependencies for a
build.gradle file, suitable for getting autocompletion working in Intellij.

# Why Jenkins?
Jenkins is incredibly powerful and configurable, able to start up with either plugin preconfigured or with Jenkins itself
preconfigured. It's rock solid and an uncounted and uncountable number of installations .

# What Workflow?
Create a project, put config files in the project, create a container from the file system, play with the container.

# Can't I just use Dockerfiles to build the config files?
For sure, and jinx is heavily tied to docker and the docker engine for managing this workflow. However, jinx also
provides some nice tooling for generating preconfigured versions of config files you know you need, and may help you
find some that you're not aware exist.

# Can I use jinx with only your project's containers? I already have jenkins containers and don't want to build your stuff.
Sure! Some commands simply execute against running containers to figure things out.
`,
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
func SetupGlobalConfig() jinxtypes.JinxGlobalRuntime {
	configFilePath := "jinx.yml"
	configFileDefault := ""
	configFileOption := configFileDefault

	JinxRuntime := jinxtypes.JinxGlobalRuntime{PullImages: true, ContainerName: "jinkies"}

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

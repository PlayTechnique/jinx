/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	jinxtypes "jinx/types"
	"log"
	"os"
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

func SetupGlobalConfig(location jinxtypes.ConfigFileLocation) (jinxtypes.JinxGlobalRuntime, error) {

	var globalRuntime jinxtypes.JinxGlobalRuntime

	data, err := os.ReadFile(location.ConfigFilePath)
	if err != nil {
		log.Fatal(err)
		return globalRuntime, err
	}

	err = yaml.Unmarshal(data, &globalRuntime)

	if err != nil {
		log.Fatal(err)
		return globalRuntime, err
	}

	return globalRuntime, err
}

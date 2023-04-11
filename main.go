/*
Copyright © 2022 Gwyn <gwyn@playtechnique.io>
*/
package main

import (
	"jinx/cmd"
	jinxtypes "jinx/types"
	"os"
)

func main() {
	configFile := jinxtypes.ConfigFileLocation{
		ConfigFilePath: "configFiles/jinx.yml",
	}

	_, err := os.Stat(configFile.ConfigFilePath)

	if err != nil {
		// If the config file doesnt exist, only the new subcommand is available
		cmd.RegisterNew()
	} else {
		err = cmd.RegisterPlugins(configFile)

		if err != nil {
			panic(err)
		}

		err = cmd.RegisterServe(configFile)

		if err != nil {
			panic(err)
		}
	}

	cmd.Execute()
}

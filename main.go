/*
Copyright © 2022 Gwyn <gwyn@playtechnique.io>
*/
package main

import (
	"jinx/cmd"
	jinxtypes "jinx/types"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
func main() {
	// Set logger to include line numbers.
	log.SetFlags(log.LstdFlags | log.Llongfile)

	configFile := jinxtypes.ConfigFileLocation{
		ConfigFilePath: "configFiles/jinx.yml",
	}

	_, err := os.Stat(configFile.ConfigFilePath)

	if err != nil {
		// If the config file doesn't exist, only the new subcommand is available
		cmd.RegisterNew()
	} else {
		err = cmd.RegisterPlugins(configFile)

		if err != nil {
			log.Print(err)
		}

		err = cmd.RegisterServe(configFile)

		if err != nil {
			log.Print(err)
		}
	}

	cmd.Execute()
}

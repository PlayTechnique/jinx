/*
Copyright © 2022 Gwyn <gwyn@playtechnique.io>
*/
package main

import (
	"jinx/cmd"
	jinxtypes "jinx/types"
)

func main() {
	configFile := jinxtypes.ConfigFileLocation{
		ConfigFilePath: "configFiles/jinx.yml",
	}

	cmd.RegisterNew()

	var err error

	err = cmd.RegisterPlugins(configFile)

	if err != nil {
		panic(err)
	}

	err = cmd.RegisterServe(configFile)

	if err != nil {
		panic(err)
	}

	cmd.Execute()
}

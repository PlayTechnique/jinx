/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"jinx/cmd"
)

func main() {
	jinxRuntime := cmd.SetupGlobalConfig()

	cmd.RegisterPlugins(jinxRuntime)
	cmd.RegisterServe(jinxRuntime)
	cmd.Execute()
}

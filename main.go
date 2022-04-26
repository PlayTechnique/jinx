/*
Copyright © 2022 Gwyn <gwyn@playtechnique.io>

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

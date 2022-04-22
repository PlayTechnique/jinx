/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"jinx/src/jenkins"
	jinxtypes "jinx/types"
)

type pluginsRuntime struct {
	globalRuntime jinxtypes.JinxData
}

// pluginsCmd represents the plugins command
func (pluginsSetup *pluginsRuntime) pluginsCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "plugins",
		Short: "Retrieve lists of plugins in various formats",
		Long: `I regularly want to freeze the plugins inside a container, so that successive containers can 
be rebuilt with the precise same plugins. However, if you are using the cli tooling for installing a Jenkins plugin then
you only know what dependencies are required after the plugins have been installed.

This command copies jinx itself into your container, where it figures out some shenanigans and returns the plugins.txt`,
		Run: func(cmd *cobra.Command, args []string) {
			jenkins.Plugins(pluginsSetup.globalRuntime)
		},
	}
}

func RegisterPlugins(jinxRunTime jinxtypes.JinxData) {
	config := pluginsRuntime{globalRuntime: jinxRunTime}

	rootCmd.AddCommand(config.pluginsCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pluginsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

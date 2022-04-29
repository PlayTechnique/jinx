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
	GlobalRuntime jinxtypes.JinxData

	TopLevelOutDir string
	RemovePlugins  bool
	OutputFormat   string
}

func (pluginsRuntime *pluginsRuntime) PluginsCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "plugins",
		Short: "Retrieve lists of plugins in various formats",
		Long: `I regularly want to freeze the plugins inside a container, so that successive containers can
be rebuilt with the precise same plugins. However, if you are using the cli tooling for installing a Jenkins plugin then
you only know what dependencies are required after the plugins have been installed.

This command copies plugins from the container into a temporary directory, and uses the copied files to generate the
plugins information you request. This temporary directory is auto-deleted after use.

If you specify a specify an output directory, that directory is not auto-deleted after use.

I also regularly use intellij to update jinkies (https://github.com/playtechnique/jinkies), where I use a build.gradle
file to control autocompletion in the IDE. The plugins command can output the right build.gradle format for the 
dependencies section of a build.gradle file, so that IDEs can autocomplete functionality for the plugins that you have.
`,
		Run: func(cmd *cobra.Command, args []string) {
			jenkins.Plugins(pluginsRuntime.GlobalRuntime, pluginsRuntime.TopLevelOutDir, pluginsRuntime.OutputFormat)
		},
	}
}

func RegisterPlugins(jinxRunTime jinxtypes.JinxData) {
	config := pluginsRuntime{GlobalRuntime: jinxRunTime, TopLevelOutDir: "", RemovePlugins: true, OutputFormat: "plugins.txt"}

	// Go invokes functions at the last possible second, so if we try to do:
	// rootCmd.AddCommand(config.PluginsCmd()) and config.PluginsCmd().Flags()...
	// then cobra receives a function to invoke, not a pointer to a cobra command, and adding the
	// help flags doesn't work. Assigning the function to a variable, though, invokes the function
	// when the variable is constructed and the variable contains the pointer, not a function.
	commander := config.PluginsCmd()

	rootCmd.AddCommand(commander)

	commander.Flags().StringVar(&config.TopLevelOutDir, "outputdir", "", "Directory to copy your plugins into.")
	commander.Flags().StringVar(&config.OutputFormat, "format", "plugins.txt", "Display format for your plugins output (plugins.txt or build.gradle)")
}

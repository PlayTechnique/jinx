package cmd

import (
	"github.com/spf13/cobra"
	"jinx/src/jinkiesengine"
)

type JinxData struct {
	containerName string
	pullImages    bool

	containerConfigPath string
	hostConfigPath      string
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Subcommands to allow you to start or stop an unconfigured jinkies",
	Long: `Why would you want an unconfigured instance of jinkies? Any time you want a jenkins instance
quickly for reasons unrelated to a specific job. Maybe you want to prototype some jcasc settings or something.

Maybe you want two instances of jinkies running at once? Use the -c flag to supply an environment variables file. To
write a blank version of this file, see the 'jinx containerconfig' subcommand.
`,
}

func (jinxData *JinxData) startSubCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start jinkies!",
		Long:  `Starts the unconfigured jinkies container`,
		Run: func(cmd *cobra.Command, args []string) {
			jinkiesengine.RunRunRun(jinxData.containerName, jinxData.pullImages, jinxData.containerConfigPath, jinxData.hostConfigPath)
		},
	}
}

func (jinxData *JinxData) stopSubCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stops your jinkies container_info.",
		Long:  `No configuration is retained after a stop, so this gets you back to a clean slate.`,
		Run: func(cmd *cobra.Command, args []string) {
			jinkiesengine.StopGirl(jinxData.containerName)
		},
	}
}

func init() {

	var jinxRuntime = JinxData{containerName: "jinkies", pullImages: true}

	rootCmd.AddCommand(serveCmd)
	serveCmd.AddCommand(jinxRuntime.startSubCommand())
	serveCmd.AddCommand(jinxRuntime.stopSubCommand())

	serveCmd.PersistentFlags().StringVarP(&jinxRuntime.containerConfigPath, "containerconfig", "c", "", "Path to config file describing your container")
	serveCmd.PersistentFlags().StringVarP(&jinxRuntime.hostConfigPath, "hostconfig", "o", "", "Path to config file describing your container host ")
}

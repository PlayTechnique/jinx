package cmd

import (
	"github.com/spf13/cobra"
	"jinx/src/jinkiesengine"
	"jinx/src/utils"
	"os"
)

type JinxData struct {
	ContainerName string
	PullImages    bool

	containerConfigPath string
	hostConfigPath      string
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Subcommands to allow you to start or stop a jinkies.",
	Long: `Why would you want an unconfigured instance of jinkies? Any time you want a jenkins instance
quickly for reasons unrelated to a specific job. Maybe you want to prototype some jcasc settings or something.

Maybe you want two instances of jinkies running at once? Use the -o flag to supply a yaml file overriding the hostconfig
(https://pkg.go.dev/github.com/docker/docker@v20.10.13+incompatible/api/types/container#HostConfig),
or use -c to supply a yaml file overriding the container config (https://pkg.go.dev/github.com/docker/docker@v20.10.13+incompatible/api/types/container#Config).
`,
}

func (jinxData *JinxData) startSubCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start jinkies!",
		Long:  `Starts the unconfigured jinkies container`,
		Run: func(cmd *cobra.Command, args []string) {
			jinkiesengine.RunRunRun(jinxData.ContainerName, jinxData.PullImages, jinxData.containerConfigPath, jinxData.hostConfigPath)
		},
	}
}

func (jinxData *JinxData) stopSubCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stops your jinkies container_info.",
		Long:  `No configuration is retained after a stop, so this gets you back to a clean slate.`,
		Run: func(cmd *cobra.Command, args []string) {
			jinkiesengine.StopGirl(jinxData.ContainerName)
		},
	}
}

func init() {
	jinxRuntime := JinxData{ContainerName: "jinkies", PullImages: true}
	defaultConfigFile := "jinx.yml"

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(serveCmd)
	serveCmd.AddCommand(jinxRuntime.startSubCommand())
	serveCmd.AddCommand(jinxRuntime.stopSubCommand())

	if _, err := os.Stat(defaultConfigFile); err == nil {
		utils.HydrateFromConfig(".jinx.yml", &jinxRuntime)
	}

	serveCmd.PersistentFlags().StringVarP(&jinxRuntime.containerConfigPath, "containerconfig", "c", "", "Path to config file describing your container")
	serveCmd.PersistentFlags().StringVarP(&jinxRuntime.hostConfigPath, "hostconfig", "o", "", "Path to config file describing your container host ")
}

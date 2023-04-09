package cmd

import (
	"github.com/spf13/cobra"
	"jinx/src/jinkiesengine"
	jinxtypes "jinx/types"
)

type ServeRuntime struct {
	GlobalRuntime jinxtypes.JinxGlobalRuntime

	ContainerConfigPath string
	HostConfigPath      string
}

// serveCmd represents the serve command.
// It's just a namespace for additional subcommands.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Subcommands to allow you to start or stop a jinkies.",
	Long: `Why would you want an unconfigured instance of jinkies? Any time you want a jenkins instance
quickly for reasons unrelated to a specific job. Maybe you want to prototype some jcasc settings or something.

Maybe you want two instances of jinkies running at once? Use the -o flag to supply a yaml file overriding the hostconfig
(https://pkg.go.dev/github.com/docker/docker@v20.10.14+incompatible/api/types/container#HostConfig),
or use -c to supply a yaml file overriding the container config (https://pkg.go.dev/github.com/docker/docker@v20.10.13+incompatible/api/types/container#Config).
`,
}

func (server *ServeRuntime) startSubCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start jinkies!",
		Long:  `Starts the jinkies container`,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := jinkiesengine.RunRunRun(server.GlobalRuntime.ContainerName, server.GlobalRuntime.PullImages, server.ContainerConfigPath, server.HostConfigPath)
			return err
		},
	}
}

func (server *ServeRuntime) stopSubCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stops your jinkies container_info.",
		Long:  `No configuration is retained after a stop, so this gets you back to a clean slate.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return jinkiesengine.StopGirl(server.GlobalRuntime.ContainerName)
		},
	}
}

func RegisterServe(configFile jinxtypes.ConfigFileLocation) error {
	jinxRunTime, err := SetupGlobalConfig(configFile)

	if err != nil {
		return err
	}

	config := ServeRuntime{GlobalRuntime: jinxRunTime}

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(serveCmd)

	serveCmd.AddCommand(config.startSubCommand())
	serveCmd.AddCommand(config.stopSubCommand())

	serveCmd.PersistentFlags().StringVarP(&config.ContainerConfigPath, "containerconfig", "c", "", "Path to config file describing your container")
	serveCmd.PersistentFlags().StringVarP(&config.HostConfigPath, "hostconfig", "o", "", "Path to config file describing your container host ")

	serveCmd.Flags().StringVarP(&config.HostConfigPath, "jenkinsfile", "e", "", "Path on the host to a Jenkinsfile to use as a seed job")

	return nil
}

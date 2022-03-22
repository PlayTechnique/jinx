package cmd

import (
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"jinx/src/jinkiesengine"
	"os"
)

type WorkflowData struct {
	containerName string
	pullImages    bool

	containerConfig container.Config
	hostConfig      container.HostConfig
}

func hydrateFromConfig[T any](configPath string, config *T) {

	viper.AddConfigPath("./")
	viper.SetConfigType("yml")
	viper.SetConfigName(configPath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	viper.Unmarshal(&config)
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

func (metadata WorkflowData) startSubCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start jinkies!",
		Long:  `Starts the unconfigured jinkies container`,
		Run: func(cmd *cobra.Command, args []string) {
			jinkiesengine.RunRunRun(metadata.containerName, metadata.pullImages, metadata.containerConfig, metadata.hostConfig)
		},
	}
}

func (metadata WorkflowData) stopSubCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stops your jinkies container_info.",
		Long:  `No configuration is retained after a stop, so this gets you back to a clean slate.`,
		Run: func(cmd *cobra.Command, args []string) {
			jinkiesengine.StopGirl(metadata.containerName)
		},
	}
}

func init() {
	var (
		genericHostConfig      container.HostConfig
		genericContainerConfig container.Config

		containerConfigPath string
		hostConfigPath      string
	)

	genericContainerConfig = container.Config{
		ExposedPorts: nat.PortSet{"8090/tcp": {}},
		Image:        "jamandbees/jinkies",
	}

	genericHostConfig = container.HostConfig{
		AutoRemove:   true,
		PortBindings: nat.PortMap{"8080/tcp": {{HostIP: "0.0.0.0", HostPort: "8090/tcp"}}},
	}

	var foo = WorkflowData{containerName: "jinkies", pullImages: true, containerConfig: genericContainerConfig, hostConfig: genericHostConfig}

	rootCmd.AddCommand(serveCmd)
	serveCmd.AddCommand(foo.startSubCommand())
	serveCmd.AddCommand(foo.stopSubCommand())

	serveCmd.PersistentFlags().StringVarP(&containerConfigPath, "containerconfig", "c", "", "Path to config file describing your container")
	serveCmd.PersistentFlags().StringVarP(&hostConfigPath, "hostconfig", "o", "", "Path to config file describing your container host ")

	fmt.Printf("%v, %v\n", containerConfigPath, hostConfigPath)

	if containerConfigPath != "" {
		hydrateFromConfig(containerConfigPath, &foo.containerConfig)
	} else {
		foo.containerConfig = genericContainerConfig
	}

	if hostConfigPath != "" {
		hydrateFromConfig(containerConfigPath, &foo.hostConfig)
	} else {
		foo.hostConfig = genericHostConfig
	}
}

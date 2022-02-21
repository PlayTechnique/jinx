package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"jinx/src/jinkiesengine"
	"os"
)

var (
	containerConfigPath = ""
)

func hydrateFromConfig(configPath string) jinkiesengine.ContainerInfo {
	var config jinkiesengine.ContainerInfo

	if configPath == "" {
		config.AutoRemove = true
		config.ImageName = "jamandbees/jinkies"
		config.ContainerName = "jinkies"
		config.ContainerPort = "8080/tcp"
		config.HostIp = "0.0.0.0"
		config.HostPort = "8090/tcp"
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigType("env")
		viper.SetConfigName(configPath)

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}
		viper.Unmarshal(&config)
	}

	return config
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Subcommands to allow you to start or stop an unconfigured jinkies",
	Long: `Why would you want an unconfigured instance of jinkies? Any time you want a jenkins instance
quickly for reasons unrelated to a specific job. Maybe you want to prototype some jcasc settings or something.`,
}

var startSubCmd = &cobra.Command{
	Use:   "start",
	Short: "start jinkies!",
	Long:  `Starts the unconfigured jinkies container`,
	Run: func(cmd *cobra.Command, args []string) {
		jinkiesengine.RunRunRun(hydrateFromConfig(containerConfigPath))
	},
}

var stopSubCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops your jinkies container_info.",
	Long:  `No configuration is retained after a stop, so this gets you back to a clean slate.`,
	Run: func(cmd *cobra.Command, args []string) {
		jinkiesengine.StopGirl(hydrateFromConfig(containerConfigPath))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.AddCommand(startSubCmd)
	serveCmd.AddCommand(stopSubCmd)

	serveCmd.PersistentFlags().StringVarP(&containerConfigPath, "containerconfig", "c", "", "Path config file describing your container")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

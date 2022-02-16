package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jinx/src/jinkiesengine"
	"os"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Subcommands to allow you to start or stop an unconfigured jinkies",
	Long: `Why would you want an unconfigured instance of jinkies? Any time you want a jenkins instance
quickly for reasons unrelated to a specific job. Maybe you want to prototype some jcasc settings or something.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {

		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Provide a subcommand or run with --help")
			os.Exit(1)
		}
	},
}

var startSubCmd = &cobra.Command{
	Use:   "start",
	Short: "start jinkies!",
	Long:  `Starts the unconfigured jinkies container`,
	Run: func(cmd *cobra.Command, args []string) {
		jinkiesengine.RunRunRun()
	},
}

var stopSubCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops your jinkies container.",
	Long:  `No configuration is retained after a stop, so this gets you back to a clean slate.`,
	Run: func(cmd *cobra.Command, args []string) {
		jinkiesengine.StopGirl()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.AddCommand(startSubCmd)
	serveCmd.AddCommand(stopSubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

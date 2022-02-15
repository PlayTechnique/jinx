package cmd

import (
	"github.com/spf13/cobra"
	"jinx/src/api"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start your jenkins",
	Long:  `jinx serve will get you an unconfigured jenkins. Sweet!`,
	Run: func(cmd *cobra.Command, args []string) {
		api.RunRunRun()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

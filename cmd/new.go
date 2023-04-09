package cmd

import (
	"github.com/spf13/cobra"
	"jinx/src/jinkiesengine"
)

type CreateLayoutRuntime struct {
}

// newCmd represents the new command
func (createLayout *CreateLayoutRuntime) newCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Create a new directory and config file layout for a jinkies!",
		Args:  cobra.ExactArgs(1),

		Long: `Run 'jinx new <dir>' to set up the skeleton of a new jinx project in a new starting dir! Use 'jinx new .' 
to use the current dir.

There are a lot of moving parts to getting Jenkins configured programatically, including init.groovy.d files,
a Dockerfile, build script stuff, the whole nine yards.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			containerName, _ := cmd.Flags().GetString("--container-name")

			_, _, err := jinkiesengine.Initialise(containerName, args[0])
			return err
		},
	}
}

func RegisterNew() {
	layout := CreateLayoutRuntime{}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.Flags().String("--container-name", "jinkies", "Set a custom container name for the docker build script.")
	rootCmd.AddCommand(layout.newCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

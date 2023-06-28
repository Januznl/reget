package cmd

import (
	"github.com/spf13/cobra"
)

var (
	verbose bool

	rootCmd = &cobra.Command{
		Use:   "reget [-v]",
		Short: "Download (latest) release from online SVN repos.",
		Long:  "Download releases from online SVN repos, you can get the latest, set a patch, minor or major version",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(githubCmd)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

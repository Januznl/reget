package cmd

import (
	"github.com/spf13/cobra"
)

var (
	verbose bool
	version = "dev"

	rootCmd = &cobra.Command{
		Short:   "Download (latest) release from online SVN repos.",
		Long:    "Download releases from online SVN repos, you can get the latest, set a patch, minor or major version",
		Version: version,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(githubCmd)
	rootCmd.AddCommand(peclCmd)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

package cmd

import (
	"fmt"
	"os"
	"reget/download"
	"reget/github"
	"runtime"

	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Download releases from Github",
	Args:  cobra.ExactArgs(1),
	Run:   githubs,
}

var flagArch string
var flagUseTags bool
var flagRelease string
var flagPinnedRelease string
var flagOutputFileName string

func init() {
	osArch := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)
	githubCmd.Flags().BoolVarP(&flagUseTags, "use-tag", "t", false, "use tags is no releases were made")
	githubCmd.Flags().StringVarP(&flagArch, "architecture", "a", osArch, "arch for example 'darwin_x64'")
	githubCmd.Flags().StringVarP(&flagRelease, "release", "r", "latest", "Release ex 1.2.1")
	githubCmd.Flags().StringVarP(&flagPinnedRelease, "pinned-release", "p", "", "Pinned release ex 1.2.0, will download 1.2.23 but not 1.3.0. Pinned release 1 will download 1.2.0 but also 1.5.0, but not 2.0.0")
	githubCmd.Flags().StringVarP(&flagOutputFileName, "output", "o", "", "local filename to save to")
}

func githubs(cmd *cobra.Command, args []string) {
	if flagUseTags {
		if downloadUrl, err := github.GetTag(args[0], flagArch, flagRelease, flagPinnedRelease); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		} else {
			download.DownloadRelease(downloadUrl, flagOutputFileName)
		}

	} else {
		if downloadUrl, err := github.GetRelease(args[0], flagArch, flagRelease, flagPinnedRelease); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		} else {
			download.DownloadRelease(downloadUrl, flagOutputFileName)
		}
	}
}

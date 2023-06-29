package cmd

import (
	"fmt"
	"os"
	"reget/download"
	"reget/pecl"

	"github.com/spf13/cobra"
)

var peclCmd = &cobra.Command{
	Use:   "pecl",
	Short: "Download releases from pecl",
	Args:  cobra.ExactArgs(1),
	Run:   pecls,
}

var flagPeclOutputFileName string

func init() {
	peclCmd.Flags().StringVarP(&flagPeclOutputFileName, "output", "o", "", "local filename to save to")
	peclCmd.Flags().StringVarP(&flagPinnedRelease, "pinned-release", "p", "", "Pinned release ex 1.2.0, will download 1.2.23 but not 1.3.0. Pinned release 1 will download 1.2.0 but also 1.5.0, but not 2.0.0")

}

func pecls(cmd *cobra.Command, args []string) {
	if downloadUrl, err := pecl.GetRelease(args[0], flagRelease, flagPinnedRelease); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		download.DownloadRelease(downloadUrl, flagOutputFileName)
	}
}

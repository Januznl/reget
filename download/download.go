package download

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

func DownloadRelease(releaseUrl string, localFileName string) {
	req, _ := http.NewRequest("GET", releaseUrl, nil)
	resp, respError := http.DefaultClient.Do(req)

	if respError != nil {
		fmt.Printf("Unable to do request, error: %s", respError.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	if localFileName == "" {
		if _, params, err := mime.ParseMediaType(resp.Header.Get("content-disposition")); err != nil {
			localFileName = "download.tar.gz"
		} else {
			localFileName = params["filename"]
		}
	}

	fmt.Printf("Downloading file %s\n", localFileName)

	f, err := os.OpenFile(localFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Unable to open file for writing, error: %s", err.Error())
		os.Exit(1)
	}

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"",
	)

	if _, err := io.Copy(io.MultiWriter(f, bar), resp.Body); err != nil {
		fmt.Printf("Unable to write to file, error: %s", err.Error())
		os.Exit(1)
	}
}

package download

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

func DownloadRelease(releaseUrl string, localFileName string) {
	req, _ := http.NewRequest("GET", releaseUrl, nil)
	resp, respError := http.DefaultClient.Do(req)

	if respError != nil {
		log.Fatalf("Unable to do request, error: %s\n", respError.Error())
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
		log.Fatalf("Unable to open file for writing, error: %s", err.Error())
	}

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"",
	)

	if _, err := io.Copy(io.MultiWriter(f, bar), resp.Body); err != nil {
		log.Fatalf("Unable to write to file, error: %s", err.Error())
	}
}

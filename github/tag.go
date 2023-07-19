package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/januznl/reget/compare"
)

type Tag struct {
	Name       string `json:"name"`
	TarBallUrl string `json:"tarball_url"`
}

func GetTag(url string, arch string, release string, pinnedRelease string) (string, error) {
	var apiUrl = fmt.Sprintf("https://api.github.com/repos/%s/tags", url)
	res, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}

	if res.StatusCode == 404 {
		return "", errors.New("could not find given repo")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var apiTags []Tag
	if err := json.Unmarshal(body, &apiTags); err != nil {
		return "", err
	}

	log.Printf("architecture:   %s\n", arch)
	log.Printf("Pinned Version: %s\n", pinnedRelease)

	for _, apiTag := range apiTags {
		log.Println(apiTag.Name)
		if pinnedRelease != "" {
			if compare.CompareReleases(pinnedRelease, apiTag.Name) {
				return apiTag.TarBallUrl, nil
			}
		} else {
			if release != "latest" {
				if compare.CompareEqualReleases(release, apiTag.Name) {
					return apiTag.TarBallUrl, nil
				}
			} else {
				return apiTag.TarBallUrl, nil
			}
		}
	}
	return "", errors.New("cannot match any download for given release")
}

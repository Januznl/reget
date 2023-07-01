package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reget/compare"
	"strings"
)

type Release struct {
	Name       string `json:"name"`
	TarBallUrl string `json:"tarball_url"`
	Assets     []Assets
}

type Assets struct {
	BrowserDownloadUrl string `json:"browser_download_url"`
}

func GetRelease(url string, arch string, release string, pinnedRelease string) (string, error) {
	var apiUrl = fmt.Sprintf("https://api.github.com/repos/%s/releases", url)
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

	var apiReleases []Release
	if err := json.Unmarshal(body, &apiReleases); err != nil {
		return "", err
	}

	fmt.Printf("architecture:   %s\n", arch)
	fmt.Printf("Pinned Version: %s\n", pinnedRelease)

	for _, apiRelease := range apiReleases {
		fmt.Println(apiRelease.Name)

		if pinnedRelease != "" {
			if compare.CompareReleases(pinnedRelease, apiRelease.Name) {
				return findReleaseUrl(apiRelease.Assets, arch, apiRelease.TarBallUrl)
			}
		} else {
			if release != "latest" {
				if compare.CompareEqualReleases(release, apiRelease.Name) {
					return findReleaseUrl(apiRelease.Assets, arch, apiRelease.TarBallUrl)
				}
			} else {
				return findReleaseUrl(apiRelease.Assets, arch, apiRelease.TarBallUrl)
			}
		}

		// if apiRelease.TarBallUrl != "" {
		// 	return apiRelease.TarBallUrl, nil
		// }
	}
	return "", errors.New("cannot match any download for given release")
}

func findReleaseUrl(assets []Assets, arch string, releaseTarballUrl string) (string, error) {
	url, err := findAssetForArch(assets, arch)

	if err != nil {
		if releaseTarballUrl != "" {
			return releaseTarballUrl, nil
		} else {
			return "", errors.New("unable to get asset for arch and no release tarball url found in release")
		}
	}

	return url, nil
}

func findAssetForArch(assets []Assets, arch string) (string, error) {
	for _, asset := range assets {
		if strings.Contains(asset.BrowserDownloadUrl, arch) {
			return asset.BrowserDownloadUrl, nil
		}
	}

	return "", errors.New("cannot find any download with given arch")
}

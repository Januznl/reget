package pecl

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reget/compare"
)

type Releases struct {
	Release []Release `xml:"r"`
}
type Release struct {
	Name string `xml:"v"`
	Type string `xml:"s"`
}

func GetRelease(url string, release string, pinnedRelease string) (string, error) {
	var apiUrl = fmt.Sprintf("https://pecl.php.net/rest/r/%s/allreleases.xml", url)
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

	var apiReleases Releases
	if err := xml.Unmarshal(body, &apiReleases); err != nil {
		return "", err
	}

	fmt.Printf("Pinned Version: %s\n", pinnedRelease)

	for _, apiRelease := range apiReleases.Release {
		fmt.Println(apiRelease.Name)

		if pinnedRelease != "" {
			if compare.CompareReleases(pinnedRelease, apiRelease.Name) {
				return fmt.Sprintf("https://pecl.php.net/get/%s-%s", url, apiRelease.Name), nil
			}
		} else {
			if release != "latest" {
				if compare.CompareEqualReleases(release, apiRelease.Name) {
					return fmt.Sprintf("https://pecl.php.net/get/%s-%s", url, apiRelease.Name), nil
				}
			} else {
				return fmt.Sprintf("https://pecl.php.net/get/%s-%s", url, apiRelease.Name), nil
			}
		}
	}
	return "", errors.New("cannot match any download for given release")
}

package pecl

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"reget/compare"
	"strings"

	"golang.org/x/mod/semver"
)

type Releases struct {
	Release []Release `xml:"r"`
}
type Release struct {
	Name string `xml:"v"`
	Type string `xml:"s"`
}

func GetRelease(url string, release string, pinnedRelease string) (string, error) {
	pinnedRelease = compare.NormalizeSemVer(pinnedRelease)

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

	var arrReleases []string
	for _, apiRelease := range apiReleases.Release {
		var semversion string
		if strings.HasPrefix(apiRelease.Name, "v") {
			semversion = apiRelease.Name
		} else {
			semversion = fmt.Sprintf("v%s", apiRelease.Name)
		}
		if semver.IsValid(semversion) {
			arrReleases = append(arrReleases, semversion)
		}
	}
	arrReleases = compare.SortReleases(arrReleases)

	for _, apiRelease := range arrReleases {
		log.Println(apiRelease)

		if pinnedRelease != "" {
			if compare.CompareReleases(pinnedRelease, apiRelease) {
				return fmt.Sprintf("https://pecl.php.net/get/%s-%s", url, getOriginalVersion(apiRelease, apiReleases.Release)), nil
			}
		} else {
			if release != "latest" {
				if compare.CompareEqualReleases(release, apiRelease) {
					return fmt.Sprintf("https://pecl.php.net/get/%s-%s", url, getOriginalVersion(apiRelease, apiReleases.Release)), nil
				}
			} else {
				return fmt.Sprintf("https://pecl.php.net/get/%s-%s", url, getOriginalVersion(apiRelease, apiReleases.Release)), nil
			}
		}
	}
	return "", errors.New("cannot match any download for given release")
}

func getOriginalVersion(version string, originalVersions []Release) string {
	for _, orgVersion := range originalVersions {
		if orgVersion.Name == version {
			return orgVersion.Name
		} else if orgVersion.Name == strings.TrimPrefix(version, "v") {
			return orgVersion.Name
		}
	}
	return version
}

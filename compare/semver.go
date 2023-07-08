package compare

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/mod/semver"
)

func CompareReleases(reqRelease string, foundRelease string) bool {
	reqRelease = NormalizeSemVer(reqRelease)
	foundRelease = NormalizeSemVer(foundRelease)

	log.Printf("Compare release %s with %s \n", reqRelease, foundRelease)

	if semver.Major(reqRelease) == semver.Major(foundRelease) {
		log.Println("found major!")
		if getMinor(reqRelease) == "0" {
			// found major, rest is wildcard
			log.Println("found major, rest is wildcard")
			return true
		} else {
			if getMinor(reqRelease) == getMinor(foundRelease) {
				if getPatch(reqRelease) == "0" {
					// found Major, Minor, patch is wildcard
					log.Println("found Major, Minor, patch is wildcard")
					return true
				} else {
					foundPatch, _ := strconv.Atoi(getPatch(foundRelease))
					reqPatch, _ := strconv.Atoi(getPatch(reqRelease))
					if foundPatch >= reqPatch {
						// found Major, Minor, patch is bigger of equal
						log.Println("found Major, Minor, patch is bigger of equal")
						return true
					}
				}
			}
		}
	}
	log.Println("No match found")
	return false
}

func CompareEqualReleases(reqRelease string, foundRelease string) bool {
	reqRelease = NormalizeSemVer(reqRelease)
	foundRelease = NormalizeSemVer(foundRelease)
	result := semver.Compare(foundRelease, reqRelease)

	log.Printf("Comparing %s with %s, result %d\n", reqRelease, foundRelease, result)

	return result == 0
}

func getMinor(release string) string {
	versionArr := strings.Split(release, ".")
	return versionArr[1]
}

func getPatch(release string) string {
	versionArr := strings.Split(release, ".")
	return versionArr[2]
}

func SortReleases(releases []string) []string {
	// Sort semver array, newest first
	sort.Sort(sort.Reverse(semver.ByVersion(releases)))
	return releases
}

func NormalizeSemVer(version string) string {
	if version != "" {
		if version == "latest" {
			return version
		}

		if !strings.HasPrefix(version, "v") {
			version = fmt.Sprintf("v%s", version)
		}

		switch strings.Count(version, ".") {
		case 0:
			return fmt.Sprintf("%s.0.0", version)
		case 1:
			return fmt.Sprintf("%s.0", version)
		}
	}
	return version
}

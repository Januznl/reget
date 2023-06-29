package compare

import (
	"fmt"
	"strconv"
	"strings"
)

func CompareReleases(reqRelease string, foundRelease string) bool {
	reqVersion := releaseToVersion(reqRelease)
	foundVersion := releaseToVersion(foundRelease)

	if reqVersion.Major == foundVersion.Major {
		if reqVersion.Minor == 0 {
			// found major, rest is wildcard
			return true
		} else {
			if foundVersion.Minor == reqVersion.Minor {
				if reqVersion.Patch == 0 {
					// found Major, Minor, patch is wildcard
					return true
				} else {
					if foundVersion.Patch >= reqVersion.Patch {
						// found Major, Minor, patch is bigger of equal
						return true
					}
				}
			}

		}
	}

	return false
}

func CompareEqualReleases(reqRelease string, foundRelease string) bool {
	reqVersion := releaseToVersion(reqRelease)
	foundVersion := releaseToVersion(foundRelease)
	if reqVersion.Major == foundVersion.Major {
		if reqVersion.Minor == foundVersion.Minor {
			if reqVersion.Patch == foundVersion.Patch {
				return true
			}
		}
	}
	return false
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func releaseToVersion(release string) Version {
	release = normalizeRelease(release)
	versionArr := strings.Split(release, ".")

	major, _ := strconv.Atoi(versionArr[0])
	minor, _ := strconv.Atoi(versionArr[1])
	patch, _ := strconv.Atoi(versionArr[2])

	return Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func normalizeRelease(version string) string {
	if strings.HasPrefix(version, "v") {
		version = strings.Replace(version, "v", "", 1)
	}

	if version != "latest" && version != "" {
		switch strings.Count(version, ".") {
		case 0:
			return fmt.Sprintf("%s.0.0", version)
		case 1:
			return fmt.Sprintf("%s.0", version)
		}
	}
	//fmt.Printf("normalized release: %s\n", version)
	return version
}

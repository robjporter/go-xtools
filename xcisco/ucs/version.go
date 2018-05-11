package ucs

import (
	"fmt"
	"strings"

	"github.com/robjporter/go-utils/web/request"
)

const (
	INDEX_POS_LATEST    = "mdfTree.addElementRel(966, 1, 'Latest', '');"
	INDEX_POS_SUGGESTED = "mdfTree.addElementRel(965, 1, 'Suggested', '');"
	INDEX_POS_RELEASES  = "mdfTree.addElementRel(967, 1, 'All Releases', '');"
	INDEX_POS_DEFERRED  = "mdfTree.addElementRel(968, 1, 'Deferred Releases', '');"
	INDEX_POS_END       = "document.write(mdfTree);"
)

var (
	lastBody      = ""
	latestData    []string
	suggestedData []string
	releasesData  []string
	deferredData  []string
)

// SETUP FUNCTION

func GetWebData() {
	code := request.New()
	resp, body, err := code.Get("https://software.cisco.com/download/release.html?mdfid=283612660&softwareid=283655658").End()
	if resp.StatusCode == 200 && err == nil {
		if body != "" {
			latestData = getReleasesFromContent(getLatestContent(body))
			suggestedData = getReleasesFromContent(getSuggestedContent(body))
			releasesData = getReleasesFromContent(getAllReleasesContent(body))
			deferredData = getReleasesFromContent(getDeferredContent(body))
		}
	}
}

// PUBLIC FUNCTIONS

func GetAllSuggestedReleases() []string {
	suggestedContent := getSuggestedContent(lastBody)
	return getReleasesFromContent(suggestedContent)
}

func GetAllLatestReleases() []string {
	latestContent := getLatestContent(lastBody)
	return getReleasesFromContent(latestContent)
}

func GetAllReleases() []string {
	allContent := getAllReleasesContent(lastBody)
	return getReleasesFromContent(allContent)
}

func GetAllDeferredReleases() []string {
	allContent := getDeferredContent(lastBody)
	return getReleasesFromContent(allContent)
}

func GetSuggestedRelease() string {
	if len(suggestedData) > 0 {
		return strings.TrimSpace(suggestedData[0])
	}
	return ""
}

func GetLatestRelease() string {
	if len(latestData) > 0 {
		return strings.TrimSpace(latestData[0])
	}
	return ""
}

func GetIsDeferredRelease(version string) bool {
	if len(deferredData) > 0 {
		for i := 0; i < len(deferredData); i++ {
			if strings.TrimSpace(deferredData[i]) == strings.TrimSpace(version) {
				return true
			}
		}
	}
	return false
}

func GetLatestReleaseTrain(train string) string {
	if isInSlice(latestData, train) {
		return "This is the latest release."
	} else {
		if len(latestData) > 0 {
			for i := 0; i < len(latestData); i++ {
				if getVersionMajor(latestData[i]) == getVersionMajor(train) {
					if getVersionMinor(latestData[i]) == getVersionMinor(train) {
						if getVersionRelease(latestData[i]) == getVersionRelease(train) {
							return strings.TrimSpace(latestData[i])
						}
						return strings.TrimSpace(latestData[i])
					}
					return strings.TrimSpace(latestData[i])
				}
			}
		}
	}
	return "No Compatible version found"
}

func IsSuggestedReleaseTrain(train string) bool {
	if isInSlice(suggestedData, train) {
		return true
	}
	return false
}

func IsLatestReleaseTrain(train string) bool {
	if isInSlice(latestData, train) {
		return true
	}
	return false
}

func GetSuggestedReleaseTrain(train string) string {
	if isInSlice(suggestedData, train) {
		return "This is a suggested release."
	} else {
		if len(suggestedData) > 0 {
			for i := 0; i < len(suggestedData); i++ {
				if getVersionMajor(suggestedData[i]) == getVersionMajor(train) {
					if getVersionMinor(suggestedData[i]) == getVersionMinor(train) {
						if getVersionRelease(suggestedData[i]) == getVersionRelease(train) {
							return strings.TrimSpace(suggestedData[i])
						}
						return strings.TrimSpace(suggestedData[i])
					}
					return strings.TrimSpace(suggestedData[i])
				}
			}
		}
	}
	return "No Compatible version found"
}

func GetIsSuggestedOrLatest(train string) bool {
	if IsSuggestedReleaseTrain(train) || IsLatestReleaseTrain(train) {
		return true
	}
	return false
}

func GetSuggestedReleases() []string {
	return suggestedData
}

func GetLatestReleases() []string {
	return latestData
}

func GetAllofReleases() []string {
	return releasesData
}

func GetDeferredReleases() []string {
	return deferredData
}

func ShowSuggestedReleases() {
	fmt.Println(suggestedData)
}

func ShowLatestReleases() {
	fmt.Println(latestData)
}

func ShowAllReleases() {
	fmt.Println(releasesData)
}

func ShowDeferredReleases() {
	fmt.Println(deferredData)
}

// PRIVATE INTERNAL HELPER FUNCTIONS

func getLatestContent(body string) string {
	latestPos := strings.Index(body, INDEX_POS_LATEST)
	allPos := strings.Index(body, INDEX_POS_RELEASES)
	if latestPos > -1 && allPos > -1 {
		return body[latestPos:allPos]
	}
	return ""
}

func getSuggestedContent(body string) string {
	suggestedPos := strings.Index(body, INDEX_POS_SUGGESTED)
	latestPos := strings.Index(body, INDEX_POS_LATEST)
	if suggestedPos > -1 && latestPos > -1 {
		return body[suggestedPos:latestPos]
	}
	return ""
}

func getAllReleasesContent(body string) string {
	allPos := strings.Index(body, INDEX_POS_RELEASES)
	deferredPos := strings.Index(body, INDEX_POS_DEFERRED)
	if allPos > -1 && deferredPos > -1 {
		return body[allPos:deferredPos]
	}
	return ""
}

func getDeferredContent(body string) string {
	deferredPos := strings.Index(body, INDEX_POS_DEFERRED)
	endPos := strings.Index(body, INDEX_POS_END)
	if deferredPos > -1 && endPos > -1 {
		return body[deferredPos:endPos]
	}
	return ""
}

func getReleasesFromContent(content string) []string {
	tmp := []string{}
	splits := strings.Split(content, "\n")
	if len(splits) > 0 {
		for i := 1; i < len(splits); i++ {
			split := strings.Split(splits[i], ",")
			if len(split) == 15 {
				tmp = append(tmp, strings.Replace(split[2], "'", "", 2))
			}
		}
	}
	return tmp
}

func getVersion(version string) string {
	major, minor, release, build := "", "", "", ""
	version = strings.TrimSpace(version)
	switch len(version) {
	case 1:
		major = version[0:1]
	case 3:
		major = version[0:1]
		minor = version[2:3]
	case 6:
		major = version[0:1]
		minor = version[2:3]
		release = version[4:5]
	case 7:
		major = version[0:1]
		minor = version[2:3]
		release = version[4:5]
		build = version[5:6]
	}
	return major + "|" + minor + "|" + release + "|" + build
}

func getVersionMajor(version string) string {
	major := ""
	version = strings.TrimSpace(version)
	if len(version) > 0 {
		major = version[0:1]
	}
	return major
}

func getVersionMinor(version string) string {
	minor := ""
	version = strings.TrimSpace(version)
	if len(version) > 2 {
		minor = version[2:3]
	}
	return minor
}

func getVersionRelease(version string) string {
	release := ""
	version = strings.TrimSpace(version)
	if len(version) > 4 {
		release = version[4:5]
	}
	return release
}

func getVersionBuild(version string) string {
	build := ""
	version = strings.TrimSpace(version)
	if len(version) > 5 {
		build = version[5:6]
	}
	return build
}

func isInSlice(str []string, train string) bool {
	for i := 0; i < len(str); i++ {
		if strings.TrimSpace(str[i]) == strings.TrimSpace(train) {
			return true
		}
	}
	return false
}

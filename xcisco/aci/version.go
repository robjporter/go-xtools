package aci

import (
	"fmt"
	"strings"

	"github.com/robjporter/go-utils/go/as"
	"github.com/robjporter/go-utils/web/request"
)

const (
	INDEX_POS_LATEST    = "mdfTree.addElementRel(469, 1, 'Latest', '');"
	INDEX_POS_SUGGESTED = "mdfTree.addElementRel(468, 1, 'Suggested', '');"
	INDEX_POS_RELEASES  = "mdfTree.addElementRel(470, 1, 'All Releases', '');"
	INDEX_POS_DEFERRED  = "mdfTree.addElementRel(471, 1, 'Deferred Releases', '');"
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
	resp, body, err := code.Get("https://software.cisco.com/download/release.html?i=!y&mdfid=285968390&softwareid=286278832").End()
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
	return "No suggested releases."
}

func GetLatestRelease() string {
	return formatVersionFromStrings(getLatestFromContent(latestData))
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

func getLatestFromContent(data []string) (string, string, string, string) {
	currentMajor := int64(0)
	currentMinor := int64(0)
	currentRelease := int64(0)
	currentBuild := int64(0)
	for i := 0; i < len(data); i++ {
		tmpMajor := as.ToInt64(getVersionMajor(data[i]))
		tmpMinor := as.ToInt64(getVersionMinor(data[i]))
		tmpRelease := as.ToInt64(getVersionRelease(data[i]))
		tmpBuild := alphabetPosition(getVersionBuild(data[i]))
		if tmpMajor > currentMajor && tmpMinor > currentMinor && tmpRelease > currentRelease && tmpBuild > currentBuild {
			currentMajor = tmpMajor
			currentMinor = tmpMinor
			currentRelease = tmpRelease
			currentBuild = tmpBuild
		}
	}
	return as.ToString(currentMajor), as.ToString(currentMinor), as.ToString(currentRelease), alphabetCharacter(currentBuild)
}

func getLatestFromContentTrain(data []string, train string) (string, string, string, string) {
	currentMajor := int64(0)
	currentMinor := int64(0)
	currentRelease := int64(0)
	currentBuild := int64(0)
	for i := 0; i < len(data); i++ {
		tmpMajor := as.ToInt64(getVersionMajor(data[i]))
		tmpMinor := as.ToInt64(getVersionMinor(data[i]))
		tmpRelease := as.ToInt64(getVersionRelease(data[i]))
		tmpBuild := alphabetPosition(getVersionBuild(data[i]))
		if tmpMajor > currentMajor && tmpMinor > currentMinor && tmpRelease > currentRelease && tmpBuild > currentBuild {
			currentMajor = tmpMajor
			currentMinor = tmpMinor
			currentRelease = tmpRelease
			currentBuild = tmpBuild
		}
	}
	return as.ToString(currentMajor), as.ToString(currentMinor), as.ToString(currentRelease), alphabetCharacter(currentBuild)
}

func formatVersionFromStrings(major string, minor string, release string, build string) string {
	return major + "." + minor + "(" + release + build + ")"
}

func alphabetPosition(char string) int64 {
	return int64(strings.Index("abcdefghijklmnopqrstuvwxyz", char) + 1)
}

func alphabetCharacter(position int64) string {
	return "abcdefghijklmnopqrstuvwxyz"[position-1 : position]
}

/*
LATEST MAJOR: > 2  LATEST MINOR: > 2  LATEST RELEASE: > 1  LATEST BUILD: > o
TRAIN MAJOR: > 2  TRAIN MINOR: > 1  TRAIN RELEASE: > 2  TRAIN BUILD: > e
CURRENT MAJOR: > 2  CURRENT MINOR: > 0  CURRENT RELEASE: > 2  CURRENT BUILD: > 2
CURRENT MAJOR: > 2  CURRENT MINOR: > 2  CURRENT RELEASE: > 1  CURRENT BUILD: > 1
CURRENT MAJOR: > 2  CURRENT MINOR: > 1  CURRENT RELEASE: > 2  CURRENT BUILD: > 2
CURRENT MAJOR: > 1  CURRENT MINOR: > 2  CURRENT RELEASE: > 3  CURRENT BUILD: > 3
*/

func GetLatestReleaseTrain(train string) string {
	if IsLatestReleaseTrain(train) {
		return "This is the latest release."
	}
	return formatVersionFromStrings(getLatestFromContentTrain(latestData, train))
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

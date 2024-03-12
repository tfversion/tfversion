package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
)

// TODO: implement automatic pagination depending on maxResults (API returns max. 20 results per page)
// API docs say to use the timestamp_created of the last list item and use query parameter `after` to get the next page

type Build struct {
	Arch string `json:"arch"`
	Os   string `json:"os"`
	Url  string `json:"url"`
}

type Status struct {
	State            string `json:"state"`
	TimestampUpdated string `json:"timestamp_updated"`
}

type Release struct {
	Builds           []Build `json:"builds"`
	Name             string  `json:"name"`
	Status           Status  `json:"status"`
	Version          string  `json:"version"`
	IsPreRelease     bool    `json:"is_prerelease"`
	TimestampCreated string  `json:"timestamp_created"`
	TimestampUpdated string  `json:"timestamp_updated"`
}

func ListVersions(maxResults int) []string {
	url := fmt.Sprintf("%s?limit=%v", download.TerraformReleasesApiUrl, maxResults)
	resp, err := http.Get(url)
	if err != nil {
		helpers.ExitWithError("getting Terraform releases from API", err)
	}
	defer resp.Body.Close()

	var releases []*Release
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		helpers.ExitWithError("getting parsing Terraform releases", err)
	}

	var availableVersions []string
	for _, r := range releases {
		availableVersions = append(availableVersions, r.Version)
	}

	return availableVersions
}

func GetVersion(version string) Release {
	resp, err := http.Get(download.TerraformReleasesApiUrl + "/" + version)
	if err != nil {
		helpers.ExitWithError("getting Terraform release from API", err)
	}
	defer resp.Body.Close()

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		helpers.ExitWithError("parsing Terraform release", err)
	}

	return release
}

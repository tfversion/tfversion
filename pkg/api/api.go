package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tfversion/tfversion/pkg/helpers"
)

const (
	// TerraformReleasesApiUrl is the URL to list available Terraform releases.
	TerraformReleasesApiUrl = "https://api.releases.hashicorp.com/v1/releases/terraform"
	ReleasesPerPage         = 20
)

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

func ListReleases(after string) ([]*Release, error) {
	url := fmt.Sprintf("%s?limit=%v&after=%s", TerraformReleasesApiUrl, ReleasesPerPage, after)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var releases []*Release
	err = json.NewDecoder(resp.Body).Decode(&releases)
	if err != nil {
		return nil, err
	}

	// append next page if there are still results
	if len(releases) > 0 {
		lastRelease := releases[len(releases)-1]
		nextPage, err := ListReleases(lastRelease.TimestampCreated)
		if err != nil {
			return nil, err
		}
		releases = append(releases, nextPage...)
	}

	return releases, nil
}

// TODO: use this method to validate a specific version before downloading and use builds[0].url for downloading
func GetRelease(version string) (*Release, error) {
	url := fmt.Sprintf("%s/%s", TerraformReleasesApiUrl, version)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var release *Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		helpers.ExitWithError("parsing Terraform release", err)
	}

	return release, nil
}

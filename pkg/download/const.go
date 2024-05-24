package download

const (
	// TerraformReleasesUrl is the URL used to download Terraform releases from.
	TerraformReleasesUrl = "https://releases.hashicorp.com/terraform"
	// MaxRetries is the maximum number of retries for a download.
	MaxRetries = 3
	// RetryTimeInSeconds is the time to wait before retrying a download.
	RetryTimeInSeconds = 2
)

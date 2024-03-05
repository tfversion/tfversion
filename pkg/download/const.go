package download

const (
	// TerraformReleasesUrl is the URL used to download Terraform releases from.
	TerraformReleasesUrl = "https://releases.hashicorp.com/terraform"
	// MaxRetries is the maximum number of retries for a download.
	MaxRetries = 3
	// RetryTimeInSeconds is the time to wait before retrying a download.
	RetryTimeInSeconds = 2
	// ApplicationDir is the directory where tfversion downloads Terraform releases.
	ApplicationDir = ".tfversion"
	// UseDir is the directory where tfversion puts the symlink to the active version.
	UseDir = "bin"
	// VersionsDir is the directory where tfversion installs all versions.
	VersionsDir = "versions"
	// AliasesDir is the directory where tfversion stores the aliases.
	AliasesDir = "aliases"
	// TerraformBinaryName is the name of the Terraform binary.
	TerraformBinaryName = "terraform"
)

package paths

const (
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

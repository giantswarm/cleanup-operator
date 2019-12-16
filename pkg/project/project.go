package project

var (
	bundleVersion = "0.0.1"
	description   = "An operator that disposes of any evidence of test stuff in your cluster."
	gitSHA        = "n/a"
	name          = "cleanup-operator"
	source        = "https://github.com/giantswarm/cleanup-operator"
	version       = "n/a"
)

func BundleVersion() string {
	return bundleVersion
}

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}

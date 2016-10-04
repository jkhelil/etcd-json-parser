package version

// vars
var (
	FullVersion        = "v0.0.0-x-sha1"
	DockerImageVersion = ""
)

// Version struct
type Version struct {
	fullversion        string
	dockerImageVersion string
}

// Get returns codebase version information
func Get() Version {
	return Version{
		fullversion:        FullVersion,
		dockerImageVersion: DockerImageVersion,
	}
}

func (v Version) String() string {
	return v.fullversion
}

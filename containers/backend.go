package containers

type Backend struct {
	Name string

	ConfigPath string // path of the config file this backend came from
	Source     string // Required

	// the following are shared, used by any source

	EnvFile string // either absolute or relative to ConfigPath

	Host string
	Port int

	Hosts []string
	Ports []int

	AwsRegion string
}

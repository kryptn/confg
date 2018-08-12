package containers

type Backend struct {
	Name string

	// should be required
	Source string

	// the following are shared, used by any source
	EnvFile string

	Host  string
	Port  int

	Hosts []string
	Ports []int
}

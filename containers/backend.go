package containers

import "errors"

type Backend struct {
	Name string

	// should be required
	Source string

	// the following are shared, used by any source
	EnvFile string

	Host string
	Port int

	Hosts []string
	Ports []int
}

func (b Backend) Validate() (bool, []error) {
	ok := true
	errs := []error{}

	if b.Source == "" {
		ok = false
		errs = append(errs, errors.New("Source must be defined"))
	}

	return ok, errs
}

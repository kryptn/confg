package containers

type Confg struct {
	version string

	Backends []*Backend
	Keys     []*Key

	Rendered map[string]map[string]interface{}
}

func (c *Confg) Validate() *[]error {
	errors := []error{}

	return &errors

}

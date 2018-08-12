package containers

type Group struct {
	Name    string
	backend string
	Backend Backend
	Keys    map[string]interface{}
}

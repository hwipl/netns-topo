package deploy

// Netns is a network namespace
type Netns struct {
	Name string
}

// NewNetns returns a new network namespace
func NewNetns() *Netns {
	return &Netns{}
}

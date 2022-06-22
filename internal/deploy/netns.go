package deploy

// Netns is a network namespace
type Netns struct {
	Name string
}

// Start starts the network namespace
func (n *Netns) Start() {
}

// Stop stops the network namespace
func (n *Netns) Stop() {
}

// NewNetns returns a new network namespace
func NewNetns() *Netns {
	return &Netns{}
}

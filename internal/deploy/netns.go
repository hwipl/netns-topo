package deploy

import (
	"fmt"
)

const (
	// netnsPrefix is the prefix used to build namespace names
	netnsPrefix = "netns-topo-"
)

// Netns is a network namespace
type Netns struct {
	Name string
}

// netnsName creates a network namespace name from topology and node name
func netnsName(topology, node string) string {
	return fmt.Sprintf("%s%s-%s", netnsPrefix, topology, node)
}

// Start starts the network namespace
func (n *Netns) Start() {
	runIP("netns", "add", n.Name)
}

// Stop stops the network namespace
func (n *Netns) Stop() {
	runIP("netns", "delete", n.Name)
}

// NewNetns returns a new network namespace
func NewNetns() *Netns {
	return &Netns{}
}

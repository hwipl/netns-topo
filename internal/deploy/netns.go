package deploy

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	// netnsPrefix is the prefix used to build namespace names
	netnsPrefix = "netns-topo-"
)

var (
	// networkNamespaces is the current list of active network namespaces
	networkNamespaces []string
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

// listNetns returns the names of active network namespaces
func listNetns() []string {
	// use buffered list of namespaces if present
	if networkNamespaces != nil {
		return networkNamespaces
	}

	// get list of namespaces
	b := &bytes.Buffer{}
	runIPStdout(b, "netns", "list")
	nses := []string{}
	for _, s := range strings.Split(b.String(), "\n") {
		fields := strings.Fields(s)
		if len(fields) == 0 {
			continue
		}
		name := fields[0]
		if strings.HasPrefix(name, "netns-topo-") {
			nses = append(nses, name)
		}
	}

	// buffer list of namespaces
	networkNamespaces = nses
	return nses
}

// Status returns the status of the network namespace
func (n *Netns) Status() Status {
	nses := listNetns()
	for _, ns := range nses {
		if n.Name == ns {
			return StatusActive
		}
	}
	return StatusInactive
}

// NewNetns returns a new network namespace
func NewNetns() *Netns {
	return &Netns{}
}

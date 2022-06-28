package deploy

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
	cmd := exec.Command("ip", "netns", "add", n.Name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}

// Stop stops the network namespace
func (n *Netns) Stop() {
	cmd := exec.Command("ip", "netns", "delete", n.Name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}

// NewNetns returns a new network namespace
func NewNetns() *Netns {
	return &Netns{}
}

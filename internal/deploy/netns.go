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

// Start starts the network namespace
func (n *Netns) Start() {
	name := fmt.Sprintf("%s%s", netnsPrefix, n.Name)
	cmd := exec.Command("ip", "netns", "add", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}

// Stop stops the network namespace
func (n *Netns) Stop() {
}

// NewNetns returns a new network namespace
func NewNetns() *Netns {
	return &Netns{}
}

package deploy

import (
	"log"
	"os"
	"os/exec"
)

const (
	tempVeth1 = "netnstopoveth0"
	tempVeth2 = "netnstopoveth1"
)

// Veth is a veth device
type Veth struct {
	Name  string
	Netns [2]string
}

// runIP runs the ip command with the parameters params
func runIP(params ...string) {
	cmd := exec.Command("ip", params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}

// Start starts the veth device
func (v *Veth) Start() {
	// add temporary veth devices
	runIP("link", "add", tempVeth1, "type", "veth", "peer", "name",
		tempVeth2)

	// move temporary veth devices into network namespaces
	runIP("link", "set", tempVeth1, "netns", v.Netns[0])
	runIP("link", "set", tempVeth2, "netns", v.Netns[1])

	// rename temporary veth devices
	runIP("netns", "exec", v.Netns[0], "ip", "link", "set", tempVeth1,
		"name", v.Name)
	runIP("netns", "exec", v.Netns[1], "ip", "link", "set", tempVeth2,
		"name", v.Name)
}

// Stop stops the veth device
func (v *Veth) Stop() {
	// delete veth devices
	runIP("netns", "exec", v.Netns[0], "ip", "link", "delete", v.Name,
		"type", "veth")
}

// NewVeth returns a new veth device
func NewVeth() *Veth {
	return &Veth{}
}

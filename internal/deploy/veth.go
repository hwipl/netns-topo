package deploy

import (
	"crypto/rand"
	"fmt"
	"log"
)

const (
	tempVeth1 = "netnstopoveth0"
	tempVeth2 = "netnstopoveth1"
)

// Veth is a veth device
type Veth struct {
	Name  string
	Netns [2]string
	MACs  [2]string
	IPs   [2]string
}

// generateMAC returns a random MAC address
func generateMAC() string {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	// set local and unicast bits
	b[0] |= 0b00000010
	b[0] &= 0b11111110

	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		b[0], b[1], b[2], b[3], b[4], b[5])
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
	runNetnsIP(v.Netns[0], "link", "set", tempVeth1, "name", v.Name)
	runNetnsIP(v.Netns[1], "link", "set", tempVeth2, "name", v.Name)

	// set MAC addresses of veth devices
	for i, mac := range v.MACs {
		if mac == "" {
			mac = generateMAC()
		}
		runNetnsIP(v.Netns[i], "link", "set", v.Name, "address", mac)
	}

	// set IP addresses of veth devices
	for i, ip := range v.IPs {
		if ip == "" {
			continue
		}
		runNetnsIP(v.Netns[i], "address", "add", ip, "dev", v.Name)
	}

	// set veth devices up
	runNetnsIP(v.Netns[0], "link", "set", v.Name, "up")
	runNetnsIP(v.Netns[1], "link", "set", v.Name, "up")
}

// Stop stops the veth device
func (v *Veth) Stop() {
	// delete veth devices
	runNetnsIP(v.Netns[0], "link", "delete", v.Name, "type", "veth")
}

// NewVeth returns a new veth device
func NewVeth() *Veth {
	return &Veth{}
}

package deploy

const (
	tempVeth1 = "netnstopoveth0"
	tempVeth2 = "netnstopoveth1"
)

// Veth is a veth device
type Veth struct {
	Name  string
	Netns [2]string
	MACs  [2]string
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
			continue
		}
		runNetnsIP(v.Netns[i], "link", "set", v.Name, "address", mac)
	}
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

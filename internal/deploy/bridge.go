package deploy

// Bridge is a bridge in a network namespace
type Bridge struct {
	Netns string
	Veths []string
}

// Start starts the bridge
func (b *Bridge) Start() {
	runNetnsIP(b.Netns, "link", "add", "br0", "type", "bridge")
	for _, v := range b.Veths {
		runNetnsIP(b.Netns, "link", "set", v, "master", "br0")
		runNetnsIP(b.Netns, "link", "set", v, "up")
		runNetnsIP(b.Netns, "link", "set", v, "promisc", "on")
	}
	runNetnsIP(b.Netns, "link", "set", "br0", "up")
	runNetnsIP(b.Netns, "link", "set", "br0", "promisc", "on")
}

// Stop stops the bridge
func (b *Bridge) Stop() {
	runNetnsIP(b.Netns, "link", "delete", "br0", "type", "bridge")
}

// NewBridge returns a new bridge
func NewBridge() *Bridge {
	return &Bridge{}
}

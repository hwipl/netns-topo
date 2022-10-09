package deploy

// Router is a router in a network namespace
type Router struct {
	Netns string
}

// Start starts the router
func (r *Router) Start() {
	// set ip forwarding on devices
	runNetns(r.Netns, "sysctl", "-w", "net.ipv4.conf.all.forwarding=1")
	runNetns(r.Netns, "sysctl", "-w", "net.ipv6.conf.all.forwarding=1")
}

// Stop stops the router
func (r *Router) Stop() {
}

// NewRouter returns a new router
func NewRouter() *Router {
	return &Router{}
}

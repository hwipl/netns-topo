package deploy

import "github.com/hwipl/netns-topo/internal/topo"

// Route is a routing table entry in a network namespace
type Route struct {
	Netns string
	Route *topo.Route
}

// Start configures the route
func (r *Route) Start() {
	runNetnsIP(r.Netns, "route", "add", r.Route.Route, "via", r.Route.Via)
}

// Stop removes the route
func (r *Route) Stop() {
	runNetnsIP(r.Netns, "route", "del", r.Route.Route, "via", r.Route.Via)
}

// NewRoute returns a new route
func NewRoute() *Route {
	return &Route{}
}

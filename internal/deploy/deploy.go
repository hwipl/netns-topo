package deploy

import "github.com/hwipl/netns-topo/internal/topo"

// Deploy is a deployment of a topology
type Deploy struct {
	t *topo.Topology
}

// Start starts the deployment
func (d *Deploy) Start() {
}

// Stop stops the deployment
func (d *Deploy) Stop() {
}

// NewDeploy returns a new deployment for t
func NewDeploy(t *topo.Topology) *Deploy {
	return &Deploy{
		t: t,
	}
}

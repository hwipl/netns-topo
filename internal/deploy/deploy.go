package deploy

import "github.com/hwipl/netns-topo/internal/topo"

// Deploy is a deployment of a topology
type Deploy struct {
	t  *topo.Topology
	ns []*Netns
}

// Start starts the deployment
func (d *Deploy) Start() {
	for _, ns := range d.ns {
		ns.Start()
	}
}

// Stop stops the deployment
func (d *Deploy) Stop() {
	for _, ns := range d.ns {
		ns.Stop()
	}
}

// createNamespaces creates namespaces from t
func createNamespaces(t *topo.Topology) (namespaces []*Netns) {
	for _, n := range t.Nodes {
		ns := NewNetns()
		ns.Name = n.Name
		namespaces = append(namespaces, ns)
	}
	return
}

// NewDeploy returns a new deployment for t
func NewDeploy(t *topo.Topology) *Deploy {
	return &Deploy{
		t:  t,
		ns: createNamespaces(t),
	}
}

package deploy

import "github.com/hwipl/netns-topo/internal/topo"

// Deploy is a deployment of a topology
type Deploy struct {
	t     *topo.Topology
	ns    []*Netns
	veths []*Veth
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

// createVeths create veths from t
func createVeths(t *topo.Topology) (veths []*Veth) {
	for _, l := range t.Links {
		v := NewVeth()
		v.Name = l.Name
		veths = append(veths, v)
	}
	return
}

// NewDeploy returns a new deployment for t
func NewDeploy(t *topo.Topology) *Deploy {
	return &Deploy{
		t:     t,
		ns:    createNamespaces(t),
		veths: createVeths(t),
	}
}

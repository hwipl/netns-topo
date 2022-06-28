package deploy

import (
	"github.com/hwipl/netns-topo/internal/topo"
)

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
	for _, v := range d.veths {
		v.Start()
	}
}

// Stop stops the deployment
func (d *Deploy) Stop() {
	for _, v := range d.veths {
		v.Stop()
	}
	for _, ns := range d.ns {
		ns.Stop()
	}
}

// createNamespaces creates namespaces from t
func (d *Deploy) createNamespaces() {
	for _, n := range d.t.Nodes {
		ns := NewNetns()
		ns.Name = netnsName(d.t.Name, n.Name)
		d.ns = append(d.ns, ns)
	}
	return
}

// createVeths create veths from t
func (d *Deploy) createVeths() {
	for _, l := range d.t.Links {
		v := NewVeth()
		v.Name = l.Name
		for i := range l.Nodes {
			v.Netns[i] = netnsName(d.t.Name, l.Nodes[i].Name)
		}
		d.veths = append(d.veths, v)
	}
	return
}

// NewDeploy returns a new deployment for t
func NewDeploy(t *topo.Topology) *Deploy {
	d := &Deploy{
		t: t,
	}
	d.createNamespaces()
	d.createVeths()
	return d
}

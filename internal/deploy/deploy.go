package deploy

import (
	"github.com/hwipl/netns-topo/internal/topo"
)

// Deploy is a deployment of a topology
type Deploy struct {
	t        *topo.Topology
	ns       []*Netns
	veths    []*Veth
	bridges  []*Bridge
	routers  []*Router
	nodeRuns []*Run
	topoRuns []*Run
}

// Start starts the deployment
func (d *Deploy) Start() {
	for _, ns := range d.ns {
		ns.Start()
	}
	for _, v := range d.veths {
		v.Start()
	}
	for _, b := range d.bridges {
		b.Start()
	}
	for _, r := range d.routers {
		r.Start()
	}
	for _, r := range d.nodeRuns {
		r.Start()
	}
	for _, r := range d.topoRuns {
		r.Start()
	}
}

// Stop stops the deployment
func (d *Deploy) Stop() {
	for _, r := range d.topoRuns {
		r.Stop()
	}
	for _, r := range d.nodeRuns {
		r.Stop()
	}
	for _, r := range d.routers {
		r.Stop()
	}
	for _, b := range d.bridges {
		b.Stop()
	}
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
}

// createVeths creates veths from t
func (d *Deploy) createVeths() {
	for _, l := range d.t.Links {
		v := NewVeth()
		v.Name = l.Name
		for i := range l.Nodes {
			v.Netns[i] = netnsName(d.t.Name, l.Nodes[i].Name)
		}
		for i, mac := range l.MACs {
			v.MACs[i] = mac
		}
		for i, ip := range l.IPs {
			v.IPs[i] = ip
		}
		d.veths = append(d.veths, v)
	}
}

// getNetnsVeths returns the veths in netns
func (d *Deploy) getNetnsVeths(netns string) []string {
	veths := []string{}
	for _, v := range d.veths {
		if v.Netns[0] == netns || v.Netns[1] == netns {
			veths = append(veths, v.Name)
		}
	}
	return veths
}

// createBridges creates bridges from t
func (d *Deploy) createBridges() {
	for _, n := range d.t.Nodes {
		if n.Type == topo.NodeTypeBridge {
			br := NewBridge()
			br.Netns = netnsName(d.t.Name, n.Name)
			br.Veths = d.getNetnsVeths(br.Netns)
			d.bridges = append(d.bridges, br)
		}
	}
}

// createRouters creates routers from t
func (d *Deploy) createRouters() {
	for _, n := range d.t.Nodes {
		if n.Type == topo.NodeTypeRouter {
			r := NewRouter()
			r.Netns = netnsName(d.t.Name, n.Name)
			d.routers = append(d.routers, r)
		}
	}
}

// createRuns creates runs from t
func (d *Deploy) createRuns() {
	for _, n := range d.t.Nodes {
		if len(n.Run) > 0 {
			r := NewRun()
			r.Netns = netnsName(d.t.Name, n.Name)
			r.Commands = n.Run
			d.nodeRuns = append(d.nodeRuns, r)
		}
	}
	for _, run := range d.t.Run {
		r := NewRun()
		r.Netns = netnsName(d.t.Name, run.Node)
		r.Commands = run.Commands
		d.topoRuns = append(d.topoRuns, r)
	}
}

// NewDeploy returns a new deployment for t
func NewDeploy(t *topo.Topology) *Deploy {
	d := &Deploy{
		t: t,
	}
	d.createNamespaces()
	d.createVeths()
	d.createBridges()
	d.createRouters()
	d.createRuns()
	return d
}

package deploy

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/hwipl/netns-topo/internal/topo"
)

// Deploy is a deployment of a topology
type Deploy struct {
	t        *topo.Topology
	ns       []*Netns
	veths    []*Veth
	routes   []*Route
	bridges  []*Bridge
	routers  []*Router
	nodeRuns []*Run
	topoRuns []*Run
}

// getNetnsStatus returns the status of all network namespaces
func (d *Deploy) getNetnsStatus() Status {
	status := StatusUnknown
	for _, n := range d.ns {
		switch n.Status() {
		case StatusInactive:
			if status == StatusUnknown {
				status = StatusInactive
			}
			if status == StatusActive {
				status = StatusPartial
			}
		case StatusActive:
			if status == StatusUnknown {
				status = StatusActive
			}
			if status == StatusInactive {
				status = StatusPartial
			}
		}
	}
	return status
}

// Status returns the status of the deployment
func (d *Deploy) Status() Status {
	netnsStatus := d.getNetnsStatus()

	// TODO: add more status checks?
	return netnsStatus
}

// getDeployDir returns the directory where active deployments are saved
func getDeployDir() string {
	dir := filepath.Join(os.TempDir(), "netns-topo", "topologies")
	return dir
}

// makeDeployDir creates and returns the directory where active deployments
// are saved
func makeDeployDir() string {
	dir := getDeployDir()
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// saveDeployFile saves the deployment in the directory where active
// deployments are saved
func (d *Deploy) saveDeployFile() {
	dir := makeDeployDir()
	file := filepath.Join(dir, d.t.Name)
	err := os.WriteFile(file, d.t.YAML(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// removeDeployFile removes the deployment from the directory where active
// deployments are saved
func (d *Deploy) removeDeployFile() {
	dir := getDeployDir()
	file := filepath.Join(dir, d.t.Name)
	err := os.Remove(file)
	if err != nil {
		log.Fatal(err)
	}
}

// Start starts the deployment
func (d *Deploy) Start() {
	if d.Status() == StatusActive {
		log.Println(d.t.Name, "already active")
		return
	}

	d.saveDeployFile()

	for _, ns := range d.ns {
		ns.Start()
	}
	for _, v := range d.veths {
		v.Start()
	}
	for _, r := range d.routes {
		r.Start()
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
	if d.Status() != StatusActive {
		log.Println(d.t.Name, "not active")
		return
	}

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
	for _, r := range d.routes {
		r.Stop()
	}
	for _, v := range d.veths {
		v.Stop()
	}
	for _, ns := range d.ns {
		ns.Stop()
	}

	d.removeDeployFile()
}

// RunCmd runs cmd on node in the deployment
func (d *Deploy) RunCmd(node, cmd string) {
	if d.Status() != StatusActive {
		log.Println(d.t.Name, "not active")
		return
	}

	runIPStdinOutErr(os.Stdin, os.Stdout, os.Stderr,
		"netns", "exec", netnsName(d.t.Name, node),
		"/bin/bash", "-c", cmd)
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

// createRoutes creates routes from t
func (d *Deploy) createRoutes() {
	for _, n := range d.t.Nodes {
		for _, route := range n.Routes {
			r := NewRoute()
			r.Netns = netnsName(d.t.Name, n.Name)
			r.Route = route
			d.routes = append(d.routes, r)
		}
	}
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
	d.createRoutes()
	d.createBridges()
	d.createRouters()
	d.createRuns()
	return d
}

// listDeployDir returns the deployments saved in the directory for
// active deployments
func listDeployDir() []*Deploy {
	deploys := []*Deploy{}

	// read content of deployment directory
	dir := getDeployDir()
	files, err := os.ReadDir(dir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return deploys
		}

		log.Fatal(err)
	}

	// read deployments
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		t := topo.NewTopologyYAMLFile(filepath.Join(dir, f.Name()))
		d := NewDeploy(t)
		deploys = append(deploys, d)
	}

	return deploys
}

// ListDeploys lists active deploys
func ListDeploys() {
	for _, d := range listDeployDir() {
		fmt.Println(d.t.Name, d.Status())
	}
}

package cmd

import (
	"flag"
	"log"

	"github.com/hwipl/netns-topo/internal/deploy"
	"github.com/hwipl/netns-topo/internal/topo"
)

// findDeploy tries to find an existing deploy identified by name, to read it
// from a saved topology identified by name, or to read it from a file
// identified by name
func findDeploy(name string) *deploy.Deploy {
	d := deploy.GetDeploy(name)
	if d == nil {
		if t := topo.GetTopology(name); t != nil {
			d = deploy.NewDeploy(t)
		}
	}
	if d == nil {
		t := topo.NewTopologyYAMLFile(name)
		d = deploy.NewDeploy(t)
	}
	return d
}

// isForce returns whether force was specified on the command line
func isForce() bool {
	if flag.Arg(2) == "force" {
		return true
	}
	return false
}

// Run is the main entry point
func Run() {
	// parse command line arguments
	flag.Parse()
	command := flag.Arg(0)
	name := flag.Arg(1)

	// handle command
	switch command {
	case "start":
		d := findDeploy(name)
		d.Start(isForce())
	case "stop":
		d := findDeploy(name)
		d.Stop(isForce())
	case "list":
		deploy.ListDeploys()
	case "run":
		d := findDeploy(name)
		node := flag.Arg(2)
		cmd := flag.Arg(3)
		d.RunCmd(node, cmd)
	case "save":
		d := findDeploy(name)
		d.Topology().SaveTopologyFile()
	default:
		log.Fatal("unknown command: ", command)
	}
}

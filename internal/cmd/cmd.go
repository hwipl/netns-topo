package cmd

import (
	"flag"
	"log"

	"github.com/hwipl/netns-topo/internal/deploy"
	"github.com/hwipl/netns-topo/internal/topo"
)

// findDeploy tries to find an existing deploy identified by name or reading
// it from a file identified by name
func findDeploy(name string) *deploy.Deploy {
	d := deploy.GetDeploy(name)
	if d == nil {
		t := topo.NewTopologyYAMLFile(name)
		d = deploy.NewDeploy(t)
	}
	return d
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
		d.Start()
	case "stop":
		d := findDeploy(name)
		d.Stop()
	case "list":
		deploy.ListDeploys()
	case "run":
		d := findDeploy(name)
		node := flag.Arg(2)
		cmd := flag.Arg(3)
		d.RunCmd(node, cmd)
	default:
		log.Fatal("unknown command: ", command)
	}
}

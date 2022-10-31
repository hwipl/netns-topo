package cmd

import (
	"flag"
	"log"

	"github.com/hwipl/netns-topo/internal/deploy"
	"github.com/hwipl/netns-topo/internal/topo"
)

// Run is the main entry point
func Run() {
	// parse command line arguments
	flag.Parse()
	command := flag.Arg(0)
	file := flag.Arg(1)

	// handle command
	switch command {
	case "start":
		t := topo.NewTopologyYAMLFile(file)
		d := deploy.NewDeploy(t)
		d.Start()
	case "stop":
		t := topo.NewTopologyYAMLFile(file)
		d := deploy.NewDeploy(t)
		d.Stop()
	default:
		log.Fatal("unknown command: ", command)
	}
}

package cmd

import (
	"flag"
	"log"
	"os"

	"github.com/hwipl/netns-topo/internal/deploy"
	"github.com/hwipl/netns-topo/internal/topo"
)

// parseTopology parses topology from file
func parseTopology(file string) *topo.Topology {
	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return topo.NewTopologyYAML(b)
}

// Run is the main entry point
func Run() {
	// parse command line arguments
	flag.Parse()
	command := flag.Arg(0)
	file := flag.Arg(1)

	// handle command
	switch command {
	case "start":
		t := parseTopology(file)
		d := deploy.NewDeploy(t)
		d.Start()
	case "stop":
		t := parseTopology(file)
		d := deploy.NewDeploy(t)
		d.Stop()
	default:
		log.Fatal("unknown command: ", command)
	}
}

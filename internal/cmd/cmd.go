package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hwipl/netns-topo/internal/deploy"
	"github.com/hwipl/netns-topo/internal/topo"
)

// usage prints the usage message
func usage() {
	out := flag.CommandLine.Output()
	fmt.Fprintf(out, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(out, "  %s <command>\n\n", os.Args[0])
	fmt.Fprintf(out, "Commands:\n")
	for _, s := range []string{
		"start <topology> [force]",
		"\tstart topology",
		"stop <topology> [force]",
		"\tstop topology",
		"list",
		"\tlist topologies",
		"run <topology> <node> <command>",
		"\trun command on node in topology",
		"save <topology>",
		"\tsave topology",
		"remove <topology>",
		"\tremove saved topology",
		"help",
		"\tshow this help",
	} {
		fmt.Fprintf(out, "  %s\n", s)
	}
}

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

// listTopologies lists topologies
func listTopologies() {
	fmt.Println("Saved Topologies:")
	fmt.Println("=================")
	for _, t := range topo.GetTopologies() {
		fmt.Println(t.Name)
	}
	fmt.Println()
}

// listDeploy lists active deploys
func listDeploys() {
	fmt.Println("Active Topologies:")
	fmt.Println("==================")
	for _, d := range deploy.GetDeploys() {
		fmt.Printf("%s\t(%s)\n", d.Topology().Name, d.Status())
	}
	fmt.Println()
}

// Run is the main entry point
func Run() {
	// parse command line arguments
	flag.Usage = usage
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
		listTopologies()
		listDeploys()
	case "run":
		d := findDeploy(name)
		node := flag.Arg(2)
		cmd := flag.Arg(3)
		d.RunCmd(node, cmd)
	case "save":
		d := findDeploy(name)
		d.Topology().SaveTopologyFile()
	case "remove":
		d := findDeploy(name)
		d.Topology().RemoveTopologyFile()
	case "help":
		flag.Usage()
	default:
		flag.Usage()
		log.Fatal("unknown command: ", command)
	}
}

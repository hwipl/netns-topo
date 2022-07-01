package cmd

import (
	"fmt"
	"testing"

	"github.com/hwipl/netns-topo/internal/deploy"
	"github.com/hwipl/netns-topo/internal/topo"
)

// TestAll tests creating, parsing and deploying a topology
func TestAll(_ *testing.T) {
	// create dummy topology for testing
	t := topo.NewTopology()
	t.Name = "Topo1"

	node1 := topo.NewNode()
	node1.Name = "Node1"
	node1.Type = topo.NodeTypeNode

	node2 := topo.NewNode()
	node2.Name = "Node2"
	node2.Type = topo.NodeTypeNode

	node3 := topo.NewNode()
	node3.Name = "Node3"
	node3.Type = topo.NodeTypeNode

	t.AddNode(node1)
	t.AddNode(node2)
	t.AddNode(node3)

	link1 := topo.NewLink()
	link1.Name = "Link1"
	link1.Type = topo.LinkTypeVeth

	link2 := topo.NewLink()
	link2.Name = "Link2"
	link2.Type = topo.LinkTypeVeth

	link1.Nodes[0] = node1
	link1.Nodes[1] = node2

	link2.Nodes[0] = node2
	link2.Nodes[1] = node3

	t.AddLink(link1)
	t.AddLink(link2)

	// dump yaml of topology
	fmt.Println("YAML Topology:")
	b := t.YAML()
	fmt.Printf("%s", b)

	// parse yaml topology
	t = topo.NewTopologyYAML(b)

	fmt.Println("Parsed Topology:")
	fmt.Println("name:", t.Name)
	fmt.Println("nodes:")
	for _, n := range t.Nodes {
		fmt.Println("  - name:", n.Name)
		fmt.Println("    type:", &n.Type)
	}
	fmt.Println("links:")
	for _, l := range t.Links {
		fmt.Println("  - name:", l.Name)
		fmt.Println("    type:", &l.Type)
		fmt.Println("    nodes:")
		for _, n := range l.Nodes {
			if n == nil {
				continue
			}
			fmt.Println("      -", n.Name)
		}
	}

	// deploy topology
	d := deploy.NewDeploy(t)
	d.Start()
	defer d.Stop()

}

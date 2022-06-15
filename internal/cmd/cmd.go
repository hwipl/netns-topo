package cmd

import (
	"fmt"
)

// Run is the main entry point
func Run() {
	// create dummy topology for testing
	topo := NewTopology()

	node1 := NewNode()
	node1.Name = "Node1"
	node1.Type = NodeTypeNode

	node2 := NewNode()
	node2.Name = "Node2"
	node2.Type = NodeTypeNode

	node3 := NewNode()
	node3.Name = "Node3"
	node3.Type = NodeTypeNode

	topo.AddNode(node1)
	topo.AddNode(node2)
	topo.AddNode(node3)

	link1 := NewLink()
	link1.Name = "Link1"

	link2 := NewLink()
	link2.Name = "Link2"

	link1.Nodes[0] = node1
	link1.Nodes[1] = node2

	link2.Nodes[0] = node2
	link2.Nodes[1] = node3

	topo.AddLink(link1)
	topo.AddLink(link2)

	// dump yaml of topology
	fmt.Printf("%s", topo.YAML())
}

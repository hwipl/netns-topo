package cmd

import (
	"fmt"
)

// Run is the main entry point
func Run() {
	// create dummy topology for testing
	topo := NewTopology()

	node1 := NewNode()
	node1.Type = NodeTypeNode

	node2 := NewNode()
	node2.Type = NodeTypeNode

	node3 := NewNode()
	node3.Type = NodeTypeNode

	topo.AddNode(node1)
	topo.AddNode(node2)
	topo.AddNode(node3)

	link1 := NewLink()
	link2 := NewLink()

	link1.Nodes[0] = node1
	link1.Nodes[1] = node2

	link2.Nodes[0] = node2
	link2.Nodes[1] = node3

	topo.AddLink(link1)
	topo.AddLink(link2)

	// dump yaml of topology
	fmt.Printf("%s", topo.YAML())
}

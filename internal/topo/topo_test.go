package topo

import (
	"testing"
)

// TestNodeTypeString tests String of NodeType
func TestNodeTypeString(t *testing.T) {
	// test nil
	var none *NodeType
	if none.String() != "<nil>" {
		t.Errorf("got %s, want <nil>", none.String())
	}

	// test types
	types := map[NodeType]string{
		NodeTypeNode:    "node",
		NodeTypeBridge:  "bridge",
		NodeTypeInvalid: "",
		123:             "",
	}
	for nt, v := range types {
		s := nt.String()
		if s != v {
			t.Errorf("got %s, want %s", s, v)
		}
	}
}

// TestNodeString tests String of Node
func TestNodeString(t *testing.T) {
	// test nil
	var none *Node
	if none.String() != "<nil>" {
		t.Errorf("got %s, want <nil>", none.String())
	}

	// test node
	node := NewNode()
	node.Name = "mynode"
	want := "{Name: mynode, Type: node}"
	got := node.String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestParseNodeType tests ParseNodeType
func TestParseNodeType(t *testing.T) {
	// test node
	want := NodeTypeNode
	got := ParseNodeType("node")
	if got != want {
		t.Errorf("got %s, want %s", &got, &want)
	}

	// test bridge
	want = NodeTypeBridge
	got = ParseNodeType("bridge")
	if got != want {
		t.Errorf("got %s, want %s", &got, &want)
	}

	// test invalid
	want = NodeTypeInvalid
	got = ParseNodeType("does not exist")
	if got != want {
		t.Errorf("got %s, want %s", &got, &want)
	}
}

// TestLinkTypeString tests String of LinkType
func TestLinkTypeString(t *testing.T) {
	// test nil
	var none *LinkType
	if none.String() != "<nil>" {
		t.Errorf("got %s, want <nil>", none.String())
	}

	// test types
	types := map[LinkType]string{
		LinkTypeVeth:    "veth",
		LinkTypeInvalid: "",
		123:             "",
	}
	for lt, v := range types {
		s := lt.String()
		if s != v {
			t.Errorf("got %s, want %s", s, v)
		}
	}
}

// TestLinkString tests String of Link
func TestLinkString(t *testing.T) {
	// test nil
	var none *Link
	if none.String() != "<nil>" {
		t.Errorf("got %s, want <nil>", none.String())
	}

	// test link
	link := NewLink()
	link.Name = "mylink"
	want := "{Name: mylink, Type: veth, Nodes: [<nil> <nil>]}"
	got := link.String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestParseLinkType tests ParseLinkType
func TestParseLinkType(t *testing.T) {
	// test veth
	want := LinkTypeVeth
	got := ParseLinkType("veth")
	if got != want {
		t.Errorf("got %s, want %s", &got, &want)
	}

	// test invalid
	want = LinkTypeInvalid
	got = ParseLinkType("does not exist")
	if got != want {
		t.Errorf("got %s, want %s", &got, &want)
	}
}

// TestTopologyString tests String of Topology
func TestTopologyString(t *testing.T) {
	// test nil
	var none *Topology
	if none.String() != "<nil>" {
		t.Errorf("got %s, want <nil>", none.String())
	}

	// test topology
	topology := NewTopology()
	topology.Name = "mytopology"
	want := "{Name: mytopology, Nodes: [], Links: []}"
	got := topology.String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestTopologyAddNode tests AddNode of Topology
func TestTopologyAddNode(t *testing.T) {
	topology := NewTopology()
	num := 0

	// add nil
	topology.AddNode(nil)
	num = len(topology.Nodes)
	if num != 0 {
		t.Errorf("got %d want 0", num)
	}

	// add node
	node1 := NewNode()
	topology.AddNode(node1)
	num = len(topology.Nodes)
	if num != 1 {
		t.Errorf("got %d want 1", num)
	}
	if topology.Nodes[0] != node1 {
		t.Errorf("got %p, want %p", topology.Nodes[0], node1)
	}

	// add node
	node2 := NewNode()
	topology.AddNode(node2)
	num = len(topology.Nodes)
	if num != 2 {
		t.Errorf("got %d want 2", num)
	}
	if topology.Nodes[1] != node2 {
		t.Errorf("got %p, want %p", topology.Nodes[1], node2)
	}
}

// TestTopologyGetNode tests GetNode of Topology
func TestTopologyGetNode(t *testing.T) {
	topology := NewTopology()

	// test empty
	if topology.GetNode("does not exist") != nil {
		t.Errorf("got != nil, want nil")
	}

	// add node
	node := NewNode()
	name := "Node1"
	node.Name = name
	topology.AddNode(node)

	// test nonexistent
	if topology.GetNode("does not exist") != nil {
		t.Errorf("got != nil, want nil")
	}

	// test existent
	n := topology.GetNode(name)
	if n != node {
		t.Errorf("got %p, want %p", n, node)
	}
}

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
		t.Errorf("got %s, want \"\"", none.String())
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

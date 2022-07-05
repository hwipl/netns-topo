package topo

import (
	"testing"
)

// TestNodeTypeString tests String of NodeType
func TestNodeTypeString(t *testing.T) {
	// test nil
	var none *NodeType
	if none.String() != "" {
		t.Errorf("got %s, want \"\"", none.String())
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

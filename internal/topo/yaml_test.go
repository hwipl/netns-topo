package topo

import (
	"bytes"
	"testing"
)

// TestParseYAMLTopology tests parsing yaml topologies
func TestParseYAMLTopology(t *testing.T) {
	want := []byte(`name: Topo1
nodes:
    - name: Node1
      type: node
    - name: Node2
      type: node
    - name: Node3
      type: node
links:
    - name: Link1
      type: veth
      nodes:
        - Node1
        - Node2
    - name: Link2
      type: veth
      nodes:
        - Node2
        - Node3
`)
	got := NewYAMLTopology(ParseYAMLTopology(want)).YAML()
	if !bytes.Equal(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

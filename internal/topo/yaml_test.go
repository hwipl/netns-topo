package topo

import (
	"testing"
)

// TestParseYAMLTopology tests parsing yaml topologies
func TestParseYAMLTopology(t *testing.T) {
	want := "{Name: Topo1, Nodes: [" +
		"{Name: Node1, Type: node, Routes: [], Run: []} " +
		"{Name: Node2, Type: node, Routes: [], Run: []} " +
		"{Name: Node3, Type: node, Routes: [], Run: []}" +
		"], " +
		"Links: [" +
		"{Name: Link1, Type: veth, Nodes: [" +
		"{Name: Node1, Type: node, Routes: [], Run: []} " +
		"{Name: Node2, Type: node, Routes: [], Run: []}" +
		"], MACs: [ ], IPs: [ ]} " +
		"{Name: Link2, Type: veth, Nodes: [" +
		"{Name: Node2, Type: node, Routes: [], Run: []} " +
		"{Name: Node3, Type: node, Routes: [], Run: []}" +
		"], MACs: [ ], IPs: [ ]}" +
		"], Run: []}"

	yaml := []byte(`name: Topo1
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
	got := ParseYAMLTopology(yaml).String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	yaml = []byte(`---
name: Topo1
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
	got = ParseYAMLTopology(yaml).String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

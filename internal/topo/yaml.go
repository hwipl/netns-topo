package topo

import (
	"log"

	"gopkg.in/yaml.v3"
)

// YAMLNodeType is a yaml representation of a node type
type YAMLNodeType string

// YAMLRoute is a yaml representation of a route on a node
type YAMLRoute Route

// YAMLNode is a yaml representation of a node
type YAMLNode struct {
	Name   string
	Type   YAMLNodeType
	Routes []*Route
	Run    []string
}

// NewYAMLNode returns a new YAMLNode
func NewYAMLNode(n *Node) *YAMLNode {
	return &YAMLNode{
		Name:   n.Name,
		Type:   YAMLNodeType(n.Type.String()),
		Routes: n.Routes,
		Run:    n.Run,
	}
}

// YAMLLinkType is a yaml representation of a link type
type YAMLLinkType string

// YAMLLink is a yaml representation of a link
type YAMLLink struct {
	Name  string
	Type  YAMLLinkType
	Nodes [2]string
	MACs  [2]string
	IPs   [2]string
}

// NewYAMLLink returns a new YAMLLink
func NewYAMLLink(l *Link) *YAMLLink {
	return &YAMLLink{
		Name: l.Name,
		Type: YAMLLinkType(l.Type.String()),
		Nodes: [2]string{
			l.Nodes[0].Name,
			l.Nodes[1].Name,
		},
		MACs: [2]string{
			l.MACs[0],
			l.MACs[1],
		},
		IPs: [2]string{
			l.IPs[0],
			l.IPs[1],
		},
	}
}

// YAMLRun is a yaml representation of commands to run on a node
type YAMLRun struct {
	Node     string
	Commands []string
}

// NewYAMLRun returns a new YAMLRun
func NewYAMLRun(r *Run) *YAMLRun {
	return &YAMLRun{
		Node:     r.Node,
		Commands: r.Commands,
	}
}

// YAMLTopology is a yaml representation of a topology
type YAMLTopology struct {
	Name  string
	Nodes []*YAMLNode
	Links []*YAMLLink
	Run   []*YAMLRun
}

// YAML returns the topology as yaml
func (t *YAMLTopology) YAML() []byte {
	b, err := yaml.Marshal(t)
	if err != nil {
		return nil
	}
	return b
}

// ParseYAMLTopology parses and returns the yaml topology in b
func ParseYAMLTopology(b []byte) *Topology {
	yt := &YAMLTopology{}
	err := yaml.Unmarshal(b, yt)
	if err != nil {
		log.Fatalf("cannot parse yaml topology: %v", err)
	}

	// create topology
	t := NewTopology()

	// set name
	t.Name = yt.Name

	// set nodes
	for _, yn := range yt.Nodes {
		n := NewNode()
		n.Name = yn.Name
		n.Type = ParseNodeType(string(yn.Type))
		n.Routes = yn.Routes
		n.Run = yn.Run
		t.AddNode(n)
	}

	// set links
	for _, yl := range yt.Links {
		l := NewLink()
		l.Name = yl.Name
		l.Type = ParseLinkType(string(yl.Type))
		for i, yn := range yl.Nodes {
			l.Nodes[i] = t.GetNode(yn)
		}
		for i, ym := range yl.MACs {
			l.MACs[i] = ym
		}
		for i, yi := range yl.IPs {
			l.IPs[i] = yi
		}
		t.AddLink(l)
	}

	// set run
	for _, yr := range yt.Run {
		r := NewRun()
		r.Node = yr.Node
		r.Commands = yr.Commands
		t.AddRun(r)
	}

	return t
}

// NewYAMLTopology returns a new YAMLTopology
func NewYAMLTopology(t *Topology) *YAMLTopology {
	yt := &YAMLTopology{}
	if t == nil {
		return yt
	}

	// set name
	yt.Name = t.Name

	// set nodes
	for _, n := range t.Nodes {
		yt.Nodes = append(yt.Nodes, NewYAMLNode(n))
	}

	// set links
	for _, l := range t.Links {
		yt.Links = append(yt.Links, NewYAMLLink(l))
	}

	// set run
	for _, r := range t.Run {
		yt.Run = append(yt.Run, NewYAMLRun(r))
	}

	return yt
}

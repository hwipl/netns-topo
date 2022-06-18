package topo

import (
	"log"

	"gopkg.in/yaml.v3"
)

// YAMLNodeType is a yaml representation of a node type
type YAMLNodeType string

// YAMLNode is a yaml representation of a node
type YAMLNode struct {
	Name string
	Type YAMLNodeType
}

// NewYAMLNode returns a new YAMLNode
func NewYAMLNode(n *Node) *YAMLNode {
	return &YAMLNode{
		Name: n.Name,
		Type: YAMLNodeType(n.Type.String()),
	}
}

// FromYAMLNode returns a new Node from yn
func FromYAMLNode(yn *YAMLNode) *Node {
	n := NewNode()
	n.Name = yn.Name
	// TODO: set type
	return n
}

// YAMLLinkType is a yaml representation of a link type
type YAMLLinkType string

// YAMLLink is a yaml representation of a link
type YAMLLink struct {
	Name  string
	Type  YAMLLinkType
	Nodes [2]string
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
	}
}

// FromYAMLLink returns a new Link from yl
func FromYAMLLink(yl *YAMLLink) *Link {
	l := NewLink()
	l.Name = yl.Name
	// TODO: set type
	// TODO: set nodes
	return l
}

// YAMLTopology is a yaml representation of a topology
type YAMLTopology struct {
	Name  string
	Nodes []*YAMLNode
	Links []*YAMLLink
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
		t.Nodes = append(t.Nodes, FromYAMLNode(yn))
	}

	// set links
	for _, yl := range yt.Links {
		t.Links = append(t.Links, FromYAMLLink(yl))
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

	return yt
}

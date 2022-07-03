package topo

import (
	"fmt"
)

// NodeType is the type of a node
type NodeType uint8

// Node types
const (
	NodeTypeNode NodeType = iota
	NodeTypeBridge
	NodeTypeInvalid
)

// String returns the node type as string
func (nt *NodeType) String() string {
	switch *nt {
	case NodeTypeNode:
		return "node"
	case NodeTypeBridge:
		return "bridge"
	}
	return ""
}

// ParseNodeType parses the node type in s
func ParseNodeType(s string) NodeType {
	switch s {
	case "node":
		return NodeTypeNode
	case "bridge":
		return NodeTypeBridge
	}
	return NodeTypeInvalid
}

// Node is a node in a topology
type Node struct {
	Name string
	Type NodeType
}

// String returns the node as string
func (n *Node) String() string {
	return fmt.Sprintf("{Name: %s, Type: %s}", n.Name, &n.Type)
}

// NewNode returns a new Node
func NewNode() *Node {
	return &Node{}
}

// LinkType is the type of a link
type LinkType uint8

// Link types
const (
	LinkTypeVeth LinkType = iota
	LinkTypeInvalid
)

// String returns the link type as string
func (lt *LinkType) String() string {
	switch *lt {
	case LinkTypeVeth:
		return "veth"
	}
	return ""
}

// ParseLinkType parses the link type in s
func ParseLinkType(s string) LinkType {
	switch s {
	case "veth":
		return LinkTypeVeth
	}
	return LinkTypeInvalid
}

// Link is a link between nodes in a topology
type Link struct {
	Name  string
	Type  LinkType
	Nodes [2]*Node
}

// String returns the link as string
func (l *Link) String() string {
	return fmt.Sprintf("{Name: %s, Type: %s, Nodes: %s}", l.Name, &l.Type,
		l.Nodes)
}

// NewLink returns a new Link
func NewLink() *Link {
	return &Link{}
}

// Topology is a network topology
type Topology struct {
	Name  string
	Nodes []*Node
	Links []*Link
}

// AddNode adds a node to the topology
func (t *Topology) AddNode(node *Node) {
	t.Nodes = append(t.Nodes, node)
}

// GetNode returns the first node with name
func (t *Topology) GetNode(name string) *Node {
	for _, n := range t.Nodes {
		if n.Name == name {
			return n
		}
	}
	return nil
}

// AddLink adds a link to the topology
func (t *Topology) AddLink(link *Link) {
	t.Links = append(t.Links, link)
}

// YAML returns the topology as yaml
func (t *Topology) YAML() []byte {
	return NewYAMLTopology(t).YAML()
}

// NewTopology returns a new Topology
func NewTopology() *Topology {
	return &Topology{}
}

// NewTopologyYAML parses and returns the yaml topology in b
func NewTopologyYAML(b []byte) *Topology {
	return ParseYAMLTopology(b)
}

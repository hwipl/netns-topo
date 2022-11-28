package topo

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// NodeType is the type of a node
type NodeType uint8

// Node types
const (
	NodeTypeNode NodeType = iota
	NodeTypeBridge
	NodeTypeRouter
	NodeTypeInvalid
)

// String returns the node type as string
func (nt *NodeType) String() string {
	if nt == nil {
		return "<nil>"
	}
	switch *nt {
	case NodeTypeNode:
		return "node"
	case NodeTypeBridge:
		return "bridge"
	case NodeTypeRouter:
		return "router"
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
	case "router":
		return NodeTypeRouter
	}
	return NodeTypeInvalid
}

// Route is a routing table entry on a node
type Route struct {
	Route string
	Via   string
}

// String returns the route as string
func (r *Route) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Route: %s, Via: %s}", r.Route, r.Via)
}

// Node is a node in a topology
type Node struct {
	Name   string
	Type   NodeType
	Routes []*Route
	Run    []string
}

// String returns the node as string
func (n *Node) String() string {
	if n == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Name: %s, Type: %s, Routes: %s, Run: %s}",
		n.Name, &n.Type, n.Routes, n.Run)
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
	if lt == nil {
		return "<nil>"
	}
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
	MACs  [2]string
	IPs   [2]string
}

// String returns the link as string
func (l *Link) String() string {
	if l == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Name: %s, Type: %s, Nodes: %s, MACs: %s, IPs: %s}",
		l.Name, &l.Type, l.Nodes, l.MACs, l.IPs)
}

// NewLink returns a new Link
func NewLink() *Link {
	return &Link{}
}

// Run is a list of command to run on a node
type Run struct {
	Node     string // TODO: use *Node?
	Commands []string
}

// String returns run as string
func (r *Run) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Node: %s, Commands: %s}", r.Node, r.Commands)
}

// NewRun returns a new Run
func NewRun() *Run {
	return &Run{}
}

// Topology is a network topology
type Topology struct {
	Name  string
	Nodes []*Node
	Links []*Link
	Run   []*Run
}

// String returns the topology as string
func (t *Topology) String() string {
	if t == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Name: %s, Nodes: %s, Links: %s, Run: %s}",
		t.Name, t.Nodes, t.Links, t.Run)
}

// AddNode adds a node to the topology
func (t *Topology) AddNode(node *Node) {
	if node == nil {
		return
	}
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
	if link == nil {
		return
	}
	t.Links = append(t.Links, link)
}

// AddRun adds a list of commands to run to the topology
func (t *Topology) AddRun(run *Run) {
	if run == nil {
		return
	}
	t.Run = append(t.Run, run)
}

// YAML returns the topology as yaml
func (t *Topology) YAML() []byte {
	return NewYAMLTopology(t).YAML()
}

// getTopologyDir returns the directory where topologies are saved
func getTopologyDir() string {
	return "/var/lib/netns-topo/topologies"
}

// makeTopologyDir creates and returns the directory where topologies are saved
func makeTopologyDir() string {
	dir := getTopologyDir()
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// SaveTopologyFile saves the topology in the topologies directory
func (t *Topology) SaveTopologyFile() {
	dir := makeTopologyDir()
	file := filepath.Join(dir, t.Name)
	err := os.WriteFile(file, t.YAML(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// RemoveTopologyFile removes the topology from the topologies directory
func (t *Topology) RemoveTopologyFile() {
	dir := getTopologyDir()
	file := filepath.Join(dir, t.Name)
	err := os.Remove(file)
	if err != nil {
		log.Fatal(err)
	}
}

// NewTopology returns a new Topology
func NewTopology() *Topology {
	return &Topology{}
}

// NewTopologyYAML parses and returns the yaml topology in b
func NewTopologyYAML(b []byte) *Topology {
	return ParseYAMLTopology(b)
}

// NewTopologyYAMLFile returns a new topology parsed from yaml file
func NewTopologyYAMLFile(file string) *Topology {
	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return NewTopologyYAML(b)
}

// listTopologyDir returns the topologies saved in the topologies directory
func listTopologyDir() []*Topology {
	topos := []*Topology{}

	// read content of topologies directory
	dir := getTopologyDir()
	files, err := os.ReadDir(dir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return topos
		}

		log.Fatal(err)
	}

	// read topologies
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		t := NewTopologyYAMLFile(filepath.Join(dir, f.Name()))
		topos = append(topos, t)
	}

	return topos
}

// GetTopologies returns all topologies
func GetTopologies() []*Topology {
	return listTopologyDir()
}

// GetTopology returns the topology identified by name
func GetTopology(name string) *Topology {
	for _, t := range GetTopologies() {
		if t.Name == name {
			return t
		}
	}
	return nil
}

// ListTopologies lists topologies
func ListTopologies() {
	fmt.Println("Saved Topologies:")
	fmt.Println("=================")
	for _, t := range GetTopologies() {
		fmt.Println(t.Name)
	}
	fmt.Println()
}

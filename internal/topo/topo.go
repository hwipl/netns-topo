package topo

// NodeType is the type of a node
type NodeType uint8

// Node types
const (
	NodeTypeNode NodeType = iota
	NodeTypeBridge
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

// Node is a node in a topology
type Node struct {
	Name string
	Type NodeType
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
)

// String returns the link type as string
func (lt *LinkType) String() string {
	switch *lt {
	case LinkTypeVeth:
		return "veth"
	}
	return ""
}

// Link is a link between nodes in a topology
type Link struct {
	Name  string
	Type  LinkType
	Nodes [2]*Node
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

package cmd

// NodeType is the type of a node
type NodeType uint8

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

// NewTopology returns a new Topology
func NewTopology() *Topology {
	return &Topology{}
}

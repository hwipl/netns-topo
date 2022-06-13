package cmd

import "gopkg.in/yaml.v3"

// YAMLNodeType is a yaml representation of a node type
type YAMLNodeType string

// YAMLNode is a yaml representation of a node
type YAMLNode struct {
	Name string
	Type YAMLNodeType
}

// YAMLLinkType is a yaml representation of a link type
type YAMLLinkType string

// YAMLLink is a yaml representation of a link
type YAMLLink struct {
	Name  string
	Type  YAMLLinkType
	Nodes [2]string
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

// NewYAMLTopology returns a new YAMLTopology
func NewYAMLTopology() *YAMLTopology {
	return &YAMLTopology{}
}

package parser

import (
	"fmt"
	"strings"
)

type NodeType string

type Node struct {
	nodeType    NodeType
	parent      *Node
	children    []*Node
	value       interface{}
	startOffset int
	length      int
}

func NewNode(typ NodeType, value interface{}, startOffset, length int) *Node {
	return &Node{
		nodeType:    typ,
		parent:      nil,
		children:    make([]*Node, 0),
		value:       value,
		startOffset: startOffset,
		length:      length,
	}
}

func (s *Node) Type() NodeType {
	return s.nodeType
}

func (s *Node) Parent() *Node {
	return s.parent
}

func (s *Node) IsRoot() bool {
	return s.parent == nil
}

func (s *Node) Root() *Node {
	if s.IsRoot() {
		return s
	}
	return s.Parent().Root()
}

func (s *Node) Depth() int {
	if s.IsRoot() {
		return 0
	}
	return s.Parent().Depth() + 1
}

func (s *Node) IsLeaf() bool {
	return len(s.children) == 0
}

func (s *Node) Children() []*Node {
	return s.children
}

func (s *Node) AddChild(typ NodeType, value interface{}, startOffset, length int) *Node {
	child := Node{
		nodeType:    typ,
		parent:      s,
		children:    make([]*Node, 0),
		value:       value,
		startOffset: startOffset,
		length:      length,
	}
	s.children = append(s.children, &child)
	return &child
}

func (s *Node) Value() interface{} {
	return s.value
}

func DumpTree(node *Node) {
	fmt.Printf("%s[%s] %v\n", strings.Repeat("  ", node.Depth()), node.nodeType, node.value)
	for _, child := range node.children {
		DumpTree(child)
	}
}

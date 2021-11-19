package parser

import (
	"fmt"
	"github.com/mtrense/parsertk/lexer"
	"strings"
)

type Bracket struct {
	StartOffset int
	Length      int
}

type INode interface {
	Parent() INode
	Root() INode
	IsRoot() bool
	Depth() int
	Children() []INode
	IsLeaf() bool
	Bracket() Bracket
	AddChild(child ...INode)
	NodeType() string
	String() string
}

type Node struct {
	parent   INode
	children []INode
	bracket  Bracket
}

func NewNode(parent INode, tok lexer.Token) Node {
	return Node{
		parent:   parent,
		children: make([]INode, 0),
		bracket:  Bracket{StartOffset: tok.Offset, Length: tok.Length()},
	}
}

func (s *Node) Parent() INode {
	return s.parent
}

func (s *Node) IsRoot() bool {
	return s.parent == nil
}

func (s *Node) Root() INode {
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

func (s *Node) Children() []INode {
	return s.children
}

func (s *Node) IsLeaf() bool {
	return len(s.children) == 0
}

func (s *Node) Bracket() Bracket {
	return s.bracket
}

func (s *Node) NodeType() string {
	return ""
}

func (s *Node) String() string {
	return ""
}

func (s *Node) AddChild(child ...INode) {
	s.children = append(s.children, child...)
}

func DumpTree(node INode) {
	fmt.Printf("%s[%s] %v\n", strings.Repeat("  ", node.Depth()), node.NodeType(), node.String())
	for _, child := range node.Children() {
		DumpTree(child)
	}
}

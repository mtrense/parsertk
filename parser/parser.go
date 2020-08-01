package parser

import (
	"fmt"

	"github.com/mtrense/parsertk/lexer"
)

type NodeFactory func(cn *Node, tok lexer.Token) *Node

type Parser struct {
	rootNode      *Node
	currentNode   *Node
	nodeFactories map[lexer.TokenType]NodeFactory
}

var Ignore = func(cn *Node, tok lexer.Token) *Node {
	return nil
}

func NewParser(rootType NodeType) Parser {
	root := NewNode(rootType, nil, 0, 0)
	return Parser{
		rootNode:      root,
		currentNode:   root,
		nodeFactories: make(map[lexer.TokenType]NodeFactory),
	}
}

func (s *Parser) RegisterFactory(typ lexer.TokenType, factory NodeFactory) *Parser {
	s.nodeFactories[typ] = factory
	return s
}

func (s *Parser) Visit(tok lexer.Token) {
	factory, ok := s.nodeFactories[tok.Typ]
	if ok {
		node := factory(s.currentNode, tok)
		if node != nil {
			s.currentNode = node
		}
	} else {
		fmt.Printf("No factory found for TokenType %s\n", tok.Typ)
	}
}

func (s *Parser) RootNode() *Node {
	return s.rootNode
}

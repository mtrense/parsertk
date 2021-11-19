package parser

import (
	"fmt"

	"github.com/mtrense/parsertk/lexer"
)

type NodeFactory func(cn INode, tok lexer.Token) INode

type Parser struct {
	rootNode      INode
	currentNode   INode
	nodeFactories map[lexer.TokenType]NodeFactory
}

var Ignore NodeFactory = func(cn INode, tok lexer.Token) INode {
	return nil
}

func NewParser(rootNode INode) *Parser {
	return &Parser{
		rootNode:      rootNode,
		currentNode:   rootNode,
		nodeFactories: make(map[lexer.TokenType]NodeFactory),
	}
}

func (s *Parser) RegisterFactory(typ lexer.TokenType, factory NodeFactory) *Parser {
	s.nodeFactories[typ] = factory
	return s
}

// Visit implements lexer.LexerVisitor
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

func (s *Parser) RootNode() INode {
	return s.rootNode
}

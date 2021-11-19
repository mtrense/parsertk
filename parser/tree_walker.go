package parser

type TreeWalker func(n *Node)

func (s *Node) Walk(walker TreeWalker) {

}


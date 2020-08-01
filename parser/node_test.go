package parser

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Node", func() {
	It("is root when it has no parent", func() {
		subject := Node{
			nodeType: "TEST",
			parent:   nil,
		}
		Expect(subject.IsRoot()).To(BeTrue())
	})
	It("is not root when it has a parent", func() {
		root := &Node{
			nodeType: "TEST",
			parent:   nil,
		}
		subject := Node{
			nodeType: "TEST",
			parent:   root,
		}
		Expect(subject.IsRoot()).To(BeFalse())
	})
	It("returns the root node", func() {
		subject := Node{
			nodeType: "TEST",
			parent:   nil,
		}
		root := &Node{
			nodeType: "TEST",
			parent:   nil,
			children: []*Node{
				&subject,
			},
		}
		subject.parent = root
		Expect(subject.Root()).To(Equal(root))
	})
	It("returns the correct depth", func() {
		subject := Node{
			nodeType: "TEST",
			parent:   nil,
		}
		root := &Node{
			nodeType: "TEST",
			parent:   nil,
			children: []*Node{
				&subject,
			},
		}
		subject.parent = root
		Expect(subject.Depth()).To(Equal(1))
		Expect(root.Depth()).To(Equal(0))
	})
	It("is leaf when it has no children", func() {
		subject := Node{
			nodeType: "TEST",
			children: make([]*Node, 0),
		}
		Expect(subject.IsLeaf()).To(BeTrue())
	})
	It("is no leaf when it has children", func() {
		child := &Node{
			nodeType: "TEST",
		}
		subject := Node{
			nodeType: "TEST",
			children: []*Node{child},
		}
		Expect(subject.IsLeaf()).To(BeFalse())
	})
	It("creates children with AddChild", func() {
		subject := &Node{
			nodeType: "TEST",
		}
		child1 := subject.AddChild("TEST", nil, 0, 0)
		child2 := subject.AddChild("TEST", nil, 0, 0)
		Expect(subject.IsLeaf()).To(BeFalse())
		Expect(subject.IsRoot()).To(BeTrue())
		Expect(len(subject.Children())).To(Equal(2))
		Expect(child1.IsLeaf()).To(BeTrue())
		Expect(child1.Parent()).To(Equal(subject))
		Expect(child2.IsLeaf()).To(BeTrue())
		Expect(child2.Parent()).To(Equal(subject))
	})
})

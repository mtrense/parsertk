package lexer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("stringReader", func() {
	It("reports EOF for an empty string", func() {
		Expect(StringReader("").EOF()).To(BeTrue())
	})
	It("returns unicode replacement char for an empty string on peek", func() {
		Expect(StringReader("").Peek()).To(Equal('\uFFFD'))
	})
	It("returns unicode replacement char for an empty string on read", func() {
		Expect(StringReader("").Read()).To(Equal('\uFFFD'))
	})
	It("does not report EOF for a non-empty string", func() {
		Expect(StringReader("abc").EOF()).To(BeFalse())
	})
	It("does not report an error on empty string", func() {
		Expect(StringReader("abc").Error()).To(BeNil())
	})
	It("increments the offset on reading the next rune", func() {
		s := StringReader("abc")
		Expect(s.Read()).To(Equal('a'))
		Expect(s.(*stringReader).offset).To(Equal(1))
	})
	It("doesn't increment the offset on peeking the next rune", func() {
		s := StringReader("abc")
		Expect(s.Peek()).To(Equal('a'))
		Expect(s.Peek()).To(Equal('a'))
		Expect(s.(*stringReader).offset).To(Equal(0))
	})
	It("marks the current offset", func() {
		s := StringReader("abcdef")
		Expect(s.Read()).To(Equal('a'))
		Expect(s.Read()).To(Equal('b'))
		Expect(s.(*stringReader).offset).To(Equal(2))
		s.Mark()
		Expect(s.Read()).To(Equal('c'))
		Expect(s.Read()).To(Equal('d'))
		Expect(s.(*stringReader).offset).To(Equal(4))
		s.Mark()
		Expect(s.Read()).To(Equal('e'))
		Expect(s.Read()).To(Equal('f'))
		Expect(s.(*stringReader).offset).To(Equal(6))
		s.Rewind()
		Expect(s.(*stringReader).offset).To(Equal(4))
		s.Rewind()
		Expect(s.(*stringReader).offset).To(Equal(2))
	})
	It("reads the unicode replacement char on EOF", func() {
		s := StringReader("abc")
		Expect(s.Read()).To(Equal('a'))
		Expect(s.Read()).To(Equal('b'))
		Expect(s.Read()).To(Equal('c'))
		Expect(s.EOF()).To(BeTrue())
		Expect(s.Read()).To(Equal('\uFFFD'))
	})
})

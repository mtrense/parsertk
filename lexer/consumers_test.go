package lexer

import (
	"regexp"
	"unicode"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ConsumeSingleRune", func() {
	It("returns a valid token on successful tokenization", func() {
		reader := StringReader("(")
		tok, valid := ConsumeSingleRune(TokenTypeStart, '(')(reader)
		Expect(tok.Typ).To(Equal(TokenTypeStart))
		Expect(valid).To(BeTrue())
	})
	It("returns an invalid token on failed tokenization", func() {
		reader := StringReader(")")
		_, valid := ConsumeSingleRune(TokenTypeStart, '(')(reader)
		Expect(valid).To(BeFalse())
	})
	It("returns a valid token on successful tokenization when given multiple runes", func() {
		reader := StringReader("(")
		tok, valid := ConsumeSingleRune(TokenTypeStart, '(', ')')(reader)
		Expect(tok.Typ).To(Equal(TokenTypeStart))
		Expect(valid).To(BeTrue())
	})
})

var _ = Describe("ConsumeCharacterClass", func() {
	It("returns a valid token on successful tokenization", func() {
		reader := StringReader(" ")
		tok, valid := ConsumeCharacterClass(TokenTypeWhitespace, unicode.White_Space)(reader)
		Expect(tok.Typ).To(Equal(TokenTypeWhitespace))
		Expect(valid).To(BeTrue())
	})
	It("returns an invalid token on failed tokenization", func() {
		reader := StringReader("x")
		_, valid := ConsumeCharacterClass(TokenTypeWhitespace, unicode.White_Space)(reader)
		Expect(valid).To(BeFalse())
	})
})

var _ = Describe("ConsumeRunes", func() {
	It("returns a valid token on successful tokenization", func() {
		reader := StringReader("abcdef")
		tok, valid := ConsumeRunes(TokenTypeSymbol, "abc")(reader)
		Expect(tok.Typ).To(Equal(TokenTypeSymbol))
		Expect(valid).To(BeTrue())
	})
	It("returns an invalid token on failed tokenization", func() {
		reader := StringReader("def")
		_, valid := ConsumeRunes(TokenTypeSymbol, "abc")(reader)
		Expect(valid).To(BeFalse())
	})
})

var _ = Describe("ConsumeRegexpValidated", func() {
	It("returns a valid token on successful tokenization", func() {
		reader := StringReader("123.4567")
		tok, valid := ConsumeRegexpValidated(regexp.MustCompile("^(0|0\\.\\d+|[1-9]\\d*(\\.\\d+)?)$"), ConsumeRunes(TokenTypeNumber, "0123456789."))(reader)
		Expect(tok.Typ).To(Equal(TokenTypeNumber))
		Expect(valid).To(BeTrue())
		Expect(tok.Value).To(Equal("123.4567"))
	})
	It("returns a valid token on successful tokenization", func() {
		reader := StringReader("0")
		tok, valid := ConsumeRegexpValidated(regexp.MustCompile("^(0|0\\.\\d+|[1-9]\\d*(\\.\\d+)?)$"), ConsumeRunes(TokenTypeNumber, "0123456789."))(reader)
		Expect(tok.Typ).To(Equal(TokenTypeNumber))
		Expect(valid).To(BeTrue())
		Expect(tok.Value).To(Equal("0"))
	})
	It("returns a valid token on successful tokenization", func() {
		reader := StringReader("0.1234")
		tok, valid := ConsumeRegexpValidated(regexp.MustCompile("^(0|0\\.\\d+|[1-9]\\d*(\\.\\d+)?)$"), ConsumeRunes(TokenTypeNumber, "0123456789."))(reader)
		Expect(tok.Typ).To(Equal(TokenTypeNumber))
		Expect(valid).To(BeTrue())
		Expect(tok.Value).To(Equal("0.1234"))
	})
	It("returns a valid token on successful tokenization", func() {
		reader := StringReader("1234")
		tok, valid := ConsumeRegexpValidated(regexp.MustCompile("^(0|0\\.\\d+|[1-9]\\d*(\\.\\d+)?)$"), ConsumeRunes(TokenTypeNumber, "0123456789."))(reader)
		Expect(tok.Typ).To(Equal(TokenTypeNumber))
		Expect(valid).To(BeTrue())
		Expect(tok.Value).To(Equal("1234"))
	})
	It("returns invalid token on failed tokenization", func() {
		reader := StringReader("123.45.67")
		_, valid := ConsumeRegexpValidated(regexp.MustCompile("^(0|0\\.\\d+|[1-9]\\d*(\\.\\d+)?)$"), ConsumeRunes(TokenTypeNumber, "0123456789."))(reader)
		Expect(valid).To(BeFalse())
	})
	It("returns invalid token on failed tokenization", func() {
		reader := StringReader("012.3456")
		_, valid := ConsumeRegexpValidated(regexp.MustCompile("^(0|0\\.\\d+|[1-9]\\d*(\\.\\d+)?)$"), ConsumeRunes(TokenTypeNumber, "0123456789."))(reader)
		Expect(valid).To(BeFalse())
	})
	It("returns invalid token on failed tokenization", func() {
		reader := StringReader("0123456")
		_, valid := ConsumeRegexpValidated(regexp.MustCompile("^(0|0\\.\\d+|[1-9]\\d*(\\.\\d+)?)$"), ConsumeRunes(TokenTypeNumber, "0123456789."))(reader)
		Expect(valid).To(BeFalse())
	})
	It("returns invalid token on failed tokenization", func() {
		reader := StringReader(".0123456")
		_, valid := ConsumeRegexpValidated(regexp.MustCompile("^(0|0\\.\\d+|[1-9]\\d*(\\.\\d+)?)$"), ConsumeRunes(TokenTypeNumber, "0123456789."))(reader)
		Expect(valid).To(BeFalse())
	})
	It("returns invalid token on failed tokenization", func() {
		reader := StringReader("123456.")
		_, valid := ConsumeRegexpValidated(regexp.MustCompile("^(0|0\\.\\d+|[1-9]\\d*(\\.\\d+)?)$"), ConsumeRunes(TokenTypeNumber, "0123456789."))(reader)
		Expect(valid).To(BeFalse())
	})
})

var _ = Describe("ConsumeText", func() {
	It("returns a valid token on successful tokenization", func() {
		reader := StringReader("abcdef")
		tok, valid := ConsumeText(TokenTypeSymbol, "abc")(reader)
		Expect(tok.Typ).To(Equal(TokenTypeSymbol))
		Expect(valid).To(BeTrue())
	})
	It("returns an invalid token on failed tokenization", func() {
		reader := StringReader("def")
		_, valid := ConsumeText(TokenTypeSymbol, "abc")(reader)
		Expect(valid).To(BeFalse())
	})
})

var _ = Describe("ConsumeRegex", func() {
	It("returns a valid token on successful tokenization", func() {
		reader := StringReader("abbbcdef")
		tok, valid := ConsumeRegex(TokenTypeSymbol, regexp.MustCompile("ab+c"))(reader)
		Expect(tok.Typ).To(Equal(TokenTypeSymbol))
		Expect(tok.Value).To(Equal("abbbc"))
		Expect(valid).To(BeTrue())
	})
	It("returns an invalid token on failed tokenization", func() {
		reader := StringReader("acdef")
		_, valid := ConsumeRegex(TokenTypeSymbol, regexp.MustCompile("ab+c"))(reader)
		Expect(valid).To(BeFalse())
	})
})

var _ = Describe("ConsumeString", func() {
	It("returns a valid token on successful tokenization", func() {
		reader := StringReader("'abcdef'")
		tok, valid := ConsumeString(TokenTypeString)(reader)
		Expect(tok.Typ).To(Equal(TokenTypeString))
		Expect(tok.Value).To(Equal("abcdef"))
		Expect(valid).To(BeTrue())
	})
	It("returns an invalid token on failed tokenization", func() {
		reader := StringReader("acdef")
		_, valid := ConsumeString(TokenTypeString)(reader)
		Expect(valid).To(BeFalse())
	})
})

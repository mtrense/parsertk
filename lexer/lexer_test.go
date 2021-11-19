package lexer

import (
	"unicode"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type RecordingVisitor struct {
	tokens []Token
}

func (s *RecordingVisitor) visit(token Token) {
	s.tokens = append(s.tokens, token)
}

const (
	TokenTypeStart      TokenType = "START"
	TokenTypeWhitespace TokenType = "WS"
	TokenTypeSymbol     TokenType = "SYMBOL"
	TokenTypeString     TokenType = "STRING"
	TokenTypeNumber     TokenType = "NUMBER"
	TokenTypeBoolean    TokenType = "BOOL"
	TokenTypeEnd        TokenType = "END"
	TokenTypeEOF        TokenType = "EOF"
	TokenTypeError      TokenType = "ERR"
)

var SexpTokens = []TokenConsumer{
	ConsumeSingleRune(TokenTypeStart, '('),
	ConsumeSingleRune(TokenTypeEnd, ')'),
	ConsumeRunes(TokenTypeSymbol, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_$#%&"),
	ConsumeCharacterClass(TokenTypeWhitespace, unicode.White_Space),
	ConsumeString(TokenTypeString),
}

var _ = Describe("Lex", func() {
	var rv RecordingVisitor
	BeforeEach(func() {
		rv = RecordingVisitor{}
	})
	It("tokenizes an empty input string", func() {
		LexStatic(StringReader(""), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(len(rv.tokens)).To(Equal(1))
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeEOF, "").
			Build()))
	})
	It("tokenizes an empty expression", func() {
		LexStatic(StringReader("()"), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(len(rv.tokens)).To(Equal(3))
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeStart, "(").
			T(TokenTypeEnd, ")").
			T(TokenTypeEOF, "").
			Build()))
	})
	It("tokenizes a nested empty expression", func() {
		LexStatic(StringReader("(())"), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(len(rv.tokens)).To(Equal(5))
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeStart, "(").
			T(TokenTypeStart, "(").
			T(TokenTypeEnd, ")").
			T(TokenTypeEnd, ")").
			T(TokenTypeEOF, "").
			Build()))
	})
	It("tokenizes a series of spaces", func() {
		LexStatic(StringReader("   "), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeWhitespace, "   ").
			T(TokenTypeEOF, "").
			Build()))
	})
	It("tokenizes a series of different whitespace characters", func() {
		LexStatic(StringReader(" \n\t "), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeWhitespace, " \n\t ").
			T(TokenTypeEOF, "").
			Build()))
	})
	It("tokenizes a single symbol", func() {
		LexStatic(StringReader("test"), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeSymbol, "test").
			T(TokenTypeEOF, "").
			Build()))
	})
	It("tokenizes an expression with a single symbol", func() {
		LexStatic(StringReader("(test)"), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeStart, "(").
			T(TokenTypeSymbol, "test").
			T(TokenTypeEnd, ")").
			T(TokenTypeEOF, "").
			Build()))
	})
	It("tokenizes an expression with a single symbol and whitespace around it", func() {
		LexStatic(StringReader("( test )"), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeStart, "(").
			T(TokenTypeWhitespace, " ").
			T(TokenTypeSymbol, "test").
			T(TokenTypeWhitespace, " ").
			T(TokenTypeEnd, ")").
			T(TokenTypeEOF, "").
			Build()))
	})
	It("tokenizes an expression with a single symbol with special characters", func() {
		LexStatic(StringReader("($t%e&s0t9)"), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeStart, "(").
			T(TokenTypeSymbol, "$t%e&s0t9").
			T(TokenTypeEnd, ")").
			T(TokenTypeEOF, "").
			Build()))
	})
	It("tokenizes a single string", func() {
		LexStatic(StringReader("\"abcdef\""), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeString, "\"abcdef\"").
			T(TokenTypeEOF, "").
			Build()))
	})
	It("tokenizes a single string with escaped characters in it", func() {
		LexStatic(StringReader("\"abc\\\"def\""), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			TL(TokenTypeString, "\"abc\"def\"", 10).
			T(TokenTypeEOF, "").
			Build()))
	})
	It("produces an error token when string is not closed", func() {
		LexStatic(StringReader("\"abcdef"), (&rv).visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
		Expect(rv.tokens).To(Equal(NewTokenGenerator().
			T(TokenTypeError, "No valid token found").
			Build()))
	})
})

var _ = Describe("consumeString", func() {
	It("returns false when string is not started with ticks", func() {
		_, ok := ConsumeString(TokenTypeString)(StringReader("abc"))
		Expect(ok).To(BeFalse())
	})
})

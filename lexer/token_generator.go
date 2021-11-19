package lexer

type TokenGenerator struct {
	tokens []Token
	offset int
}

func NewTokenGenerator() *TokenGenerator {
	return &TokenGenerator{
		tokens: make([]Token, 0),
	}
}

func (s *TokenGenerator) T(typ TokenType, value string) *TokenGenerator {
	return s.TL(typ, value, len(value))
}

func (s *TokenGenerator) TL(typ TokenType, value string, lengthOverride int) *TokenGenerator {
	s.tokens = append(s.tokens, Token{
		Typ:    typ,
		Value:  value,
		Offset: s.offset,
		Line:   0,
		Column: 0,
	})
	s.offset += lengthOverride
	return s
}

func (s *TokenGenerator) Build() []Token {
	return s.tokens
}

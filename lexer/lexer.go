package lexer

type TokenType string

type Token struct {
	Typ    TokenType
	Value  string
	Offset int
	Line   int
	Column int
}

func (s *Token) Length() int {
	return len(s.Value)
}

func (s *Token) End() int {
	return s.Offset + len(s.Value)
}

type Visitor func(token Token)

type TokenConsumer func(BufferedRuneReader) (Token, bool)

// LexStatic scans the input using one fixed set of valid Tokens.
func LexStatic(input BufferedRuneReader, visitor Visitor, eofToken, errorToken TokenType, validTokens ...TokenConsumer) {
	for {
		var (
			tok   Token
			valid bool = false
		)
		if input.EOF() {
			tok = t(eofToken, "")
			tok.Offset = input.Offset()
			visitor(tok)
			break
		}
		for _, tokenConsumer := range validTokens {
			offset := input.Mark()
			tok, valid = tokenConsumer(input)
			if valid {
				tok.Offset = offset
				visitor(tok)
				break
			} else {
				input.Rewind()
			}
		}
		if !valid {
			tok = t(errorToken, "No valid token found")
			tok.Offset = input.Offset()
			visitor(tok)
			break
		}
	}
}

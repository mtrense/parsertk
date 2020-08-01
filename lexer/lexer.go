package lexer

type TokenType string

type Token struct {
	Typ    TokenType
	Value  string
	Offset uint
	Line   uint
	Column uint
}

type LexerVisitor func(token Token)

type TokenConsumer func(BufferedRuneReader) (Token, bool)

type LexerState interface {
}

func LexStatic(input BufferedRuneReader, visitor LexerVisitor, eofToken, errorToken TokenType, validTokens ...TokenConsumer) {
	for {
		var (
			tok   Token
			valid bool = false
		)
		if input.EOF() {
			visitor(t(eofToken, ""))
			break
		}
		for _, tokenConsumer := range validTokens {
			input.Mark()
			tok, valid = tokenConsumer(input)
			if valid {
				visitor(tok)
				break
			} else {
				input.Rewind()
			}
		}
		if !valid {
			visitor(t(errorToken, "No valid token found"))
			break
		}
	}
}

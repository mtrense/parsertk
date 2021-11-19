package lexer

import (
	"regexp"
	"strings"
	"unicode"
)

func ConsumeSingleRune(typ TokenType, expected ...rune) TokenConsumer {
	return func(input BufferedRuneReader) (Token, bool) {
		r := input.Read()
		for _, ex := range expected {
			if r == ex {
				return t(typ, string(r)), true
			}
		}
		return Token{}, false
	}
}

func ConsumeCharacterClass(typ TokenType, classes ...*unicode.RangeTable) TokenConsumer {
	return func(input BufferedRuneReader) (Token, bool) {
		var value strings.Builder
		for {
			if input.EOF() {
				break
			}
			r := input.Peek()
			if unicode.In(r, classes...) {
				value.WriteRune(input.Read())
			} else {
				break
			}
		}
		if value.Len() > 0 {
			return t(typ, value.String()), true
		}
		return Token{}, false
	}
}

func ConsumeRunes(typ TokenType, runeString string) TokenConsumer {
	return func(input BufferedRuneReader) (Token, bool) {
		var value strings.Builder
		for {
			if input.EOF() {
				break
			}
			r := input.Peek()
			if strings.IndexRune(runeString, r) != -1 {
				value.WriteRune(input.Read())
			} else {
				break
			}
		}
		if value.Len() > 0 {
			return t(typ, value.String()), true
		}
		return Token{}, false
	}
}

func ConsumeRegexpValidated(re *regexp.Regexp, consumer TokenConsumer) TokenConsumer {
	return func(input BufferedRuneReader) (Token, bool) {
		tok, valid := consumer(input)
		if !valid {
			return tok, valid
		}
		if re.MatchString(tok.Value) {
			return tok, true
		}
		return Token{}, false
	}
}

func ConsumeText(typ TokenType, text string) TokenConsumer {
	return func(input BufferedRuneReader) (Token, bool) {
		var value strings.Builder
		for _, expected := range []rune(text) {
			if input.Peek() == expected {
				value.WriteRune(input.Read())
			} else {
				return Token{}, false
			}
		}
		return t(typ, value.String()), true
	}
}

func ConsumeRegex(typ TokenType, re *regexp.Regexp) TokenConsumer {
	return func(input BufferedRuneReader) (Token, bool) {
		var value strings.Builder
		for {
			if input.EOF() {
				break
			}
			value.WriteRune(input.Read())
			if re.MatchString(value.String()) {
				return t(typ, value.String()), true
			}
		}
		return Token{}, false
	}
}

func ConsumeString(typ TokenType) TokenConsumer {
	return func(input BufferedRuneReader) (Token, bool) {
		var value strings.Builder
		delimiter := input.Read()
		if !(delimiter == '"' || delimiter == '\'') {
			return Token{}, false
		}
		value.WriteRune(delimiter)
		var escaped bool = false
		for {
			if input.EOF() {
				return Token{}, false
			}
			r := input.Read()
			if r == '\\' {
				escaped = true
				continue
			}
			if escaped {
				escaped = false
				if r == delimiter {
					value.WriteRune(r)
					continue
				}
			} else {
				if r == delimiter {
					value.WriteRune(delimiter)
					return t(typ, value.String()), true
				}
			}
			value.WriteRune(r)
		}
	}
}

func t(typ TokenType, value string) Token {
	return Token{
		Typ:   typ,
		Value: value,
	}
}

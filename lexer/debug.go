package lexer

import (
	"fmt"
	"github.com/fatih/color"
)

// Printer takes a token and returns a formatted string representing that Token.
type Printer func(tok Token) string

type DebugPrintingVisitor struct {
	printer Printer
}

func NewDebugPrintingVisitor(printer Printer) *DebugPrintingVisitor {
	return &DebugPrintingVisitor{
		printer: printer,
	}
}

// Visit implements LexerVisitor
func (s *DebugPrintingVisitor) Visit(token Token) {
	var printer Printer
	if s.printer != nil {
		printer = s.printer
	} else {
		printer = func(tok Token) string {
			return fmt.Sprintf("%s: '%#v'  [%d:%d]\n", tok.Typ, tok.Value, tok.Offset, tok.End())
		}
	}
	formattedToken := printer(token)
	if formattedToken != "" {
		fmt.Print(formattedToken)
	}
}

type ColorPrinter struct {
	colors map[TokenType]*color.Color
}

func NewColorPrinter() *ColorPrinter {
	return &ColorPrinter{
		colors: make(map[TokenType]*color.Color),
	}
}

func (s *ColorPrinter) Define(typ TokenType, c *color.Color) *ColorPrinter {
	s.colors[typ] = c
	return s
}

func (s *ColorPrinter) Printer(tok Token) string {
	if c, ok := s.colors[tok.Typ]; ok {
		return c.Sprint(tok.Value)
	}
	return tok.Value
}

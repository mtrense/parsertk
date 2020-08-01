package lexer

import (
	"fmt"
)

type Printer func(tok Token) string

type DebugPrintingVisitor struct {
	printer Printer
}

func NewDebugPrintingVisitor(printer Printer) *DebugPrintingVisitor {
	return &DebugPrintingVisitor{
		printer: printer,
	}
}

func (s *DebugPrintingVisitor) Visit(token Token) {
	var printer Printer
	if s.printer != nil {
		printer = s.printer
	} else {
		printer = func(tok Token) string {
			return fmt.Sprintf("%s: '%s'", tok.Typ, tok.Value)
		}
	}
	formattedToken := printer(token)
	if formattedToken != "" {
		fmt.Println(formattedToken)
	}
}

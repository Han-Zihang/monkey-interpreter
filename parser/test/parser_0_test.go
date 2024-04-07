package test

import (
	"monkey-interpreter/parser"
	"testing"
)

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errs := p.Errors()
	if len(errs) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errs))
	for _, msg := range errs {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

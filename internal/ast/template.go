package ast

import (
	"bytes"
	"github.com/inkbytefo/go-minus/internal/token"
	"strings"
)

// TemplateStatement, bir şablon tanımını temsil eder.
// Örnek: template<T> func add(a T, b T) T { return a + b; }
type TemplateStatement struct {
	Token          token.Token // token.TEMPLATE token'ı
	TypeParameters []*Identifier
	Node           Node // ClassStatement veya FunctionStatement
}

func (ts *TemplateStatement) statementNode()       {}
func (ts *TemplateStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *TemplateStatement) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range ts.TypeParameters {
		params = append(params, p.String())
	}

	out.WriteString(ts.TokenLiteral() + "<")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString("> ")
	out.WriteString(ts.Node.String())

	return out.String()
}

// Pos, düğümün konumunu döndürür.
func (ts *TemplateStatement) Pos() token.Position {
	return ts.Token.Position
}

// End, düğümün bitiş konumunu döndürür.
func (ts *TemplateStatement) End() token.Position {
	switch node := ts.Node.(type) {
	case *ClassStatement:
		return node.End()
	case *FunctionStatement:
		return node.End()
	default:
		return ts.Token.Position
	}
}

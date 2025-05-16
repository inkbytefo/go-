package ast

import (
	"bytes"
	"github.com/inkbytefo/go-minus/internal/token"
	"strings"
)

// FunctionStatement, bir fonksiyon tanımını temsil eder.
// Örnek: func add(x, y int) int { return x + y; }
type FunctionStatement struct {
	Token      token.Token // token.FUNCTION token'ı
	Name       *Identifier
	Parameters []*Identifier
	ReturnType Expression // Opsiyonel dönüş tipi
	Body       *BlockStatement
}

func (fs *FunctionStatement) statementNode()       {}
func (fs *FunctionStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *FunctionStatement) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fs.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fs.TokenLiteral() + " ")
	out.WriteString(fs.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")

	if fs.ReturnType != nil {
		out.WriteString(fs.ReturnType.String() + " ")
	}

	out.WriteString(fs.Body.String())

	return out.String()
}

// Pos, düğümün konumunu döndürür.
func (fs *FunctionStatement) Pos() token.Position {
	return fs.Token.Position
}

// End, düğümün bitiş konumunu döndürür.
func (fs *FunctionStatement) End() token.Position {
	if fs.Body != nil {
		return fs.Body.End()
	}
	return fs.Token.Position
}

package ast

import (
	"bytes"
	"github.com/inkbytefo/go-minus/internal/token"
)

// TryExpression, bir try ifadesini temsil eder.
// Örnek: try expr
type TryExpression struct {
	Token      token.Token // token.TRY token'ı
	Expression Expression
}

func (te *TryExpression) expressionNode()      {}
func (te *TryExpression) TokenLiteral() string { return te.Token.Literal }
func (te *TryExpression) String() string {
	var out bytes.Buffer

	out.WriteString("try ")
	out.WriteString(te.Expression.String())

	return out.String()
}

// Pos, düğümün konumunu döndürür.
func (te *TryExpression) Pos() token.Position {
	return te.Token.Position
}

// End, düğümün bitiş konumunu döndürür.
func (te *TryExpression) End() token.Position {
	return te.Expression.End()
}

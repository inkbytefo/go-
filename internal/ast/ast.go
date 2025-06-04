package ast

import (
	"bytes"
	"strings"

	"github.com/inkbytefo/go-minus/internal/token"
)

// Node, AST'deki her düğümün sahip olması gereken temel arayüzdür.
type Node interface {
	TokenLiteral() string // Düğümle ilişkili token'ın değişmez değerini döndürür
	String() string       // Hata ayıklama ve test için AST düğümünün okunabilir bir temsilini döndürür
	Pos() token.Position  // Düğümün konumunu döndürür
	End() token.Position  // Düğümün bitiş konumunu döndürür
}

// Statement, bir ifadeyi temsil eden bir AST düğümüdür.
type Statement interface {
	Node
	statementNode()
}

// Expression, bir değeri temsil eden bir AST düğümüdür.
type Expression interface {
	Node
	expressionNode()
}

// Program, bir GO+ programının kök AST düğümüdür.
// Her geçerli GO+ programı bir dizi ifadeden (Statement) oluşur.
type Program struct {
	Statements []Statement
}

// TokenLiteral, programın ilk ifadesinin token değişmez değerini döndürür (eğer varsa).
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String, programdaki tüm ifadelerin okunabilir bir temsilini döndürür.
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Pos, programın konumunu döndürür.
func (p *Program) Pos() token.Position {
	if len(p.Statements) > 0 {
		return p.Statements[0].Pos()
	}
	return token.Position{}
}

// End, programın bitiş konumunu döndürür.
func (p *Program) End() token.Position {
	if len(p.Statements) > 0 {
		return p.Statements[len(p.Statements)-1].End()
	}
	return token.Position{}
}

// Identifier, bir tanımlayıcıyı (değişken adı, fonksiyon adı vb.) temsil eder.
type Identifier struct {
	Token token.Token // token.IDENT token'ı
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
func (i *Identifier) Pos() token.Position  { return i.Token.Position }
func (i *Identifier) End() token.Position  { return i.Token.Position }

// VarStatement, bir değişken tanımlama ifadesini temsil eder.
// Örnek: var x int = 5
type VarStatement struct {
	Token token.Token // token.VAR token'ı
	Name  *Identifier
	Type  Expression // Opsiyonel tip
	Value Expression // Opsiyonel değer
}

func (vs *VarStatement) statementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }
func (vs *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vs.TokenLiteral() + " ")
	out.WriteString(vs.Name.String())

	if vs.Type != nil {
		out.WriteString(" " + vs.Type.String())
	}

	if vs.Value != nil {
		out.WriteString(" = ")
		out.WriteString(vs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
func (vs *VarStatement) Pos() token.Position { return vs.Token.Position }
func (vs *VarStatement) End() token.Position {
	if vs.Value != nil {
		return vs.Value.End()
	}
	if vs.Type != nil {
		return vs.Type.End()
	}
	return vs.Name.End()
}

// ConstStatement, bir sabit tanımlama ifadesini temsil eder.
// Örnek: const x = 5
type ConstStatement struct {
	Token token.Token // token.CONST token'ı
	Name  *Identifier
	Type  Expression // Opsiyonel tip
	Value Expression
}

func (cs *ConstStatement) statementNode()       {}
func (cs *ConstStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ConstStatement) String() string {
	var out bytes.Buffer

	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.Name.String())

	if cs.Type != nil {
		out.WriteString(" " + cs.Type.String())
	}

	out.WriteString(" = ")
	out.WriteString(cs.Value.String())
	out.WriteString(";")

	return out.String()
}
func (cs *ConstStatement) Pos() token.Position { return cs.Token.Position }
func (cs *ConstStatement) End() token.Position { return cs.Value.End() }

// ReturnStatement, bir dönüş ifadesini temsil eder.
// Örnek: return 5
type ReturnStatement struct {
	Token       token.Token // token.RETURN token'ı
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
func (rs *ReturnStatement) Pos() token.Position { return rs.Token.Position }
func (rs *ReturnStatement) End() token.Position {
	if rs.ReturnValue != nil {
		return rs.ReturnValue.End()
	}
	return rs.Token.Position
}

// ExpressionStatement, bir ifade cümlesini temsil eder.
// Örnek: x + 5
type ExpressionStatement struct {
	Token      token.Token // İfadenin ilk token'ı
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) Pos() token.Position { return es.Token.Position }
func (es *ExpressionStatement) End() token.Position {
	if es.Expression != nil {
		return es.Expression.End()
	}
	return es.Token.Position
}

// BlockStatement, bir blok ifadesini temsil eder.
// Örnek: { x = 5; y = 10; }
type BlockStatement struct {
	Token      token.Token // token.LBRACE token'ı
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	out.WriteString("{ ")
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	out.WriteString(" }")

	return out.String()
}
func (bs *BlockStatement) Pos() token.Position { return bs.Token.Position }
func (bs *BlockStatement) End() token.Position {
	if len(bs.Statements) > 0 {
		return bs.Statements[len(bs.Statements)-1].End()
	}
	return bs.Token.Position
}

// IntegerLiteral, bir tamsayı değişmez değerini temsil eder.
// Örnek: 5
type IntegerLiteral struct {
	Token token.Token // token.INT token'ı
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
func (il *IntegerLiteral) Pos() token.Position  { return il.Token.Position }
func (il *IntegerLiteral) End() token.Position  { return il.Token.Position }

// FloatLiteral, bir ondalık sayı değişmez değerini temsil eder.
// Örnek: 3.14
type FloatLiteral struct {
	Token token.Token // token.FLOAT token'ı
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return fl.Token.Literal }
func (fl *FloatLiteral) Pos() token.Position  { return fl.Token.Position }
func (fl *FloatLiteral) End() token.Position  { return fl.Token.Position }

// StringLiteral, bir string değişmez değerini temsil eder.
// Örnek: "hello"
type StringLiteral struct {
	Token token.Token // token.STRING token'ı
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return "\"" + sl.Value + "\"" }
func (sl *StringLiteral) Pos() token.Position  { return sl.Token.Position }
func (sl *StringLiteral) End() token.Position  { return sl.Token.Position }

// CharLiteral, bir karakter değişmez değerini temsil eder.
// Örnek: 'a'
type CharLiteral struct {
	Token token.Token // token.CHAR token'ı
	Value rune
}

func (cl *CharLiteral) expressionNode()      {}
func (cl *CharLiteral) TokenLiteral() string { return cl.Token.Literal }
func (cl *CharLiteral) String() string       { return "'" + string(cl.Value) + "'" }
func (cl *CharLiteral) Pos() token.Position  { return cl.Token.Position }
func (cl *CharLiteral) End() token.Position  { return cl.Token.Position }

// BooleanLiteral, bir boolean değişmez değerini temsil eder.
// Örnek: true, false
type BooleanLiteral struct {
	Token token.Token // token.TRUE veya token.FALSE token'ı
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) String() string {
	if bl.Value {
		return "true"
	}
	return "false"
}
func (bl *BooleanLiteral) Pos() token.Position { return bl.Token.Position }
func (bl *BooleanLiteral) End() token.Position { return bl.Token.Position }

// NullLiteral, bir null değişmez değerini temsil eder.
// Örnek: nil
type NullLiteral struct {
	Token token.Token // token.NULL token'ı
}

func (nl *NullLiteral) expressionNode()      {}
func (nl *NullLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NullLiteral) String() string       { return "nil" }
func (nl *NullLiteral) Pos() token.Position  { return nl.Token.Position }
func (nl *NullLiteral) End() token.Position  { return nl.Token.Position }

// PrefixExpression, bir önek ifadesini temsil eder.
// Örnek: !true, -5
type PrefixExpression struct {
	Token    token.Token // Önek operatörü token'ı
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
func (pe *PrefixExpression) Pos() token.Position { return pe.Token.Position }
func (pe *PrefixExpression) End() token.Position { return pe.Right.End() }

// InfixExpression, bir araek ifadesini temsil eder.
// Örnek: 5 + 5
type InfixExpression struct {
	Token    token.Token // Operatör token'ı
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
func (ie *InfixExpression) Pos() token.Position { return ie.Left.Pos() }
func (ie *InfixExpression) End() token.Position { return ie.Right.End() }

// PostfixExpression, bir postfix ifadesini temsil eder.
// Örnek: i++, j--
type PostfixExpression struct {
	Token    token.Token // Operatör token'ı
	Left     Expression
	Operator string
}

func (pe *PostfixExpression) expressionNode()      {}
func (pe *PostfixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PostfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Left.String())
	out.WriteString(pe.Operator)
	out.WriteString(")")

	return out.String()
}
func (pe *PostfixExpression) Pos() token.Position { return pe.Left.Pos() }
func (pe *PostfixExpression) End() token.Position { return pe.Token.Position }

// IfExpression, bir if ifadesini temsil eder.
// Örnek: if (x > y) { x } else { y }
type IfExpression struct {
	Token       token.Token // token.IF token'ı
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
func (ie *IfExpression) Pos() token.Position { return ie.Token.Position }
func (ie *IfExpression) End() token.Position {
	if ie.Alternative != nil {
		return ie.Alternative.End()
	}
	return ie.Consequence.End()
}

// FunctionLiteral, bir fonksiyon değişmez değerini temsil eder.
// Örnek: func(x, y) { return x + y; }
type FunctionLiteral struct {
	Token      token.Token // token.FUNCTION token'ı
	Parameters []*Identifier
	Body       *BlockStatement
	ReturnType Expression // Opsiyonel dönüş tipi
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")

	if fl.ReturnType != nil {
		out.WriteString(fl.ReturnType.String() + " ")
	}

	out.WriteString(fl.Body.String())

	return out.String()
}
func (fl *FunctionLiteral) Pos() token.Position { return fl.Token.Position }
func (fl *FunctionLiteral) End() token.Position { return fl.Body.End() }

// CallExpression, bir fonksiyon çağrısını temsil eder.
// Örnek: add(1, 2)
type CallExpression struct {
	Token     token.Token // token.LPAREN token'ı
	Function  Expression  // Identifier veya FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
func (ce *CallExpression) Pos() token.Position { return ce.Function.Pos() }
func (ce *CallExpression) End() token.Position {
	if len(ce.Arguments) > 0 {
		return ce.Arguments[len(ce.Arguments)-1].End()
	}
	return ce.Token.Position
}

// ArrayLiteral, bir dizi değişmez değerini temsil eder.
// Örnek: [1, 2, 3]
type ArrayLiteral struct {
	Token    token.Token // token.LBRACKET token'ı
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
func (al *ArrayLiteral) Pos() token.Position { return al.Token.Position }
func (al *ArrayLiteral) End() token.Position {
	if len(al.Elements) > 0 {
		return al.Elements[len(al.Elements)-1].End()
	}
	return al.Token.Position
}

// IndexExpression, bir dizin ifadesini temsil eder.
// Örnek: myArray[1]
type IndexExpression struct {
	Token token.Token // token.LBRACKET token'ı
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
func (ie *IndexExpression) Pos() token.Position { return ie.Left.Pos() }
func (ie *IndexExpression) End() token.Position { return ie.Index.End() }

// ArrayType, bir dizi tipini temsil eder.
// Örnek: [5]int, []int
type ArrayType struct {
	Token       token.Token // token.LBRACKET token'ı
	Size        Expression  // Array size (nil for slices)
	ElementType Expression  // Element type
}

func (at *ArrayType) expressionNode()      {}
func (at *ArrayType) TokenLiteral() string { return at.Token.Literal }
func (at *ArrayType) String() string {
	var out bytes.Buffer

	out.WriteString("[")
	if at.Size != nil {
		out.WriteString(at.Size.String())
	}
	out.WriteString("]")
	out.WriteString(at.ElementType.String())

	return out.String()
}
func (at *ArrayType) Pos() token.Position { return at.Token.Position }
func (at *ArrayType) End() token.Position { return at.ElementType.End() }

// HashLiteral, bir hash değişmez değerini temsil eder.
// Örnek: {"one": 1, "two": 2}
type HashLiteral struct {
	Token token.Token // token.LBRACE token'ı
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
func (hl *HashLiteral) Pos() token.Position { return hl.Token.Position }
func (hl *HashLiteral) End() token.Position { return hl.Token.Position }

// ForStatement, bir for döngüsünü temsil eder.
// Örnek: for i := 0; i < 10; i++ { ... }
type ForStatement struct {
	Token     token.Token // token.FOR token'ı
	Init      Statement   // Opsiyonel başlangıç ifadesi
	Condition Expression  // Opsiyonel koşul
	Post      Statement   // Opsiyonel sonrası ifadesi
	Body      *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out bytes.Buffer

	out.WriteString("for ")

	if fs.Init != nil {
		out.WriteString(fs.Init.String())
	}
	out.WriteString("; ")

	if fs.Condition != nil {
		out.WriteString(fs.Condition.String())
	}
	out.WriteString("; ")

	if fs.Post != nil {
		out.WriteString(fs.Post.String())
	}

	out.WriteString(" ")
	out.WriteString(fs.Body.String())

	return out.String()
}
func (fs *ForStatement) Pos() token.Position { return fs.Token.Position }
func (fs *ForStatement) End() token.Position { return fs.Body.End() }

// WhileStatement, bir while döngüsünü temsil eder.
// Örnek: while (x < 10) { ... }
type WhileStatement struct {
	Token     token.Token // token.WHILE token'ı
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WhileStatement) String() string {
	var out bytes.Buffer

	out.WriteString("while ")
	out.WriteString(ws.Condition.String())
	out.WriteString(" ")
	out.WriteString(ws.Body.String())

	return out.String()
}
func (ws *WhileStatement) Pos() token.Position { return ws.Token.Position }
func (ws *WhileStatement) End() token.Position { return ws.Body.End() }

// SwitchStatement, bir switch ifadesini temsil eder.
// Örnek: switch x { case 1: ... case 2: ... default: ... }
type SwitchStatement struct {
	Token token.Token // token.SWITCH token'ı
	Tag   Expression  // Opsiyonel switch ifadesi (nil olabilir)
	Cases []*CaseClause
}

func (ss *SwitchStatement) statementNode()       {}
func (ss *SwitchStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SwitchStatement) String() string {
	var out bytes.Buffer

	out.WriteString("switch ")
	if ss.Tag != nil {
		out.WriteString(ss.Tag.String())
		out.WriteString(" ")
	}
	out.WriteString("{\n")

	for _, c := range ss.Cases {
		out.WriteString(c.String())
	}

	out.WriteString("}")

	return out.String()
}
func (ss *SwitchStatement) Pos() token.Position { return ss.Token.Position }
func (ss *SwitchStatement) End() token.Position {
	if len(ss.Cases) > 0 {
		return ss.Cases[len(ss.Cases)-1].End()
	}
	return ss.Token.Position
}

// CaseClause, bir case veya default bloğunu temsil eder.
// Örnek: case 1, 2: ... veya default: ...
type CaseClause struct {
	Token  token.Token  // token.CASE veya token.DEFAULT token'ı
	Values []Expression // Case değerleri (default için nil)
	Body   []Statement  // Case gövdesi
}

func (cc *CaseClause) statementNode()       {}
func (cc *CaseClause) TokenLiteral() string { return cc.Token.Literal }
func (cc *CaseClause) String() string {
	var out bytes.Buffer

	if cc.Token.Type == token.DEFAULT {
		out.WriteString("default:")
	} else {
		out.WriteString("case ")
		for i, v := range cc.Values {
			if i > 0 {
				out.WriteString(", ")
			}
			out.WriteString(v.String())
		}
		out.WriteString(":")
	}

	out.WriteString("\n")
	for _, stmt := range cc.Body {
		out.WriteString("\t")
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}

	return out.String()
}
func (cc *CaseClause) Pos() token.Position { return cc.Token.Position }
func (cc *CaseClause) End() token.Position {
	if len(cc.Body) > 0 {
		return cc.Body[len(cc.Body)-1].End()
	}
	return cc.Token.Position
}

// ClassStatement, bir sınıf tanımını temsil eder.
// Örnek: class Person { ... }
type ClassStatement struct {
	Token      token.Token // token.CLASS token'ı
	Name       *Identifier
	Extends    *Identifier   // Opsiyonel kalıtım
	Implements []*Identifier // Opsiyonel arayüz uygulamaları
	Body       *BlockStatement
}

func (cs *ClassStatement) statementNode()       {}
func (cs *ClassStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ClassStatement) String() string {
	var out bytes.Buffer

	out.WriteString("class ")
	out.WriteString(cs.Name.String())

	if cs.Extends != nil {
		out.WriteString(" extends ")
		out.WriteString(cs.Extends.String())
	}

	if len(cs.Implements) > 0 {
		out.WriteString(" implements ")
		impls := []string{}
		for _, impl := range cs.Implements {
			impls = append(impls, impl.String())
		}
		out.WriteString(strings.Join(impls, ", "))
	}

	out.WriteString(" ")
	out.WriteString(cs.Body.String())

	return out.String()
}
func (cs *ClassStatement) Pos() token.Position { return cs.Token.Position }
func (cs *ClassStatement) End() token.Position { return cs.Body.End() }

// MethodStatement, bir metot tanımını temsil eder.
// Örnek: func (p Person) sayHello() { ... }
type MethodStatement struct {
	Token      token.Token // token.FUNC token'ı
	Receiver   *Identifier
	Name       *Identifier
	Parameters []*Identifier
	ReturnType Expression // Opsiyonel dönüş tipi
	Body       *BlockStatement
}

func (ms *MethodStatement) statementNode()       {}
func (ms *MethodStatement) TokenLiteral() string { return ms.Token.Literal }
func (ms *MethodStatement) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range ms.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(ms.TokenLiteral())
	out.WriteString(" (")
	out.WriteString(ms.Receiver.String())
	out.WriteString(") ")
	out.WriteString(ms.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")

	if ms.ReturnType != nil {
		out.WriteString(ms.ReturnType.String() + " ")
	}

	out.WriteString(ms.Body.String())

	return out.String()
}
func (ms *MethodStatement) Pos() token.Position { return ms.Token.Position }
func (ms *MethodStatement) End() token.Position { return ms.Body.End() }

// TryCatchStatement, bir try-catch ifadesini temsil eder.
// Örnek: try { ... } catch (e Error) { ... } finally { ... }
type TryCatchStatement struct {
	Token   token.Token // token.TRY token'ı
	Try     *BlockStatement
	Catches []*CatchClause
	Finally *BlockStatement // Opsiyonel finally bloğu
}

func (tcs *TryCatchStatement) statementNode()       {}
func (tcs *TryCatchStatement) TokenLiteral() string { return tcs.Token.Literal }
func (tcs *TryCatchStatement) String() string {
	var out bytes.Buffer

	out.WriteString("try ")
	out.WriteString(tcs.Try.String())

	for _, catch := range tcs.Catches {
		out.WriteString(" ")
		out.WriteString(catch.String())
	}

	if tcs.Finally != nil {
		out.WriteString(" finally ")
		out.WriteString(tcs.Finally.String())
	}

	return out.String()
}
func (tcs *TryCatchStatement) Pos() token.Position { return tcs.Token.Position }
func (tcs *TryCatchStatement) End() token.Position {
	if tcs.Finally != nil {
		return tcs.Finally.End()
	}
	if len(tcs.Catches) > 0 {
		return tcs.Catches[len(tcs.Catches)-1].End()
	}
	return tcs.Try.End()
}

// CatchClause, bir catch bloğunu temsil eder.
// Örnek: catch (e Error) { ... }
type CatchClause struct {
	Token     token.Token // token.CATCH token'ı
	Parameter *Identifier // Opsiyonel parametre
	Type      Expression  // Opsiyonel tip
	Body      *BlockStatement
}

func (cc *CatchClause) expressionNode()      {}
func (cc *CatchClause) TokenLiteral() string { return cc.Token.Literal }
func (cc *CatchClause) String() string {
	var out bytes.Buffer

	out.WriteString("catch")

	if cc.Parameter != nil {
		out.WriteString(" (")
		out.WriteString(cc.Parameter.String())

		if cc.Type != nil {
			out.WriteString(" ")
			out.WriteString(cc.Type.String())
		}

		out.WriteString(")")
	}

	out.WriteString(" ")
	out.WriteString(cc.Body.String())

	return out.String()
}
func (cc *CatchClause) Pos() token.Position { return cc.Token.Position }
func (cc *CatchClause) End() token.Position { return cc.Body.End() }

// ThrowStatement, bir throw ifadesini temsil eder.
// Örnek: throw new Error("message")
type ThrowStatement struct {
	Token token.Token // token.THROW token'ı
	Value Expression
}

func (ts *ThrowStatement) statementNode()       {}
func (ts *ThrowStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *ThrowStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ts.TokenLiteral() + " ")
	out.WriteString(ts.Value.String())
	out.WriteString(";")

	return out.String()
}
func (ts *ThrowStatement) Pos() token.Position { return ts.Token.Position }
func (ts *ThrowStatement) End() token.Position { return ts.Value.End() }

// ScopeStatement, bir scope ifadesini temsil eder.
// Örnek: scope { ... }
type ScopeStatement struct {
	Token token.Token // token.SCOPE token'ı
	Body  *BlockStatement
}

func (ss *ScopeStatement) statementNode()       {}
func (ss *ScopeStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *ScopeStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ss.TokenLiteral() + " ")
	out.WriteString(ss.Body.String())

	return out.String()
}
func (ss *ScopeStatement) Pos() token.Position { return ss.Token.Position }
func (ss *ScopeStatement) End() token.Position { return ss.Body.End() }

// TemplateExpression, bir şablon ifadesini temsil eder.
// Örnek: template<T> func add(a T, b T) T { return a + b; }
type TemplateExpression struct {
	Token      token.Token // token.TEMPLATE token'ı
	Parameters []*Identifier
	Body       Expression // FunctionLiteral, ClassStatement vb.
}

func (te *TemplateExpression) expressionNode()      {}
func (te *TemplateExpression) TokenLiteral() string { return te.Token.Literal }
func (te *TemplateExpression) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range te.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("template<")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString("> ")
	out.WriteString(te.Body.String())

	return out.String()
}
func (te *TemplateExpression) Pos() token.Position { return te.Token.Position }
func (te *TemplateExpression) End() token.Position { return te.Body.End() }

// NewExpression, bir new ifadesini temsil eder.
// Örnek: new Person("John", 30)
type NewExpression struct {
	Token     token.Token // token.NEW token'ı
	Class     Expression  // Sınıf adı
	Arguments []Expression
}

func (ne *NewExpression) expressionNode()      {}
func (ne *NewExpression) TokenLiteral() string { return ne.Token.Literal }
func (ne *NewExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ne.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ne.TokenLiteral() + " ")
	out.WriteString(ne.Class.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
func (ne *NewExpression) Pos() token.Position { return ne.Token.Position }
func (ne *NewExpression) End() token.Position {
	if len(ne.Arguments) > 0 {
		return ne.Arguments[len(ne.Arguments)-1].End()
	}
	return ne.Class.End()
}

// MemberExpression, bir üye erişim ifadesini temsil eder.
// Örnek: person.name, person->name
type MemberExpression struct {
	Token  token.Token // token.DOT veya token.ARROW token'ı
	Object Expression
	Member Expression // Genellikle bir Identifier
}

func (me *MemberExpression) expressionNode()      {}
func (me *MemberExpression) TokenLiteral() string { return me.Token.Literal }
func (me *MemberExpression) String() string {
	var out bytes.Buffer

	out.WriteString(me.Object.String())

	if me.Token.Literal == "->" {
		out.WriteString("->")
	} else {
		out.WriteString(".")
	}

	out.WriteString(me.Member.String())

	return out.String()
}
func (me *MemberExpression) Pos() token.Position { return me.Object.Pos() }
func (me *MemberExpression) End() token.Position { return me.Member.End() }

// PackageStatement, bir paket bildirimini temsil eder.
// Örnek: package main
type PackageStatement struct {
	Token token.Token // token.PACKAGE token'ı
	Name  *Identifier
}

func (ps *PackageStatement) statementNode()       {}
func (ps *PackageStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *PackageStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ps.TokenLiteral() + " ")
	out.WriteString(ps.Name.String())
	return out.String()
}
func (ps *PackageStatement) Pos() token.Position { return ps.Token.Position }
func (ps *PackageStatement) End() token.Position { return ps.Name.End() }

// ImportStatement, bir import bildirimini temsil eder.
// Örnek: import "fmt"
type ImportStatement struct {
	Token token.Token // token.IMPORT token'ı
	Path  *StringLiteral
}

func (is *ImportStatement) statementNode()       {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Literal }
func (is *ImportStatement) String() string {
	var out bytes.Buffer
	out.WriteString(is.TokenLiteral() + " ")
	out.WriteString(is.Path.String())
	return out.String()
}
func (is *ImportStatement) Pos() token.Position { return is.Token.Position }
func (is *ImportStatement) End() token.Position { return is.Path.End() }

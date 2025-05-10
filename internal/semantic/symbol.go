package semantic

import (
	"fmt"
	"goplus/internal/token"
)

// SymbolType, bir sembolün tipini temsil eder.
type SymbolType int

const (
	UNKNOWN_TYPE SymbolType = iota
	INTEGER_TYPE
	FLOAT_TYPE
	STRING_TYPE
	BOOLEAN_TYPE
	CHAR_TYPE
	NULL_TYPE
	ARRAY_TYPE
	HASH_TYPE
	FUNCTION_TYPE
	CLASS_TYPE
	INTERFACE_TYPE
	STRUCT_TYPE
	MAP_TYPE
	CHAN_TYPE
	TEMPLATE_TYPE
	VOID_TYPE
)

// String, sembol tipinin string temsilini döndürür.
func (st SymbolType) String() string {
	switch st {
	case INTEGER_TYPE:
		return "int"
	case FLOAT_TYPE:
		return "float"
	case STRING_TYPE:
		return "string"
	case BOOLEAN_TYPE:
		return "bool"
	case CHAR_TYPE:
		return "char"
	case NULL_TYPE:
		return "null"
	case ARRAY_TYPE:
		return "array"
	case HASH_TYPE:
		return "hash"
	case FUNCTION_TYPE:
		return "function"
	case CLASS_TYPE:
		return "class"
	case INTERFACE_TYPE:
		return "interface"
	case STRUCT_TYPE:
		return "struct"
	case MAP_TYPE:
		return "map"
	case CHAN_TYPE:
		return "chan"
	case TEMPLATE_TYPE:
		return "template"
	case VOID_TYPE:
		return "void"
	default:
		return "unknown"
	}
}

// Symbol, bir sembolü temsil eder.
type Symbol struct {
	Name      string
	Type      SymbolType
	Scope     *Scope
	Token     token.Token
	IsConst   bool
	Value     interface{}
	Signature *FunctionSignature // Fonksiyonlar için
	Class     *ClassInfo         // Sınıflar için
}

// FunctionSignature, bir fonksiyonun imzasını temsil eder.
type FunctionSignature struct {
	Parameters []*Symbol
	ReturnType SymbolType
}

// ClassInfo, bir sınıfın bilgilerini temsil eder.
type ClassInfo struct {
	Fields     map[string]*Symbol
	Methods    map[string]*Symbol
	Extends    *Symbol
	Implements []*Symbol
}

// Scope, bir kapsamı temsil eder.
type Scope struct {
	Parent    *Scope
	Symbols   map[string]*Symbol
	Children  []*Scope
	IsGlobal  bool
	IsClass   bool
	ClassName string
}

// NewScope, yeni bir kapsam oluşturur.
func NewScope(parent *Scope) *Scope {
	s := &Scope{
		Parent:   parent,
		Symbols:  make(map[string]*Symbol),
		Children: []*Scope{},
		IsGlobal: false,
	}

	if parent != nil {
		parent.Children = append(parent.Children, s)
	}

	return s
}

// Define, bir sembolü tanımlar.
func (s *Scope) Define(name string, symbolType SymbolType, tok token.Token) *Symbol {
	symbol := &Symbol{
		Name:  name,
		Type:  symbolType,
		Scope: s,
		Token: tok,
	}

	s.Symbols[name] = symbol
	return symbol
}

// Resolve, bir sembolü çözümler.
func (s *Scope) Resolve(name string) *Symbol {
	symbol, ok := s.Symbols[name]
	if ok {
		return symbol
	}

	if s.Parent != nil {
		return s.Parent.Resolve(name)
	}

	return nil
}

// String, kapsamın string temsilini döndürür.
func (s *Scope) String() string {
	var scopeType string
	if s.IsGlobal {
		scopeType = "Global"
	} else if s.IsClass {
		scopeType = fmt.Sprintf("Class(%s)", s.ClassName)
	} else {
		scopeType = "Local"
	}

	result := fmt.Sprintf("Scope(%s):\n", scopeType)

	for name, symbol := range s.Symbols {
		result += fmt.Sprintf("  %s: %s\n", name, symbol.Type)
	}

	for _, child := range s.Children {
		result += child.String()
	}

	return result
}

package semantic

import (
	"fmt"
)

// Type, bir tipi temsil eder.
type Type interface {
	String() string
	Equals(Type) bool
}

// BasicType, temel bir tipi temsil eder.
type BasicType struct {
	Name string
	Kind SymbolType
}

// String, temel tipin string temsilini döndürür.
func (bt *BasicType) String() string {
	return bt.Name
}

// Equals, iki temel tipin eşit olup olmadığını kontrol eder.
func (bt *BasicType) Equals(other Type) bool {
	if otherBasic, ok := other.(*BasicType); ok {
		return bt.Kind == otherBasic.Kind
	}
	return false
}

// ArrayType, bir dizi tipini temsil eder.
type ArrayType struct {
	ElementType Type
}

// String, dizi tipinin string temsilini döndürür.
func (at *ArrayType) String() string {
	return fmt.Sprintf("[]%s", at.ElementType.String())
}

// Equals, iki dizi tipinin eşit olup olmadığını kontrol eder.
func (at *ArrayType) Equals(other Type) bool {
	if otherArray, ok := other.(*ArrayType); ok {
		return at.ElementType.Equals(otherArray.ElementType)
	}
	return false
}

// MapType, bir map tipini temsil eder.
type MapType struct {
	KeyType   Type
	ValueType Type
}

// String, map tipinin string temsilini döndürür.
func (mt *MapType) String() string {
	return fmt.Sprintf("map[%s]%s", mt.KeyType.String(), mt.ValueType.String())
}

// Equals, iki map tipinin eşit olup olmadığını kontrol eder.
func (mt *MapType) Equals(other Type) bool {
	if otherMap, ok := other.(*MapType); ok {
		return mt.KeyType.Equals(otherMap.KeyType) && mt.ValueType.Equals(otherMap.ValueType)
	}
	return false
}

// FunctionType, bir fonksiyon tipini temsil eder.
type FunctionType struct {
	ParameterTypes []Type
	ReturnType     Type
}

// String, fonksiyon tipinin string temsilini döndürür.
func (ft *FunctionType) String() string {
	result := "func("

	for i, paramType := range ft.ParameterTypes {
		if i > 0 {
			result += ", "
		}
		result += paramType.String()
	}

	result += ")"

	if ft.ReturnType != nil {
		result += " " + ft.ReturnType.String()
	}

	return result
}

// Equals, iki fonksiyon tipinin eşit olup olmadığını kontrol eder.
func (ft *FunctionType) Equals(other Type) bool {
	if otherFunc, ok := other.(*FunctionType); ok {
		if len(ft.ParameterTypes) != len(otherFunc.ParameterTypes) {
			return false
		}

		for i, paramType := range ft.ParameterTypes {
			if !paramType.Equals(otherFunc.ParameterTypes[i]) {
				return false
			}
		}

		if ft.ReturnType == nil && otherFunc.ReturnType == nil {
			return true
		}

		if ft.ReturnType == nil || otherFunc.ReturnType == nil {
			return false
		}

		return ft.ReturnType.Equals(otherFunc.ReturnType)
	}
	return false
}

// ClassType, bir sınıf tipini temsil eder.
type ClassType struct {
	Name       string
	Fields     map[string]Type
	Methods    map[string]*FunctionType
	Extends    *ClassType
	Implements []*InterfaceType
}

// String, sınıf tipinin string temsilini döndürür.
func (ct *ClassType) String() string {
	return ct.Name
}

// Equals, iki sınıf tipinin eşit olup olmadığını kontrol eder.
func (ct *ClassType) Equals(other Type) bool {
	if otherClass, ok := other.(*ClassType); ok {
		return ct.Name == otherClass.Name
	}
	return false
}

// InterfaceType, bir arayüz tipini temsil eder.
type InterfaceType struct {
	Name    string
	Methods map[string]*FunctionType
}

// String, arayüz tipinin string temsilini döndürür.
func (it *InterfaceType) String() string {
	return it.Name
}

// Equals, iki arayüz tipinin eşit olup olmadığını kontrol eder.
func (it *InterfaceType) Equals(other Type) bool {
	if otherInterface, ok := other.(*InterfaceType); ok {
		return it.Name == otherInterface.Name
	}
	return false
}

// TemplateType, bir şablon tipini temsil eder.
type TemplateType struct {
	Name       string
	Parameters []string
	BaseType   Type
}

// String, şablon tipinin string temsilini döndürür.
func (tt *TemplateType) String() string {
	result := tt.Name + "<"

	for i, param := range tt.Parameters {
		if i > 0 {
			result += ", "
		}
		result += param
	}

	result += ">"

	return result
}

// Equals, iki şablon tipinin eşit olup olmadığını kontrol eder.
func (tt *TemplateType) Equals(other Type) bool {
	if otherTemplate, ok := other.(*TemplateType); ok {
		return tt.Name == otherTemplate.Name && len(tt.Parameters) == len(otherTemplate.Parameters)
	}
	return false
}

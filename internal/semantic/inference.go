package semantic

import (
	"github.com/inkbytefo/go-minus/internal/ast"
)

// TypeInference, tip çıkarımı işlemlerini gerçekleştirir.
type TypeInference struct {
	analyzer *Analyzer
}

// NewTypeInference, yeni bir TypeInference oluşturur.
func NewTypeInference(analyzer *Analyzer) *TypeInference {
	return &TypeInference{
		analyzer: analyzer,
	}
}

// InferType, bir ifadenin tipini çıkarır.
func (ti *TypeInference) InferType(expr ast.Expression) Type {
	switch e := expr.(type) {
	case *ast.Identifier:
		return ti.inferIdentifierType(e)
	case *ast.IntegerLiteral:
		return &BasicType{Name: "int", Kind: INTEGER_TYPE}
	case *ast.FloatLiteral:
		return &BasicType{Name: "float", Kind: FLOAT_TYPE}
	case *ast.StringLiteral:
		return &BasicType{Name: "string", Kind: STRING_TYPE}
	case *ast.CharLiteral:
		return &BasicType{Name: "char", Kind: CHAR_TYPE}
	case *ast.BooleanLiteral:
		return &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
	case *ast.NullLiteral:
		return &BasicType{Name: "null", Kind: NULL_TYPE}
	case *ast.PrefixExpression:
		return ti.inferPrefixExpressionType(e)
	case *ast.InfixExpression:
		return ti.inferInfixExpressionType(e)
	case *ast.IfExpression:
		return ti.inferIfExpressionType(e)
	case *ast.FunctionLiteral:
		return ti.inferFunctionLiteralType(e)
	case *ast.CallExpression:
		return ti.inferCallExpressionType(e)
	case *ast.ArrayLiteral:
		return ti.inferArrayLiteralType(e)
	case *ast.IndexExpression:
		return ti.inferIndexExpressionType(e)
	case *ast.HashLiteral:
		return ti.inferHashLiteralType(e)
	case *ast.MemberExpression:
		return ti.inferMemberExpressionType(e)
	case *ast.NewExpression:
		return ti.inferNewExpressionType(e)
	case *ast.TemplateExpression:
		return ti.inferTemplateExpressionType(e)
	default:
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

// inferIdentifierType, bir tanımlayıcının tipini çıkarır.
func (ti *TypeInference) inferIdentifierType(expr *ast.Identifier) Type {
	// Tanımlayıcıyı çözümle
	symbol := ti.analyzer.currentScope.Resolve(expr.Value)
	if symbol == nil {
		ti.analyzer.reportError(expr.Token, "Tanımlanmamış tanımlayıcı: %s", expr.Value)
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}

	// Sembol tipini döndür
	switch symbol.Type {
	case INTEGER_TYPE:
		return &BasicType{Name: "int", Kind: INTEGER_TYPE}
	case FLOAT_TYPE:
		return &BasicType{Name: "float", Kind: FLOAT_TYPE}
	case STRING_TYPE:
		return &BasicType{Name: "string", Kind: STRING_TYPE}
	case BOOLEAN_TYPE:
		return &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
	case CHAR_TYPE:
		return &BasicType{Name: "char", Kind: CHAR_TYPE}
	case NULL_TYPE:
		return &BasicType{Name: "null", Kind: NULL_TYPE}
	case FUNCTION_TYPE:
		// Fonksiyon tipini oluştur
		funcType := &FunctionType{
			ParameterTypes: make([]Type, 0),
			ReturnType:     &BasicType{Name: "void", Kind: VOID_TYPE},
		}

		// Fonksiyon imzası varsa, parametre ve dönüş tiplerini ekle
		if symbol.Signature != nil {
			for _, param := range symbol.Signature.Parameters {
				paramType := symbolTypeToType(param.Type)
				funcType.ParameterTypes = append(funcType.ParameterTypes, paramType)
			}

			funcType.ReturnType = symbolTypeToType(symbol.Signature.ReturnType)
		}

		return funcType
	case CLASS_TYPE:
		// Sınıf tipini oluştur
		return &ClassType{
			Name:       symbol.Name,
			Fields:     make(map[string]Type),
			Methods:    make(map[string]*FunctionType),
			Implements: []*InterfaceType{},
		}
	default:
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

// inferPrefixExpressionType, bir önek ifadesinin tipini çıkarır.
func (ti *TypeInference) inferPrefixExpressionType(expr *ast.PrefixExpression) Type {
	// Sağ tarafı analiz et
	rightType := ti.InferType(expr.Right)

	// Operatöre göre tip kontrolü yap
	switch expr.Operator {
	case "!":
		// ! operatörü boolean tipinde olmalıdır
		if basicType, ok := rightType.(*BasicType); !ok || basicType.Kind != BOOLEAN_TYPE {
			ti.analyzer.reportError(expr.Token, "! operatörü boolean tipinde olmalıdır")
		}
		return &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
	case "-":
		// - operatörü sayısal tipte olmalıdır
		if basicType, ok := rightType.(*BasicType); !ok || (basicType.Kind != INTEGER_TYPE && basicType.Kind != FLOAT_TYPE) {
			ti.analyzer.reportError(expr.Token, "- operatörü sayısal tipte olmalıdır")
		}
		return rightType
	default:
		ti.analyzer.reportError(expr.Token, "Bilinmeyen önek operatörü: %s", expr.Operator)
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

// inferInfixExpressionType, bir araek ifadesinin tipini çıkarır.
func (ti *TypeInference) inferInfixExpressionType(expr *ast.InfixExpression) Type {
	// Sol ve sağ tarafı analiz et
	leftType := ti.InferType(expr.Left)
	rightType := ti.InferType(expr.Right)

	// Operatöre göre tip kontrolü yap
	switch expr.Operator {
	case "-", "*", "/", "%":
		// Aritmetik operatörler sayısal tipte olmalıdır
		if basicLeftType, ok := leftType.(*BasicType); !ok || (basicLeftType.Kind != INTEGER_TYPE && basicLeftType.Kind != FLOAT_TYPE) {
			ti.analyzer.reportError(expr.Token, "Aritmetik operatörün sol tarafı sayısal tipte olmalıdır")
		}
		if basicRightType, ok := rightType.(*BasicType); !ok || (basicRightType.Kind != INTEGER_TYPE && basicRightType.Kind != FLOAT_TYPE) {
			ti.analyzer.reportError(expr.Token, "Aritmetik operatörün sağ tarafı sayısal tipte olmalıdır")
		}

		// Eğer herhangi bir taraf FLOAT_TYPE ise, sonuç FLOAT_TYPE olur
		var basicLeftType, basicRightType *BasicType
		var ok1, ok2 bool
		basicLeftType, ok1 = leftType.(*BasicType)
		basicRightType, ok2 = rightType.(*BasicType)
		if (ok1 && basicLeftType.Kind == FLOAT_TYPE) || (ok2 && basicRightType.Kind == FLOAT_TYPE) {
			return &BasicType{Name: "float", Kind: FLOAT_TYPE}
		}
		return &BasicType{Name: "int", Kind: INTEGER_TYPE}
	case "+":
		// + operatörü string birleştirme için de kullanılabilir
		if basicLeftType, ok1 := leftType.(*BasicType); ok1 && basicLeftType.Kind == STRING_TYPE {
			if basicRightType, ok2 := rightType.(*BasicType); ok2 && basicRightType.Kind == STRING_TYPE {
				return &BasicType{Name: "string", Kind: STRING_TYPE}
			}
		}

		// + operatörü sayısal tipler için de kullanılabilir
		var basicLeftType2, basicRightType2 *BasicType
		var ok3, ok4 bool

		basicLeftType2, ok3 = leftType.(*BasicType)
		if !ok3 || (basicLeftType2.Kind != INTEGER_TYPE && basicLeftType2.Kind != FLOAT_TYPE) {
			ti.analyzer.reportError(expr.Token, "Aritmetik operatörün sol tarafı sayısal tipte olmalıdır")
		}

		basicRightType2, ok4 = rightType.(*BasicType)
		if !ok4 || (basicRightType2.Kind != INTEGER_TYPE && basicRightType2.Kind != FLOAT_TYPE) {
			ti.analyzer.reportError(expr.Token, "Aritmetik operatörün sağ tarafı sayısal tipte olmalıdır")
		}

		// Eğer herhangi bir taraf FLOAT_TYPE ise, sonuç FLOAT_TYPE olur
		if (ok3 && basicLeftType2.Kind == FLOAT_TYPE) || (ok4 && basicRightType2.Kind == FLOAT_TYPE) {
			return &BasicType{Name: "float", Kind: FLOAT_TYPE}
		}
		return &BasicType{Name: "int", Kind: INTEGER_TYPE}
	case "<", ">", "<=", ">=", "==", "!=":
		// Karşılaştırma operatörleri aynı tipte olmalıdır
		if !leftType.Equals(rightType) {
			ti.analyzer.reportError(expr.Token, "Karşılaştırma operatörünün sol ve sağ tarafı aynı tipte olmalıdır")
		}
		return &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
	case "&&", "||":
		// Mantıksal operatörler boolean tipinde olmalıdır
		if basicLeftType, ok := leftType.(*BasicType); !ok || basicLeftType.Kind != BOOLEAN_TYPE {
			ti.analyzer.reportError(expr.Token, "Mantıksal operatörün sol tarafı boolean tipinde olmalıdır")
		}
		if basicRightType, ok := rightType.(*BasicType); !ok || basicRightType.Kind != BOOLEAN_TYPE {
			ti.analyzer.reportError(expr.Token, "Mantıksal operatörün sağ tarafı boolean tipinde olmalıdır")
		}
		return &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
	case "=", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=", "<<=", ">>=":
		// Atama operatörleri için sol taraf bir değişken olmalıdır
		if _, ok := expr.Left.(*ast.Identifier); !ok {
			ti.analyzer.reportError(expr.Token, "Atama operatörünün sol tarafı bir değişken olmalıdır")
		}
		// Sağ taraf sol tarafla aynı tipte olmalıdır
		if !leftType.Equals(rightType) {
			ti.analyzer.reportError(expr.Token, "Atama operatörünün sağ tarafı sol tarafla aynı tipte olmalıdır")
		}
		return leftType
	case ":=":
		// Kısa değişken tanımlama operatörü
		// Sol taraf bir tanımlayıcı olmalıdır
		if ident, ok := expr.Left.(*ast.Identifier); ok {
			// Tanımlayıcıyı tanımla
			ti.analyzer.currentScope.Define(ident.Value, symbolTypeFromType(rightType), ident.Token)
		} else {
			ti.analyzer.reportError(expr.Token, "Kısa değişken tanımlama operatörünün sol tarafı bir tanımlayıcı olmalıdır")
		}
		return rightType
	default:
		ti.analyzer.reportError(expr.Token, "Bilinmeyen araek operatörü: %s", expr.Operator)
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

// symbolTypeToType, bir SymbolType'ı Type'a dönüştürür.
func symbolTypeToType(symbolType SymbolType) Type {
	switch symbolType {
	case INTEGER_TYPE:
		return &BasicType{Name: "int", Kind: INTEGER_TYPE}
	case FLOAT_TYPE:
		return &BasicType{Name: "float", Kind: FLOAT_TYPE}
	case STRING_TYPE:
		return &BasicType{Name: "string", Kind: STRING_TYPE}
	case BOOLEAN_TYPE:
		return &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
	case CHAR_TYPE:
		return &BasicType{Name: "char", Kind: CHAR_TYPE}
	case NULL_TYPE:
		return &BasicType{Name: "null", Kind: NULL_TYPE}
	case VOID_TYPE:
		return &BasicType{Name: "void", Kind: VOID_TYPE}
	default:
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

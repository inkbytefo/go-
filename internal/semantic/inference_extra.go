package semantic

import (
	"github.com/inkbytefo/go-minus/internal/ast"
)

// inferIfExpressionType, bir if ifadesinin tipini çıkarır.
func (ti *TypeInference) inferIfExpressionType(expr *ast.IfExpression) Type {
	// Koşul boolean tipinde olmalıdır
	conditionType := ti.InferType(expr.Condition)
	if basicType, ok := conditionType.(*BasicType); !ok || basicType.Kind != BOOLEAN_TYPE {
		ti.analyzer.reportError(expr.Token, "If ifadesinin koşulu boolean tipinde olmalıdır")
	}

	// Consequence ve alternative bloklarının tiplerini çıkar
	consequenceType := ti.inferBlockStatementType(expr.Consequence)

	// Alternative blok varsa, tipini çıkar
	if expr.Alternative != nil {
		alternativeType := ti.inferBlockStatementType(expr.Alternative)

		// Consequence ve alternative bloklarının tipleri aynı olmalıdır
		if !consequenceType.Equals(alternativeType) {
			ti.analyzer.reportError(expr.Token, "If ifadesinin consequence ve alternative bloklarının tipleri aynı olmalıdır")
		}
	}

	return consequenceType
}

// inferBlockStatementType, bir blok ifadesinin tipini çıkarır.
func (ti *TypeInference) inferBlockStatementType(block *ast.BlockStatement) Type {
	// Blok boşsa, void tipini döndür
	if len(block.Statements) == 0 {
		return &BasicType{Name: "void", Kind: VOID_TYPE}
	}

	// Bloktaki son ifadenin tipini döndür
	lastStmt := block.Statements[len(block.Statements)-1]

	// Son ifade bir return ifadesi ise, döndürülen değerin tipini döndür
	if returnStmt, ok := lastStmt.(*ast.ReturnStatement); ok {
		if returnStmt.ReturnValue != nil {
			return ti.InferType(returnStmt.ReturnValue)
		}
		return &BasicType{Name: "void", Kind: VOID_TYPE}
	}

	// Son ifade bir expression ifadesi ise, ifadenin tipini döndür
	if exprStmt, ok := lastStmt.(*ast.ExpressionStatement); ok {
		return ti.InferType(exprStmt.Expression)
	}

	// Diğer durumlarda void tipini döndür
	return &BasicType{Name: "void", Kind: VOID_TYPE}
}

// inferFunctionLiteralType, bir fonksiyon değişmez değerinin tipini çıkarır.
func (ti *TypeInference) inferFunctionLiteralType(expr *ast.FunctionLiteral) Type {
	// Fonksiyon tipini oluştur
	funcType := &FunctionType{
		ParameterTypes: make([]Type, 0),
		ReturnType:     &BasicType{Name: "void", Kind: VOID_TYPE},
	}

	// Parametrelerin tiplerini ekle
	for _, param := range expr.Parameters {
		// Parametre tipini belirle (varsayılan olarak int)
		paramType := &BasicType{Name: "int", Kind: INTEGER_TYPE}

		// Parametreyi sembol tablosuna ekle
		ti.analyzer.currentScope.Define(param.Value, INTEGER_TYPE, param.Token)

		// Parametre tipini fonksiyon tipine ekle
		funcType.ParameterTypes = append(funcType.ParameterTypes, paramType)
	}

	// Dönüş tipini belirle
	if expr.ReturnType != nil {
		// Dönüş tipi belirtilmişse, bu tipi kullan
		if typeIdent, ok := expr.ReturnType.(*ast.Identifier); ok {
			switch typeIdent.Value {
			case "int":
				funcType.ReturnType = &BasicType{Name: "int", Kind: INTEGER_TYPE}
			case "float":
				funcType.ReturnType = &BasicType{Name: "float", Kind: FLOAT_TYPE}
			case "string":
				funcType.ReturnType = &BasicType{Name: "string", Kind: STRING_TYPE}
			case "bool":
				funcType.ReturnType = &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
			case "char":
				funcType.ReturnType = &BasicType{Name: "char", Kind: CHAR_TYPE}
			case "void":
				funcType.ReturnType = &BasicType{Name: "void", Kind: VOID_TYPE}
			default:
				// Bilinmeyen tip
				ti.analyzer.reportError(typeIdent.Token, "Bilinmeyen dönüş tipi: %s", typeIdent.Value)
			}
		}
	} else if expr.Body != nil {
		// Dönüş tipi belirtilmemişse, gövdeden çıkar
		funcType.ReturnType = ti.inferBlockStatementType(expr.Body)
	}

	return funcType
}

// inferCallExpressionType, bir fonksiyon çağrısının tipini çıkarır.
func (ti *TypeInference) inferCallExpressionType(expr *ast.CallExpression) Type {
	// Fonksiyonun tipini çıkar
	funcType := ti.InferType(expr.Function)

	// Fonksiyon tipi kontrolü
	if ft, ok := funcType.(*FunctionType); ok {
		// Variadic function kontrolü için member expression'ı kontrol et
		isVariadic := false
		if memberExpr, ok := expr.Function.(*ast.MemberExpression); ok {
			if objectIdent, ok := memberExpr.Object.(*ast.Identifier); ok {
				if packageSymbol := ti.analyzer.currentScope.Resolve(objectIdent.Value); packageSymbol != nil && packageSymbol.Type == PACKAGE_TYPE {
					if memberIdent, ok := memberExpr.Member.(*ast.Identifier); ok {
						if packageSymbol.Class != nil && packageSymbol.Class.Methods != nil {
							if methodSymbol, exists := packageSymbol.Class.Methods[memberIdent.Value]; exists {
								isVariadic = methodSymbol.Signature.IsVariadic
							}
						}
					}
				}
			}
		}

		// Argüman sayısı kontrolü (variadic functions için farklı)
		if isVariadic {
			// Variadic function: minimum parametre sayısını kontrol et
			if len(expr.Arguments) < len(ft.ParameterTypes) {
				ti.analyzer.reportError(expr.Token, "Fonksiyon çağrısı için yetersiz argüman sayısı: en az %d bekleniyor, %d alındı", len(ft.ParameterTypes), len(expr.Arguments))
			}
		} else {
			// Normal function: tam parametre sayısını kontrol et
			if len(expr.Arguments) != len(ft.ParameterTypes) {
				ti.analyzer.reportError(expr.Token, "Fonksiyon çağrısı için yanlış argüman sayısı: %d bekleniyor, %d alındı", len(ft.ParameterTypes), len(expr.Arguments))
			}
		}

		// Argüman tiplerini kontrol et
		for i, arg := range expr.Arguments {
			if i < len(ft.ParameterTypes) {
				argType := ti.InferType(arg)
				if !argType.Equals(ft.ParameterTypes[i]) {
					// UNKNOWN_TYPE parametreler için tip kontrolü yapmayalım (variadic için)
					if basicType, ok := ft.ParameterTypes[i].(*BasicType); !ok || basicType.Kind != UNKNOWN_TYPE {
						ti.analyzer.reportError(expr.Token, "Fonksiyon çağrısı için yanlış argüman tipi: %s bekleniyor, %s alındı", ft.ParameterTypes[i].String(), argType.String())
					}
				}
			}
		}

		// Fonksiyonun dönüş tipini döndür
		return ft.ReturnType
	} else {
		ti.analyzer.reportError(expr.Token, "Çağrılabilir olmayan ifade")
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

// symbolTypeFromType, bir Type'ı SymbolType'a dönüştürür.
func symbolTypeFromType(t Type) SymbolType {
	if basicType, ok := t.(*BasicType); ok {
		return basicType.Kind
	}
	return UNKNOWN_TYPE
}

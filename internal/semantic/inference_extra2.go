package semantic

import (
	"github.com/inkbytefo/go-minus/internal/ast"
)

// inferArrayLiteralType, bir dizi değişmez değerinin tipini çıkarır.
func (ti *TypeInference) inferArrayLiteralType(expr *ast.ArrayLiteral) Type {
	// Dizi boşsa, varsayılan olarak int dizisi döndür
	if len(expr.Elements) == 0 {
		return &ArrayType{
			ElementType: &BasicType{Name: "int", Kind: INTEGER_TYPE},
		}
	}

	// İlk elemanın tipini al
	elementType := ti.InferType(expr.Elements[0])

	// Diğer elemanların tiplerini kontrol et
	for _, element := range expr.Elements[1:] {
		elemType := ti.InferType(element)
		if !elemType.Equals(elementType) {
			ti.analyzer.reportError(expr.Token, "Dizi elemanları aynı tipte olmalıdır")
			break
		}
	}

	// Dizi tipini döndür
	return &ArrayType{
		ElementType: elementType,
	}
}

// inferIndexExpressionType, bir indeks ifadesinin tipini çıkarır.
func (ti *TypeInference) inferIndexExpressionType(expr *ast.IndexExpression) Type {
	// Sol tarafın tipini çıkar
	leftType := ti.InferType(expr.Left)

	// İndeks ifadesinin tipini çıkar
	indexType := ti.InferType(expr.Index)

	// İndeks ifadesi int tipinde olmalıdır
	if basicType, ok := indexType.(*BasicType); !ok || basicType.Kind != INTEGER_TYPE {
		ti.analyzer.reportError(expr.Token, "İndeks ifadesi int tipinde olmalıdır")
	}

	// Sol taraf bir dizi ise, eleman tipini döndür
	if arrayType, ok := leftType.(*ArrayType); ok {
		return arrayType.ElementType
	}

	// Sol taraf bir string ise, char tipini döndür
	if basicType, ok := leftType.(*BasicType); ok && basicType.Kind == STRING_TYPE {
		return &BasicType{Name: "char", Kind: CHAR_TYPE}
	}

	// Diğer durumlarda hata ver
	ti.analyzer.reportError(expr.Token, "İndeks operatörü dizi veya string tipinde olmalıdır")
	return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
}

// inferHashLiteralType, bir hash değişmez değerinin tipini çıkarır.
func (ti *TypeInference) inferHashLiteralType(expr *ast.HashLiteral) Type {
	// Hash boşsa, varsayılan olarak string->int hash döndür
	if len(expr.Pairs) == 0 {
		return &HashType{
			KeyType:   &BasicType{Name: "string", Kind: STRING_TYPE},
			ValueType: &BasicType{Name: "int", Kind: INTEGER_TYPE},
		}
	}

	// İlk çiftin tiplerini al
	var keyType, valueType Type
	for k, v := range expr.Pairs {
		keyType = ti.InferType(k)
		valueType = ti.InferType(v)
		break
	}

	// Diğer çiftlerin tiplerini kontrol et
	for k, v := range expr.Pairs {
		kType := ti.InferType(k)
		vType := ti.InferType(v)

		if !kType.Equals(keyType) {
			ti.analyzer.reportError(expr.Token, "Hash anahtarları aynı tipte olmalıdır")
		}

		if !vType.Equals(valueType) {
			ti.analyzer.reportError(expr.Token, "Hash değerleri aynı tipte olmalıdır")
		}
	}

	// Hash tipini döndür
	return &HashType{
		KeyType:   keyType,
		ValueType: valueType,
	}
}

// inferMemberExpressionType, bir üye erişim ifadesinin tipini çıkarır.
func (ti *TypeInference) inferMemberExpressionType(expr *ast.MemberExpression) Type {
	// Üye adını al
	var memberName string
	if memberIdent, ok := expr.Member.(*ast.Identifier); ok {
		memberName = memberIdent.Value
	} else {
		ti.analyzer.reportError(expr.Token, "Üye adı bir tanımlayıcı olmalıdır")
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}

	// Package erişimi kontrolü
	if objectIdent, ok := expr.Object.(*ast.Identifier); ok {
		if packageSymbol := ti.analyzer.currentScope.Resolve(objectIdent.Value); packageSymbol != nil && packageSymbol.Type == PACKAGE_TYPE {
			// Package.function erişimi
			if packageSymbol.Class != nil && packageSymbol.Class.Methods != nil {
				if methodSymbol, exists := packageSymbol.Class.Methods[memberName]; exists {
					// Function type'ını döndür
					return &FunctionType{
						ParameterTypes: make([]Type, len(methodSymbol.Signature.Parameters)),
						ReturnType:     symbolTypeToType(methodSymbol.Signature.ReturnType),
					}
				}
			}
			ti.analyzer.reportError(expr.Token, "Package %s'de %s fonksiyonu bulunamadı", objectIdent.Value, memberName)
			return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
		}
	}

	// Nesnenin tipini çıkar
	objectType := ti.InferType(expr.Object)

	// Nesne bir sınıf ise, üye tipini döndür
	if classType, ok := objectType.(*ClassType); ok {
		// Üye bir alan ise
		if fieldType, ok := classType.Fields[memberName]; ok {
			return fieldType
		}

		// Üye bir metot ise
		if methodType, ok := classType.Methods[memberName]; ok {
			return methodType
		}

		// Üye bulunamadı
		ti.analyzer.reportError(expr.Token, "Sınıfta '%s' adında bir üye bulunamadı", memberName)
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}

	// Nesne bir arayüz ise, üye tipini döndür
	if interfaceType, ok := objectType.(*InterfaceType); ok {
		// Üye bir metot ise
		if methodType, ok := interfaceType.Methods[memberName]; ok {
			return methodType
		}

		// Üye bulunamadı
		ti.analyzer.reportError(expr.Token, "Arayüzde '%s' adında bir metot bulunamadı", memberName)
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}

	// Diğer durumlarda hata ver
	ti.analyzer.reportError(expr.Token, "Üye erişimi için nesne bir sınıf, arayüz veya package olmalıdır")
	return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
}

// inferNewExpressionType, bir new ifadesinin tipini çıkarır.
func (ti *TypeInference) inferNewExpressionType(expr *ast.NewExpression) Type {
	// Sınıf adını al
	var className string
	if classIdent, ok := expr.Class.(*ast.Identifier); ok {
		className = classIdent.Value
	} else {
		ti.analyzer.reportError(expr.Token, "Sınıf adı bir tanımlayıcı olmalıdır")
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}

	// Sınıfı sembol tablosundan bul
	symbol := ti.analyzer.currentScope.Resolve(className)
	if symbol == nil || symbol.Type != CLASS_TYPE {
		ti.analyzer.reportError(expr.Token, "Tanımlanmamış sınıf: %s", className)
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}

	// Sınıf tipini döndür
	return &ClassType{
		Name:       className,
		Fields:     make(map[string]Type),
		Methods:    make(map[string]*FunctionType),
		Implements: []*InterfaceType{},
	}
}

// inferTemplateExpressionType, bir şablon ifadesinin tipini çıkarır.
func (ti *TypeInference) inferTemplateExpressionType(expr *ast.TemplateExpression) Type {
	// Şablon gövdesinin tipini çıkar
	bodyType := ti.InferType(expr.Body)

	// Şablon tipini döndür
	templateParams := make([]string, len(expr.Parameters))
	for i, param := range expr.Parameters {
		templateParams[i] = param.Value
	}

	return &TemplateType{
		Name:       "template",
		Parameters: templateParams,
		BaseType:   bodyType,
	}
}

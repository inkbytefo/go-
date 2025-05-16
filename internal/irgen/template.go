package irgen

import (
	"fmt"
	"strings"

	"github.com/inkbytefo/go-minus/internal/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// TemplateInfo, bir şablon hakkında bilgi tutar.
type TemplateInfo struct {
	Name           string
	TypeParameters []string
	Node           ast.Node
	Instances      map[string]interface{} // Örneklenmiş şablonlar (sınıf veya fonksiyon)
}

// generateTemplateStatement, bir şablon tanımlaması için IR üretir.
func (g *IRGenerator) generateTemplateStatement(stmt *ast.TemplateStatement) {
	// Şablon adını al
	var templateName string

	// Şablonun türüne göre adını belirle
	switch node := stmt.Node.(type) {
	case *ast.ClassStatement:
		templateName = node.Name.Value
	case *ast.FunctionStatement:
		templateName = node.Name.Value
	default:
		g.ReportError("Desteklenmeyen şablon türü: %T", stmt.Node)
		return
	}

	// Tip parametrelerini al
	typeParams := make([]string, len(stmt.TypeParameters))
	for i, param := range stmt.TypeParameters {
		typeParams[i] = param.Value
	}

	// Şablon bilgisini oluştur
	templateInfo := &TemplateInfo{
		Name:           templateName,
		TypeParameters: typeParams,
		Node:           stmt.Node,
		Instances:      make(map[string]interface{}),
	}

	// Şablon bilgisini kaydet
	g.templateTable[templateName] = templateInfo
}

// generateTemplateExpression, bir şablon ifadesi için IR üretir.
func (g *IRGenerator) generateTemplateExpression(expr *ast.TemplateExpression) value.Value {
	// Şablon adını al
	templateName := expr.Name.Value

	// Şablon bilgisini bul
	templateInfo, exists := g.templateTable[templateName]
	if !exists {
		g.ReportError("Şablon bulunamadı: %s", templateName)
		return nil
	}

	// Tip argümanlarını değerlendir
	typeArgs := make([]types.Type, len(expr.TypeArguments))
	for i, arg := range expr.TypeArguments {
		if typeIdent, ok := arg.(*ast.Identifier); ok {
			if t, exists := g.typeTable[typeIdent.Value]; exists {
				typeArgs[i] = t
			} else {
				g.ReportError("Bilinmeyen tip: %s", typeIdent.Value)
				return nil
			}
		} else {
			g.ReportError("Desteklenmeyen tip argümanı: %T", arg)
			return nil
		}
	}

	// Tip argümanlarından bir anahtar oluştur
	key := g.getTemplateInstanceKey(templateName, typeArgs)

	// Şablon zaten örneklenmişse, onu kullan
	if instance, exists := templateInfo.Instances[key]; exists {
		return g.useTemplateInstance(templateInfo, instance, expr)
	}

	// Şablonu örnekle
	instance := g.instantiateTemplate(templateInfo, typeArgs)
	if instance == nil {
		return nil
	}

	// Örneklenmiş şablonu kaydet
	templateInfo.Instances[key] = instance

	// Örneklenmiş şablonu kullan
	return g.useTemplateInstance(templateInfo, instance, expr)
}

// getTemplateInstanceKey, şablon örneği için bir anahtar oluşturur.
func (g *IRGenerator) getTemplateInstanceKey(templateName string, typeArgs []types.Type) string {
	parts := make([]string, len(typeArgs)+1)
	parts[0] = templateName

	for i, arg := range typeArgs {
		parts[i+1] = arg.String()
	}

	return strings.Join(parts, "_")
}

// instantiateTemplate, bir şablonu belirtilen tip argümanlarıyla örnekler.
func (g *IRGenerator) instantiateTemplate(templateInfo *TemplateInfo, typeArgs []types.Type) interface{} {
	// Tip parametrelerinin sayısını kontrol et
	if len(templateInfo.TypeParameters) != len(typeArgs) {
		g.ReportError("Şablon için yanlış sayıda tip argümanı: %s", templateInfo.Name)
		return nil
	}

	// Tip eşlemesi oluştur
	typeMap := make(map[string]types.Type)
	for i, param := range templateInfo.TypeParameters {
		typeMap[param] = typeArgs[i]
	}

	// Şablonun türüne göre örnekleme yap
	switch node := templateInfo.Node.(type) {
	case *ast.ClassStatement:
		return g.instantiateTemplateClass(templateInfo, node, typeMap)
	case *ast.FunctionStatement:
		return g.instantiateTemplateFunction(templateInfo, node, typeMap)
	default:
		g.ReportError("Desteklenmeyen şablon türü: %T", templateInfo.Node)
		return nil
	}
}

// instantiateTemplateClass, bir sınıf şablonunu örnekler.
func (g *IRGenerator) instantiateTemplateClass(templateInfo *TemplateInfo, classStmt *ast.ClassStatement, typeMap map[string]types.Type) *ClassInfo {
	// Örneklenmiş sınıf adını oluştur
	instanceName := g.getTemplateInstanceName(templateInfo.Name, typeMap)

	// Sınıf bilgisi oluştur
	classInfo := &ClassInfo{
		Name:       instanceName,
		Methods:    make(map[string]*MethodInfo),
		Fields:     make(map[string]FieldInfo),
		Interfaces: make([]*ClassInfo, 0),
	}

	// Sınıf alanlarını ve metotlarını topla
	fieldTypes := make([]types.Type, 0)
	fieldNames := make([]string, 0)

	// VTable işaretçisi ekle (ilk alan olarak)
	vtablePtrType := types.NewPointer(types.Void)
	fieldTypes = append(fieldTypes, vtablePtrType)
	fieldNames = append(fieldNames, "_vtable")

	// Sınıf gövdesindeki ifadeleri işle
	if classStmt.Body != nil {
		// Önce alanları işle
		for _, s := range classStmt.Body.Statements {
			// Değişken üyeleri işle
			if varStmt, ok := s.(*ast.VarStatement); ok {
				fieldName := varStmt.Name.Value

				// Alan tipini belirle
				var fieldType types.Type = types.I32 // Varsayılan olarak int32
				if varStmt.Type != nil {
					if typeIdent, ok := varStmt.Type.(*ast.Identifier); ok {
						// Tip parametresi mi kontrol et
						if t, exists := typeMap[typeIdent.Value]; exists {
							fieldType = t
						} else if t, exists := g.typeTable[typeIdent.Value]; exists {
							fieldType = t
						} else {
							g.ReportError("Bilinmeyen tip: %s", typeIdent.Value)
							fieldType = types.I32 // Varsayılan olarak int32
						}
					}
				}

				// Alanı ekle
				isPrivate := false // Varsayılan olarak public
				// TODO: Erişim belirleyicilerini kontrol et

				classInfo.Fields[fieldName] = FieldInfo{
					Name:      fieldName,
					Type:      fieldType,
					Index:     len(fieldTypes),
					IsPrivate: isPrivate,
				}

				fieldTypes = append(fieldTypes, fieldType)
				fieldNames = append(fieldNames, fieldName)
			}
		}

		// Sonra metotları işle
		for _, s := range classStmt.Body.Statements {
			// Metotları işle
			if funcStmt, ok := s.(*ast.FunctionStatement); ok {
				methodName := funcStmt.Name.Value

				// Metot imzasını oluştur
				paramTypes := make([]types.Type, 0)
				paramTypes = append(paramTypes, types.NewPointer(types.Void)) // this işaretçisi

				for _, param := range funcStmt.Parameters {
					paramType := types.I32 // Varsayılan olarak int32
					// TODO: Parametre tiplerini belirle
					paramTypes = append(paramTypes, paramType)
				}

				returnType := types.I32 // Varsayılan olarak int32
				// TODO: Dönüş tipini belirle

				// Fonksiyon tipini oluştur
				funcType := types.NewFunc(returnType, paramTypes...)

				// Metot adını oluştur (sınıf adı + metot adı)
				fullMethodName := fmt.Sprintf("%s_%s", instanceName, methodName)

				// Metodu oluştur
				method := g.module.NewFunc(fullMethodName, returnType, ir.NewParam("this", types.NewPointer(types.Void)))

				// Parametreleri ekle
				for i, param := range funcStmt.Parameters {
					paramName := param.Value
					paramType := paramTypes[i+1] // +1 çünkü ilk parametre this
					method.Params = append(method.Params, ir.NewParam(paramName, paramType))
				}

				// Metot bilgisini ekle
				isVirtual := false // Varsayılan olarak virtual değil
				// TODO: Virtual metotları belirle

				methodInfo := &MethodInfo{
					Name:        methodName,
					Function:    method,
					IsVirtual:   isVirtual,
					VTableIndex: -1, // Henüz belirlenmedi
					Signature:   funcType,
				}

				classInfo.Methods[methodName] = methodInfo

				// Metot gövdesini işle
				// TODO: Metot gövdesini işle
			}
		}
	}

	// Struct tipini oluştur
	structType := types.NewStruct(fieldTypes...)
	classInfo.StructType = structType

	// VTable tipini oluştur
	vtableType := types.NewStruct() // Boş VTable
	classInfo.VTableType = vtableType

	// VTable örneğini oluştur
	vtableInit := constant.NewStruct(vtableType)
	vtableGlobal := g.module.NewGlobalDef(fmt.Sprintf("%s_vtable", instanceName), vtableInit)
	classInfo.VTableInstance = vtableGlobal

	// Sınıf bilgisini kaydet
	g.classTable[instanceName] = classInfo

	// Tipi kaydet
	g.typeTable[instanceName] = structType

	return classInfo
}

// instantiateTemplateFunction, bir fonksiyon şablonunu örnekler.
func (g *IRGenerator) instantiateTemplateFunction(templateInfo *TemplateInfo, funcStmt *ast.FunctionStatement, typeMap map[string]types.Type) *ir.Func {
	// Örneklenmiş fonksiyon adını oluştur
	instanceName := g.getTemplateInstanceName(templateInfo.Name, typeMap)

	// Parametre tiplerini belirle
	paramTypes := make([]types.Type, len(funcStmt.Parameters))
	for i, param := range funcStmt.Parameters {
		// Parametre tipini belirle
		if param.Type != nil {
			if typeIdent, ok := param.Type.(*ast.Identifier); ok {
				// Tip parametresi mi kontrol et
				if t, exists := typeMap[typeIdent.Value]; exists {
					paramTypes[i] = t
				} else if t, exists := g.typeTable[typeIdent.Value]; exists {
					paramTypes[i] = t
				} else {
					g.ReportError("Bilinmeyen tip: %s", typeIdent.Value)
					paramTypes[i] = types.I32 // Varsayılan olarak int32
				}
			}
		} else {
			paramTypes[i] = types.I32 // Varsayılan olarak int32
		}
	}

	// Dönüş tipini belirle
	var returnType types.Type = types.I32 // Varsayılan olarak int32
	if funcStmt.ReturnType != nil {
		if typeIdent, ok := funcStmt.ReturnType.(*ast.Identifier); ok {
			// Tip parametresi mi kontrol et
			if t, exists := typeMap[typeIdent.Value]; exists {
				returnType = t
			} else if t, exists := g.typeTable[typeIdent.Value]; exists {
				returnType = t
			} else {
				g.ReportError("Bilinmeyen tip: %s", typeIdent.Value)
			}
		}
	}

	// Fonksiyonu oluştur
	fn := g.module.NewFunc(instanceName, returnType)

	// Parametreleri ekle
	for i, param := range funcStmt.Parameters {
		paramName := param.Value
		fn.Params = append(fn.Params, ir.NewParam(paramName, paramTypes[i]))
	}

	// Fonksiyon gövdesini işle
	// TODO: Fonksiyon gövdesini işle

	// Fonksiyonu kaydet
	g.symbolTable[instanceName] = fn

	return fn
}

// getTemplateInstanceName, şablon örneği için bir ad oluşturur.
func (g *IRGenerator) getTemplateInstanceName(templateName string, typeMap map[string]types.Type) string {
	parts := make([]string, len(typeMap)+1)
	parts[0] = templateName

	i := 1
	for _, t := range typeMap {
		parts[i] = t.String()
		i++
	}

	return strings.Join(parts, "_")
}

// useTemplateInstance, örneklenmiş bir şablonu kullanır.
func (g *IRGenerator) useTemplateInstance(templateInfo *TemplateInfo, instance interface{}, expr *ast.TemplateExpression) value.Value {
	switch templateInfo.Node.(type) {
	case *ast.ClassStatement:
		// Sınıf şablonu
		if classInfo, ok := instance.(*ClassInfo); ok {
			// New ifadesi mi?
			if expr.IsNew {
				// Bellek ayır
				if g.currentBB == nil {
					g.ReportError("Geçerli bir blok yok, new ifadesi değerlendirilemiyor")
					return nil
				}

				mallocFunc := g.getMallocFunction()
				size := g.currentBB.NewPtrToInt(constant.NewGetElementPtr(classInfo.StructType, constant.NewNull(types.NewPointer(classInfo.StructType)), constant.NewInt(types.I32, 1)), types.I64)
				allocPtr := g.currentBB.NewCall(mallocFunc, size)

				// Tipi dönüştür
				objPtr := g.currentBB.NewBitCast(allocPtr, types.NewPointer(classInfo.StructType))

				// VTable'ı ayarla
				vtablePtr := g.currentBB.NewGetElementPtr(classInfo.StructType, objPtr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
				vtableAddr := g.currentBB.NewBitCast(classInfo.VTableInstance, types.NewPointer(types.Void))
				g.currentBB.NewStore(vtableAddr, vtablePtr)

				// Yapıcı metodu çağır (varsa)
				if constructorInfo, exists := classInfo.Methods["constructor"]; exists {
					thisPtr := g.currentBB.NewBitCast(objPtr, types.NewPointer(types.Void))

					// Argümanları değerlendir
					args := make([]value.Value, 0, len(expr.Arguments)+1)
					args = append(args, thisPtr) // this işaretçisi

					for _, arg := range expr.Arguments {
						argVal := g.generateExpression(arg)
						if argVal != nil {
							args = append(args, argVal)
						}
					}

					// Yapıcı metodu çağır
					g.currentBB.NewCall(constructorInfo.Function, args...)
				}

				return objPtr
			} else {
				// Sınıf tipini döndür
				return nil // LLVM IR'da tip değerleri yok
			}
		}
	case *ast.FunctionStatement:
		// Fonksiyon şablonu
		if fn, ok := instance.(*ir.Func); ok {
			// Fonksiyon çağrısı mı?
			if len(expr.Arguments) > 0 {
				if g.currentBB == nil {
					g.ReportError("Geçerli bir blok yok, fonksiyon çağrısı yapılamıyor")
					return nil
				}

				// Argümanları değerlendir
				args := make([]value.Value, 0, len(expr.Arguments))
				for _, arg := range expr.Arguments {
					argVal := g.generateExpression(arg)
					if argVal != nil {
						args = append(args, argVal)
					}
				}

				// Fonksiyon çağrısı yap
				return g.currentBB.NewCall(fn, args...)
			} else {
				// Fonksiyon işaretçisini döndür
				return fn
			}
		}
	}

	g.ReportError("Geçersiz şablon örneği türü: %T", instance)
	return nil
}

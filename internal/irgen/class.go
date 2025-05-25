package irgen

import (
	"fmt"

	"github.com/inkbytefo/go-minus/internal/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// ClassInfo, bir sınıf hakkında bilgi tutar.
type ClassInfo struct {
	Name           string
	StructType     *types.StructType
	VTableType     *types.StructType
	VTableInstance *ir.Global
	Methods        map[string]*MethodInfo
	Fields         map[string]FieldInfo
	Parent         *ClassInfo
	Interfaces     []*ClassInfo
}

// MethodInfo, bir metot hakkında bilgi tutar.
type MethodInfo struct {
	Name        string
	Function    *ir.Func
	IsVirtual   bool
	VTableIndex int
	Signature   *types.FuncType
}

// FieldInfo, bir alan hakkında bilgi tutar.
type FieldInfo struct {
	Name      string
	Type      types.Type
	Index     int
	IsPrivate bool
}

// generateClassStatement, bir sınıf tanımlaması için IR üretir.
func (g *IRGenerator) generateClassStatement(stmt *ast.ClassStatement) {
	// Sınıf adını al
	className := stmt.Name.Value

	// Sınıf bilgisi oluştur
	classInfo := &ClassInfo{
		Name:       className,
		Methods:    make(map[string]*MethodInfo),
		Fields:     make(map[string]FieldInfo),
		Interfaces: make([]*ClassInfo, 0),
	}

	// Ebeveyn sınıfı varsa, onu işle
	if stmt.Extends != nil {
		parentName := stmt.Extends.Value
		if parentInfo, exists := g.classTable[parentName]; exists {
			classInfo.Parent = parentInfo
		} else {
			g.ReportError("Ebeveyn sınıf bulunamadı: %s", parentName)
		}
	}

	// Arayüzleri işle
	for _, iface := range stmt.Implements {
		ifaceName := iface.Value
		if ifaceInfo, exists := g.classTable[ifaceName]; exists {
			classInfo.Interfaces = append(classInfo.Interfaces, ifaceInfo)
		} else {
			g.ReportError("Arayüz bulunamadı: %s", ifaceName)
		}
	}

	// Sınıf alanlarını ve metotlarını topla
	fieldTypes := make([]types.Type, 0)
	fieldNames := make([]string, 0)

	// Ebeveyn sınıfın alanlarını ekle
	if classInfo.Parent != nil {
		for fieldName, fieldInfo := range classInfo.Parent.Fields {
			classInfo.Fields[fieldName] = FieldInfo{
				Name:      fieldInfo.Name,
				Type:      fieldInfo.Type,
				Index:     len(fieldTypes),
				IsPrivate: fieldInfo.IsPrivate,
			}
			fieldTypes = append(fieldTypes, fieldInfo.Type)
			fieldNames = append(fieldNames, fieldInfo.Name)
		}
	}

	// VTable işaretçisi ekle (ilk alan olarak)
	vtablePtrType := types.NewPointer(types.Void)
	fieldTypes = append(fieldTypes, vtablePtrType)
	fieldNames = append(fieldNames, "_vtable")

	// Sınıf gövdesindeki ifadeleri işle
	if stmt.Body != nil {
		// Önce alanları işle
		for _, s := range stmt.Body.Statements {
			// Değişken üyeleri işle
			if varStmt, ok := s.(*ast.VarStatement); ok {
				fieldName := varStmt.Name.Value

				// Alan tipini belirle
				var fieldType types.Type = types.I32 // Varsayılan olarak int32
				if varStmt.Type != nil {
					if typeIdent, ok := varStmt.Type.(*ast.Identifier); ok {
						if t, exists := g.typeTable[typeIdent.Value]; exists {
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
		for _, s := range stmt.Body.Statements {
			// Metotları işle
			if funcStmt, ok := s.(*ast.FunctionStatement); ok {
				methodName := funcStmt.Name.Value

				// Metot imzasını oluştur
				paramTypes := make([]types.Type, 0)
				paramTypes = append(paramTypes, types.NewPointer(types.Void)) // this işaretçisi

				for range funcStmt.Parameters {
					paramType := types.I32 // Varsayılan olarak int32
					// TODO: Parametre tiplerini belirle
					paramTypes = append(paramTypes, paramType)
				}

				returnType := types.I32 // Varsayılan olarak int32
				// TODO: Dönüş tipini belirle

				// Fonksiyon tipini oluştur
				funcType := types.NewFunc(returnType, paramTypes...)

				// Metot adını oluştur (sınıf adı + metot adı)
				fullMethodName := fmt.Sprintf("%s_%s", className, methodName)

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
	vtableTypes := make([]types.Type, 0)
	vtableFuncs := make([]*ir.Func, 0)

	// Ebeveyn sınıfın VTable'ını miras al
	if classInfo.Parent != nil {
		for _, methodInfo := range classInfo.Parent.Methods {
			if methodInfo.IsVirtual {
				vtableTypes = append(vtableTypes, methodInfo.Signature)
				vtableFuncs = append(vtableFuncs, methodInfo.Function)
			}
		}
	}

	// Sanal metotları VTable'a ekle
	for _, methodInfo := range classInfo.Methods {
		if methodInfo.IsVirtual {
			// Ebeveyn sınıfta aynı isimde bir metot var mı kontrol et
			overridden := false
			if classInfo.Parent != nil {
				if parentMethod, exists := classInfo.Parent.Methods[methodInfo.Name]; exists && parentMethod.IsVirtual {
					// Metot zaten VTable'da, indeksini kullan
					methodInfo.VTableIndex = parentMethod.VTableIndex
					// VTable'daki fonksiyonu güncelle
					vtableFuncs[methodInfo.VTableIndex] = methodInfo.Function
					overridden = true
				}
			}

			if !overridden {
				// Yeni bir VTable girişi ekle
				methodInfo.VTableIndex = len(vtableTypes)
				vtableTypes = append(vtableTypes, methodInfo.Signature)
				vtableFuncs = append(vtableFuncs, methodInfo.Function)
			}
		}
	}

	// VTable tipini oluştur
	vtableType := types.NewStruct(vtableTypes...)
	classInfo.VTableType = vtableType

	// VTable örneğini oluştur
	vtableInit := constant.NewStruct(vtableType)
	vtableGlobal := g.module.NewGlobalDef(fmt.Sprintf("%s_vtable", className), vtableInit)
	classInfo.VTableInstance = vtableGlobal

	// Sınıf bilgisini kaydet
	g.classTable[className] = classInfo

	// Tipi kaydet
	g.typeTable[className] = structType
}

// generateNewExpression, bir new ifadesi için IR üretir.
func (g *IRGenerator) generateNewExpression(expr *ast.NewExpression) value.Value {
	// Sınıf adını al
	var className string
	if classIdent, ok := expr.Class.(*ast.Identifier); ok {
		className = classIdent.Value
	} else {
		g.ReportError("Sınıf adı bir tanımlayıcı olmalıdır")
		return nil
	}

	// Sınıf bilgisini bul
	classInfo, exists := g.classTable[className]
	if !exists {
		g.ReportError("Sınıf bulunamadı: %s", className)
		return nil
	}

	if g.currentBB == nil {
		g.ReportError("Geçerli bir blok yok, new ifadesi değerlendirilemiyor")
		return nil
	}

	// Bellek ayır
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
}

// getMallocFunction, malloc fonksiyonunu döndürür.
func (g *IRGenerator) getMallocFunction() *ir.Func {
	// Malloc fonksiyonunu bul veya oluştur
	mallocFunc := g.getFunction("malloc")
	if mallocFunc == nil {
		// Malloc fonksiyonunu tanımla
		mallocFunc = g.module.NewFunc("malloc", types.NewPointer(types.I8), ir.NewParam("size", types.I64))
		g.symbolTable["malloc"] = mallocFunc
	}
	return mallocFunc
}

// generateMemberExpression, bir üye erişim ifadesi için IR üretir.
func (g *IRGenerator) generateMemberExpression(expr *ast.MemberExpression) value.Value {
	// Nesneyi değerlendir
	obj := g.generateExpression(expr.Object)
	if obj == nil {
		return nil
	}

	// Üye adını al
	var memberName string
	if memberIdent, ok := expr.Member.(*ast.Identifier); ok {
		memberName = memberIdent.Value
	} else {
		g.ReportError("Üye adı bir tanımlayıcı olmalıdır")
		return nil
	}

	// Nesnenin tipini kontrol et
	objType := obj.Type()
	if !types.IsPointer(objType) {
		g.ReportError("Üye erişimi için nesne bir işaretçi olmalıdır")
		return nil
	}

	// Struct tipini al
	structType, ok := objType.(*types.PointerType).ElemType.(*types.StructType)
	if !ok {
		g.ReportError("Üye erişimi için nesne bir struct işaretçisi olmalıdır")
		return nil
	}

	// Sınıf adını bul
	var className string
	for name, typ := range g.typeTable {
		if typ == structType {
			className = name
			break
		}
	}

	if className == "" {
		g.ReportError("Sınıf adı bulunamadı")
		return nil
	}

	// Sınıf bilgisini bul
	classInfo, exists := g.classTable[className]
	if !exists {
		g.ReportError("Sınıf bilgisi bulunamadı: %s", className)
		return nil
	}

	// Üye bir alan mı yoksa metot mu?
	if fieldInfo, exists := classInfo.Fields[memberName]; exists {
		// Alan erişimi
		if g.currentBB == nil {
			g.ReportError("Geçerli bir blok yok, alan erişimi değerlendirilemiyor")
			return nil
		}

		// Alan işaretçisini al
		fieldPtr := g.currentBB.NewGetElementPtr(structType, obj, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, int64(fieldInfo.Index)))

		// Alanın değerini yükle
		return g.currentBB.NewLoad(fieldInfo.Type, fieldPtr)
	} else if methodInfo, exists := classInfo.Methods[memberName]; exists {
		// Metot erişimi
		// Metot çağrısı için bir fonksiyon işaretçisi döndür
		return methodInfo.Function
	}

	g.ReportError("Üye bulunamadı: %s", memberName)
	return nil
}

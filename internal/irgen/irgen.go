package irgen

import (
	"fmt"
	"goplus/internal/ast"
	"goplus/internal/semantic"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// IRGenerator, AST'yi LLVM IR'ına veya benzer bir ara koda dönüştürür.
type IRGenerator struct {
	errors      []string
	moduleName  string
	module      *ir.Module
	currentFunc *ir.Func
	currentBB   *ir.Block
	symbolTable map[string]value.Value // Sembol tablosu
	typeTable   map[string]types.Type  // Tip tablosu
	analyzer    *semantic.Analyzer     // Semantik analizci
}

// New, yeni bir IRGenerator oluşturur.
func New() *IRGenerator {
	module := ir.NewModule()
	return &IRGenerator{
		errors:      []string{},
		module:      module,
		moduleName:  "goplus_module",
		symbolTable: make(map[string]value.Value),
		typeTable:   make(map[string]types.Type),
	}
}

// NewWithAnalyzer, semantik analizci ile yeni bir IRGenerator oluşturur.
func NewWithAnalyzer(analyzer *semantic.Analyzer) *IRGenerator {
	module := ir.NewModule()
	return &IRGenerator{
		errors:      []string{},
		module:      module,
		moduleName:  "goplus_module",
		symbolTable: make(map[string]value.Value),
		typeTable:   make(map[string]types.Type),
		analyzer:    analyzer,
	}
}

// Errors, IR üretimi sırasında karşılaşılan hataları döndürür.
func (g *IRGenerator) Errors() []string {
	return g.errors
}

// ReportError, bir hata mesajı ekler.
func (g *IRGenerator) ReportError(format string, args ...any) {
	g.errors = append(g.errors, fmt.Sprintf(format, args...))
}

// GenerateProgram, programın AST'sinden IR üretir.
func (g *IRGenerator) GenerateProgram(program *ast.Program) (string, error) {
	// Modülü sıfırla
	g.module = ir.NewModule()
	g.module.SourceFilename = g.moduleName

	// Temel tipleri tanımla
	g.defineBasicTypes()

	// Hata kontrolü
	if len(g.Errors()) > 0 {
		return "", fmt.Errorf("IR üretimi sırasında hatalar oluştu: %v", g.Errors())
	}

	// AST düğümlerini gezerek IR üretme
	for _, stmt := range program.Statements {
		g.generateStatement(stmt)
	}

	// Main fonksiyonu yoksa oluştur
	if g.getFunction("main") == nil {
		g.createMainFunction()
	}

	// Hata kontrolü
	if len(g.Errors()) > 0 {
		return "", fmt.Errorf("IR üretimi sırasında hatalar oluştu: %v", g.Errors())
	}

	// Optimizasyon geçişleri uygula
	g.applyOptimizations()

	// Modülü string olarak döndür
	return g.module.String(), nil
}

// applyOptimizations, IR koduna optimizasyon geçişleri uygular.
func (g *IRGenerator) applyOptimizations() {
	// Şu anda optimizasyon işlemleri optimizer paketi tarafından yapılıyor
	// Bu metod, ileride doğrudan IR üzerinde optimizasyon yapmak için kullanılabilir
}

// defineBasicTypes, temel tipleri tanımlar.
func (g *IRGenerator) defineBasicTypes() {
	// Temel tipleri tanımla
	g.typeTable["int"] = types.I32
	g.typeTable["int8"] = types.I8
	g.typeTable["int16"] = types.I16
	g.typeTable["int32"] = types.I32
	g.typeTable["int64"] = types.I64
	g.typeTable["uint"] = types.I32
	g.typeTable["uint8"] = types.I8
	g.typeTable["uint16"] = types.I16
	g.typeTable["uint32"] = types.I32
	g.typeTable["uint64"] = types.I64
	g.typeTable["float32"] = types.Float
	g.typeTable["float64"] = types.Double
	g.typeTable["bool"] = types.I1
	g.typeTable["byte"] = types.I8
	g.typeTable["rune"] = types.I32
	g.typeTable["string"] = types.NewPointer(types.I8) // Basitleştirilmiş string temsili
}

// getFunction, belirtilen isimde bir fonksiyonu döndürür.
func (g *IRGenerator) getFunction(name string) *ir.Func {
	for _, f := range g.module.Funcs {
		if f.Name() == name {
			return f
		}
	}
	return nil
}

// createMainFunction, main fonksiyonunu oluşturur.
func (g *IRGenerator) createMainFunction() {
	// Main fonksiyonu oluştur
	mainFunc := g.module.NewFunc("main", types.I32)
	entryBlock := mainFunc.NewBlock("entry")

	// Return 0
	entryBlock.NewRet(constant.NewInt(types.I32, 0))
}

// generateStatement, bir deyim için IR üretir.
func (g *IRGenerator) generateStatement(stmt ast.Statement) {
	switch s := stmt.(type) {
	case *ast.PackageStatement:
		// Paket bildirimi için özel bir işlem yapmıyoruz
		// Sadece modül adını ayarlıyoruz
		g.moduleName = s.Name.Value
	case *ast.ExpressionStatement:
		g.generateExpression(s.Expression)
	case *ast.VarStatement:
		g.generateVarStatement(s)
	case *ast.ReturnStatement:
		g.generateReturnStatement(s)
	case *ast.BlockStatement:
		g.generateBlockStatement(s)
	case *ast.WhileStatement:
		g.generateWhileStatement(s)
	// Fonksiyon tanımlamaları için özel bir durum eklenebilir
	// Şimdilik bu durumu atlıyoruz
	case *ast.ClassStatement:
		g.generateClassStatement(s)
	default:
		g.ReportError("Desteklenmeyen deyim türü: %T", s)
	}
}

// generateExpression, bir ifade için IR üretir ve değeri döndürür.
func (g *IRGenerator) generateExpression(expr ast.Expression) value.Value {
	switch e := expr.(type) {
	case *ast.Identifier:
		return g.generateIdentifier(e)
	case *ast.IntegerLiteral:
		return g.generateIntegerLiteral(e)
	case *ast.FloatLiteral:
		return g.generateFloatLiteral(e)
	case *ast.StringLiteral:
		return g.generateStringLiteral(e)
	case *ast.BooleanLiteral:
		return g.generateBooleanLiteral(e)
	case *ast.PrefixExpression:
		return g.generatePrefixExpression(e)
	case *ast.InfixExpression:
		return g.generateInfixExpression(e)
	case *ast.CallExpression:
		return g.generateCallExpression(e)
	case *ast.FunctionLiteral:
		return g.generateFunctionLiteral(e)
	case *ast.IfExpression:
		return g.generateIfExpression(e)
	default:
		g.ReportError("Desteklenmeyen ifade türü: %T", e)
		return nil
	}
}

// getExpressionType, bir ifadenin tipini döndürür.
func (g *IRGenerator) getExpressionType(expr ast.Expression) types.Type {
	switch e := expr.(type) {
	case *ast.Identifier:
		// Tanımlayıcının tipini bul
		if val, exists := g.symbolTable[e.Value]; exists {
			return g.getValueType(val)
		}
		return nil
	case *ast.IntegerLiteral:
		return types.I32 // Varsayılan olarak int32
	case *ast.FloatLiteral:
		return types.Double // Varsayılan olarak float64
	case *ast.StringLiteral:
		return types.NewPointer(types.I8) // Basitleştirilmiş string temsili
	case *ast.BooleanLiteral:
		return types.I1
	case *ast.PrefixExpression:
		return g.getExpressionType(e.Right)
	case *ast.InfixExpression:
		// Aritmetik operatörler için
		if e.Operator == "+" || e.Operator == "-" || e.Operator == "*" || e.Operator == "/" {
			leftType := g.getExpressionType(e.Left)
			rightType := g.getExpressionType(e.Right)
			// Tip yükseltme (type promotion)
			if leftType == types.Double || rightType == types.Double {
				return types.Double
			}
			return types.I32
		}
		// Karşılaştırma operatörleri için
		if e.Operator == "==" || e.Operator == "!=" || e.Operator == "<" || e.Operator == ">" || e.Operator == "<=" || e.Operator == ">=" {
			return types.I1
		}
		return types.I32
	default:
		g.ReportError("Desteklenmeyen ifade türü (tip belirlenemiyor): %T", e)
		return nil
	}
}

// getValueType, bir değerin tipini döndürür.
func (g *IRGenerator) getValueType(val value.Value) types.Type {
	if val == nil {
		return nil
	}
	return val.Type()
}

// generateConstantExpression, sabit bir ifade için IR üretir.
func (g *IRGenerator) generateConstantExpression(expr ast.Expression) constant.Constant {
	switch e := expr.(type) {
	case *ast.IntegerLiteral:
		return constant.NewInt(types.I32, e.Value)
	case *ast.FloatLiteral:
		return constant.NewFloat(types.Double, e.Value)
	case *ast.BooleanLiteral:
		if e.Value {
			return constant.NewInt(types.I1, 1)
		}
		return constant.NewInt(types.I1, 0)
	case *ast.StringLiteral:
		// String sabitleri için global değişken oluştur
		strConst := g.module.NewGlobalDef("", constant.NewCharArrayFromString(e.Value+"\x00"))
		return constant.NewGetElementPtr(strConst.ContentType, strConst, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	default:
		g.ReportError("Desteklenmeyen sabit ifade türü: %T", e)
		return nil
	}
}

// Temel ifade türleri için IR üretme fonksiyonları

func (g *IRGenerator) generateIdentifier(ident *ast.Identifier) value.Value {
	// Tanımlayıcının değerini sembol tablosundan bul
	if val, exists := g.symbolTable[ident.Value]; exists {
		// Eğer değer bir pointer ise (örn. alloca), yükle
		if ptr, ok := val.(value.Value); ok && types.IsPointer(ptr.Type()) {
			if g.currentBB != nil {
				return g.currentBB.NewLoad(ptr.Type().(*types.PointerType).ElemType, ptr)
			}
		}
		return val
	}

	g.ReportError("Tanımlanmamış tanımlayıcı: %s", ident.Value)
	return nil
}

func (g *IRGenerator) generateIntegerLiteral(lit *ast.IntegerLiteral) value.Value {
	return constant.NewInt(types.I32, lit.Value)
}

func (g *IRGenerator) generateFloatLiteral(lit *ast.FloatLiteral) value.Value {
	return constant.NewFloat(types.Double, lit.Value)
}

func (g *IRGenerator) generateStringLiteral(lit *ast.StringLiteral) value.Value {
	// String sabitleri için global değişken oluştur
	strConst := g.module.NewGlobalDef("", constant.NewCharArrayFromString(lit.Value+"\x00"))
	return constant.NewGetElementPtr(strConst.ContentType, strConst, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
}

func (g *IRGenerator) generateBooleanLiteral(lit *ast.BooleanLiteral) value.Value {
	if lit.Value {
		return constant.NewInt(types.I1, 1)
	}
	return constant.NewInt(types.I1, 0)
}

// Karmaşık ifade türleri için IR üretme fonksiyonları

func (g *IRGenerator) generatePrefixExpression(expr *ast.PrefixExpression) value.Value {
	right := g.generateExpression(expr.Right)
	if right == nil {
		return nil
	}

	if g.currentBB == nil {
		g.ReportError("Geçerli bir blok yok, önek ifadesi değerlendirilemiyor")
		return nil
	}

	switch expr.Operator {
	case "!":
		// Boolean değil
		return g.currentBB.NewXor(right, constant.NewInt(types.I1, 1))
	case "-":
		// Sayısal negatif
		if types.IsInt(right.Type()) {
			return g.currentBB.NewSub(constant.NewInt(types.I32, 0), right)
		} else if types.IsFloat(right.Type()) {
			return g.currentBB.NewFSub(constant.NewFloat(types.Double, 0), right)
		}
	}

	g.ReportError("Desteklenmeyen önek operatörü: %s", expr.Operator)
	return nil
}

func (g *IRGenerator) generateInfixExpression(expr *ast.InfixExpression) value.Value {
	// Sol ve sağ ifadeleri değerlendir
	left := g.generateExpression(expr.Left)
	right := g.generateExpression(expr.Right)

	if left == nil || right == nil {
		return nil
	}

	if g.currentBB == nil {
		g.ReportError("Geçerli bir blok yok, araek ifadesi değerlendirilemiyor")
		return nil
	}

	// Tip uyumluluğunu kontrol et ve gerekirse dönüşüm yap
	leftType := left.Type()
	rightType := right.Type()

	// Aritmetik ve atama operatörleri
	switch expr.Operator {
	case "+":
		if types.IsInt(leftType) && types.IsInt(rightType) {
			return g.currentBB.NewAdd(left, right)
		} else if types.IsFloat(leftType) && types.IsFloat(rightType) {
			return g.currentBB.NewFAdd(left, right)
		}
	case "-":
		if types.IsInt(leftType) && types.IsInt(rightType) {
			return g.currentBB.NewSub(left, right)
		} else if types.IsFloat(leftType) && types.IsFloat(rightType) {
			return g.currentBB.NewFSub(left, right)
		}
	case "*":
		if types.IsInt(leftType) && types.IsInt(rightType) {
			return g.currentBB.NewMul(left, right)
		} else if types.IsFloat(leftType) && types.IsFloat(rightType) {
			return g.currentBB.NewFMul(left, right)
		}
	case "/":
		if types.IsInt(leftType) && types.IsInt(rightType) {
			return g.currentBB.NewSDiv(left, right) // Signed division
		} else if types.IsFloat(leftType) && types.IsFloat(rightType) {
			return g.currentBB.NewFDiv(left, right)
		}
	case "=":
		// Atama operatörü
		// Sol taraf bir tanımlayıcı olmalı
		if ident, ok := expr.Left.(*ast.Identifier); ok {
			// Tanımlayıcının değerini sembol tablosundan bul
			if val, exists := g.symbolTable[ident.Value]; exists {
				// Değeri ata
				g.currentBB.NewStore(right, val)
				return right
			} else {
				g.ReportError("Tanımlanmamış tanımlayıcı: %s", ident.Value)
				return nil
			}
		} else {
			g.ReportError("Atama operatörünün sol tarafı bir tanımlayıcı olmalıdır")
			return nil
		}
	// Karşılaştırma operatörleri
	case "==":
		if types.IsInt(leftType) && types.IsInt(rightType) {
			return g.currentBB.NewICmp(enum.IPredEQ, left, right)
		} else if types.IsFloat(leftType) && types.IsFloat(rightType) {
			return g.currentBB.NewFCmp(enum.FPredOEQ, left, right)
		}
	case "!=":
		if types.IsInt(leftType) && types.IsInt(rightType) {
			return g.currentBB.NewICmp(enum.IPredNE, left, right)
		} else if types.IsFloat(leftType) && types.IsFloat(rightType) {
			return g.currentBB.NewFCmp(enum.FPredONE, left, right)
		}
	case "<":
		if types.IsInt(leftType) && types.IsInt(rightType) {
			return g.currentBB.NewICmp(enum.IPredSLT, left, right) // Signed less than
		} else if types.IsFloat(leftType) && types.IsFloat(rightType) {
			return g.currentBB.NewFCmp(enum.FPredOLT, left, right)
		}
	case ">":
		if types.IsInt(leftType) && types.IsInt(rightType) {
			return g.currentBB.NewICmp(enum.IPredSGT, left, right) // Signed greater than
		} else if types.IsFloat(leftType) && types.IsFloat(rightType) {
			return g.currentBB.NewFCmp(enum.FPredOGT, left, right)
		}
	case "<=":
		if types.IsInt(leftType) && types.IsInt(rightType) {
			return g.currentBB.NewICmp(enum.IPredSLE, left, right) // Signed less equal
		} else if types.IsFloat(leftType) && types.IsFloat(rightType) {
			return g.currentBB.NewFCmp(enum.FPredOLE, left, right)
		}
	case ">=":
		if types.IsInt(leftType) && types.IsInt(rightType) {
			return g.currentBB.NewICmp(enum.IPredSGE, left, right) // Signed greater equal
		} else if types.IsFloat(leftType) && types.IsFloat(rightType) {
			return g.currentBB.NewFCmp(enum.FPredOGE, left, right)
		}
	// Mantıksal operatörler
	case "&&":
		// Kısa devre değerlendirme için bloklar oluştur
		currFunc := g.currentFunc
		if currFunc == nil {
			g.ReportError("Geçerli bir fonksiyon yok, mantıksal AND değerlendirilemiyor")
			return nil
		}

		// Bloklar oluştur
		rightBlock := currFunc.NewBlock("")
		mergeBlock := currFunc.NewBlock("")

		// Sol ifade doğruysa sağ ifadeyi değerlendir, değilse direkt false döndür
		g.currentBB.NewCondBr(left, rightBlock, mergeBlock)

		// Sağ ifadeyi değerlendir
		g.currentBB = rightBlock
		g.generateExpression(expr.Right) // Sonucu kullanmıyoruz
		g.currentBB.NewBr(mergeBlock)

		// Sonuç bloğu
		g.currentBB = mergeBlock
		// Basitleştirilmiş yaklaşım: Sadece false döndür
		return constant.NewInt(types.I1, 0)
	case "||":
		// Kısa devre değerlendirme için bloklar oluştur
		currFunc := g.currentFunc
		if currFunc == nil {
			g.ReportError("Geçerli bir fonksiyon yok, mantıksal OR değerlendirilemiyor")
			return nil
		}

		// Bloklar oluştur
		rightBlock := currFunc.NewBlock("")
		mergeBlock := currFunc.NewBlock("")

		// Sol ifade yanlışsa sağ ifadeyi değerlendir, doğruysa direkt true döndür
		g.currentBB.NewCondBr(left, mergeBlock, rightBlock)

		// Sağ ifadeyi değerlendir
		g.currentBB = rightBlock
		g.generateExpression(expr.Right) // Sonucu kullanmıyoruz
		g.currentBB.NewBr(mergeBlock)

		// Sonuç bloğu
		g.currentBB = mergeBlock
		// Basitleştirilmiş yaklaşım: Sadece true döndür
		return constant.NewInt(types.I1, 1)
	}

	g.ReportError("Desteklenmeyen araek operatörü: %s", expr.Operator)
	return nil
}

func (g *IRGenerator) generateCallExpression(expr *ast.CallExpression) value.Value {
	// Fonksiyon adını al
	var funcName string
	if ident, ok := expr.Function.(*ast.Identifier); ok {
		funcName = ident.Value
	} else {
		g.ReportError("Desteklenmeyen fonksiyon çağrısı türü: %T", expr.Function)
		return nil
	}

	// Fonksiyonu bul
	var fn value.Value
	if val, exists := g.symbolTable[funcName]; exists {
		fn = val
	} else {
		// Fonksiyon bulunamadıysa, dış fonksiyon olarak tanımla
		fn = g.module.NewFunc(funcName, types.I32)
		g.symbolTable[funcName] = fn
	}

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
}

func (g *IRGenerator) generateFunctionLiteral(expr *ast.FunctionLiteral) value.Value {
	// Fonksiyon adını belirle
	funcName := "anonymous_func"

	// Parametre tiplerini belirle
	paramTypes := make([]types.Type, len(expr.Parameters))
	for i := range expr.Parameters {
		paramTypes[i] = types.I32 // Varsayılan olarak int32
	}

	// Dönüş tipini belirle
	var returnType types.Type = types.I32 // Varsayılan olarak int32

	// Fonksiyonu oluştur
	fn := g.module.NewFunc(funcName, returnType)

	// Parametreleri ekle
	for _, paramType := range paramTypes {
		param := ir.NewParam("", paramType)
		fn.Params = append(fn.Params, param)
	}

	// Önceki durumu kaydet
	prevFunc := g.currentFunc
	prevBB := g.currentBB

	// Yeni durumu ayarla
	g.currentFunc = fn
	entryBlock := fn.NewBlock("entry")
	g.currentBB = entryBlock

	// Parametreleri sembol tablosuna ekle
	if len(expr.Parameters) > 0 && len(fn.Params) > 0 {
		for i, param := range expr.Parameters {
			if i < len(fn.Params) {
				paramName := param.Value
				paramVal := fn.Params[i]
				paramVal.SetName(paramName)

				// Parametre için yerel değişken oluştur
				alloca := entryBlock.NewAlloca(paramTypes[i])
				alloca.SetName(paramName + ".addr")
				entryBlock.NewStore(paramVal, alloca)

				g.symbolTable[paramName] = alloca
			}
		}
	}

	// Fonksiyon gövdesini işle
	if expr.Body != nil {
		g.generateBlockStatement(expr.Body)
	}

	// Eğer son blok bir dönüş ifadesi ile bitmiyorsa, varsayılan dönüş ekle
	if g.currentBB.Term == nil {
		g.currentBB.NewRet(constant.NewInt(types.I32, 0))
	}

	// Önceki durumu geri yükle
	g.currentFunc = prevFunc
	g.currentBB = prevBB

	return fn
}

// Deyim türleri için IR üretme fonksiyonları

func (g *IRGenerator) generateVarStatement(stmt *ast.VarStatement) {
	// Değişken adını al
	varName := stmt.Name.Value

	// Değişken tipini belirle
	var varType types.Type
	if stmt.Type != nil {
		// Tip belirtilmişse, bu tipi kullan
		if typeIdent, ok := stmt.Type.(*ast.Identifier); ok {
			if t, exists := g.typeTable[typeIdent.Value]; exists {
				varType = t
			} else {
				g.ReportError("Bilinmeyen tip: %s", typeIdent.Value)
				return
			}
		} else {
			g.ReportError("Desteklenmeyen tip ifadesi: %T", stmt.Type)
			return
		}
	} else if stmt.Value != nil {
		// Tip belirtilmemişse ve değer varsa, değerin tipini kullan
		exprType := g.getExpressionType(stmt.Value)
		if exprType != nil {
			varType = exprType
		} else {
			g.ReportError("Değişken tipi belirlenemedi: %s", varName)
			return
		}
	} else {
		g.ReportError("Değişken tipi belirtilmemiş ve değer atanmamış: %s", varName)
		return
	}

	// Değişken global mi yoksa lokal mi?
	if g.currentFunc == nil {
		// Global değişken
		globalVar := g.module.NewGlobalDef(varName, constant.NewZeroInitializer(varType))
		g.symbolTable[varName] = globalVar

		// Değer atanmışsa, değeri ata
		if stmt.Value != nil {
			if constVal := g.generateConstantExpression(stmt.Value); constVal != nil {
				globalVar.Init = constVal
			}
		}
	} else {
		// Lokal değişken
		if g.currentBB == nil {
			g.ReportError("Geçerli bir blok yok, değişken tanımlanamıyor: %s", varName)
			return
		}

		// Değişken için bellek ayır
		alloca := g.currentBB.NewAlloca(varType)
		alloca.SetName(varName)
		g.symbolTable[varName] = alloca

		// Değer atanmışsa, değeri ata
		if stmt.Value != nil {
			val := g.generateExpression(stmt.Value)
			if val != nil {
				g.currentBB.NewStore(val, alloca)
			}
		}
	}
}

func (g *IRGenerator) generateReturnStatement(stmt *ast.ReturnStatement) {
	if g.currentFunc == nil {
		g.ReportError("Fonksiyon dışında dönüş deyimi kullanılamaz")
		return
	}

	if g.currentBB == nil {
		g.ReportError("Geçerli bir blok yok, dönüş deyimi değerlendirilemiyor")
		return
	}

	// Dönüş değeri varsa değerlendir
	if stmt.ReturnValue != nil {
		retVal := g.generateExpression(stmt.ReturnValue)
		if retVal != nil {
			g.currentBB.NewRet(retVal)
		} else {
			g.currentBB.NewRet(constant.NewInt(types.I32, 0)) // Varsayılan dönüş değeri
		}
	} else {
		// Dönüş değeri yoksa void dönüş
		g.currentBB.NewRet(constant.NewInt(types.I32, 0)) // Varsayılan dönüş değeri
	}
}

func (g *IRGenerator) generateBlockStatement(stmt *ast.BlockStatement) {
	// Blok içindeki tüm deyimleri değerlendir
	for _, s := range stmt.Statements {
		g.generateStatement(s)

		// Eğer bir dönüş deyimi ile karşılaşıldıysa, sonraki deyimleri değerlendirme
		if g.currentBB != nil && g.currentBB.Term != nil {
			break
		}
	}
}

// Fonksiyon tanımlamaları için özel bir fonksiyon eklenebilir
// Şimdilik bu fonksiyonu kaldırıyoruz

// generateIfExpression, bir if ifadesi için IR üretir.
func (g *IRGenerator) generateIfExpression(expr *ast.IfExpression) value.Value {
	// Koşul ifadesini değerlendir
	condition := g.generateExpression(expr.Condition)
	if condition == nil {
		return nil
	}

	if g.currentFunc == nil {
		g.ReportError("Geçerli bir fonksiyon yok, if ifadesi değerlendirilemiyor")
		return nil
	}

	// Bloklar oluştur
	thenBlock := g.currentFunc.NewBlock("if.then")
	elseBlock := g.currentFunc.NewBlock("if.else")
	mergeBlock := g.currentFunc.NewBlock("if.end")

	// Koşula göre dallanma
	g.currentBB.NewCondBr(condition, thenBlock, elseBlock)

	// Then bloğunu işle
	g.currentBB = thenBlock
	if expr.Consequence != nil {
		g.generateBlockStatement(expr.Consequence)
		// Eğer blok bir dönüş ifadesi ile bitmiyorsa, merge bloğuna git
		if g.currentBB.Term == nil {
			g.currentBB.NewBr(mergeBlock)
		}
	} else {
		g.currentBB.NewBr(mergeBlock)
	}

	// Else bloğunu işle
	g.currentBB = elseBlock
	if expr.Alternative != nil {
		g.generateBlockStatement(expr.Alternative)
		// Eğer blok bir dönüş ifadesi ile bitmiyorsa, merge bloğuna git
		if g.currentBB.Term == nil {
			g.currentBB.NewBr(mergeBlock)
		}
	} else {
		g.currentBB.NewBr(mergeBlock)
	}

	// Merge bloğuna geç
	g.currentBB = mergeBlock

	// If ifadesi bir değer döndürmez, sadece kontrol akışını değiştirir
	return nil
}

// generateWhileStatement, bir while döngüsü için IR üretir.
func (g *IRGenerator) generateWhileStatement(stmt *ast.WhileStatement) {
	if g.currentFunc == nil {
		g.ReportError("Geçerli bir fonksiyon yok, while döngüsü değerlendirilemiyor")
		return
	}

	// Bloklar oluştur
	condBlock := g.currentFunc.NewBlock("while.cond")
	bodyBlock := g.currentFunc.NewBlock("while.body")
	endBlock := g.currentFunc.NewBlock("while.end")

	// Koşul bloğuna git
	g.currentBB.NewBr(condBlock)

	// Koşul bloğunu işle
	g.currentBB = condBlock
	condition := g.generateExpression(stmt.Condition)
	if condition == nil {
		return
	}

	// Koşula göre dallanma
	g.currentBB.NewCondBr(condition, bodyBlock, endBlock)

	// Döngü gövdesini işle
	g.currentBB = bodyBlock
	if stmt.Body != nil {
		g.generateBlockStatement(stmt.Body)
	}

	// Koşul bloğuna geri dön
	if g.currentBB.Term == nil {
		g.currentBB.NewBr(condBlock)
	}

	// Döngü sonrası bloğa geç
	g.currentBB = endBlock
}

func (g *IRGenerator) generateClassStatement(stmt *ast.ClassStatement) {
	// Sınıf adını al
	className := stmt.Name.Value

	// Şimdilik sınıfları basit bir struct olarak temsil ediyoruz
	// Sınıf üyelerini topla
	memberTypes := make([]types.Type, 0)
	memberNames := make([]string, 0)

	// Sınıf gövdesindeki ifadeleri işle
	if stmt.Body != nil {
		for _, s := range stmt.Body.Statements {
			// Şimdilik sadece değişken üyeleri destekliyoruz
			if varStmt, ok := s.(*ast.VarStatement); ok {
				memberName := varStmt.Name.Value

				// Üye tipini belirle
				var memberType types.Type = types.I32 // Varsayılan olarak int32
				if varStmt.Type != nil {
					if typeIdent, ok := varStmt.Type.(*ast.Identifier); ok {
						if t, exists := g.typeTable[typeIdent.Value]; exists {
							memberType = t
						}
					}
				}

				memberTypes = append(memberTypes, memberType)
				memberNames = append(memberNames, memberName)
			}
		}
	}

	// Struct tipini oluştur
	structType := types.NewStruct(memberTypes...)

	// Tipi kaydet
	g.typeTable[className] = structType

	// Sınıf metotlarını işle
	if stmt.Body != nil {
		for _, s := range stmt.Body.Statements {
			if funcLit, ok := s.(*ast.ExpressionStatement); ok {
				if fn, ok := funcLit.Expression.(*ast.FunctionLiteral); ok {
					// Metot adını al
					methodName := className + "_method"
					_ = methodName // Şimdilik kullanmıyoruz

					// Metodu oluştur
					g.generateFunctionLiteral(fn)
				}
			}
		}
	}
}

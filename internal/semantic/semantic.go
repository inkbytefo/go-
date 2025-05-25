package semantic

import (
	"fmt"

	"github.com/inkbytefo/go-minus/internal/ast"

	"github.com/inkbytefo/go-minus/internal/token"
)

// Analyzer, semantik analiz işlemlerini gerçekleştirir.
type Analyzer struct {
	errorReporter *ErrorReporter
	currentScope  *Scope
	globalScope   *Scope
	packageName   string
	imports       []string
	typeInference bool // Tip çıkarımı etkin mi?
	inferencer    *TypeInference
}

// New, yeni bir Analyzer oluşturur.
func New() *Analyzer {
	globalScope := NewScope(nil)
	globalScope.IsGlobal = true

	a := &Analyzer{
		errorReporter: NewErrorReporter(),
		currentScope:  globalScope,
		globalScope:   globalScope,
		packageName:   "",
		imports:       []string{},
		typeInference: true, // Varsayılan olarak tip çıkarımı etkin
	}

	a.inferencer = NewTypeInference(a)

	// Built-in functions ve packages'ları ekle
	a.initializeBuiltins()

	return a
}

// PrintGlobalScope, global scope'daki sembolleri yazdırır (debug için).
func (a *Analyzer) PrintGlobalScope() {
	fmt.Printf("Global scope symbols (%d):\n", len(a.globalScope.Symbols))
	for name, symbol := range a.globalScope.Symbols {
		fmt.Printf("  %s: %s\n", name, symbol.Type.String())
	}
}

// initializeBuiltins, built-in functions ve packages'ları global scope'a ekler.
func (a *Analyzer) initializeBuiltins() {
	// Built-in functions
	a.addBuiltinFunction("println", []SymbolType{}, VOID_TYPE)
	a.addBuiltinFunction("print", []SymbolType{}, VOID_TYPE)
	a.addBuiltinFunction("panic", []SymbolType{UNKNOWN_TYPE}, VOID_TYPE)
	a.addBuiltinFunction("recover", []SymbolType{}, UNKNOWN_TYPE)
	a.addBuiltinFunction("len", []SymbolType{UNKNOWN_TYPE}, INTEGER_TYPE)
	a.addBuiltinFunction("cap", []SymbolType{UNKNOWN_TYPE}, INTEGER_TYPE)
	a.addBuiltinFunction("make", []SymbolType{UNKNOWN_TYPE}, UNKNOWN_TYPE)
	a.addBuiltinFunction("new", []SymbolType{UNKNOWN_TYPE}, UNKNOWN_TYPE)

	// Standard library packages
	a.addStandardPackage("fmt")
	a.addStandardPackage("os")
	a.addStandardPackage("io")
	a.addStandardPackage("strings")
	a.addStandardPackage("math")
}

// addBuiltinFunction, bir built-in function'ı global scope'a ekler.
func (a *Analyzer) addBuiltinFunction(name string, paramTypes []SymbolType, returnType SymbolType) {
	symbol := a.globalScope.Define(name, FUNCTION_TYPE, token.Token{})
	symbol.Signature = &FunctionSignature{
		Parameters: make([]*Symbol, len(paramTypes)),
		ReturnType: returnType,
	}

	// Parametre sembollerini oluştur
	for i, paramType := range paramTypes {
		symbol.Signature.Parameters[i] = &Symbol{
			Name: "",
			Type: paramType,
		}
	}
}

// addStandardPackage, bir standard library package'ını global scope'a ekler.
func (a *Analyzer) addStandardPackage(name string) {
	symbol := a.globalScope.Define(name, PACKAGE_TYPE, token.Token{})

	// Package'a özgü fonksiyonları ekle
	switch name {
	case "fmt":
		a.addVariadicPackageFunction(symbol, "Println", []SymbolType{}, VOID_TYPE)              // Variadic function
		a.addVariadicPackageFunction(symbol, "Printf", []SymbolType{STRING_TYPE}, VOID_TYPE)    // Format string + variadic
		a.addVariadicPackageFunction(symbol, "Print", []SymbolType{}, VOID_TYPE)                // Variadic function
		a.addVariadicPackageFunction(symbol, "Sprintf", []SymbolType{STRING_TYPE}, STRING_TYPE) // Format string + variadic
	case "os":
		a.addPackageFunction(symbol, "Exit", []SymbolType{INTEGER_TYPE}, VOID_TYPE)
		a.addPackageFunction(symbol, "Getenv", []SymbolType{STRING_TYPE}, STRING_TYPE)
		a.addPackageFunction(symbol, "Setenv", []SymbolType{STRING_TYPE, STRING_TYPE}, UNKNOWN_TYPE)
	case "math":
		a.addPackageFunction(symbol, "Max", []SymbolType{FLOAT_TYPE, FLOAT_TYPE}, FLOAT_TYPE)
		a.addPackageFunction(symbol, "Min", []SymbolType{FLOAT_TYPE, FLOAT_TYPE}, FLOAT_TYPE)
		a.addPackageFunction(symbol, "Abs", []SymbolType{FLOAT_TYPE}, FLOAT_TYPE)
	}
}

// addPackageFunction, bir package'a function ekler.
func (a *Analyzer) addPackageFunction(packageSymbol *Symbol, funcName string, paramTypes []SymbolType, returnType SymbolType) {
	if packageSymbol.Class == nil {
		packageSymbol.Class = &ClassInfo{
			Fields:     make(map[string]*Symbol),
			Methods:    make(map[string]*Symbol),
			Implements: []*Symbol{},
		}
	}

	funcSymbol := &Symbol{
		Name: funcName,
		Type: FUNCTION_TYPE,
		Signature: &FunctionSignature{
			Parameters: make([]*Symbol, len(paramTypes)),
			ReturnType: returnType,
		},
	}

	// Parametre sembollerini oluştur
	for i, paramType := range paramTypes {
		funcSymbol.Signature.Parameters[i] = &Symbol{
			Name: "",
			Type: paramType,
		}
	}

	packageSymbol.Class.Methods[funcName] = funcSymbol
}

// addVariadicPackageFunction, bir package'a variadic function ekler.
func (a *Analyzer) addVariadicPackageFunction(packageSymbol *Symbol, funcName string, paramTypes []SymbolType, returnType SymbolType) {
	if packageSymbol.Class == nil {
		packageSymbol.Class = &ClassInfo{
			Fields:     make(map[string]*Symbol),
			Methods:    make(map[string]*Symbol),
			Implements: []*Symbol{},
		}
	}

	funcSymbol := &Symbol{
		Name: funcName,
		Type: FUNCTION_TYPE,
		Signature: &FunctionSignature{
			Parameters: make([]*Symbol, len(paramTypes)),
			ReturnType: returnType,
			IsVariadic: true, // Variadic function flag'i
		},
	}

	// Parametre sembollerini oluştur
	for i, paramType := range paramTypes {
		funcSymbol.Signature.Parameters[i] = &Symbol{
			Name: "",
			Type: paramType,
		}
	}

	packageSymbol.Class.Methods[funcName] = funcSymbol
}

// EnableTypeInference, tip çıkarımını etkinleştirir.
func (a *Analyzer) EnableTypeInference() {
	a.typeInference = true
}

// DisableTypeInference, tip çıkarımını devre dışı bırakır.
func (a *Analyzer) DisableTypeInference() {
	a.typeInference = false
}

// Analyze, bir AST'yi analiz eder.
func (a *Analyzer) Analyze(program *ast.Program) {
	// Ön analiz: Tüm fonksiyon ve sınıf tanımlarını topla
	a.collectDeclarations(program)

	// Ana analiz: Tüm ifadeleri analiz et
	for _, stmt := range program.Statements {
		a.analyzeStatement(stmt)
	}
}

// Errors, analiz sırasında karşılaşılan hataları döndürür.
func (a *Analyzer) Errors() []string {
	return a.errorReporter.GetAllMessages()
}

// HasErrors, hata olup olmadığını kontrol eder.
func (a *Analyzer) HasErrors() bool {
	return a.errorReporter.HasErrors()
}

// HasWarnings, uyarı olup olmadığını kontrol eder.
func (a *Analyzer) HasWarnings() bool {
	return a.errorReporter.HasWarnings()
}

// PrintErrors, tüm hataları yazdırır.
func (a *Analyzer) PrintErrors() {
	a.errorReporter.PrintAllMessages()
}

// collectDeclarations, tüm fonksiyon ve sınıf tanımlarını toplar.
func (a *Analyzer) collectDeclarations(program *ast.Program) {
	for _, stmt := range program.Statements {
		switch s := stmt.(type) {
		case *ast.ExpressionStatement:
			if fn, ok := s.Expression.(*ast.FunctionLiteral); ok {
				a.collectFunctionDeclaration(fn)
			}
		case *ast.ClassStatement:
			a.collectClassDeclaration(s)
		}
	}
}

// collectFunctionDeclaration, bir fonksiyon tanımını toplar.
func (a *Analyzer) collectFunctionDeclaration(fn *ast.FunctionLiteral) {
	// Fonksiyon adı için bir alan eklenebilir
	// name := fn.Name.Value

	// Parametre tipleri
	paramTypes := make([]SymbolType, len(fn.Parameters))
	for i := range fn.Parameters {
		// Parametre tipi için bir alan eklenebilir
		// paramTypes[i] = a.resolveType(fn.Parameters[i].Type)
		paramTypes[i] = UNKNOWN_TYPE
	}

	// Dönüş tipi
	// var returnType SymbolType = VOID_TYPE
	if fn.ReturnType != nil {
		// returnType = a.resolveType(fn.ReturnType)
		// returnType = UNKNOWN_TYPE
	}

	// Fonksiyon sembolü oluştur
	// symbol := a.currentScope.Define(name, FUNCTION_TYPE, fn.Token)
	// symbol.Signature = &FunctionSignature{
	// 	Parameters: make([]*Symbol, len(paramTypes)),
	// 	ReturnType: returnType,
	// }
}

// collectClassDeclaration, bir sınıf tanımını toplar.
func (a *Analyzer) collectClassDeclaration(class *ast.ClassStatement) {
	name := class.Name.Value

	// Sınıf sembolü oluştur
	symbol := a.currentScope.Define(name, CLASS_TYPE, class.Token)
	symbol.Class = &ClassInfo{
		Fields:     make(map[string]*Symbol),
		Methods:    make(map[string]*Symbol),
		Implements: []*Symbol{},
	}

	// Kalıtım
	if class.Extends != nil {
		// Kalıtım alınan sınıfı çözümle
		// extendsSymbol := a.currentScope.Resolve(class.Extends.Value)
		// if extendsSymbol != nil && extendsSymbol.Type == CLASS_TYPE {
		// 	symbol.Class.Extends = extendsSymbol
		// } else {
		// 	a.reportError(class.Extends.Token, "Kalıtım alınan sınıf bulunamadı: %s", class.Extends.Value)
		// }
	}

	// Arayüz uygulamaları
	for range class.Implements {
		// Uygulanan arayüzü çözümle
		// implSymbol := a.currentScope.Resolve(impl.Value)
		// if implSymbol != nil && implSymbol.Type == INTERFACE_TYPE {
		// 	symbol.Class.Implements = append(symbol.Class.Implements, implSymbol)
		// } else {
		// 	a.reportError(impl.Token, "Uygulanan arayüz bulunamadı: %s", impl.Value)
		// }
	}

	// Sınıf kapsamı oluştur
	classScope := NewScope(a.currentScope)
	classScope.IsClass = true
	classScope.ClassName = name

	// Sınıf üyelerini analiz et
	prevScope := a.currentScope
	a.currentScope = classScope

	// Sınıf gövdesini analiz et
	// for _, stmt := range class.Body.Statements {
	// 	a.analyzeStatement(stmt)
	// }

	a.currentScope = prevScope
}

// analyzeStatement, bir ifadeyi analiz eder.
func (a *Analyzer) analyzeStatement(stmt ast.Statement) Type {
	switch s := stmt.(type) {
	case *ast.VarStatement:
		return a.analyzeVarStatement(s)
	case *ast.ConstStatement:
		return a.analyzeConstStatement(s)
	case *ast.ReturnStatement:
		return a.analyzeReturnStatement(s)
	case *ast.ExpressionStatement:
		return a.analyzeExpression(s.Expression)
	case *ast.BlockStatement:
		return a.analyzeBlockStatement(s)
	case *ast.ForStatement:
		return a.analyzeForStatement(s)
	case *ast.WhileStatement:
		return a.analyzeWhileStatement(s)
	case *ast.ClassStatement:
		return a.analyzeClassStatement(s)
	case *ast.MethodStatement:
		return a.analyzeMethodStatement(s)
	case *ast.TryCatchStatement:
		return a.analyzeTryCatchStatement(s)
	case *ast.ThrowStatement:
		return a.analyzeThrowStatement(s)
	case *ast.ScopeStatement:
		return a.analyzeScopeStatement(s)
	case *ast.PackageStatement:
		return a.analyzePackageStatement(s)
	case *ast.ImportStatement:
		return a.analyzeImportStatement(s)
	default:
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

// analyzeExpression, bir ifadeyi analiz eder.
func (a *Analyzer) analyzeExpression(expr ast.Expression) Type {
	// Tip çıkarımı etkinse, inferencer'ı kullan
	if a.typeInference {
		return a.inferencer.InferType(expr)
	}

	// Tip çıkarımı etkin değilse, manuel analiz yap
	switch e := expr.(type) {
	case *ast.Identifier:
		return a.analyzeIdentifier(e)
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
		return a.analyzePrefixExpression(e)
	case *ast.InfixExpression:
		return a.analyzeInfixExpression(e)
	case *ast.IfExpression:
		return a.analyzeIfExpression(e)
	case *ast.FunctionLiteral:
		return a.analyzeFunctionLiteral(e)
	case *ast.CallExpression:
		return a.analyzeCallExpression(e)
	case *ast.ArrayLiteral:
		return a.analyzeArrayLiteral(e)
	case *ast.IndexExpression:
		return a.analyzeIndexExpression(e)
	case *ast.HashLiteral:
		return a.analyzeHashLiteral(e)
	case *ast.MemberExpression:
		return a.analyzeMemberExpression(e)
	case *ast.NewExpression:
		return a.analyzeNewExpression(e)
	case *ast.TemplateExpression:
		return a.analyzeTemplateExpression(e)
	default:
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

// reportError, bir hata rapor eder.
func (a *Analyzer) reportError(tok token.Token, format string, args ...interface{}) *SemanticError {
	return a.errorReporter.ReportError(tok, format, args...)
}

// reportWarning, bir uyarı rapor eder.
func (a *Analyzer) reportWarning(tok token.Token, format string, args ...interface{}) *SemanticError {
	return a.errorReporter.ReportWarning(tok, format, args...)
}

// reportInfo, bir bilgi rapor eder.
func (a *Analyzer) reportInfo(tok token.Token, format string, args ...interface{}) *SemanticError {
	return a.errorReporter.ReportInfo(tok, format, args...)
}

// Temel analiz fonksiyonları
func (a *Analyzer) analyzeVarStatement(stmt *ast.VarStatement) Type {
	// Değişken tipini belirle
	var varType Type = &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}

	// Değişken değerini analiz et
	if stmt.Value != nil {
		valueType := a.analyzeExpression(stmt.Value)

		// Tip belirtilmişse, tip kontrolü yap
		if stmt.Type != nil {
			// declaredType := a.analyzeExpression(stmt.Type)
			// if !declaredType.Equals(valueType) {
			// 	a.reportError(stmt.Token, "Tip uyuşmazlığı: %s tipindeki değer %s tipindeki değişkene atanamaz", valueType.String(), declaredType.String())
			// }
			// varType = declaredType
		} else {
			// Tip belirtilmemişse, değerin tipini kullan
			varType = valueType
		}
	} else if stmt.Type != nil {
		// Değer yoksa ama tip belirtilmişse, belirtilen tipi kullan
		// varType = a.analyzeExpression(stmt.Type)
	}

	// Değişkeni tanımla
	a.currentScope.Define(stmt.Name.Value, symbolTypeFromType(varType), stmt.Token)

	return varType
}

func (a *Analyzer) analyzeConstStatement(stmt *ast.ConstStatement) Type {
	// Sabit tipini belirle
	var constType Type = &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}

	// Sabit değerini analiz et
	if stmt.Value != nil {
		valueType := a.analyzeExpression(stmt.Value)

		// Tip belirtilmişse, tip kontrolü yap
		if stmt.Type != nil {
			// declaredType := a.analyzeExpression(stmt.Type)
			// if !declaredType.Equals(valueType) {
			// 	a.reportError(stmt.Token, "Tip uyuşmazlığı: %s tipindeki değer %s tipindeki sabite atanamaz", valueType.String(), declaredType.String())
			// }
			// constType = declaredType
		} else {
			// Tip belirtilmemişse, değerin tipini kullan
			constType = valueType
		}
	} else {
		a.reportError(stmt.Token, "Sabit tanımında değer belirtilmelidir")
	}

	// Sabiti tanımla
	symbol := a.currentScope.Define(stmt.Name.Value, symbolTypeFromType(constType), stmt.Token)
	symbol.IsConst = true

	return constType
}

// Artık inference_extra.go dosyasında tanımlandı

// Diğer analiz fonksiyonları buraya eklenecek
func (a *Analyzer) analyzeReturnStatement(stmt *ast.ReturnStatement) Type {
	if stmt.ReturnValue != nil {
		return a.analyzeExpression(stmt.ReturnValue)
	}
	return &BasicType{Name: "void", Kind: VOID_TYPE}
}

func (a *Analyzer) analyzeBlockStatement(stmt *ast.BlockStatement) Type {
	// Yeni bir kapsam oluştur
	blockScope := NewScope(a.currentScope)
	prevScope := a.currentScope
	a.currentScope = blockScope

	var lastType Type = &BasicType{Name: "void", Kind: VOID_TYPE}

	// Blok içindeki ifadeleri analiz et
	for _, s := range stmt.Statements {
		lastType = a.analyzeStatement(s)
	}

	// Önceki kapsama geri dön
	a.currentScope = prevScope

	return lastType
}

func (a *Analyzer) analyzeForStatement(stmt *ast.ForStatement) Type {
	// Yeni bir kapsam oluştur
	forScope := NewScope(a.currentScope)
	prevScope := a.currentScope
	a.currentScope = forScope

	// Başlangıç ifadesini analiz et
	if stmt.Init != nil {
		a.analyzeStatement(stmt.Init)
	}

	// Koşulu analiz et
	if stmt.Condition != nil {
		condType := a.analyzeExpression(stmt.Condition)
		if basicType, ok := condType.(*BasicType); !ok || basicType.Kind != BOOLEAN_TYPE {
			a.reportError(stmt.Token, "For döngüsü koşulu boolean tipinde olmalıdır")
		}
	}

	// Sonrası ifadesini analiz et
	if stmt.Post != nil {
		a.analyzeStatement(stmt.Post)
	}

	// Gövdeyi analiz et
	a.analyzeBlockStatement(stmt.Body)

	// Önceki kapsama geri dön
	a.currentScope = prevScope

	return &BasicType{Name: "void", Kind: VOID_TYPE}
}

func (a *Analyzer) analyzeWhileStatement(stmt *ast.WhileStatement) Type {
	// Koşulu analiz et
	condType := a.analyzeExpression(stmt.Condition)
	if basicType, ok := condType.(*BasicType); !ok || basicType.Kind != BOOLEAN_TYPE {
		a.reportError(stmt.Token, "While döngüsü koşulu boolean tipinde olmalıdır")
	}

	// Gövdeyi analiz et
	a.analyzeBlockStatement(stmt.Body)

	return &BasicType{Name: "void", Kind: VOID_TYPE}
}

// Diğer analiz fonksiyonları
func (a *Analyzer) analyzeClassStatement(stmt *ast.ClassStatement) Type {
	// Sınıf tipini oluştur
	classType := &ClassType{
		Name:       stmt.Name.Value,
		Fields:     make(map[string]Type),
		Methods:    make(map[string]*FunctionType),
		Implements: []*InterfaceType{},
	}

	// Sınıf kapsamı oluştur
	classScope := NewScope(a.currentScope)
	classScope.IsClass = true
	classScope.ClassName = stmt.Name.Value

	// Sınıf üyelerini analiz et
	prevScope := a.currentScope
	a.currentScope = classScope

	// Sınıf gövdesini analiz et
	if stmt.Body != nil {
		a.analyzeBlockStatement(stmt.Body)
	}

	// Önceki kapsama geri dön
	a.currentScope = prevScope

	return classType
}

func (a *Analyzer) analyzeMethodStatement(stmt *ast.MethodStatement) Type {
	// Metot tipini oluştur
	funcType := &FunctionType{
		ParameterTypes: make([]Type, len(stmt.Parameters)),
		ReturnType:     &BasicType{Name: "void", Kind: VOID_TYPE},
	}

	// Parametre tiplerini belirle
	for i := range stmt.Parameters {
		// Parametre tipi için bir alan eklenebilir
		// funcType.ParameterTypes[i] = a.analyzeExpression(stmt.Parameters[i].Type)
		funcType.ParameterTypes[i] = &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}

	// Dönüş tipini belirle
	if stmt.ReturnType != nil {
		// funcType.ReturnType = a.analyzeExpression(stmt.ReturnType)
	}

	// Metot gövdesini analiz et
	if stmt.Body != nil {
		a.analyzeBlockStatement(stmt.Body)
	}

	return funcType
}

func (a *Analyzer) analyzeTryCatchStatement(stmt *ast.TryCatchStatement) Type {
	// Try bloğunu analiz et
	if stmt.Try != nil {
		a.analyzeBlockStatement(stmt.Try)
	}

	// Catch bloklarını analiz et
	for _, catch := range stmt.Catches {
		// Catch parametresini tanımla
		if catch.Parameter != nil {
			catchScope := NewScope(a.currentScope)
			prevScope := a.currentScope
			a.currentScope = catchScope

			// Parametre tipini belirle
			var paramType Type = &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
			if catch.Type != nil {
				// paramType = a.analyzeExpression(catch.Type)
			}

			// Parametreyi tanımla
			a.currentScope.Define(catch.Parameter.Value, symbolTypeFromType(paramType), catch.Parameter.Token)

			// Catch bloğunu analiz et
			if catch.Body != nil {
				a.analyzeBlockStatement(catch.Body)
			}

			// Önceki kapsama geri dön
			a.currentScope = prevScope
		} else {
			// Catch bloğunu analiz et
			if catch.Body != nil {
				a.analyzeBlockStatement(catch.Body)
			}
		}
	}

	// Finally bloğunu analiz et
	if stmt.Finally != nil {
		a.analyzeBlockStatement(stmt.Finally)
	}

	return &BasicType{Name: "void", Kind: VOID_TYPE}
}

func (a *Analyzer) analyzeThrowStatement(stmt *ast.ThrowStatement) Type {
	// Fırlatılan ifadeyi analiz et
	if stmt.Value != nil {
		a.analyzeExpression(stmt.Value)
	}

	return &BasicType{Name: "void", Kind: VOID_TYPE}
}

func (a *Analyzer) analyzeScopeStatement(stmt *ast.ScopeStatement) Type {
	// Scope bloğunu analiz et
	if stmt.Body != nil {
		return a.analyzeBlockStatement(stmt.Body)
	}

	return &BasicType{Name: "void", Kind: VOID_TYPE}
}

func (a *Analyzer) analyzePackageStatement(stmt *ast.PackageStatement) Type {
	// Paket adını kaydet
	// a.packageName = stmt.Name.Value

	return &BasicType{Name: "void", Kind: VOID_TYPE}
}

func (a *Analyzer) analyzeImportStatement(stmt *ast.ImportStatement) Type {
	// Import yolunu kaydet
	importPath := stmt.Path.Value
	a.imports = append(a.imports, importPath)

	// Standard library package'larını handle et
	switch importPath {
	case "fmt":
		// fmt package'ı zaten initializeBuiltins'de eklendi
		// Burada ek bir şey yapmaya gerek yok
	case "os":
		// os package'ı zaten initializeBuiltins'de eklendi
	case "io":
		// io package'ı zaten initializeBuiltins'de eklendi
	case "strings":
		// strings package'ı zaten initializeBuiltins'de eklendi
	case "math":
		// math package'ı zaten initializeBuiltins'de eklendi
	default:
		// Bilinmeyen package için warning verebiliriz
		// a.reportError(stmt.Token, "Bilinmeyen package: %s", importPath)
	}

	return &BasicType{Name: "void", Kind: VOID_TYPE}
}

func (a *Analyzer) analyzeIdentifier(expr *ast.Identifier) Type {
	// Tanımlayıcıyı çözümle
	symbol := a.currentScope.Resolve(expr.Value)
	if symbol == nil {
		a.reportError(expr.Token, "Tanımlanmamış tanımlayıcı: %s", expr.Value)
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

		// Parametre tiplerini ekle
		if symbol.Signature != nil {
			for _, param := range symbol.Signature.Parameters {
				switch param.Type {
				case INTEGER_TYPE:
					funcType.ParameterTypes = append(funcType.ParameterTypes, &BasicType{Name: "int", Kind: INTEGER_TYPE})
				case FLOAT_TYPE:
					funcType.ParameterTypes = append(funcType.ParameterTypes, &BasicType{Name: "float", Kind: FLOAT_TYPE})
				case STRING_TYPE:
					funcType.ParameterTypes = append(funcType.ParameterTypes, &BasicType{Name: "string", Kind: STRING_TYPE})
				case BOOLEAN_TYPE:
					funcType.ParameterTypes = append(funcType.ParameterTypes, &BasicType{Name: "bool", Kind: BOOLEAN_TYPE})
				case CHAR_TYPE:
					funcType.ParameterTypes = append(funcType.ParameterTypes, &BasicType{Name: "char", Kind: CHAR_TYPE})
				default:
					funcType.ParameterTypes = append(funcType.ParameterTypes, &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE})
				}
			}

			// Dönüş tipini belirle
			switch symbol.Signature.ReturnType {
			case INTEGER_TYPE:
				funcType.ReturnType = &BasicType{Name: "int", Kind: INTEGER_TYPE}
			case FLOAT_TYPE:
				funcType.ReturnType = &BasicType{Name: "float", Kind: FLOAT_TYPE}
			case STRING_TYPE:
				funcType.ReturnType = &BasicType{Name: "string", Kind: STRING_TYPE}
			case BOOLEAN_TYPE:
				funcType.ReturnType = &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
			case CHAR_TYPE:
				funcType.ReturnType = &BasicType{Name: "char", Kind: CHAR_TYPE}
			case VOID_TYPE:
				funcType.ReturnType = &BasicType{Name: "void", Kind: VOID_TYPE}
			default:
				funcType.ReturnType = &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
			}
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

func (a *Analyzer) analyzePrefixExpression(expr *ast.PrefixExpression) Type {
	// Sağ tarafı analiz et
	rightType := a.analyzeExpression(expr.Right)

	// Operatöre göre tip kontrolü yap
	switch expr.Operator {
	case "!":
		// ! operatörü boolean tipinde olmalıdır
		if basicType, ok := rightType.(*BasicType); !ok || basicType.Kind != BOOLEAN_TYPE {
			a.reportError(expr.Token, "! operatörü boolean tipinde olmalıdır")
		}
		return &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
	case "-":
		// - operatörü sayısal tipte olmalıdır
		if basicType, ok := rightType.(*BasicType); !ok || (basicType.Kind != INTEGER_TYPE && basicType.Kind != FLOAT_TYPE) {
			a.reportError(expr.Token, "- operatörü sayısal tipte olmalıdır")
		}
		return rightType
	default:
		a.reportError(expr.Token, "Bilinmeyen önek operatörü: %s", expr.Operator)
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

func (a *Analyzer) analyzeInfixExpression(expr *ast.InfixExpression) Type {
	// Sol ve sağ tarafı analiz et
	leftType := a.analyzeExpression(expr.Left)
	rightType := a.analyzeExpression(expr.Right)

	// Operatöre göre tip kontrolü yap
	switch expr.Operator {
	case "+", "-", "*", "/", "%":
		// Aritmetik operatörler sayısal tipte olmalıdır
		if basicLeftType, okLeft := leftType.(*BasicType); !okLeft || (basicLeftType.Kind != INTEGER_TYPE && basicLeftType.Kind != FLOAT_TYPE) {
			a.reportError(expr.Token, "Aritmetik operatörün sol tarafı sayısal tipte olmalıdır")
		}

		if basicRightType, okRight := rightType.(*BasicType); !okRight || (basicRightType.Kind != INTEGER_TYPE && basicRightType.Kind != FLOAT_TYPE) {
			a.reportError(expr.Token, "Aritmetik operatörün sağ tarafı sayısal tipte olmalıdır")
		}

		// Eğer herhangi bir taraf float ise, sonuç float olur
		leftBasicType, leftOk := leftType.(*BasicType)
		rightBasicType, rightOk := rightType.(*BasicType)

		if (leftOk && leftBasicType.Kind == FLOAT_TYPE) || (rightOk && rightBasicType.Kind == FLOAT_TYPE) {
			return &BasicType{Name: "float", Kind: FLOAT_TYPE}
		}

		return &BasicType{Name: "int", Kind: INTEGER_TYPE}
	case "==", "!=", "<", ">", "<=", ">=":
		// Karşılaştırma operatörleri aynı tipte olmalıdır
		if !leftType.Equals(rightType) {
			a.reportError(expr.Token, "Karşılaştırma operatörünün sol ve sağ tarafı aynı tipte olmalıdır")
		}

		return &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
	case "&&", "||":
		// Mantıksal operatörler boolean tipinde olmalıdır
		if basicLeftType, okLeft := leftType.(*BasicType); !okLeft || basicLeftType.Kind != BOOLEAN_TYPE {
			a.reportError(expr.Token, "Mantıksal operatörün sol tarafı boolean tipinde olmalıdır")
		}

		if basicRightType, okRight := rightType.(*BasicType); !okRight || basicRightType.Kind != BOOLEAN_TYPE {
			a.reportError(expr.Token, "Mantıksal operatörün sağ tarafı boolean tipinde olmalıdır")
		}

		return &BasicType{Name: "bool", Kind: BOOLEAN_TYPE}
	case "=":
		// Atama operatörü aynı tipte olmalıdır
		if !leftType.Equals(rightType) {
			a.reportError(expr.Token, "Atama operatörünün sol ve sağ tarafı aynı tipte olmalıdır")
		}

		return leftType
	case ":=":
		// Kısa değişken tanımlama operatörü
		// Sol taraf bir tanımlayıcı olmalıdır
		if ident, ok := expr.Left.(*ast.Identifier); ok {
			// Tanımlayıcıyı tanımla
			a.currentScope.Define(ident.Value, symbolTypeFromType(rightType), ident.Token)
		} else {
			a.reportError(expr.Token, "Kısa değişken tanımlama operatörünün sol tarafı bir tanımlayıcı olmalıdır")
		}

		return rightType
	default:
		a.reportError(expr.Token, "Bilinmeyen araek operatörü: %s", expr.Operator)
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

func (a *Analyzer) analyzeIfExpression(expr *ast.IfExpression) Type {
	// Koşulu analiz et
	condType := a.analyzeExpression(expr.Condition)
	if basicType, ok := condType.(*BasicType); !ok || basicType.Kind != BOOLEAN_TYPE {
		a.reportError(expr.Token, "If ifadesinin koşulu boolean tipinde olmalıdır")
	}

	// Consequence bloğunu analiz et
	var consequenceType Type
	if expr.Consequence != nil {
		consequenceType = a.analyzeBlockStatement(expr.Consequence)
	} else {
		consequenceType = &BasicType{Name: "void", Kind: VOID_TYPE}
	}

	// Alternative bloğunu analiz et
	var alternativeType Type
	if expr.Alternative != nil {
		alternativeType = a.analyzeBlockStatement(expr.Alternative)
	} else {
		alternativeType = &BasicType{Name: "void", Kind: VOID_TYPE}
	}

	// Eğer her iki blok da aynı tipte ise, o tipi döndür
	if consequenceType.Equals(alternativeType) {
		return consequenceType
	}

	// Aksi takdirde, void döndür
	return &BasicType{Name: "void", Kind: VOID_TYPE}
}

func (a *Analyzer) analyzeFunctionLiteral(expr *ast.FunctionLiteral) Type {
	// Fonksiyon tipini oluştur
	funcType := &FunctionType{
		ParameterTypes: make([]Type, len(expr.Parameters)),
		ReturnType:     &BasicType{Name: "void", Kind: VOID_TYPE},
	}

	// Fonksiyon kapsamı oluştur
	funcScope := NewScope(a.currentScope)
	prevScope := a.currentScope
	a.currentScope = funcScope

	// Parametre tiplerini belirle
	for i, param := range expr.Parameters {
		// Parametre tipi için bir alan eklenebilir
		// funcType.ParameterTypes[i] = a.analyzeExpression(param.Type)
		funcType.ParameterTypes[i] = &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}

		// Parametreyi tanımla
		a.currentScope.Define(param.Value, symbolTypeFromType(funcType.ParameterTypes[i]), param.Token)
	}

	// Dönüş tipini belirle
	if expr.ReturnType != nil {
		// funcType.ReturnType = a.analyzeExpression(expr.ReturnType)
	}

	// Fonksiyon gövdesini analiz et
	if expr.Body != nil {
		a.analyzeBlockStatement(expr.Body)
	}

	// Önceki kapsama geri dön
	a.currentScope = prevScope

	return funcType
}

func (a *Analyzer) analyzeCallExpression(expr *ast.CallExpression) Type {
	// Fonksiyonu analiz et
	funcType := a.analyzeExpression(expr.Function)

	// Fonksiyon tipini kontrol et
	if ft, ok := funcType.(*FunctionType); ok {
		// Argüman sayısını kontrol et
		if len(expr.Arguments) != len(ft.ParameterTypes) {
			a.reportError(expr.Token, "Fonksiyon çağrısında yanlış sayıda argüman: %d bekleniyor, %d alındı", len(ft.ParameterTypes), len(expr.Arguments))
		}

		// Argüman tiplerini kontrol et
		for i, arg := range expr.Arguments {
			argType := a.analyzeExpression(arg)
			if i < len(ft.ParameterTypes) && !argType.Equals(ft.ParameterTypes[i]) {
				a.reportError(expr.Token, "Fonksiyon çağrısında yanlış argüman tipi: %s bekleniyor, %s alındı", ft.ParameterTypes[i].String(), argType.String())
			}
		}

		// Dönüş tipini döndür
		return ft.ReturnType
	} else {
		a.reportError(expr.Token, "Çağrılabilir olmayan ifade")
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

func (a *Analyzer) analyzeArrayLiteral(expr *ast.ArrayLiteral) Type {
	// Dizi elemanlarını analiz et
	if len(expr.Elements) > 0 {
		// İlk elemanın tipini al
		elemType := a.analyzeExpression(expr.Elements[0])

		// Diğer elemanların tiplerini kontrol et
		for i := 1; i < len(expr.Elements); i++ {
			otherElemType := a.analyzeExpression(expr.Elements[i])
			if !elemType.Equals(otherElemType) {
				a.reportError(expr.Token, "Dizi elemanları aynı tipte olmalıdır")
				break
			}
		}

		// Dizi tipini döndür
		return &ArrayType{ElementType: elemType}
	}

	// Boş dizi
	return &ArrayType{ElementType: &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}}
}

func (a *Analyzer) analyzeIndexExpression(expr *ast.IndexExpression) Type {
	// Sol tarafı analiz et
	leftType := a.analyzeExpression(expr.Left)

	// İndeksi analiz et
	indexType := a.analyzeExpression(expr.Index)

	// İndeks tipini kontrol et
	if basicType, ok := indexType.(*BasicType); !ok || basicType.Kind != INTEGER_TYPE {
		a.reportError(expr.Token, "İndeks ifadesi tamsayı tipinde olmalıdır")
	}

	// Sol taraf tipini kontrol et
	if arrayType, ok := leftType.(*ArrayType); ok {
		// Dizi elemanı tipini döndür
		return arrayType.ElementType
	} else if mapType, ok := leftType.(*MapType); ok {
		// Map değeri tipini döndür
		return mapType.ValueType
	} else {
		a.reportError(expr.Token, "İndekslenebilir olmayan ifade")
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

func (a *Analyzer) analyzeHashLiteral(expr *ast.HashLiteral) Type {
	// Hash elemanlarını analiz et
	if len(expr.Pairs) > 0 {
		// İlk anahtar ve değer tiplerini al
		var keyType Type
		var valueType Type

		for k, v := range expr.Pairs {
			keyType = a.analyzeExpression(k)
			valueType = a.analyzeExpression(v)
			break
		}

		// Diğer anahtar ve değer tiplerini kontrol et
		for k, v := range expr.Pairs {
			otherKeyType := a.analyzeExpression(k)
			otherValueType := a.analyzeExpression(v)

			if !keyType.Equals(otherKeyType) {
				a.reportError(expr.Token, "Hash anahtarları aynı tipte olmalıdır")
			}

			if !valueType.Equals(otherValueType) {
				a.reportError(expr.Token, "Hash değerleri aynı tipte olmalıdır")
			}
		}

		// Map tipini döndür
		return &MapType{KeyType: keyType, ValueType: valueType}
	}

	// Boş hash
	return &MapType{
		KeyType:   &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE},
		ValueType: &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE},
	}
}

func (a *Analyzer) analyzeMemberExpression(expr *ast.MemberExpression) Type {
	// Member adını al
	memberIdent, ok := expr.Member.(*ast.Identifier)
	if !ok {
		a.reportError(expr.Token, "Üye erişimi için tanımlayıcı bekleniyor")
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
	memberName := memberIdent.Value

	// Package erişimi kontrolü
	if objectIdent, ok := expr.Object.(*ast.Identifier); ok {
		if packageSymbol := a.currentScope.Resolve(objectIdent.Value); packageSymbol != nil && packageSymbol.Type == PACKAGE_TYPE {
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
			a.reportError(expr.Token, "Package %s'de %s fonksiyonu bulunamadı", objectIdent.Value, memberName)
			return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
		}
	}

	// Nesneyi analiz et
	objectType := a.analyzeExpression(expr.Object)

	// Nesne tipini kontrol et
	if classType, ok := objectType.(*ClassType); ok {
		// Üye tipini bul
		if fieldType, ok := classType.Fields[memberName]; ok {
			return fieldType
		} else if methodType, ok := classType.Methods[memberName]; ok {
			return methodType
		} else {
			a.reportError(expr.Token, "Sınıfta tanımlanmamış üye: %s", memberName)
			return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
		}
	} else {
		a.reportError(expr.Token, "Üye erişimi için sınıf tipinde nesne veya package bekleniyor")
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

func (a *Analyzer) analyzeNewExpression(expr *ast.NewExpression) Type {
	// Sınıfı analiz et
	classType := a.analyzeExpression(expr.Class)

	// Sınıf tipini kontrol et
	if ct, ok := classType.(*ClassType); ok {
		// Argümanları analiz et
		for _, arg := range expr.Arguments {
			a.analyzeExpression(arg)
		}

		// Sınıf tipini döndür
		return ct
	} else {
		a.reportError(expr.Token, "new operatörü için sınıf tipinde ifade bekleniyor")
		return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
	}
}

func (a *Analyzer) analyzeTemplateExpression(expr *ast.TemplateExpression) Type {
	// Şablon parametrelerini tanımla
	templateScope := NewScope(a.currentScope)
	prevScope := a.currentScope
	a.currentScope = templateScope

	for _, param := range expr.Parameters {
		a.currentScope.Define(param.Value, TEMPLATE_TYPE, param.Token)
	}

	// Şablon gövdesini analiz et
	bodyType := a.analyzeExpression(expr.Body)

	// Önceki kapsama geri dön
	a.currentScope = prevScope

	// Şablon tipini döndür
	return &TemplateType{
		Name:       "template",
		Parameters: make([]string, len(expr.Parameters)),
		BaseType:   bodyType,
	}
}

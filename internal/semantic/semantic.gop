package semantic

// import (
// 	"c:\Users\tpoyr\OneDrive\Masaüstü\go+\internal\ast" // AST paketi tamamlandığında eklenecek
// 	"c:\Users\tpoyr\OneDrive\Masaüstü\go+\internal\token" // Token paketi veya türleri tanımlandığında eklenecek
// )

// Analyzer, AST üzerinde tip kontrolü, isim çözümlemesi ve diğer anlamsal doğrulamaları yapar.
// TODO: Semantic Analyzer struct'ını ve analiz metotlarını implemente et.
type Analyzer struct {
	errors []string
	// TODO: Sembol tablosu (symbol table) ve kapsam (scope) yönetimi için yapılar eklenecek.
}

// New, yeni bir Analyzer oluşturur.
func New() *Analyzer {
	return &Analyzer{errors: []string{}}
}

// Errors, analiz sırasında karşılaşılan hataları döndürür.
func (a *Analyzer) Errors() []string {
	return a.errors
}

// AnalyzeProgram, programın AST'sini analiz eder.
// TODO: Bu ana analiz fonksiyonudur ve tüm AST düğümlerini gezerek anlamsal kontroller yapmalıdır.
// func (a *Analyzer) AnalyzeProgram(node ast.Node) { // ast.Node tanımlandıktan sonra
// 	// switch node := node.(type) {
// 	// case *ast.Program:
// 	// 	for _, stmt := range node.Statements {
// 	// 		a.AnalyzeProgram(stmt)
// 	// 	}
// 	// case *ast.LetStatement:
// 	// 	// Tip kontrolü, değişken tanımlama vb.
// 	// 	break
// 	// // Diğer AST düğüm türleri için case'ler eklenecek
// 	// }
// }

// TODO: Tip kontrolü, kapsam yönetimi, isim çözümlemesi gibi yardımcı fonksiyonlar eklenecek.
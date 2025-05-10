package irgen

// import (
// 	"c:\Users\tpoyr\OneDrive\Masaüstü\go+\internal\ast" // AST paketi tamamlandığında eklenecek
// 	// "llvm.org/llvm/bindings/go/llvm" // LLVM kütüphanesi entegre edildiğinde eklenecek
// )

// IRGenerator, AST'yi LLVM IR'ına veya benzer bir ara koda dönüştürür.
// TODO: IRGenerator struct'ını ve IR üretme metotlarını implemente et.
type IRGenerator struct {
	errors []string
	// TODO: LLVM modülü, builder, context gibi LLVM özgü yapılar eklenecek.
	// module  llvm.Module
	// builder llvm.Builder
	// context llvm.Context
}

// New, yeni bir IRGenerator oluşturur.
func New() *IRGenerator {
	return &IRGenerator{errors: []string{}}
}

// Errors, IR üretimi sırasında karşılaşılan hataları döndürür.
func (g *IRGenerator) Errors() []string {
	return g.errors
}

// GenerateProgram, programın AST'sinden IR üretir.
// TODO: Bu ana IR üretme fonksiyonudur ve tüm AST düğümlerini gezerek IR kodunu oluşturmalıdır.
// func (g *IRGenerator) GenerateProgram(node ast.Node) llvm.Module { // ast.Node ve llvm.Module tanımlandıktan sonra
// 	// // LLVM context ve module oluşturma
// 	// g.context = llvm.NewContext()
// 	// defer g.context.Dispose()
// 	// g.module = g.context.NewModule("goplus_module")
// 	// g.builder = g.context.NewBuilder()
// 	// defer g.builder.Dispose()

// 	// // AST düğümlerini gezerek IR üretme
// 	// g.generate(node)

// 	// return g.module
// }

// // generate, belirli bir AST düğümü için IR üretir.
// func (g *IRGenerator) generate(node ast.Node) llvm.Value {
// 	// switch node := node.(type) {
// 	// case *ast.Program:
// 	// 	for _, stmt := range node.Statements {
// 	// 		g.generate(stmt)
// 	// 	}
// 	// // Diğer AST düğüm türleri için case'ler eklenecek
// 	// default:
// 	// 	// Bilinmeyen düğüm tipi için hata veya varsayılan işlem
// 	// 	return llvm.Value{}
// 	// }
// 	// return llvm.Value{} // Placeholder
// }

// TODO: Fonksiyonlar, ifadeler, değişkenler vb. için IR üretme yardımcı fonksiyonları eklenecek.
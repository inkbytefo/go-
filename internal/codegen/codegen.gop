package codegen

// import (
// 	"llvm.org/llvm/bindings/go/llvm" // LLVM kütüphanesi entegre edildiğinde eklenecek
// )

// CodeGenerator, optimize edilmiş LLVM IR'ını hedef platform için makine koduna dönüştürür.
// TODO: CodeGenerator struct'ını ve makine kodu üretme metotlarını implemente et.
type CodeGenerator struct {
	errors []string
	// TODO: Hedef makine (target machine), veri yerleşimi (data layout) gibi
	// LLVM özgü yapılar eklenecek.
	// targetMachine llvm.TargetMachine
}

// New, yeni bir CodeGenerator oluşturur.
func New() *CodeGenerator {
	return &CodeGenerator{errors: []string{}}
}

// Errors, kod üretimi sırasında karşılaşılan hataları döndürür.
func (cg *CodeGenerator) Errors() []string {
	return cg.errors
}

// GenerateMachineCode, LLVM modülünden hedef platform için makine kodu üretir.
// TODO: Bu fonksiyon, LLVM'in kod üretme yeteneklerini kullanarak makine kodunu oluşturmalı
// ve bir çalıştırılabilir dosya veya nesne dosyası olarak kaydetmelidir.
// func (cg *CodeGenerator) GenerateMachineCode(module llvm.Module, outputPath string) error {
// 	// // Hedef üçlüsünü (target triple) al veya ayarla
// 	// triple := llvm.DefaultTargetTriple()
// 	// target, err := llvm.GetTargetFromTriple(triple)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// // Hedef makineyi oluştur
// 	// cg.targetMachine = target.CreateTargetMachine(triple, "generic", "", llvm.CodeGenLevelDefault, llvm.RelocDefault, llvm.CodeModelDefault)
// 	// defer cg.targetMachine.Dispose()

// 	// // Modül için veri yerleşimini ve hedef üçlüsünü ayarla
// 	// module.SetDataLayout(cg.targetMachine.CreateTargetDataLayout().String())
// 	// module.SetTarget(triple)

// 	// // Nesne dosyasını diske yaz
// 	// // err = cg.targetMachine.EmitToFile(module, outputPath, llvm.ObjectFile)
// 	// // if err != nil {
// 	// // 	return err
// 	// // }

// 	// return nil
// }

// TODO: Farklı çıktı formatları (örn: assembly, object file, executable)
// veya farklı hedef platformlar için yardımcı fonksiyonlar eklenebilir.
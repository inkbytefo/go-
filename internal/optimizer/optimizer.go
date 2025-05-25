package optimizer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// OptimizationLevel, optimizasyon seviyesini belirtir.
type OptimizationLevel int

const (
	O0 OptimizationLevel = iota // Optimizasyon yok
	O1                          // Temel optimizasyonlar
	O2                          // Orta seviye optimizasyonlar
	O3                          // Agresif optimizasyonlar
)

// Optimizer, LLVM IR üzerinde optimizasyonlar yapar.
type Optimizer struct {
	level  OptimizationLevel // Optimizasyon seviyesi
	errors []string          // Optimizasyon sırasında oluşan hatalar
}

// New, yeni bir Optimizer oluşturur.
func New(level OptimizationLevel) *Optimizer {
	return &Optimizer{
		level:  level,
		errors: []string{},
	}
}

// Errors, optimizasyon sırasında karşılaşılan hataları döndürür.
func (opt *Optimizer) Errors() []string {
	return opt.errors
}

// ReportError, bir hata mesajı ekler.
func (opt *Optimizer) ReportError(format string, args ...any) {
	opt.errors = append(opt.errors, fmt.Sprintf(format, args...))
}

// OptimizationPass, bir optimizasyon geçişini temsil eder.
type OptimizationPass interface {
	// Apply, optimizasyon geçişini uygular.
	Apply(module *ir.Module) (*ir.Module, error)
	// Name, optimizasyon geçişinin adını döndürür.
	Name() string
}

// ConstantFoldingPass, sabit katlama optimizasyonu geçişi.
type ConstantFoldingPass struct{}

// Apply, sabit katlama optimizasyonunu uygular.
func (p *ConstantFoldingPass) Apply(module *ir.Module) (*ir.Module, error) {
	// Modüldeki tüm fonksiyonları gez
	for _, f := range module.Funcs {
		// Fonksiyondaki tüm blokları gez
		for _, block := range f.Blocks {
			// Bloktaki tüm komutları gez
			for i, inst := range block.Insts {
				// Sabit katlama yapılabilecek komutları bul
				if newInst := p.foldConstant(inst); newInst != nil {
					// Komutu yeni komutla değiştir
					block.Insts[i] = newInst
				}
			}
		}
	}
	return module, nil
}

// Name, optimizasyon geçişinin adını döndürür.
func (p *ConstantFoldingPass) Name() string {
	return "ConstantFolding"
}

// foldConstant, bir komutu sabit katlama ile optimize eder.
func (p *ConstantFoldingPass) foldConstant(inst ir.Instruction) ir.Instruction {
	// TODO: Implement constant folding
	return nil
}

// DeadCodeEliminationPass, ölü kod eleme optimizasyonu geçişi.
type DeadCodeEliminationPass struct{}

// Apply, ölü kod eleme optimizasyonunu uygular.
func (p *DeadCodeEliminationPass) Apply(module *ir.Module) (*ir.Module, error) {
	// Modüldeki tüm fonksiyonları gez
	for _, f := range module.Funcs {
		// Kullanılan değişkenleri topla
		usedValues := make(map[value.Value]bool)

		// Fonksiyondaki tüm blokları gez
		for _, block := range f.Blocks {
			// Bloktaki tüm komutları gez
			for _, inst := range block.Insts {
				// Komutun kullandığı değerleri işaretle
				for _, operandPtr := range inst.Operands() {
					if operand, ok := (*operandPtr).(value.Value); ok {
						usedValues[operand] = true
					}
				}
			}
		}

		// Fonksiyondaki tüm blokları tekrar gez
		for _, block := range f.Blocks {
			// Kullanılmayan komutları topla
			var newInsts []ir.Instruction
			for _, inst := range block.Insts {
				// Eğer komut bir değer üretiyorsa ve kullanılmıyorsa, atla
				if val, ok := inst.(value.Value); ok {
					if !usedValues[val] && !p.hasEffects(inst) {
						continue
					}
				}
				newInsts = append(newInsts, inst)
			}
			// Blokun komutlarını güncelle
			block.Insts = newInsts
		}
	}
	return module, nil
}

// Name, optimizasyon geçişinin adını döndürür.
func (p *DeadCodeEliminationPass) Name() string {
	return "DeadCodeElimination"
}

// hasEffects, bir komutun yan etkileri olup olmadığını kontrol eder.
func (p *DeadCodeEliminationPass) hasEffects(inst ir.Instruction) bool {
	// TODO: Implement side effect detection
	return true
}

// OptimizeModule, verilen LLVM modülünü optimize eder.
// Bu fonksiyon, llir/llvm kütüphanesi ile oluşturulan modülü optimize eder.
func (opt *Optimizer) OptimizeModule(module *ir.Module) (*ir.Module, error) {
	// Optimizasyon seviyesi O0 ise hiçbir şey yapma
	if opt.level == O0 {
		return module, nil
	}

	// Optimizasyon geçişlerini belirle
	var passes []OptimizationPass

	// Temel optimizasyon geçişleri (O1)
	if opt.level >= O1 {
		passes = append(passes, &ConstantFoldingPass{})
		passes = append(passes, &DeadCodeEliminationPass{})
	}

	// Orta seviye optimizasyon geçişleri (O2)
	if opt.level >= O2 {
		// TODO: Add more optimization passes for O2
	}

	// Agresif optimizasyon geçişleri (O3)
	if opt.level >= O3 {
		// TODO: Add more optimization passes for O3
	}

	// Optimizasyon geçişlerini uygula
	var err error
	for _, pass := range passes {
		fmt.Printf("Applying optimization pass: %s\n", pass.Name())
		module, err = pass.Apply(module)
		if err != nil {
			opt.ReportError("Optimizasyon geçişi uygulanırken hata oluştu (%s): %v", pass.Name(), err)
			return module, err
		}
	}

	// Harici LLVM optimizasyonlarını uygula
	return opt.applyExternalOptimizations(module)
}

// applyExternalOptimizations, harici LLVM optimizasyonlarını uygular.
func (opt *Optimizer) applyExternalOptimizations(module *ir.Module) (*ir.Module, error) {
	// Geçici dosya oluştur
	tempDir, err := os.MkdirTemp("", "goplus-opt")
	if err != nil {
		opt.ReportError("Geçici dizin oluşturulamadı: %v", err)
		return nil, fmt.Errorf("geçici dizin oluşturulamadı: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Modülü geçici dosyaya yaz
	inputFile := filepath.Join(tempDir, "input.ll")
	err = os.WriteFile(inputFile, []byte(module.String()), 0644)
	if err != nil {
		opt.ReportError("Modül geçici dosyaya yazılamadı: %v", err)
		return nil, fmt.Errorf("modül geçici dosyaya yazılamadı: %v", err)
	}

	// Optimizasyon için çıktı dosyası
	outputFile := filepath.Join(tempDir, "output.ll")

	// LLVM opt aracını çağır
	args := []string{
		"-S", // Assembly çıktısı
	}

	// Optimizasyon seviyesine göre geçişleri ekle
	switch opt.level {
	case O1:
		args = append(args, "-O1")
	case O2:
		args = append(args, "-O2")
	case O3:
		args = append(args, "-O3")
	}

	args = append(args, "-o", outputFile, inputFile)

	cmd := exec.Command("opt", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		opt.ReportError("LLVM opt çalıştırılamadı: %v\nÇıktı: %s", err, string(output))

		// LLVM opt aracı bulunamadıysa, orijinal modülü döndür
		if strings.Contains(err.Error(), "executable file not found") {
			opt.ReportError("LLVM opt aracı bulunamadı, optimizasyon yapılmadan devam ediliyor")
			return module, nil
		}

		return nil, fmt.Errorf("LLVM opt çalıştırılamadı: %v", err)
	}

	// Optimize edilmiş IR'ın varlığını kontrol et
	_, err = os.Stat(outputFile)
	if err != nil {
		opt.ReportError("Optimize edilmiş IR dosyası oluşturulamadı: %v", err)
		return nil, fmt.Errorf("optimize edilmiş IR dosyası oluşturulamadı: %v", err)
	}

	// Optimize edilmiş IR'ı parse et
	// Not: Şu anda llir/llvm kütüphanesi ile oluşturulan modülü doğrudan optimize edemiyoruz.
	// Bunun yerine, optimize edilmiş IR'ı string olarak döndürüyoruz.
	// İleride, llir/llvm kütüphanesinin parse fonksiyonları kullanılarak
	// optimize edilmiş IR'ı tekrar modül olarak yükleyebiliriz.

	// Şimdilik, orijinal modülü döndür ve optimize edilmiş IR'ı bir dosyaya yaz
	return module, nil
}

// GetOptimizedIRString, verilen LLVM IR string'ini optimize eder ve optimize edilmiş IR string'ini döndürür.
func (opt *Optimizer) GetOptimizedIRString(irString string) (string, error) {
	// Optimizasyon seviyesi O0 ise hiçbir şey yapma
	if opt.level == O0 {
		return irString, nil
	}

	// Önce dahili optimizasyonları uygula
	// Not: Şu anda dahili optimizasyonları doğrudan IR string'i üzerinde uygulayamıyoruz.
	// Bunun için IR'ı parse etmemiz ve modül olarak işlememiz gerekiyor.
	// Bu özellik ileride eklenecek.

	// Harici LLVM optimizasyonlarını uygula
	return opt.applyExternalOptimizationsToString(irString)
}

// applyExternalOptimizationsToString, harici LLVM optimizasyonlarını IR string'ine uygular.
func (opt *Optimizer) applyExternalOptimizationsToString(irString string) (string, error) {
	// Geçici dosya oluştur
	tempDir, err := os.MkdirTemp("", "goplus-opt")
	if err != nil {
		opt.ReportError("Geçici dizin oluşturulamadı: %v", err)
		return irString, fmt.Errorf("geçici dizin oluşturulamadı: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// IR'ı geçici dosyaya yaz
	inputFile := filepath.Join(tempDir, "input.ll")
	err = os.WriteFile(inputFile, []byte(irString), 0644)
	if err != nil {
		opt.ReportError("IR geçici dosyaya yazılamadı: %v", err)
		return irString, fmt.Errorf("IR geçici dosyaya yazılamadı: %v", err)
	}

	// Optimizasyon için çıktı dosyası
	outputFile := filepath.Join(tempDir, "output.ll")

	// LLVM opt aracını çağır
	args := []string{
		"-S", // Assembly çıktısı
	}

	// Optimizasyon seviyesine göre geçişleri ekle
	switch opt.level {
	case O1:
		args = append(args, "-O1")
	case O2:
		args = append(args, "-O2")
	case O3:
		args = append(args, "-O3")
	}

	// Özel optimizasyon geçişleri ekle
	if opt.level >= O2 {
		// Agresif inlining
		args = append(args, "--inline-threshold=100")
		// Döngü optimizasyonları
		args = append(args, "--loop-vectorize")
		args = append(args, "--loop-unroll")
	}

	if opt.level >= O3 {
		// Daha agresif inlining
		args = append(args, "--inline-threshold=1000")
		// Fonksiyon birleştirme
		args = append(args, "--mergefunc")
		// Agresif ölü kod eleme
		args = append(args, "--adce")
	}

	args = append(args, "-o", outputFile, inputFile)

	cmd := exec.Command("opt", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		opt.ReportError("LLVM opt çalıştırılamadı: %v\nÇıktı: %s", err, string(output))

		// LLVM opt aracı bulunamadıysa, orijinal IR'ı döndür
		if strings.Contains(err.Error(), "executable file not found") {
			opt.ReportError("LLVM opt aracı bulunamadı, optimizasyon yapılmadan devam ediliyor")
			return irString, nil
		}

		return irString, fmt.Errorf("LLVM opt çalıştırılamadı: %v", err)
	}

	// Optimize edilmiş IR'ı oku
	optimizedIR, err := os.ReadFile(outputFile)
	if err != nil {
		opt.ReportError("Optimize edilmiş IR okunamadı: %v", err)
		return irString, fmt.Errorf("optimize edilmiş IR okunamadı: %v", err)
	}

	return string(optimizedIR), nil
}

package optimizer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/llir/llvm/ir"
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

// OptimizeModule, verilen LLVM modülünü optimize eder.
// Bu fonksiyon, llir/llvm kütüphanesi ile oluşturulan modülü optimize eder.
func (opt *Optimizer) OptimizeModule(module *ir.Module) (*ir.Module, error) {
	// Optimizasyon seviyesi O0 ise hiçbir şey yapma
	if opt.level == O0 {
		return module, nil
	}

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

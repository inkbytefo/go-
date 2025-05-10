package codegen

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// OutputFormat, çıktı formatını belirtir.
type OutputFormat int

const (
	Assembly   OutputFormat = iota // Assembly çıktısı (.s)
	Object                         // Nesne dosyası (.o)
	Executable                     // Çalıştırılabilir dosya
)

// TargetArch, hedef mimarisi belirtir.
type TargetArch string

const (
	X86_64 TargetArch = "x86_64"  // x86_64 mimarisi
	ARM64  TargetArch = "aarch64" // ARM64 mimarisi
	RISCV  TargetArch = "riscv64" // RISC-V mimarisi
)

// TargetOS, hedef işletim sistemini belirtir.
type TargetOS string

const (
	Linux   TargetOS = "linux"   // Linux işletim sistemi
	Windows TargetOS = "windows" // Windows işletim sistemi
	MacOS   TargetOS = "darwin"  // macOS işletim sistemi
)

// CodeGenerator, optimize edilmiş LLVM IR'ını hedef platform için makine koduna dönüştürür.
type CodeGenerator struct {
	errors     []string     // Kod üretimi sırasında oluşan hatalar
	targetArch TargetArch   // Hedef mimari
	targetOS   TargetOS     // Hedef işletim sistemi
	format     OutputFormat // Çıktı formatı
}

// New, yeni bir CodeGenerator oluşturur.
func New(targetArch TargetArch, targetOS TargetOS, format OutputFormat) *CodeGenerator {
	return &CodeGenerator{
		errors:     []string{},
		targetArch: targetArch,
		targetOS:   targetOS,
		format:     format,
	}
}

// NewWithCurrentPlatform, mevcut platform için bir CodeGenerator oluşturur.
func NewWithCurrentPlatform(format OutputFormat) *CodeGenerator {
	var targetArch TargetArch
	var targetOS TargetOS

	// Mevcut mimariye göre hedef mimarisi belirle
	switch runtime.GOARCH {
	case "amd64":
		targetArch = X86_64
	case "arm64":
		targetArch = ARM64
	default:
		targetArch = X86_64 // Varsayılan olarak x86_64
	}

	// Mevcut işletim sistemine göre hedef işletim sistemi belirle
	switch runtime.GOOS {
	case "linux":
		targetOS = Linux
	case "windows":
		targetOS = Windows
	case "darwin":
		targetOS = MacOS
	default:
		targetOS = Linux // Varsayılan olarak Linux
	}

	return New(targetArch, targetOS, format)
}

// Errors, kod üretimi sırasında karşılaşılan hataları döndürür.
func (cg *CodeGenerator) Errors() []string {
	return cg.errors
}

// ReportError, bir hata mesajı ekler.
func (cg *CodeGenerator) ReportError(format string, args ...any) {
	cg.errors = append(cg.errors, fmt.Sprintf(format, args...))
}

// GetTargetTriple, hedef üçlüsünü (target triple) döndürür.
func (cg *CodeGenerator) GetTargetTriple() string {
	return fmt.Sprintf("%s-%s", cg.targetArch, cg.targetOS)
}

// GenerateMachineCode, LLVM IR'ından hedef platform için makine kodu üretir.
func (cg *CodeGenerator) GenerateMachineCode(irString, outputPath string) error {
	// Geçici dosya oluştur
	tempDir, err := os.MkdirTemp("", "goplus-codegen")
	if err != nil {
		cg.ReportError("Geçici dizin oluşturulamadı: %v", err)
		return fmt.Errorf("geçici dizin oluşturulamadı: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// IR'ı geçici dosyaya yaz
	inputFile := filepath.Join(tempDir, "input.ll")
	err = os.WriteFile(inputFile, []byte(irString), 0644)
	if err != nil {
		cg.ReportError("IR geçici dosyaya yazılamadı: %v", err)
		return fmt.Errorf("IR geçici dosyaya yazılamadı: %v", err)
	}

	// Hedef üçlüsünü belirle
	targetTriple := cg.GetTargetTriple()

	// Çıktı formatına göre işlem yap
	switch cg.format {
	case Assembly:
		return cg.generateAssembly(inputFile, outputPath, targetTriple)
	case Object:
		return cg.generateObjectFile(inputFile, outputPath, targetTriple)
	case Executable:
		return cg.generateExecutable(inputFile, outputPath, targetTriple)
	default:
		cg.ReportError("Desteklenmeyen çıktı formatı")
		return fmt.Errorf("desteklenmeyen çıktı formatı")
	}
}

// generateAssembly, LLVM IR'ından assembly kodu üretir.
func (cg *CodeGenerator) generateAssembly(inputFile, outputPath, targetTriple string) error {
	// LLVM llc aracını çağır
	args := []string{
		"-march=" + string(cg.targetArch),
		"-mtriple=" + targetTriple,
		"-filetype=asm",
		"-o", outputPath,
		inputFile,
	}

	cmd := exec.Command("llc", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		cg.ReportError("LLVM llc çalıştırılamadı: %v\nÇıktı: %s", err, string(output))

		// LLVM llc aracı bulunamadıysa
		if strings.Contains(err.Error(), "executable file not found") {
			cg.ReportError("LLVM llc aracı bulunamadı")
		}

		return fmt.Errorf("LLVM llc çalıştırılamadı: %v", err)
	}

	return nil
}

// generateObjectFile, LLVM IR'ından nesne dosyası üretir.
func (cg *CodeGenerator) generateObjectFile(inputFile, outputPath, targetTriple string) error {
	// LLVM llc aracını çağır
	args := []string{
		"-march=" + string(cg.targetArch),
		"-mtriple=" + targetTriple,
		"-filetype=obj",
		"-o", outputPath,
		inputFile,
	}

	cmd := exec.Command("llc", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		cg.ReportError("LLVM llc çalıştırılamadı: %v\nÇıktı: %s", err, string(output))

		// LLVM llc aracı bulunamadıysa
		if strings.Contains(err.Error(), "executable file not found") {
			cg.ReportError("LLVM llc aracı bulunamadı")
		}

		return fmt.Errorf("LLVM llc çalıştırılamadı: %v", err)
	}

	return nil
}

// generateExecutable, LLVM IR'ından çalıştırılabilir dosya üretir.
func (cg *CodeGenerator) generateExecutable(inputFile, outputPath, targetTriple string) error {
	// Önce nesne dosyası oluştur
	objFile := filepath.Join(filepath.Dir(inputFile), "output.o")
	err := cg.generateObjectFile(inputFile, objFile, targetTriple)
	if err != nil {
		return err
	}

	// İşletim sistemine göre bağlayıcı (linker) seç
	var linker string
	var args []string

	switch cg.targetOS {
	case Windows:
		linker = "gcc" // MinGW veya benzeri bir GCC kurulumu gerektirir
		args = []string{objFile, "-o", outputPath}
	case Linux, MacOS:
		linker = "gcc"
		args = []string{objFile, "-o", outputPath}
	default:
		cg.ReportError("Desteklenmeyen işletim sistemi: %s", cg.targetOS)
		return fmt.Errorf("desteklenmeyen işletim sistemi: %s", cg.targetOS)
	}

	cmd := exec.Command(linker, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		cg.ReportError("Bağlayıcı çalıştırılamadı: %v\nÇıktı: %s", err, string(output))

		// Bağlayıcı bulunamadıysa
		if strings.Contains(err.Error(), "executable file not found") {
			cg.ReportError("Bağlayıcı (%s) bulunamadı", linker)
		}

		return fmt.Errorf("bağlayıcı çalıştırılamadı: %v", err)
	}

	return nil
}

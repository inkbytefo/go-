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
	errors            []string     // Kod üretimi sırasında oluşan hatalar
	targetArch        TargetArch   // Hedef mimari
	targetOS          TargetOS     // Hedef işletim sistemi
	format            OutputFormat // Çıktı formatı
	optimizationLevel int          // Optimizasyon seviyesi (0-3)
	generateDebug     bool         // Debug bilgisi üret
}

// New, yeni bir CodeGenerator oluşturur.
func New(targetArch TargetArch, targetOS TargetOS, format OutputFormat) *CodeGenerator {
	return &CodeGenerator{
		errors:            []string{},
		targetArch:        targetArch,
		targetOS:          targetOS,
		format:            format,
		optimizationLevel: 2, // Varsayılan optimizasyon seviyesi
		generateDebug:     false,
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

// SetOptimizationLevel, optimizasyon seviyesini ayarlar (0-3).
func (cg *CodeGenerator) SetOptimizationLevel(level int) {
	if level < 0 {
		level = 0
	} else if level > 3 {
		level = 3
	}
	cg.optimizationLevel = level
}

// SetDebugInfo, debug bilgisi üretimini etkinleştirir/devre dışı bırakır.
func (cg *CodeGenerator) SetDebugInfo(enable bool) {
	cg.generateDebug = enable
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
	// LLC path'ini belirle
	llcPath := "llc"
	if cg.targetOS == Windows {
		windowsLlcPath := "C:\\Program Files\\LLVM\\bin\\llc.exe"
		if _, err := os.Stat(windowsLlcPath); err == nil {
			llcPath = windowsLlcPath
		}
	}

	// LLVM llc aracını çağır
	args := []string{
		"-march=" + string(cg.targetArch),
		"-mtriple=" + targetTriple,
		"-filetype=obj",
		"-o", outputPath,
		inputFile,
	}

	// Debug: LLC komutunu yazdır
	// fmt.Printf("Debug: LLC komutu: %s %v\n", llcPath, args)

	cmd := exec.Command(llcPath, args...)
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
	// Clang kullanarak doğrudan LLVM IR'dan executable oluştur
	// Bu LLC'ye ihtiyaç duymaz ve daha basittir

	// Windows'ta .exe uzantısı ekle
	if cg.targetOS == Windows && !strings.HasSuffix(outputPath, ".exe") {
		outputPath += ".exe"
	}

	// Clang path'ini belirle
	clangPath := "clang"
	if cg.targetOS == Windows {
		windowsClangPath := "C:\\Program Files\\LLVM\\bin\\clang.exe"
		if _, err := os.Stat(windowsClangPath); err == nil {
			clangPath = windowsClangPath
		}
	}

	var args []string

	// Temel clang argumentları
	args = append(args, inputFile)
	args = append(args, "-o", outputPath)

	// Target triple belirt (eğer varsa)
	if targetTriple != "" {
		args = append(args, "-target", targetTriple)
	}

	// İşletim sistemine göre özel ayarlar
	switch cg.targetOS {
	case Windows:
		// Windows için C runtime linking - daha basit yaklaşım
		// MSVC runtime library'yi link et
		args = append(args, "-lmsvcrt")
		// Legacy stdio functions için
		args = append(args, "-llegacy_stdio_definitions")
	case Linux:
		// Linux için C runtime library
		args = append(args, "-lc")
	case MacOS:
		// macOS için system libraries
		args = append(args, "-lSystem")
	default:
		return fmt.Errorf("desteklenmeyen işletim sistemi: %s", cg.targetOS)
	}

	// Optimizasyon seviyesi ekle
	switch cg.optimizationLevel {
	case 0:
		args = append(args, "-O0")
	case 1:
		args = append(args, "-O1")
	case 2:
		args = append(args, "-O2")
	case 3:
		args = append(args, "-O3")
	}

	// Debug bilgisi ekle (gerekirse)
	if cg.generateDebug {
		args = append(args, "-g")
	}

	// Debug: Clang komutunu yazdır
	// fmt.Printf("Debug: Clang komutu: %s %v\n", clangPath, args)

	cmd := exec.Command(clangPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		cg.ReportError("Clang çalıştırılamadı: %v\nÇıktı: %s", err, string(output))

		// Clang bulunamadıysa
		if strings.Contains(err.Error(), "executable file not found") {
			cg.ReportError("Clang bulunamadı. LLVM kurulumunu kontrol edin.")
			cg.ReportError("Kurulum rehberi: docs/llvm-setup.md")
		}

		return fmt.Errorf("clang çalıştırılamadı: %v", err)
	}

	return nil
}

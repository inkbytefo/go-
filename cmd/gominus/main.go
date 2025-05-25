package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/inkbytefo/go-minus/internal/codegen"
	"github.com/inkbytefo/go-minus/internal/irgen"
	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/optimizer"
	"github.com/inkbytefo/go-minus/internal/parser"
	"github.com/inkbytefo/go-minus/internal/semantic"
)

func main() {
	// Komut satırı bayraklarını tanımla
	var (
		optimizationLevel = flag.Int("O", 0, "Optimizasyon seviyesi (0-3)")
		outputFormat      = flag.String("output-format", "ll", "Çıktı formatı (ll, s, o, exe)")
		targetArch        = flag.String("target-arch", "", "Hedef mimari (x86_64, aarch64, riscv64)")
		targetOS          = flag.String("target-os", "", "Hedef işletim sistemi (linux, windows, darwin)")
		outputFile        = flag.String("o", "", "Çıktı dosyası")
		showHelp          = flag.Bool("help", false, "Yardım mesajını göster")
		showVersion       = flag.Bool("version", false, "Sürüm bilgisini göster")
	)

	// Bayrakları ayrıştır
	flag.Parse()

	// Yardım mesajını göster
	if *showHelp {
		printHelp()
		os.Exit(0)
	}

	// Sürüm bilgisini göster
	if *showVersion {
		fmt.Println("GO-Minus Derleyicisi v0.1.0")
		os.Exit(0)
	}

	// Giriş dosyasını kontrol et
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Hata: Giriş dosyası belirtilmedi")
		printHelp()
		os.Exit(1)
	}

	// Dosyayı oku
	filename := args[0]
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Hata: %s dosyası okunamadı: %v\n", filename, err)
		os.Exit(1)
	}

	// Lexer oluştur
	l := lexer.New(string(input))

	// Parser oluştur
	p := parser.New(l)

	// Programı ayrıştır
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printErrors("Ayrıştırma hataları:", p.Errors())
		os.Exit(1)
	}

	// AST'yi yazdır (verbose mod için)
	// fmt.Println("AST:")
	// fmt.Println(program.String())

	// Semantik analiz
	analyzer := semantic.New()
	analyzer.Analyze(program)
	if len(analyzer.Errors()) != 0 {
		printErrors("Semantik analiz hataları:", analyzer.Errors())
		// Semantik analiz hatalarını göster ama devam et
		// os.Exit(1)
	} else {
		fmt.Println("Semantik analiz başarılı!")
	}

	// IR üretimi
	generator := irgen.NewWithAnalyzer(analyzer)
	ir, err := generator.GenerateProgram(program)
	if err != nil {
		fmt.Printf("IR üretimi sırasında hata oluştu: %v\n", err)
		os.Exit(1)
	}

	// Optimizasyon seviyesini belirle
	optLevel := optimizer.OptimizationLevel(*optimizationLevel)
	if optLevel < optimizer.O0 || optLevel > optimizer.O3 {
		fmt.Printf("Uyarı: Geçersiz optimizasyon seviyesi: %d, varsayılan olarak O0 kullanılıyor\n", *optimizationLevel)
		optLevel = optimizer.O0
	}

	// IR'ı optimize et
	opt := optimizer.New(optLevel)
	optimizedIR, err := opt.GetOptimizedIRString(ir)
	if err != nil {
		fmt.Printf("Uyarı: IR optimizasyonu sırasında hata oluştu: %v\n", err)
		fmt.Println("Optimizasyon yapılmadan devam ediliyor...")
		optimizedIR = ir
	}

	// Çıktı dosyasını belirle
	var outputPath string
	if *outputFile != "" {
		outputPath = *outputFile
	} else {
		baseName := filepath.Base(filename)
		baseName = baseName[:len(baseName)-len(filepath.Ext(baseName))]

		// Çıktı formatına göre uzantı belirle
		var ext string
		switch strings.ToLower(*outputFormat) {
		case "ll":
			ext = ".ll"
		case "s":
			ext = ".s"
		case "o":
			ext = ".o"
		case "exe":
			if strings.ToLower(*targetOS) == "windows" {
				ext = ".exe"
			} else {
				ext = ""
			}
		default:
			fmt.Printf("Uyarı: Geçersiz çıktı formatı: %s, varsayılan olarak 'll' kullanılıyor\n", *outputFormat)
			ext = ".ll"
			*outputFormat = "ll"
		}

		outputPath = baseName + ext
	}

	// Çıktı formatına göre işlem yap
	switch strings.ToLower(*outputFormat) {
	case "ll":
		// IR dosyasını kaydet
		err = os.WriteFile(outputPath, []byte(optimizedIR), 0644)
		if err != nil {
			fmt.Printf("IR dosyası kaydedilirken hata oluştu: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("IR dosyası başarıyla oluşturuldu: %s\n", outputPath)
	case "s", "o", "exe":
		// Hedef kod üretimi
		var format codegen.OutputFormat
		switch strings.ToLower(*outputFormat) {
		case "s":
			format = codegen.Assembly
		case "o":
			format = codegen.Object
		case "exe":
			format = codegen.Executable
		}

		// Hedef mimari ve işletim sistemini belirle
		var targetArchVal codegen.TargetArch
		var targetOSVal codegen.TargetOS

		if *targetArch != "" {
			switch strings.ToLower(*targetArch) {
			case "x86_64":
				targetArchVal = codegen.X86_64
			case "aarch64":
				targetArchVal = codegen.ARM64
			case "riscv64":
				targetArchVal = codegen.RISCV
			default:
				fmt.Printf("Uyarı: Geçersiz hedef mimari: %s, mevcut platform kullanılıyor\n", *targetArch)
				targetArchVal = ""
			}
		}

		if *targetOS != "" {
			switch strings.ToLower(*targetOS) {
			case "linux":
				targetOSVal = codegen.Linux
			case "windows":
				targetOSVal = codegen.Windows
			case "darwin":
				targetOSVal = codegen.MacOS
			default:
				fmt.Printf("Uyarı: Geçersiz hedef işletim sistemi: %s, mevcut platform kullanılıyor\n", *targetOS)
				targetOSVal = ""
			}
		}

		var cg *codegen.CodeGenerator
		if targetArchVal != "" && targetOSVal != "" {
			cg = codegen.New(targetArchVal, targetOSVal, format)
		} else {
			cg = codegen.NewWithCurrentPlatform(format)
		}

		// Optimizasyon seviyesini ayarla
		cg.SetOptimizationLevel(*optimizationLevel)

		err = cg.GenerateMachineCode(optimizedIR, outputPath)
		if err != nil {
			fmt.Printf("Makine kodu üretimi sırasında hata oluştu: %v\n", err)
			if len(cg.Errors()) > 0 {
				printErrors("Kod üretimi hataları:", cg.Errors())
			}
			os.Exit(1)
		}

		fmt.Printf("%s dosyası başarıyla oluşturuldu: %s\n", strings.ToUpper(*outputFormat), outputPath)
	default:
		fmt.Printf("Hata: Desteklenmeyen çıktı formatı: %s\n", *outputFormat)
		os.Exit(1)
	}
}

func printErrors(title string, errors []string) {
	fmt.Println(title)
	for _, msg := range errors {
		fmt.Println("\t" + msg)
	}
}

func printHelp() {
	fmt.Println("Kullanım: gominus [bayraklar] <dosya.gom>")
	fmt.Println("\nBayraklar:")
	flag.PrintDefaults()
	fmt.Println("\nÖrnekler:")
	fmt.Println("  gominus test.gom                    # LLVM IR üret (test.ll)")
	fmt.Println("  gominus -O2 test.gom                # Optimize edilmiş LLVM IR üret (test.ll)")
	fmt.Println("  gominus -output-format=s test.gom   # Assembly kodu üret (test.s)")
	fmt.Println("  gominus -output-format=o test.gom   # Nesne dosyası üret (test.o)")
	fmt.Println("  gominus -output-format=exe test.gom # Çalıştırılabilir dosya üret (test veya test.exe)")
	fmt.Println("  gominus -o output.exe -output-format=exe test.gom # Belirtilen isimle çalıştırılabilir dosya üret")
}

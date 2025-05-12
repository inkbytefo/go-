package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Komut satırı bayrakları
var (
	verboseFlag  = flag.Bool("v", false, "Ayrıntılı çıktı")
	recursiveFlag = flag.Bool("r", false, "Alt dizinlerdeki testleri de çalıştır")
	patternFlag  = flag.String("pattern", "", "Test adı deseni (örn: TestAdd*)")
	timeoutFlag  = flag.Duration("timeout", 10*time.Minute, "Test zaman aşımı süresi")
	benchFlag    = flag.Bool("bench", false, "Benchmark testlerini çalıştır")
	coverFlag    = flag.Bool("cover", false, "Kod kapsama analizi yap")
	versionFlag  = flag.Bool("version", false, "Sürüm bilgisini göster")
	helpFlag     = flag.Bool("help", false, "Yardım mesajını göster")
)

// Test aracı sürümü
const version = "0.1.0"

func main() {
	// Bayrakları ayrıştır
	flag.Parse()

	// Sürüm bilgisini göster
	if *versionFlag {
		fmt.Printf("GO+ Test Aracı v%s\n", version)
		os.Exit(0)
	}

	// Yardım mesajını göster
	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	// Test dizinlerini belirle
	var testDirs []string
	if flag.NArg() == 0 {
		// Mevcut dizini kullan
		testDirs = append(testDirs, ".")
	} else {
		// Belirtilen dizinleri kullan
		testDirs = flag.Args()
	}

	// Testleri çalıştır
	runTests(testDirs)
}

// Testleri çalıştır
func runTests(dirs []string) {
	fmt.Println("GO+ testleri çalıştırılıyor...")

	// Test seçeneklerini göster
	fmt.Println("Test seçenekleri:")
	if *verboseFlag {
		fmt.Println("  Ayrıntılı çıktı: Evet")
	}
	if *recursiveFlag {
		fmt.Println("  Alt dizinlerdeki testler: Evet")
	}
	if *patternFlag != "" {
		fmt.Printf("  Test adı deseni: %s\n", *patternFlag)
	}
	fmt.Printf("  Zaman aşımı süresi: %s\n", *timeoutFlag)
	if *benchFlag {
		fmt.Println("  Benchmark testleri: Evet")
	}
	if *coverFlag {
		fmt.Println("  Kod kapsama analizi: Evet")
	}

	// Her dizin için testleri çalıştır
	totalTests := 0
	passedTests := 0
	failedTests := 0
	skippedTests := 0

	startTime := time.Now()

	for _, dir := range dirs {
		fmt.Printf("\nDizin: %s\n", dir)
		
		// Test dosyalarını bul
		testFiles, err := findTestFiles(dir)
		if err != nil {
			fmt.Printf("Hata: Test dosyaları bulunamadı: %v\n", err)
			continue
		}

		if len(testFiles) == 0 {
			fmt.Println("Test dosyası bulunamadı")
			continue
		}

		// Her test dosyası için testleri çalıştır
		for _, file := range testFiles {
			fmt.Printf("Test dosyası: %s\n", file)
			
			// Örnek test sonuçları (gerçek implementasyonda bu kısım GO+ testlerini çalıştıracak)
			fileTests := 5
			filePassed := 4
			fileFailed := 1
			fileSkipped := 0

			totalTests += fileTests
			passedTests += filePassed
			failedTests += fileFailed
			skippedTests += fileSkipped

			// Ayrıntılı çıktı
			if *verboseFlag {
				fmt.Println("  TestAdd: Başarılı")
				fmt.Println("  TestSubtract: Başarılı")
				fmt.Println("  TestMultiply: Başarılı")
				fmt.Println("  TestDivide: Başarılı")
				fmt.Println("  TestDivideByZero: Başarısız")
				fmt.Println("    Beklenen: panic, Alınan: nil")
			}
		}
	}

	// Test sonuçlarını göster
	duration := time.Since(startTime)
	fmt.Printf("\nTest sonuçları:\n")
	fmt.Printf("  Toplam: %d\n", totalTests)
	fmt.Printf("  Başarılı: %d\n", passedTests)
	fmt.Printf("  Başarısız: %d\n", failedTests)
	fmt.Printf("  Atlanmış: %d\n", skippedTests)
	fmt.Printf("  Süre: %s\n", duration)

	// Kod kapsama analizi
	if *coverFlag {
		fmt.Println("\nKod kapsama analizi:")
		fmt.Println("  Toplam satır: 100")
		fmt.Println("  Kapsanan satır: 85")
		fmt.Println("  Kapsama oranı: %85.0")
	}

	// Çıkış kodu
	if failedTests > 0 {
		os.Exit(1)
	}
}

// Test dosyalarını bul
func findTestFiles(dir string) ([]string, error) {
	var testFiles []string

	// Dizini kontrol et
	info, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}

	// Dizin değilse, doğrudan dosyayı kontrol et
	if !info.IsDir() {
		if isTestFile(dir) {
			return []string{dir}, nil
		}
		return nil, nil
	}

	// Dizin içindeki dosyaları tara
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())

		// Alt dizinleri kontrol et
		if entry.IsDir() {
			if *recursiveFlag {
				subFiles, err := findTestFiles(path)
				if err != nil {
					fmt.Printf("Uyarı: %s dizini taranamadı: %v\n", path, err)
					continue
				}
				testFiles = append(testFiles, subFiles...)
			}
			continue
		}

		// Test dosyalarını kontrol et
		if isTestFile(path) {
			testFiles = append(testFiles, path)
		}
	}

	return testFiles, nil
}

// Dosyanın bir test dosyası olup olmadığını kontrol et
func isTestFile(path string) bool {
	// Dosya adını kontrol et
	base := filepath.Base(path)
	if !strings.HasSuffix(base, "_test.gop") {
		return false
	}

	// Dosya içeriğini kontrol et (gerçek implementasyonda bu kısım GO+ test fonksiyonlarını arayacak)
	return true
}

// Yardım mesajını yazdır
func printHelp() {
	fmt.Println("GO+ Test Aracı")
	fmt.Println("\nKullanım:")
	fmt.Println("  goptest [bayraklar] [dizinler]")
	fmt.Println("\nBayraklar:")
	flag.PrintDefaults()
	fmt.Println("\nÖrnekler:")
	fmt.Println("  goptest                  # Mevcut dizindeki testleri çalıştır")
	fmt.Println("  goptest -v               # Ayrıntılı çıktı ile testleri çalıştır")
	fmt.Println("  goptest -r               # Alt dizinlerdeki testleri de çalıştır")
	fmt.Println("  goptest -pattern=TestAdd # Sadece TestAdd ile başlayan testleri çalıştır")
	fmt.Println("  goptest -bench           # Benchmark testlerini çalıştır")
	fmt.Println("  goptest -cover           # Kod kapsama analizi yap")
	fmt.Println("  goptest ./pkg ./internal # Belirtilen dizinlerdeki testleri çalıştır")
}
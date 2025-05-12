package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Komut satırı bayrakları
var (
	writeFlag    = flag.Bool("w", false, "Değişiklikleri dosyalara yaz")
	diffFlag     = flag.Bool("d", false, "Değişiklikleri diff formatında göster")
	listFlag     = flag.Bool("l", false, "Biçimlendirilmesi gereken dosyaları listele")
	recursiveFlag = flag.Bool("r", false, "Alt dizinlerdeki dosyaları da biçimlendir")
	simplifyFlag = flag.Bool("s", false, "Kodu basitleştir")
	versionFlag  = flag.Bool("version", false, "Sürüm bilgisini göster")
	helpFlag     = flag.Bool("help", false, "Yardım mesajını göster")
)

// Kod biçimlendirme aracı sürümü
const version = "0.1.0"

func main() {
	// Bayrakları ayrıştır
	flag.Parse()

	// Sürüm bilgisini göster
	if *versionFlag {
		fmt.Printf("GO+ Kod Biçimlendirme Aracı v%s\n", version)
		os.Exit(0)
	}

	// Yardım mesajını göster
	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	// Biçimlendirilecek dosyaları belirle
	var files []string
	if flag.NArg() == 0 {
		// Mevcut dizindeki tüm GO+ dosyalarını kullan
		files = findGoPlusFiles(".", *recursiveFlag)
	} else {
		// Belirtilen dosya veya dizinleri kullan
		for _, path := range flag.Args() {
			info, err := os.Stat(path)
			if err != nil {
				fmt.Printf("Hata: %s dosyası veya dizini bulunamadı: %v\n", path, err)
				continue
			}

			if info.IsDir() {
				// Dizin içindeki GO+ dosyalarını bul
				dirFiles := findGoPlusFiles(path, *recursiveFlag)
				files = append(files, dirFiles...)
			} else {
				// Dosyayı kontrol et
				if isGoPlusFile(path) {
					files = append(files, path)
				} else {
					fmt.Printf("Uyarı: %s bir GO+ dosyası değil, atlanıyor\n", path)
				}
			}
		}
	}

	// Dosyaları biçimlendir
	formatFiles(files)
}

// GO+ dosyalarını bul
func findGoPlusFiles(dir string, recursive bool) []string {
	var files []string

	// Dizini kontrol et
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Hata: %s dizini okunamadı: %v\n", dir, err)
		return files
	}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())

		// Alt dizinleri kontrol et
		if entry.IsDir() {
			if recursive {
				subFiles := findGoPlusFiles(path, recursive)
				files = append(files, subFiles...)
			}
			continue
		}

		// GO+ dosyalarını kontrol et
		if isGoPlusFile(path) {
			files = append(files, path)
		}
	}

	return files
}

// Dosyanın bir GO+ dosyası olup olmadığını kontrol et
func isGoPlusFile(path string) bool {
	return strings.HasSuffix(path, ".gop")
}

// Dosyaları biçimlendir
func formatFiles(files []string) {
	if len(files) == 0 {
		fmt.Println("Biçimlendirilecek GO+ dosyası bulunamadı")
		return
	}

	fmt.Printf("%d GO+ dosyası biçimlendiriliyor...\n", len(files))

	// Biçimlendirme seçeneklerini göster
	fmt.Println("Biçimlendirme seçenekleri:")
	if *writeFlag {
		fmt.Println("  Değişiklikleri dosyalara yaz: Evet")
	}
	if *diffFlag {
		fmt.Println("  Değişiklikleri diff formatında göster: Evet")
	}
	if *listFlag {
		fmt.Println("  Biçimlendirilmesi gereken dosyaları listele: Evet")
	}
	if *simplifyFlag {
		fmt.Println("  Kodu basitleştir: Evet")
	}

	// Her dosyayı biçimlendir
	formattedCount := 0
	for _, file := range files {
		fmt.Printf("Biçimlendiriliyor: %s\n", file)

		// Dosyayı oku
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Hata: %s dosyası okunamadı: %v\n", file, err)
			continue
		}

		// Dosyayı biçimlendir (gerçek implementasyonda GO+ parser ve formatter kullanılacak)
		formatted, changed := formatGoPlusCode(string(content), *simplifyFlag)

		// Değişiklik yoksa devam et
		if !changed {
			fmt.Printf("  Değişiklik yok: %s\n", file)
			continue
		}

		formattedCount++

		// Biçimlendirilmesi gereken dosyaları listele
		if *listFlag {
			fmt.Printf("  Biçimlendirilmesi gerekiyor: %s\n", file)
		}

		// Değişiklikleri diff formatında göster
		if *diffFlag {
			diff := generateDiff(string(content), formatted, file)
			fmt.Println(diff)
		}

		// Değişiklikleri dosyalara yaz
		if *writeFlag {
			err := os.WriteFile(file, []byte(formatted), 0644)
			if err != nil {
				fmt.Printf("Hata: %s dosyası yazılamadı: %v\n", file, err)
				continue
			}
			fmt.Printf("  Dosya güncellendi: %s\n", file)
		}
	}

	// Sonuçları göster
	fmt.Printf("\nBiçimlendirme tamamlandı:\n")
	fmt.Printf("  Toplam dosya: %d\n", len(files))
	fmt.Printf("  Biçimlendirilen dosya: %d\n", formattedCount)
	fmt.Printf("  Değişiklik olmayan dosya: %d\n", len(files)-formattedCount)
}

// GO+ kodunu biçimlendir
func formatGoPlusCode(content string, simplify bool) (string, bool) {
	// Gerçek implementasyonda GO+ parser ve formatter kullanılacak
	// Şimdilik örnek bir biçimlendirme yapalım

	// Örnek olarak, fazla boşlukları kaldıralım
	formatted := strings.ReplaceAll(content, "  ", " ")
	
	// Değişiklik olup olmadığını kontrol et
	return formatted, formatted != content
}

// İki metin arasındaki farkı göster
func generateDiff(original, formatted, filename string) string {
	// Gerçek implementasyonda diff algoritması kullanılacak
	// Şimdilik basit bir diff gösterimi yapalım
	var diff strings.Builder
	diff.WriteString(fmt.Sprintf("--- %s (original)\n", filename))
	diff.WriteString(fmt.Sprintf("+++ %s (formatted)\n", filename))
	diff.WriteString("@@ -1,1 +1,1 @@\n")
	diff.WriteString("- " + original[:min(len(original), 40)] + "...\n")
	diff.WriteString("+ " + formatted[:min(len(formatted), 40)] + "...\n")
	return diff.String()
}

// İki sayının minimumunu döndür
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Yardım mesajını yazdır
func printHelp() {
	fmt.Println("GO+ Kod Biçimlendirme Aracı")
	fmt.Println("\nKullanım:")
	fmt.Println("  gopfmt [bayraklar] [dosyalar/dizinler]")
	fmt.Println("\nBayraklar:")
	flag.PrintDefaults()
	fmt.Println("\nÖrnekler:")
	fmt.Println("  gopfmt                  # Mevcut dizindeki GO+ dosyalarını biçimlendir")
	fmt.Println("  gopfmt -w               # Değişiklikleri dosyalara yaz")
	fmt.Println("  gopfmt -d               # Değişiklikleri diff formatında göster")
	fmt.Println("  gopfmt -l               # Biçimlendirilmesi gereken dosyaları listele")
	fmt.Println("  gopfmt -r               # Alt dizinlerdeki dosyaları da biçimlendir")
	fmt.Println("  gopfmt -s               # Kodu basitleştir")
	fmt.Println("  gopfmt file1.gop file2.gop  # Belirtilen dosyaları biçimlendir")
	fmt.Println("  gopfmt ./pkg ./internal     # Belirtilen dizinlerdeki dosyaları biçimlendir")
}
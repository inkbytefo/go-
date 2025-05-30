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
	writeFlag     = flag.Bool("w", false, "Değişiklikleri dosyalara yaz")
	diffFlag      = flag.Bool("d", false, "Değişiklikleri diff formatında göster")
	listFlag      = flag.Bool("l", false, "Biçimlendirilmesi gereken dosyaları listele")
	recursiveFlag = flag.Bool("r", false, "Alt dizinlerdeki dosyaları da biçimlendir")
	simplifyFlag  = flag.Bool("s", false, "Kodu basitleştir")
	versionFlag   = flag.Bool("version", false, "Sürüm bilgisini göster")
	helpFlag      = flag.Bool("help", false, "Yardım mesajını göster")
)

// Kod biçimlendirme aracı sürümü
const version = "0.1.0"

func main() {
	// Bayrakları ayrıştır
	flag.Parse()

	// Sürüm bilgisini göster
	if *versionFlag {
		fmt.Printf("GO-Minus Kod Biçimlendirme Aracı v%s\n", version)
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
		// Mevcut dizindeki tüm GO-Minus dosyalarını kullan
		files = findGoMinusFiles(".", *recursiveFlag)
	} else {
		// Belirtilen dosya veya dizinleri kullan
		for _, path := range flag.Args() {
			info, err := os.Stat(path)
			if err != nil {
				fmt.Printf("Hata: %s dosyası veya dizini bulunamadı: %v\n", path, err)
				continue
			}

			if info.IsDir() {
				// Dizin içindeki GO-Minus dosyalarını bul
				dirFiles := findGoMinusFiles(path, *recursiveFlag)
				files = append(files, dirFiles...)
			} else {
				// Dosyayı kontrol et
				if isGoMinusFile(path) {
					files = append(files, path)
				} else {
					fmt.Printf("Uyarı: %s bir GO-Minus dosyası değil, atlanıyor\n", path)
				}
			}
		}
	}

	// Dosyaları biçimlendir
	formatFiles(files)
}

// GO-Minus dosyalarını bul
func findGoMinusFiles(dir string, recursive bool) []string {
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
				subFiles := findGoMinusFiles(path, recursive)
				files = append(files, subFiles...)
			}
			continue
		}

		// GO-Minus dosyalarını kontrol et
		if isGoMinusFile(path) {
			files = append(files, path)
		}
	}

	return files
}

// Dosyanın bir GO-Minus dosyası olup olmadığını kontrol et
func isGoMinusFile(path string) bool {
	return strings.HasSuffix(path, ".gom")
}

// Dosyaları biçimlendir
func formatFiles(files []string) {
	if len(files) == 0 {
		fmt.Println("Biçimlendirilecek GO-Minus dosyası bulunamadı")
		return
	}

	fmt.Printf("%d GO-Minus dosyası biçimlendiriliyor...\n", len(files))

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

		// Dosyayı biçimlendir (gerçek implementasyonda GO-Minus parser ve formatter kullanılacak)
		formatted, changed := formatGoMinusCode(string(content), *simplifyFlag)

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

// GO-Minus kodunu biçimlendir
func formatGoMinusCode(content string, simplify bool) (string, bool) {
	// Gerçek implementasyonda GO-Minus parser ve formatter kullanılacak
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
	fmt.Println("GO-Minus Kod Biçimlendirme Aracı")
	fmt.Println("\nKullanım:")
	fmt.Println("  gomfmt [bayraklar] [dosyalar/dizinler]")
	fmt.Println("\nBayraklar:")
	flag.PrintDefaults()
	fmt.Println("\nÖrnekler:")
	fmt.Println("  gomfmt                  # Mevcut dizindeki GO-Minus dosyalarını biçimlendir")
	fmt.Println("  gomfmt -w               # Değişiklikleri dosyalara yaz")
	fmt.Println("  gomfmt -d               # Değişiklikleri diff formatında göster")
	fmt.Println("  gomfmt -l               # Biçimlendirilmesi gereken dosyaları listele")
	fmt.Println("  gomfmt -r               # Alt dizinlerdeki dosyaları da biçimlendir")
	fmt.Println("  gomfmt -s               # Kodu basitleştir")
	fmt.Println("  gomfmt file1.gom file2.gom  # Belirtilen dosyaları biçimlendir")
	fmt.Println("  gomfmt ./pkg ./internal     # Belirtilen dizinlerdeki dosyaları biçimlendir")
}

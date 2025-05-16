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
	htmlFlag     = flag.Bool("html", false, "HTML formatında belgelendirme oluştur")
	markdownFlag = flag.Bool("markdown", false, "Markdown formatında belgelendirme oluştur")
	outputFlag   = flag.String("output", "", "Çıktı dizini veya dosyası")
	serverFlag   = flag.Bool("server", false, "Belgelendirme sunucusu başlat")
	portFlag     = flag.Int("port", 6060, "Sunucu portu")
	allFlag      = flag.Bool("all", false, "Tüm paketleri belgele (özel ve dışa aktarılmayan öğeler dahil)")
	versionFlag  = flag.Bool("version", false, "Sürüm bilgisini göster")
	helpFlag     = flag.Bool("help", false, "Yardım mesajını göster")
)

// Belgelendirme aracı sürümü
const version = "0.1.0"

func main() {
	// Bayrakları ayrıştır
	flag.Parse()

	// Sürüm bilgisini göster
	if *versionFlag {
		fmt.Printf("GO-Minus Belgelendirme Aracı v%s\n", version)
		os.Exit(0)
	}

	// Yardım mesajını göster
	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	// Sunucu modunu kontrol et
	if *serverFlag {
		startServer()
		return
	}

	// Belgelendirilecek paketleri belirle
	var packages []string
	if flag.NArg() == 0 {
		// Mevcut dizini kullan
		packages = append(packages, ".")
	} else {
		// Belirtilen paketleri kullan
		packages = flag.Args()
	}

	// Belgelendirme formatını belirle
	var format string
	switch {
	case *htmlFlag:
		format = "html"
	case *markdownFlag:
		format = "markdown"
	default:
		format = "text"
	}

	// Belgelendirme oluştur
	generateDocs(packages, format)
}

// Belgelendirme sunucusu başlat
func startServer() {
	fmt.Printf("GO-Minus belgelendirme sunucusu başlatılıyor (port: %d)...\n", *portFlag)
	fmt.Printf("Tarayıcınızda http://localhost:%d adresini açın\n", *portFlag)
	fmt.Println("Sunucuyu durdurmak için Ctrl+C tuşlarına basın")

	// Sunucu başlatma işlemleri (gerçek implementasyonda HTTP sunucusu başlatılacak)
	select {}
}

// Belgelendirme oluştur
func generateDocs(packages []string, format string) {
	fmt.Printf("GO-Minus belgelendirmesi oluşturuluyor (format: %s)...\n", format)

	// Çıktı dizinini belirle
	outputDir := *outputFlag
	if outputDir == "" {
		outputDir = "docs"
	}

	// Çıktı dizinini oluştur
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Hata: Çıktı dizini oluşturulamadı: %v\n", err)
		os.Exit(1)
	}

	// Her paket için belgelendirme oluştur
	for _, pkg := range packages {
		fmt.Printf("Paket belgelendiriliyor: %s\n", pkg)

		// Paket dosyalarını bul
		files, err := findPackageFiles(pkg)
		if err != nil {
			fmt.Printf("Hata: Paket dosyaları bulunamadı: %v\n", err)
			continue
		}

		if len(files) == 0 {
			fmt.Println("Paket dosyası bulunamadı")
			continue
		}

		// Paket belgelendirmesi oluştur
		outputPath := filepath.Join(outputDir, filepath.Base(pkg))
		switch format {
		case "html":
			outputPath += ".html"
		case "markdown":
			outputPath += ".md"
		case "text":
			outputPath += ".txt"
		}

		// Belgelendirme içeriği oluştur (gerçek implementasyonda GO-Minus dosyalarını ayrıştırıp belgelendirme oluşturacak)
		content := generatePackageDoc(pkg, files, format)

		// Belgelendirme dosyasını yaz
		if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
			fmt.Printf("Hata: Belgelendirme dosyası yazılamadı: %v\n", err)
			continue
		}

		fmt.Printf("Belgelendirme oluşturuldu: %s\n", outputPath)
	}
}

// Paket dosyalarını bul
func findPackageFiles(pkg string) ([]string, error) {
	var files []string

	// Dizini kontrol et
	info, err := os.Stat(pkg)
	if err != nil {
		return nil, err
	}

	// Dizin değilse, doğrudan dosyayı kontrol et
	if !info.IsDir() {
		if isGoPlusFile(pkg) {
			return []string{pkg}, nil
		}
		return nil, nil
	}

	// Dizin içindeki dosyaları tara
	entries, err := os.ReadDir(pkg)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(pkg, entry.Name())
		if isGoPlusFile(path) {
			files = append(files, path)
		}
	}

	return files, nil
}

// Dosyanın bir GO-Minus dosyası olup olmadığını kontrol et
func isGoPlusFile(path string) bool {
	return strings.HasSuffix(path, ".gom")
}

// Paket belgelendirmesi oluştur
func generatePackageDoc(pkg string, files []string, format string) string {
	var content strings.Builder

	// Başlık
	switch format {
	case "html":
		content.WriteString(fmt.Sprintf("<!DOCTYPE html>\n<html>\n<head>\n<title>Paket %s</title>\n</head>\n<body>\n", pkg))
		content.WriteString(fmt.Sprintf("<h1>Paket %s</h1>\n", pkg))
	case "markdown":
		content.WriteString(fmt.Sprintf("# Paket %s\n\n", pkg))
	case "text":
		content.WriteString(fmt.Sprintf("PAKET %s\n\n", pkg))
	}

	// Paket açıklaması (gerçek implementasyonda GO-Minus dosyalarından çıkarılacak)
	content.WriteString("Paket açıklaması burada yer alacak.\n\n")

	// İçindekiler
	switch format {
	case "html":
		content.WriteString("<h2>İçindekiler</h2>\n<ul>\n")
		content.WriteString("<li><a href=\"#constants\">Sabitler</a></li>\n")
		content.WriteString("<li><a href=\"#variables\">Değişkenler</a></li>\n")
		content.WriteString("<li><a href=\"#functions\">Fonksiyonlar</a></li>\n")
		content.WriteString("<li><a href=\"#types\">Tipler</a></li>\n")
		content.WriteString("</ul>\n")
	case "markdown":
		content.WriteString("## İçindekiler\n\n")
		content.WriteString("- [Sabitler](#sabitler)\n")
		content.WriteString("- [Değişkenler](#değişkenler)\n")
		content.WriteString("- [Fonksiyonlar](#fonksiyonlar)\n")
		content.WriteString("- [Tipler](#tipler)\n\n")
	case "text":
		content.WriteString("İÇİNDEKİLER\n\n")
		content.WriteString("  Sabitler\n")
		content.WriteString("  Değişkenler\n")
		content.WriteString("  Fonksiyonlar\n")
		content.WriteString("  Tipler\n\n")
	}

	// Sabitler
	switch format {
	case "html":
		content.WriteString("<h2 id=\"constants\">Sabitler</h2>\n")
		content.WriteString("<pre>const (\n    Pi = 3.14159\n    E = 2.71828\n)</pre>\n")
	case "markdown":
		content.WriteString("## Sabitler\n\n")
		content.WriteString("```go\nconst (\n    Pi = 3.14159\n    E = 2.71828\n)\n```\n\n")
	case "text":
		content.WriteString("SABİTLER\n\n")
		content.WriteString("const (\n    Pi = 3.14159\n    E = 2.71828\n)\n\n")
	}

	// Değişkenler
	switch format {
	case "html":
		content.WriteString("<h2 id=\"variables\">Değişkenler</h2>\n")
		content.WriteString("<pre>var (\n    DefaultTimeout = 30 * time.Second\n    MaxRetries = 3\n)</pre>\n")
	case "markdown":
		content.WriteString("## Değişkenler\n\n")
		content.WriteString("```go\nvar (\n    DefaultTimeout = 30 * time.Second\n    MaxRetries = 3\n)\n```\n\n")
	case "text":
		content.WriteString("DEĞİŞKENLER\n\n")
		content.WriteString("var (\n    DefaultTimeout = 30 * time.Second\n    MaxRetries = 3\n)\n\n")
	}

	// Fonksiyonlar
	switch format {
	case "html":
		content.WriteString("<h2 id=\"functions\">Fonksiyonlar</h2>\n")
		content.WriteString("<h3>func Add</h3>\n")
		content.WriteString("<pre>func Add(a, b int) int</pre>\n")
		content.WriteString("<p>Add fonksiyonu iki sayıyı toplar ve sonucu döndürür.</p>\n")
	case "markdown":
		content.WriteString("## Fonksiyonlar\n\n")
		content.WriteString("### func Add\n\n")
		content.WriteString("```go\nfunc Add(a, b int) int\n```\n\n")
		content.WriteString("Add fonksiyonu iki sayıyı toplar ve sonucu döndürür.\n\n")
	case "text":
		content.WriteString("FONKSİYONLAR\n\n")
		content.WriteString("func Add(a, b int) int\n")
		content.WriteString("    Add fonksiyonu iki sayıyı toplar ve sonucu döndürür.\n\n")
	}

	// Tipler
	switch format {
	case "html":
		content.WriteString("<h2 id=\"types\">Tipler</h2>\n")
		content.WriteString("<h3>type Person</h3>\n")
		content.WriteString("<pre>class Person {\n    public var name string\n    public var age int\n\n    func (p Person) String() string\n}</pre>\n")
		content.WriteString("<p>Person sınıfı, bir kişiyi temsil eder.</p>\n")
		content.WriteString("</body>\n</html>\n")
	case "markdown":
		content.WriteString("## Tipler\n\n")
		content.WriteString("### type Person\n\n")
		content.WriteString("```go\nclass Person {\n    public var name string\n    public var age int\n\n    func (p Person) String() string\n}\n```\n\n")
		content.WriteString("Person sınıfı, bir kişiyi temsil eder.\n")
	case "text":
		content.WriteString("TİPLER\n\n")
		content.WriteString("class Person {\n    public var name string\n    public var age int\n\n    func (p Person) String() string\n}\n")
		content.WriteString("    Person sınıfı, bir kişiyi temsil eder.\n")
	}

	return content.String()
}

// Yardım mesajını yazdır
func printHelp() {
	fmt.Println("GO-Minus Belgelendirme Aracı")
	fmt.Println("\nKullanım:")
	fmt.Println("  gomdoc [bayraklar] [paketler]")
	fmt.Println("\nBayraklar:")
	flag.PrintDefaults()
	fmt.Println("\nÖrnekler:")
	fmt.Println("  gomdoc                      # Mevcut dizindeki paketi belgele")
	fmt.Println("  gomdoc -html                # HTML formatında belgelendirme oluştur")
	fmt.Println("  gomdoc -markdown            # Markdown formatında belgelendirme oluştur")
	fmt.Println("  gomdoc -output=docs         # Belgelendirmeyi docs dizinine kaydet")
	fmt.Println("  gomdoc -server              # Belgelendirme sunucusu başlat")
	fmt.Println("  gomdoc -server -port=8080   # Belgelendirme sunucusunu 8080 portunda başlat")
	fmt.Println("  gomdoc ./pkg ./internal     # Belirtilen paketleri belgele")
}
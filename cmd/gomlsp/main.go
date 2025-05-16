package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sourcegraph/jsonrpc2"
)

// Dil sunucusu sürümü
const version = "0.1.0"

// Komut satırı bayrakları
var (
	versionFlag = flag.Bool("version", false, "Sürüm bilgisini göster")
	helpFlag    = flag.Bool("help", false, "Yardım mesajını göster")
	logFlag     = flag.String("log", "", "Log dosyası (varsayılan: stderr)")
	modeFlag    = flag.String("mode", "stdio", "İletişim modu (stdio, tcp)")
	addrFlag    = flag.String("addr", ":8080", "TCP sunucu adresi (mode=tcp ise)")
)

func main() {
	// Bayrakları ayrıştır
	flag.Parse()

	// Sürüm bilgisini göster
	if *versionFlag {
		fmt.Printf("GO-Minus Dil Sunucusu v%s\n", version)
		os.Exit(0)
	}

	// Yardım mesajını göster
	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	// Log dosyasını ayarla
	if *logFlag != "" {
		f, err := os.Create(*logFlag)
		if err != nil {
			log.Fatalf("Log dosyası oluşturulamadı: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	// Sinyal işleyicisini ayarla
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
	}()

	// İletişim moduna göre sunucuyu başlat
	switch *modeFlag {
	case "stdio":
		log.Println("GO-Minus Dil Sunucusu başlatılıyor (stdio modu)...")
		serveStdio(ctx)
	case "tcp":
		log.Printf("GO-Minus Dil Sunucusu başlatılıyor (tcp modu, adres: %s)...\n", *addrFlag)
		serveTCP(ctx, *addrFlag)
	default:
		log.Fatalf("Bilinmeyen mod: %s", *modeFlag)
	}
}

// Stdio üzerinden sunucu başlat
func serveStdio(ctx context.Context) {
	stream := jsonrpc2.NewBufferedStream(stdrwc{}, jsonrpc2.VSCodeObjectCodec{})
	handler := &langHandler{}
	conn := jsonrpc2.NewConn(ctx, stream, handler)
	<-conn.DisconnectNotify()
	log.Println("Bağlantı kapatıldı")
}

// TCP üzerinden sunucu başlat
func serveTCP(ctx context.Context, addr string) {
	// TODO: TCP sunucu implementasyonu
	log.Println("TCP sunucu henüz implementasyonu tamamlanmadı")
}

// Standart giriş/çıkış için okuma/yazma kapatıcı
type stdrwc struct{}

func (stdrwc) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

func (stdrwc) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func (stdrwc) Close() error {
	if err := os.Stdin.Close(); err != nil {
		return err
	}
	return os.Stdout.Close()
}

// Dil sunucusu işleyicisi
type langHandler struct{}

// Handle, JSON-RPC isteklerini işler
func (h *langHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	switch req.Method {
	case "initialize":
		// TODO: İstemci başlatma işlemleri
		if err := conn.Reply(ctx, req.ID, map[string]interface{}{
			"capabilities": map[string]interface{}{
				"textDocumentSync": 1, // Full
				"hoverProvider":    true,
				"completionProvider": map[string]interface{}{
					"triggerCharacters": []string{".", ":"},
				},
				"definitionProvider":              true,
				"referencesProvider":              true,
				"documentSymbolProvider":          true,
				"workspaceSymbolProvider":         true,
				"documentFormattingProvider":      true,
				"documentRangeFormattingProvider": true,
			},
		}); err != nil {
			log.Printf("initialize yanıtı gönderilemedi: %v", err)
		}
	case "initialized":
		// İstemci başlatıldı
		log.Println("İstemci başlatıldı")
	case "shutdown":
		// Sunucu kapatılıyor
		log.Println("Sunucu kapatılıyor")
		if err := conn.Reply(ctx, req.ID, nil); err != nil {
			log.Printf("shutdown yanıtı gönderilemedi: %v", err)
		}
	case "exit":
		// Çıkış
		log.Println("Çıkış yapılıyor")
		os.Exit(0)
	case "textDocument/didOpen":
		// Belge açıldı
		log.Println("Belge açıldı")
		// TODO: Belge açma işlemleri
	case "textDocument/didChange":
		// Belge değişti
		log.Println("Belge değişti")
		// TODO: Belge değişikliği işlemleri
	case "textDocument/didClose":
		// Belge kapandı
		log.Println("Belge kapandı")
		// TODO: Belge kapanma işlemleri
	case "textDocument/completion":
		// Kod tamamlama
		log.Println("Kod tamamlama istendi")
		// TODO: Kod tamamlama işlemleri
		if err := conn.Reply(ctx, req.ID, map[string]interface{}{
			"isIncomplete": false,
			"items": []map[string]interface{}{
				{
					"label":  "fmt",
					"kind":   9, // Module
					"detail": "fmt paketi",
				},
				{
					"label":  "Println",
					"kind":   3, // Function
					"detail": "func Println(a ...interface{}) (n int, err error)",
				},
			},
		}); err != nil {
			log.Printf("completion yanıtı gönderilemedi: %v", err)
		}
	case "textDocument/hover":
		// Fare üzerinde bilgi
		log.Println("Fare üzerinde bilgi istendi")
		// TODO: Fare üzerinde bilgi işlemleri
		if err := conn.Reply(ctx, req.ID, map[string]interface{}{
			"contents": "GO-Minus dil sunucusu tarafından sağlanan bilgi",
		}); err != nil {
			log.Printf("hover yanıtı gönderilemedi: %v", err)
		}
	case "textDocument/definition":
		// Tanıma gitme
		log.Println("Tanıma gitme istendi")
		// TODO: Tanıma gitme işlemleri
		if err := conn.Reply(ctx, req.ID, nil); err != nil {
			log.Printf("definition yanıtı gönderilemedi: %v", err)
		}
	case "textDocument/references":
		// Referanslar
		log.Println("Referanslar istendi")
		// TODO: Referanslar işlemleri
		if err := conn.Reply(ctx, req.ID, nil); err != nil {
			log.Printf("references yanıtı gönderilemedi: %v", err)
		}
	case "textDocument/documentSymbol":
		// Belge sembolleri
		log.Println("Belge sembolleri istendi")
		// TODO: Belge sembolleri işlemleri
		if err := conn.Reply(ctx, req.ID, nil); err != nil {
			log.Printf("documentSymbol yanıtı gönderilemedi: %v", err)
		}
	case "workspace/symbol":
		// Çalışma alanı sembolleri
		log.Println("Çalışma alanı sembolleri istendi")
		// TODO: Çalışma alanı sembolleri işlemleri
		if err := conn.Reply(ctx, req.ID, nil); err != nil {
			log.Printf("workspaceSymbol yanıtı gönderilemedi: %v", err)
		}
	case "textDocument/formatting":
		// Belge biçimlendirme
		log.Println("Belge biçimlendirme istendi")
		// TODO: Belge biçimlendirme işlemleri
		if err := conn.Reply(ctx, req.ID, nil); err != nil {
			log.Printf("formatting yanıtı gönderilemedi: %v", err)
		}
	case "textDocument/rangeFormatting":
		// Belge aralığı biçimlendirme
		log.Println("Belge aralığı biçimlendirme istendi")
		// TODO: Belge aralığı biçimlendirme işlemleri
		if err := conn.Reply(ctx, req.ID, nil); err != nil {
			log.Printf("rangeFormatting yanıtı gönderilemedi: %v", err)
		}
	default:
		// Bilinmeyen metot
		log.Printf("Bilinmeyen metot: %s", req.Method)
		if req.ID != nil {
			if err := conn.Reply(ctx, req.ID, nil); err != nil {
				log.Printf("Bilinmeyen metot yanıtı gönderilemedi: %v", err)
			}
		}
	}
}

// Yardım mesajını yazdır
func printHelp() {
	fmt.Println("GO-Minus Dil Sunucusu")
	fmt.Println("\nKullanım:")
	fmt.Println("  gomlsp [bayraklar]")
	fmt.Println("\nBayraklar:")
	flag.PrintDefaults()
	fmt.Println("\nÖrnekler:")
	fmt.Println("  gomlsp                  # Standart giriş/çıkış üzerinden çalıştır")
	fmt.Println("  gomlsp -mode=tcp        # TCP sunucu olarak çalıştır (varsayılan port: 8080)")
	fmt.Println("  gomlsp -mode=tcp -addr=:9090  # TCP sunucu olarak belirtilen portta çalıştır")
	fmt.Println("  gomlsp -log=gomlsp.log  # Log dosyasına yaz")
}
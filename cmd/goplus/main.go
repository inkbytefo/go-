package main

import (
	"fmt"
	"goplus/internal/lexer"
	"goplus/internal/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: goplus <dosya.gop>")
		os.Exit(1)
	}

	// Dosyayı oku
	filename := os.Args[1]
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
		printParserErrors(p.Errors())
		os.Exit(1)
	}

	// AST'yi yazdır
	fmt.Println("AST:")
	fmt.Println(program.String())
}

func printParserErrors(errors []string) {
	fmt.Println("Ayrıştırma hataları:")
	for _, msg := range errors {
		fmt.Println("\t" + msg)
	}
}
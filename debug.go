package main

import (
	"fmt"

	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/parser"
	"github.com/inkbytefo/go-minus/internal/semantic"
)

func main() {
	input := `package main

import "fmt"

func main() {
    fmt.Println("Hello")
}`

	fmt.Println("Input:", input)
	fmt.Println("\nParser test:")

	// Parser test
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	errors := p.Errors()
	if len(errors) > 0 {
		fmt.Println("Parser errors:")
		for _, err := range errors {
			fmt.Println("\t", err)
		}
		return
	} else {
		fmt.Println("Parsing successful!")
	}

	fmt.Println("\nSemantic analysis test:")

	// Semantic analysis test
	analyzer := semantic.New()

	// Global scope'daki sembolleri kontrol et
	fmt.Println("Global scope symbols:")
	analyzer.PrintGlobalScope()

	analyzer.Analyze(program)

	semanticErrors := analyzer.Errors()
	if len(semanticErrors) > 0 {
		fmt.Println("Semantic errors:")
		for _, err := range semanticErrors {
			fmt.Println("\t", err)
		}
	} else {
		fmt.Println("Semantic analysis successful!")
	}
}

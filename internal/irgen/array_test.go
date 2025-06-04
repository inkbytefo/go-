package irgen

import (
	"testing"

	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/parser"
	"github.com/inkbytefo/go-minus/internal/semantic"
)

func TestArrayLiteralIRGeneration(t *testing.T) {
	input := `
		func main() {
			arr := [1, 2, 3]
			x := arr[0]
		}
	`

	// Parse the input
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	// Check for parsing errors
	errors := p.Errors()
	if len(errors) > 0 {
		t.Fatalf("Parser errors: %v", errors)
	}

	// Semantic analysis
	analyzer := semantic.New()
	analyzer.Analyze(program)

	// Check for semantic errors
	if len(analyzer.Errors()) > 0 {
		t.Fatalf("Semantic errors: %v", analyzer.Errors())
	}

	// IR generation
	generator := NewWithAnalyzer(analyzer)
	ir, err := generator.GenerateProgram(program)

	// Check for IR generation errors
	if err != nil {
		t.Fatalf("IR generation error: %v", err)
	}

	// Check for IR generation errors
	if len(generator.Errors()) > 0 {
		t.Fatalf("IR generation errors: %v", generator.Errors())
	}

	// Print the generated IR for debugging
	t.Logf("Generated IR:\n%s", ir)
}

func TestArrayIndexingIRGeneration(t *testing.T) {
	input := `
		func main() {
			arr := [10, 20, 30]
			first := arr[0]
			second := arr[1]
		}
	`

	// Parse the input
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	// Check for parsing errors
	errors := p.Errors()
	if len(errors) > 0 {
		t.Fatalf("Parser errors: %v", errors)
	}

	// Semantic analysis
	analyzer := semantic.New()
	analyzer.Analyze(program)

	// Check for semantic errors
	if len(analyzer.Errors()) > 0 {
		t.Fatalf("Semantic errors: %v", analyzer.Errors())
	}

	// IR generation
	generator := NewWithAnalyzer(analyzer)
	ir, err := generator.GenerateProgram(program)

	// Check for IR generation errors
	if err != nil {
		t.Fatalf("IR generation error: %v", err)
	}

	// Check for IR generation errors
	if len(generator.Errors()) > 0 {
		t.Fatalf("IR generation errors: %v", generator.Errors())
	}

	// Print the generated IR for debugging
	t.Logf("Generated IR:\n%s", ir)
}

func TestEmptyArrayIRGeneration(t *testing.T) {
	input := `
		func main() {
			arr := []
		}
	`

	// Parse the input
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	// Check for parsing errors
	errors := p.Errors()
	if len(errors) > 0 {
		t.Fatalf("Parser errors: %v", errors)
	}

	// Semantic analysis
	analyzer := semantic.New()
	analyzer.Analyze(program)

	// Check for semantic errors
	if len(analyzer.Errors()) > 0 {
		t.Fatalf("Semantic errors: %v", analyzer.Errors())
	}

	// IR generation
	generator := NewWithAnalyzer(analyzer)
	ir, err := generator.GenerateProgram(program)

	// Check for IR generation errors
	if err != nil {
		t.Fatalf("IR generation error: %v", err)
	}

	// Check for IR generation errors
	if len(generator.Errors()) > 0 {
		t.Fatalf("IR generation errors: %v", generator.Errors())
	}

	// Print the generated IR for debugging
	t.Logf("Generated IR:\n%s", ir)
}

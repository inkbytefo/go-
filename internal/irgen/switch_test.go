package irgen

import (
	"testing"

	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/parser"
	"github.com/inkbytefo/go-minus/internal/semantic"
)

func TestSwitchStatementIRGeneration(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name: "Simple switch with integer cases",
			input: `
				func main() {
					var x int = 2
					switch x {
					case 1:
						fmt.Println("one")
					case 2:
						fmt.Println("two")
					default:
						fmt.Println("other")
					}
				}
			`,
		},
		{
			name: "Switch without tag (boolean cases)",
			input: `
				func main() {
					var x int = 5
					switch {
					case x > 0:
						fmt.Println("positive")
					case x < 0:
						fmt.Println("negative")
					default:
						fmt.Println("zero")
					}
				}
			`,
		},
		{
			name: "Switch with multiple case values",
			input: `
				func main() {
					var x int = 3
					switch x {
					case 1, 2, 3:
						fmt.Println("small")
					case 4, 5:
						fmt.Println("medium")
					default:
						fmt.Println("large")
					}
				}
			`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the input
			l := lexer.New(tt.input)
			p := parser.New(l)
			program := p.ParseProgram()

			// Check for parsing errors
			if len(p.Errors()) > 0 {
				t.Fatalf("Parser errors: %v", p.Errors())
			}

			// Perform semantic analysis
			analyzer := semantic.New()
			analyzer.Analyze(program)
			if len(analyzer.Errors()) > 0 {
				t.Fatalf("Semantic analysis errors: %v", analyzer.Errors())
			}

			// Generate IR
			generator := NewWithAnalyzer(analyzer)
			ir, err := generator.GenerateProgram(program)
			if err != nil {
				t.Fatalf("IR generation error: %v", err)
			}

			// Check that IR was generated without errors
			if len(generator.Errors()) > 0 {
				t.Errorf("IR generation errors: %v", generator.Errors())
			}

			// Check that module was created
			if generator.module == nil {
				t.Errorf("Expected module to be created, got nil")
			}

			// Basic validation - check that main function exists
			mainFunc := generator.module.Funcs[0]
			if mainFunc.Name() != "main" {
				t.Errorf("Expected main function, got %s", mainFunc.Name())
			}

			// Check that switch blocks were created
			if len(mainFunc.Blocks) < 3 { // At least entry, case, and end blocks
				t.Errorf("Expected at least 3 blocks for switch statement, got %d", len(mainFunc.Blocks))
			}

			// Print IR for debugging (optional)
			t.Logf("Generated IR:\n%s", ir)
		})
	}
}

func TestSwitchWithoutDefault(t *testing.T) {
	input := `
		func main() {
			var x int = 1
			switch x {
			case 1:
				fmt.Println("one")
			case 2:
				fmt.Println("two")
			}
		}
	`

	// Parse the input
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	// Check for parsing errors
	if len(p.Errors()) > 0 {
		t.Fatalf("Parser errors: %v", p.Errors())
	}

	// Perform semantic analysis
	analyzer := semantic.New()
	analyzer.Analyze(program)
	if len(analyzer.Errors()) > 0 {
		t.Fatalf("Semantic analysis errors: %v", analyzer.Errors())
	}

	// Generate IR
	generator := NewWithAnalyzer(analyzer)
	ir, err := generator.GenerateProgram(program)
	if err != nil {
		t.Fatalf("IR generation error: %v", err)
	}

	// Check that IR was generated without errors
	if len(generator.Errors()) > 0 {
		t.Errorf("IR generation errors: %v", generator.Errors())
	}

	// Print IR for debugging
	t.Logf("Generated IR:\n%s", ir)
}

func TestEmptySwitch(t *testing.T) {
	input := `
		func main() {
			var x int = 1
			switch x {
			}
		}
	`

	// Parse the input
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	// Check for parsing errors
	if len(p.Errors()) > 0 {
		t.Fatalf("Parser errors: %v", p.Errors())
	}

	// Perform semantic analysis
	analyzer := semantic.New()
	analyzer.Analyze(program)
	if len(analyzer.Errors()) > 0 {
		t.Fatalf("Semantic analysis errors: %v", analyzer.Errors())
	}

	// Generate IR
	generator := NewWithAnalyzer(analyzer)
	ir, err := generator.GenerateProgram(program)
	if err != nil {
		t.Fatalf("IR generation error: %v", err)
	}

	// Check that IR was generated without errors
	if len(generator.Errors()) > 0 {
		t.Errorf("IR generation errors: %v", generator.Errors())
	}

	// Print IR for debugging
	t.Logf("Generated IR:\n%s", ir)
}

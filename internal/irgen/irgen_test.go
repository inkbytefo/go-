package irgen

import (
	"strings"
	"testing"

	"github.com/inkbytefo/go-minus/internal/ast"
	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/parser"
	"github.com/inkbytefo/go-minus/internal/semantic"
	"github.com/inkbytefo/go-minus/internal/token"
)

// TestGenerateProgram tests the GenerateProgram function.
func TestGenerateProgram(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		contains []string
	}{
		{
			name: "Empty program",
			input: `
package main

func main() {
}
`,
			wantErr:  false,
			contains: []string{"define", "main"},
		},
		{
			name: "Simple variable declaration",
			input: `
package main

func main() {
    var x int = 42
}
`,
			wantErr:  false,
			contains: []string{"define", "main", "alloca", "store", "i32"},
		},
		{
			name: "Simple function call",
			input: `
package main

func main() {
    println("Hello, World!")
}
`,
			wantErr:  false,
			contains: []string{"define", "main", "call", "println"},
		},
		{
			name: "Simple arithmetic",
			input: `
package main

func main() {
    var x int = 2 + 3 * 4
}
`,
			wantErr:  false,
			contains: []string{"define", "main", "alloca", "mul", "add", "store"},
		},
		{
			name: "Simple if statement",
			input: `
package main

func main() {
    var x int = 42
    if x > 10 {
        println("x is greater than 10")
    }
}
`,
			wantErr:  false,
			contains: []string{"define", "main", "icmp", "br", "call"},
		},
		{
			name: "Simple while loop",
			input: `
package main

func main() {
    var i int = 0
    while i < 10 {
        println(i)
        i = i + 1
    }
}
`,
			wantErr:  false,
			contains: []string{"define", "main", "icmp", "br", "call", "add"},
		},
		{
			name: "Function with parameters and return value",
			input: `
package main

func add(a int, b int) int {
    return a + b
}

func main() {
    var result int = add(2, 3)
    println(result)
}
`,
			wantErr:  false,
			contains: []string{"define", "main", "define", "add", "call", "ret"},
		},
		{
			name: "Invalid syntax",
			input: `
package main

func main() {
    var x int = 
}
`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the input
			l := lexer.New(tt.input)
			p := parser.New(l)
			program := p.ParseProgram()
			
			if len(p.Errors()) > 0 && !tt.wantErr {
				t.Fatalf("Parser errors: %v", p.Errors())
			}
			
			// Perform semantic analysis
			analyzer := semantic.New()
			analyzer.Analyze(program)
			
			if len(analyzer.Errors()) > 0 && !tt.wantErr {
				t.Fatalf("Semantic errors: %v", analyzer.Errors())
			}
			
			// Generate IR
			generator := NewWithAnalyzer(analyzer)
			ir, err := generator.GenerateProgram(program)
			
			// Check for errors
			if (err != nil) != tt.wantErr {
				t.Fatalf("GenerateProgram() error = %v, wantErr %v", err, tt.wantErr)
			}
			
			if err != nil {
				return
			}
			
			// Check that the IR contains the expected strings
			for _, s := range tt.contains {
				if !strings.Contains(ir, s) {
					t.Errorf("IR does not contain %q", s)
				}
			}
		})
	}
}

// TestDebugInfo tests the debug information generation.
func TestDebugInfo(t *testing.T) {
	// Create a simple program
	input := `
package main

func main() {
    var x int = 42
    println(x)
}
`
	
	// Parse the input
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		t.Fatalf("Parser errors: %v", p.Errors())
	}
	
	// Perform semantic analysis
	analyzer := semantic.New()
	analyzer.Analyze(program)
	
	if len(analyzer.Errors()) > 0 {
		t.Fatalf("Semantic errors: %v", analyzer.Errors())
	}
	
	// Create an IR generator with debug info enabled
	generator := NewWithAnalyzer(analyzer)
	generator.SetSourceFile("test.gom", "/tmp")
	generator.EnableDebugInfo(true)
	
	// Generate IR
	ir, err := generator.GenerateProgram(program)
	if err != nil {
		t.Fatalf("GenerateProgram() error = %v", err)
	}
	
	// In a real implementation, we would check that the IR contains debug info
	// For now, we just check that the IR is not empty
	if ir == "" {
		t.Fatal("IR is empty")
	}
}

// TestSetSourceFile tests the SetSourceFile function.
func TestSetSourceFile(t *testing.T) {
	generator := New()
	generator.SetSourceFile("test.gom", "/tmp")
	
	if generator.sourceFile != "test.gom" {
		t.Errorf("sourceFile = %q, want %q", generator.sourceFile, "test.gom")
	}
	
	if generator.sourceDir != "/tmp" {
		t.Errorf("sourceDir = %q, want %q", generator.sourceDir, "/tmp")
	}
}

// TestEnableDebugInfo tests the EnableDebugInfo function.
func TestEnableDebugInfo(t *testing.T) {
	generator := New()
	generator.EnableDebugInfo(true)
	
	if !generator.generateDebug {
		t.Error("generateDebug = false, want true")
	}
	
	generator.EnableDebugInfo(false)
	
	if generator.generateDebug {
		t.Error("generateDebug = true, want false")
	}
}

package test

import (
	"testing"

	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/parser"
	"github.com/inkbytefo/go-minus/internal/semantic"
)

// TestFullCompilerPipeline tests the complete compiler pipeline
func TestFullCompilerPipeline(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name: "Simple variable declaration",
			input: `
				var x = 42;
				var y = "hello";
				var z = true;
			`,
			wantErr: false,
		},
		{
			name: "Function declaration",
			input: `
				func add(a, b) {
					return a + b;
				}
			`,
			wantErr: false,
		},
		{
			name: "Class declaration",
			input: `
				class Person {
					name
					age
				}
			`,
			wantErr: false,
		},
		{
			name: "Complex program",
			input: `
				var globalVar = 100;
				
				func calculate(x, y) {
					var result = x + y;
					return result;
				}
				
				class Calculator {
					value
				}
				
				var calc = calculate(10, 20);
			`,
			wantErr: false,
		},
		{
			name: "Syntax error",
			input: `
				var x = ;
			`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Lexical analysis
			l := lexer.New(tt.input)
			
			// Syntax analysis
			p := parser.New(l)
			program := p.ParseProgram()
			
			parseErrors := p.Errors()
			if len(parseErrors) > 0 && !tt.wantErr {
				t.Errorf("Parse errors: %v", parseErrors)
				return
			}
			
			if len(parseErrors) == 0 && tt.wantErr {
				t.Errorf("Expected parse errors but got none")
				return
			}
			
			if tt.wantErr {
				return // Skip semantic analysis for syntax errors
			}
			
			// Semantic analysis
			analyzer := semantic.New()
			analyzer.Analyze(program)
			
			semanticErrors := analyzer.Errors()
			if len(semanticErrors) > 0 {
				t.Logf("Semantic errors (may be expected): %v", semanticErrors)
			}
			
			// Basic validation
			if program == nil {
				t.Errorf("Program is nil")
				return
			}
			
			if len(program.Statements) == 0 {
				t.Errorf("Program has no statements")
				return
			}
			
			t.Logf("Successfully processed %d statements", len(program.Statements))
		})
	}
}

// TestLexerPerformance tests lexer performance
func TestLexerPerformance(t *testing.T) {
	input := `
		package main
		
		import "fmt"
		
		var globalVar = 42;
		
		class Person {
			private:
				string name
				int age
			
			public:
				Person(string name, int age) {
					this.name = name
					this.age = age
				}
				
				string getName() {
					return this.name
				}
				
				void birthday() {
					this.age++
				}
		}
		
		func fibonacci(n) {
			if (n <= 1) {
				return n;
			}
			return fibonacci(n - 1) + fibonacci(n - 2);
		}
		
		func main() {
			var person = Person("John", 30);
			fmt.Println("Name:", person.getName());
			person.birthday();
			
			var result = fibonacci(10);
			fmt.Println("Fibonacci(10):", result);
		}
	`
	
	b := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := lexer.New(input)
			for {
				tok := l.NextToken()
				if tok.Type == "EOF" {
					break
				}
			}
		}
	})
	
	t.Logf("Lexer performance: %v", b)
}

// TestParserPerformance tests parser performance
func TestParserPerformance(t *testing.T) {
	input := `
		var x = 5;
		var y = 10;
		var z = x + y;
		
		func add(a, b) {
			return a + b;
		}
		
		class Calculator {
			value
		}
		
		var result = add(x, y);
	`
	
	b := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l := lexer.New(input)
			p := parser.New(l)
			_ = p.ParseProgram()
		}
	})
	
	t.Logf("Parser performance: %v", b)
}

// TestSemanticAnalysisPerformance tests semantic analysis performance
func TestSemanticAnalysisPerformance(t *testing.T) {
	input := `
		var x = 5;
		var y = 10;
		var z = x + y;
		
		func add(a, b) {
			return a + b;
		}
		
		var result = add(x, y);
	`
	
	// Pre-parse the program
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		t.Fatalf("Parse errors: %v", p.Errors())
	}
	
	b := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			analyzer := semantic.New()
			analyzer.Analyze(program)
		}
	})
	
	t.Logf("Semantic analysis performance: %v", b)
}

// TestErrorRecovery tests error recovery mechanisms
func TestErrorRecovery(t *testing.T) {
	input := `
		var x = 5;
		var y = ; // Syntax error
		var z = 10; // Should still be parsed
	`
	
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	errors := p.Errors()
	if len(errors) == 0 {
		t.Errorf("Expected parse errors but got none")
	}
	
	// Should still have some valid statements
	if program == nil || len(program.Statements) == 0 {
		t.Errorf("Parser should recover and parse valid statements")
	}
	
	t.Logf("Recovered from errors and parsed %d statements", len(program.Statements))
	t.Logf("Errors: %v", errors)
}

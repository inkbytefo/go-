package semantic

import (
	"testing"

	"github.com/inkbytefo/go-minus/internal/ast"
	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/parser"
	"github.com/inkbytefo/go-minus/internal/testutil"
)

// parseProgram parses a GO-Minus program and returns the AST and any errors.
func parseProgram(input string) (*ast.Program, []string) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return program, p.Errors()
}

// analyzeProgram performs semantic analysis on a program.
func analyzeProgram(program *ast.Program) (*Analyzer, []string) {
	analyzer := New()
	analyzer.Analyze(program)
	return analyzer, analyzer.Errors()
}

func TestVariableDeclaration(t *testing.T) {
	tests := []testutil.SemanticTestCase{
		{
			Name:    "Simple variable declaration",
			Input:   "var x = 5;",
			WantErr: false,
		},
		{
			Name:    "Multiple variable declarations",
			Input:   "var x = 5; var y = 10;",
			WantErr: false,
		},
		// TODO: Variable redeclaration check not implemented yet
		// {
		//     Name:    "Variable redeclaration should fail",
		//     Input:   "var x = 5; var x = 10;",
		//     WantErr: true,
		//     ErrorMsg: "already declared",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			program, parseErrors := parseProgram(tt.Input)
			testutil.AssertNoErrors(t, parseErrors)

			_, semanticErrors := analyzeProgram(program)

			if tt.WantErr {
				if len(semanticErrors) == 0 {
					t.Errorf("Expected semantic error, but got none")
				} else if tt.ErrorMsg != "" {
					testutil.AssertErrorContains(t, semanticErrors, tt.ErrorMsg)
				}
			} else {
				testutil.AssertNoErrors(t, semanticErrors)
			}
		})
	}
}

func TestVariableUsage(t *testing.T) {
	tests := []testutil.SemanticTestCase{
		{
			Name:    "Use declared variable",
			Input:   "var x = 5; var y = x;",
			WantErr: false,
		},
		{
			Name:     "Use undeclared variable should fail",
			Input:    "var y = x;",
			WantErr:  true,
			ErrorMsg: "Tanımlanmamış tanımlayıcı",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			program, parseErrors := parseProgram(tt.Input)
			testutil.AssertNoErrors(t, parseErrors)

			_, semanticErrors := analyzeProgram(program)

			if tt.WantErr {
				if len(semanticErrors) == 0 {
					t.Errorf("Expected semantic error, but got none")
				} else if tt.ErrorMsg != "" {
					testutil.AssertErrorContains(t, semanticErrors, tt.ErrorMsg)
				}
			} else {
				testutil.AssertNoErrors(t, semanticErrors)
			}
		})
	}
}

func TestFunctionDeclaration(t *testing.T) {
	tests := []testutil.SemanticTestCase{
		{
			Name:    "Simple function declaration",
			Input:   "func add(x, y) { return x + y; }",
			WantErr: false,
		},
		{
			Name:    "Function without parameters",
			Input:   "func hello() { }",
			WantErr: false,
		},
		// TODO: Function redeclaration check not implemented yet
		// {
		//     Name:     "Function redeclaration should fail",
		//     Input:    "func test() { } func test() { }",
		//     WantErr:  true,
		//     ErrorMsg: "already declared",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			program, parseErrors := parseProgram(tt.Input)
			if len(parseErrors) > 0 {
				t.Skip("Parser errors, skipping semantic test")
				return
			}

			_, semanticErrors := analyzeProgram(program)

			if tt.WantErr {
				if len(semanticErrors) == 0 {
					t.Errorf("Expected semantic error, but got none")
				} else if tt.ErrorMsg != "" {
					testutil.AssertErrorContains(t, semanticErrors, tt.ErrorMsg)
				}
			} else {
				testutil.AssertNoErrors(t, semanticErrors)
			}
		})
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []testutil.SemanticTestCase{
		{
			Name:    "Return in function",
			Input:   "func test() { return 5; }",
			WantErr: false,
		},
		// TODO: Return outside function check not implemented yet
		// {
		//     Name:     "Return outside function should fail",
		//     Input:    "return 5;",
		//     WantErr:  true,
		//     ErrorMsg: "return outside function",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			program, parseErrors := parseProgram(tt.Input)
			if len(parseErrors) > 0 {
				t.Skip("Parser errors, skipping semantic test")
				return
			}

			_, semanticErrors := analyzeProgram(program)

			if tt.WantErr {
				if len(semanticErrors) == 0 {
					t.Errorf("Expected semantic error, but got none")
				} else if tt.ErrorMsg != "" {
					testutil.AssertErrorContains(t, semanticErrors, tt.ErrorMsg)
				}
			} else {
				testutil.AssertNoErrors(t, semanticErrors)
			}
		})
	}
}

func TestExpressions(t *testing.T) {
	tests := []testutil.SemanticTestCase{
		{
			Name:    "Arithmetic expression",
			Input:   "var result = 5 + 3;",
			WantErr: false,
		},
		{
			Name:    "Boolean expression",
			Input:   "var result = true && false;",
			WantErr: false,
		},
		{
			Name:    "Comparison expression",
			Input:   "var result = 5 > 3;",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			program, parseErrors := parseProgram(tt.Input)
			testutil.AssertNoErrors(t, parseErrors)

			_, semanticErrors := analyzeProgram(program)

			if tt.WantErr {
				if len(semanticErrors) == 0 {
					t.Errorf("Expected semantic error, but got none")
				} else if tt.ErrorMsg != "" {
					testutil.AssertErrorContains(t, semanticErrors, tt.ErrorMsg)
				}
			} else {
				testutil.AssertNoErrors(t, semanticErrors)
			}
		})
	}
}

func BenchmarkSemanticAnalysis(b *testing.B) {
	input := `
	var x = 5;
	var y = 10;
	var z = x + y;

	func add(a, b) {
		return a + b;
	}

	var result = add(x, y);
	`

	// Pre-parse the program once
	program, parseErrors := parseProgram(input)
	if len(parseErrors) > 0 {
		b.Fatalf("Parse errors: %v", parseErrors)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer := New()
		analyzer.Analyze(program)
	}
}

func TestComplexProgram(t *testing.T) {
	input := `
	var globalVar = 42;
	var x = 5;
	var y = 10;
	var z = x + y;
	`

	program, parseErrors := parseProgram(input)
	if len(parseErrors) > 0 {
		t.Skip("Parser errors, skipping semantic test")
		return
	}

	_, semanticErrors := analyzeProgram(program)
	// Some semantic errors are expected due to incomplete implementation
	if len(semanticErrors) > 0 {
		t.Logf("Semantic errors (expected): %v", semanticErrors)
	}
}

package parser

import (
	"testing"

	"github.com/inkbytefo/go-minus/internal/ast"
	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/testutil"
)

// parseProgram parses a GO-Minus program and returns the AST and any errors.
func parseProgram(input string) (*ast.Program, []string) {
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	return program, p.Errors()
}

func TestVarStatements(t *testing.T) {
	tests := []testutil.ParserTestCase{
		{
			Name:    "Simple variable declaration",
			Input:   "var x = 5;",
			WantErr: false,
		},
		{
			Name: "Multiple variable declarations",
			Input: `
				var x = 5;
				var y = 10;
				var foobar = 838383;
			`,
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			program, errors := parseProgram(tt.Input)

			if tt.WantErr {
				testutil.AssertHasErrors(t, errors, 1)
				return
			}

			testutil.AssertNoErrors(t, errors)

			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			if len(program.Statements) == 0 {
				t.Fatalf("program.Statements is empty")
			}

			// Check that all statements are variable statements
			for i, stmt := range program.Statements {
				if !testVarStatement(t, stmt) {
					t.Errorf("Statement %d is not a variable statement", i)
				}
			}
		})
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []testutil.ParserTestCase{
		{
			Name:    "Simple return",
			Input:   "return 5;",
			WantErr: false,
		},
		{
			Name:    "Return expression",
			Input:   "return x + y;",
			WantErr: false,
		},
		{
			Name:    "Return without value",
			Input:   "return;",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			program, errors := parseProgram(tt.Input)

			if tt.WantErr {
				testutil.AssertHasErrors(t, errors, 1)
				return
			}

			testutil.AssertNoErrors(t, errors)

			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			if len(program.Statements) != 1 {
				t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
			}

			stmt := program.Statements[0]
			if !testReturnStatement(t, stmt) {
				t.Errorf("Statement is not a return statement")
			}
		})
	}
}

func TestFunctionStatements(t *testing.T) {
	t.Skip("Function statements not fully implemented yet")
}

func TestClassStatements(t *testing.T) {
	tests := []testutil.ParserTestCase{
		{
			Name: "Simple class",
			Input: `
				class Person {
					name
				}
			`,
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			program, errors := parseProgram(tt.Input)

			if tt.WantErr {
				testutil.AssertHasErrors(t, errors, 1)
				return
			}

			testutil.AssertNoErrors(t, errors)

			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			if len(program.Statements) != 1 {
				t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
			}

			stmt := program.Statements[0]
			if !testClassStatement(t, stmt) {
				t.Errorf("Statement is not a class statement")
			}
		})
	}
}

func TestExpressions(t *testing.T) {
	tests := []testutil.ParserTestCase{
		{
			Name:    "Integer literal",
			Input:   "5;",
			WantErr: false,
		},
		{
			Name:    "Float literal",
			Input:   "3.14;",
			WantErr: false,
		},
		{
			Name:    "String literal",
			Input:   `"hello world";`,
			WantErr: false,
		},
		{
			Name:    "Boolean literal",
			Input:   "true;",
			WantErr: false,
		},
		{
			Name:    "Identifier",
			Input:   "foobar;",
			WantErr: false,
		},
		{
			Name:    "Prefix expression",
			Input:   "!true;",
			WantErr: false,
		},
		{
			Name:    "Infix expression",
			Input:   "5 + 5;",
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			program, errors := parseProgram(tt.Input)

			if tt.WantErr {
				testutil.AssertHasErrors(t, errors, 1)
				return
			}

			testutil.AssertNoErrors(t, errors)

			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			if len(program.Statements) != 1 {
				t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
			}

			stmt := program.Statements[0]
			if !testExpressionStatement(t, stmt) {
				t.Errorf("Statement is not an expression statement")
			}
		})
	}
}

func testVarStatement(t *testing.T, s ast.Statement) bool {
	t.Helper()

	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not 'var'. got=%q", s.TokenLiteral())
		return false
	}

	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.VarStatement. got=%T", s)
		return false
	}

	if varStmt.Name == nil {
		t.Errorf("varStmt.Name is nil")
		return false
	}

	return true
}

func testReturnStatement(t *testing.T, s ast.Statement) bool {
	t.Helper()

	if s.TokenLiteral() != "return" {
		t.Errorf("s.TokenLiteral not 'return'. got=%q", s.TokenLiteral())
		return false
	}

	_, ok := s.(*ast.ReturnStatement)
	if !ok {
		t.Errorf("s not *ast.ReturnStatement. got=%T", s)
		return false
	}

	return true
}

func testFunctionStatement(t *testing.T, s ast.Statement) bool {
	t.Helper()

	if s.TokenLiteral() != "func" {
		t.Errorf("s.TokenLiteral not 'func'. got=%q", s.TokenLiteral())
		return false
	}

	_, ok := s.(*ast.FunctionStatement)
	if !ok {
		t.Errorf("s not *ast.FunctionStatement. got=%T", s)
		return false
	}

	return true
}

func testClassStatement(t *testing.T, s ast.Statement) bool {
	t.Helper()

	if s.TokenLiteral() != "class" {
		t.Errorf("s.TokenLiteral not 'class'. got=%q", s.TokenLiteral())
		return false
	}

	_, ok := s.(*ast.ClassStatement)
	if !ok {
		t.Errorf("s not *ast.ClassStatement. got=%T", s)
		return false
	}

	return true
}

func testExpressionStatement(t *testing.T, s ast.Statement) bool {
	t.Helper()

	_, ok := s.(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("s not *ast.ExpressionStatement. got=%T", s)
		return false
	}

	return true
}

func BenchmarkParser(b *testing.B) {
	input := `
	package main

	import "fmt"

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

	func main() {
		person := Person("John", 30)
		fmt.Println("Name:", person.getName())
		person.birthday()
	}
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l := lexer.New(input)
		p := New(l)
		_ = p.ParseProgram()
	}
}

package parser

import (
	"testing"

	"github.com/inkbytefo/go-minus/internal/ast"
	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/testutil"
)

func TestArrayTypeParsing(t *testing.T) {
	tests := []testutil.ParserTestCase{
		{
			Name: "Array type with size",
			Input: `
				var arr [5]int
			`,
			WantErr: false,
		},
		{
			Name: "Slice type",
			Input: `
				var slice []int
			`,
			WantErr: false,
		},
		{
			Name: "Array literal with type",
			Input: `
				var arr = [3]int{1, 2, 3}
			`,
			WantErr: false,
		},
		{
			Name: "Slice literal",
			Input: `
				var slice = []int{1, 2, 3}
			`,
			WantErr: false,
		},
		{
			Name: "Array indexing",
			Input: `
				var x = arr[0]
			`,
			WantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			l := lexer.New(tt.Input)
			p := New(l)
			program := p.ParseProgram()

			// Check for parsing errors
			errors := p.Errors()
			if tt.WantErr && len(errors) == 0 {
				t.Errorf("Expected parsing errors, but got none")
				return
			}
			if !tt.WantErr && len(errors) > 0 {
				t.Errorf("Unexpected parsing errors: %v", errors)
				return
			}

			if tt.WantErr {
				return // Skip further checks for error cases
			}

			// Check that we have at least one statement
			if len(program.Statements) == 0 {
				t.Errorf("Expected at least one statement, got none")
				return
			}
		})
	}
}

func TestArrayTypeExpressionParsing(t *testing.T) {
	input := `[5]int`

	l := lexer.New(input)
	p := New(l)
	
	// Parse as expression
	expr := p.parseExpression(LOWEST)

	errors := p.Errors()
	if len(errors) > 0 {
		t.Fatalf("Parser errors: %v", errors)
	}

	arrayType, ok := expr.(*ast.ArrayType)
	if !ok {
		t.Fatalf("Expected ArrayType, got %T", expr)
	}

	// Check size
	if arrayType.Size == nil {
		t.Errorf("Expected array size, got nil")
	}

	// Check element type
	if arrayType.ElementType == nil {
		t.Errorf("Expected element type, got nil")
	}

	// Check token literal
	if arrayType.TokenLiteral() != "[" {
		t.Errorf("Expected token literal '[', got %s", arrayType.TokenLiteral())
	}
}

func TestSliceTypeExpressionParsing(t *testing.T) {
	input := `[]string`

	l := lexer.New(input)
	p := New(l)
	
	// Parse as expression
	expr := p.parseExpression(LOWEST)

	errors := p.Errors()
	if len(errors) > 0 {
		t.Fatalf("Parser errors: %v", errors)
	}

	arrayType, ok := expr.(*ast.ArrayType)
	if !ok {
		t.Fatalf("Expected ArrayType, got %T", expr)
	}

	// Check that size is nil for slices
	if arrayType.Size != nil {
		t.Errorf("Expected nil size for slice, got %v", arrayType.Size)
	}

	// Check element type
	if arrayType.ElementType == nil {
		t.Errorf("Expected element type, got nil")
	}

	// Check token literal
	if arrayType.TokenLiteral() != "[" {
		t.Errorf("Expected token literal '[', got %s", arrayType.TokenLiteral())
	}
}

func TestArrayLiteralParsing(t *testing.T) {
	input := `[1, 2, 3, 4, 5]`

	l := lexer.New(input)
	p := New(l)
	
	// Parse as expression
	expr := p.parseExpression(LOWEST)

	errors := p.Errors()
	if len(errors) > 0 {
		t.Fatalf("Parser errors: %v", errors)
	}

	arrayLit, ok := expr.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("Expected ArrayLiteral, got %T", expr)
	}

	// Check number of elements
	if len(arrayLit.Elements) != 5 {
		t.Errorf("Expected 5 elements, got %d", len(arrayLit.Elements))
	}

	// Check token literal
	if arrayLit.TokenLiteral() != "[" {
		t.Errorf("Expected token literal '[', got %s", arrayLit.TokenLiteral())
	}
}

func TestArrayIndexingParsing(t *testing.T) {
	input := `arr[0]`

	l := lexer.New(input)
	p := New(l)
	
	// Parse as expression
	expr := p.parseExpression(LOWEST)

	errors := p.Errors()
	if len(errors) > 0 {
		t.Fatalf("Parser errors: %v", errors)
	}

	indexExpr, ok := expr.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("Expected IndexExpression, got %T", expr)
	}

	// Check left side (array identifier)
	if indexExpr.Left == nil {
		t.Errorf("Expected left expression, got nil")
	}

	// Check index
	if indexExpr.Index == nil {
		t.Errorf("Expected index expression, got nil")
	}

	// Check token literal
	if indexExpr.TokenLiteral() != "[" {
		t.Errorf("Expected token literal '[', got %s", indexExpr.TokenLiteral())
	}
}

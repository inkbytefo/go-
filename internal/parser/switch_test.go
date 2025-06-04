package parser

import (
	"testing"

	"github.com/inkbytefo/go-minus/internal/ast"
	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/testutil"
)

func TestSwitchStatements(t *testing.T) {
	tests := []testutil.ParserTestCase{
		{
			Name: "Simple switch with cases",
			Input: `
				switch x {
				case 1:
					fmt.Println("one")
				case 2:
					fmt.Println("two")
				default:
					fmt.Println("other")
				}
			`,
			WantErr: false,
		},
		{
			Name: "Switch without tag",
			Input: `
				switch {
				case x > 0:
					fmt.Println("positive")
				case x < 0:
					fmt.Println("negative")
				default:
					fmt.Println("zero")
				}
			`,
			WantErr: false,
		},
		{
			Name: "Switch with multiple case values",
			Input: `
				switch x {
				case 1, 2, 3:
					fmt.Println("small")
				case 4, 5:
					fmt.Println("medium")
				default:
					fmt.Println("large")
				}
			`,
			WantErr: false,
		},
		{
			Name: "Switch without default",
			Input: `
				switch x {
				case 1:
					fmt.Println("one")
				case 2:
					fmt.Println("two")
				}
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

			// Check that the first statement is a switch statement
			stmt, ok := program.Statements[0].(*ast.SwitchStatement)
			if !ok {
				t.Errorf("Expected SwitchStatement, got %T", program.Statements[0])
				return
			}

			// Basic validation
			if stmt.TokenLiteral() != "switch" {
				t.Errorf("Expected token literal 'switch', got %s", stmt.TokenLiteral())
			}

			// Check that we have cases
			if len(stmt.Cases) == 0 {
				t.Errorf("Expected at least one case, got none")
			}
		})
	}
}

func TestSwitchStatementParsing(t *testing.T) {
	input := `
		switch x {
		case 1:
			y = 1
		case 2, 3:
			y = 2
		default:
			y = 0
		}
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	errors := p.Errors()
	if len(errors) > 0 {
		t.Fatalf("Parser errors: %v", errors)
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.SwitchStatement)
	if !ok {
		t.Fatalf("Expected SwitchStatement, got %T", program.Statements[0])
	}

	// Check switch tag
	if stmt.Tag == nil {
		t.Errorf("Expected switch tag, got nil")
	}

	// Check number of cases
	if len(stmt.Cases) != 3 {
		t.Errorf("Expected 3 cases, got %d", len(stmt.Cases))
	}

	// Check first case
	case1 := stmt.Cases[0]
	if len(case1.Values) != 1 {
		t.Errorf("Expected 1 value in first case, got %d", len(case1.Values))
	}

	// Check second case (multiple values)
	case2 := stmt.Cases[1]
	if len(case2.Values) != 2 {
		t.Errorf("Expected 2 values in second case, got %d", len(case2.Values))
	}

	// Check default case
	case3 := stmt.Cases[2]
	if len(case3.Values) != 0 {
		t.Errorf("Expected 0 values in default case, got %d", len(case3.Values))
	}
	if case3.TokenLiteral() != "default" {
		t.Errorf("Expected 'default' token literal, got %s", case3.TokenLiteral())
	}
}

func TestSwitchWithoutTag(t *testing.T) {
	input := `
		switch {
		case x > 0:
			fmt.Println("positive")
		case x == 0:
			fmt.Println("zero")
		default:
			fmt.Println("negative")
		}
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	errors := p.Errors()
	if len(errors) > 0 {
		t.Fatalf("Parser errors: %v", errors)
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.SwitchStatement)
	if !ok {
		t.Fatalf("Expected SwitchStatement, got %T", program.Statements[0])
	}

	// Check that tag is nil for tagless switch
	if stmt.Tag != nil {
		t.Errorf("Expected nil tag for tagless switch, got %T", stmt.Tag)
	}

	// Check number of cases
	if len(stmt.Cases) != 3 {
		t.Errorf("Expected 3 cases, got %d", len(stmt.Cases))
	}
}

// Package testutil provides utilities for testing the GO-Minus compiler.
package testutil

import (
	"strings"
	"testing"

	"github.com/inkbytefo/go-minus/internal/token"
)

// TestCase represents a test case for compiler components.
type TestCase struct {
	Name     string
	Input    string
	Expected any
	WantErr  bool
	ErrorMsg string
}

// LexerTestCase represents a test case for the lexer.
type LexerTestCase struct {
	Name     string
	Input    string
	Expected []token.Token
}

// ParserTestCase represents a test case for the parser.
type ParserTestCase struct {
	Name     string
	Input    string
	WantErr  bool
	ErrorMsg string
}

// SemanticTestCase represents a test case for semantic analysis.
type SemanticTestCase struct {
	Name     string
	Input    string
	WantErr  bool
	ErrorMsg string
}

// AssertNoErrors checks that no errors occurred during parsing.
func AssertNoErrors(t *testing.T, errors []string) {
	t.Helper()
	if len(errors) > 0 {
		t.Errorf("Expected no errors, but got %d errors:", len(errors))
		for _, err := range errors {
			t.Errorf("  %s", err)
		}
	}
}

// AssertHasErrors checks that errors occurred during parsing.
func AssertHasErrors(t *testing.T, errors []string, expectedCount int) {
	t.Helper()
	if len(errors) != expectedCount {
		t.Errorf("Expected %d errors, but got %d errors:", expectedCount, len(errors))
		for _, err := range errors {
			t.Errorf("  %s", err)
		}
	}
}

// AssertErrorContains checks that at least one error contains the expected message.
func AssertErrorContains(t *testing.T, errors []string, expectedMsg string) {
	t.Helper()
	for _, err := range errors {
		if strings.Contains(err, expectedMsg) {
			return
		}
	}
	t.Errorf("Expected error containing '%s', but got errors: %v", expectedMsg, errors)
}

// RunLexerTests runs a series of lexer test cases.
// Note: This function requires the lexer package to be imported in the test file.
func RunLexerTests(t *testing.T, tests []LexerTestCase, newLexerFunc func(string) interface{}) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// This is a placeholder - actual implementation should be in test files
			// to avoid import cycles
			t.Skip("RunLexerTests should be implemented in individual test files")
		})
	}
}

// CreateTestToken creates a token for testing purposes.
func CreateTestToken(tokenType token.TokenType, literal string, line, column int) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: literal,
		Line:    line,
		Column:  column,
		Pos:     0,
		End:     len(literal),
	}
}

// CreateTestPosition creates a position for testing purposes.
func CreateTestPosition(line, column int) token.Position {
	return token.Position{
		Line:   line,
		Column: column,
		Offset: 0,
	}
}

// AssertTokenEqual checks if two tokens are equal.
func AssertTokenEqual(t *testing.T, expected, actual token.Token) {
	t.Helper()
	if expected.Type != actual.Type {
		t.Errorf("Token type: expected %q, got %q", expected.Type, actual.Type)
	}
	if expected.Literal != actual.Literal {
		t.Errorf("Token literal: expected %q, got %q", expected.Literal, actual.Literal)
	}
	if expected.Line != actual.Line {
		t.Errorf("Token line: expected %d, got %d", expected.Line, actual.Line)
	}
	if expected.Column != actual.Column {
		t.Errorf("Token column: expected %d, got %d", expected.Column, actual.Column)
	}
}

// AssertPositionEqual checks if two positions are equal.
func AssertPositionEqual(t *testing.T, expected, actual token.Position) {
	t.Helper()
	if expected.Line != actual.Line {
		t.Errorf("Position line: expected %d, got %d", expected.Line, actual.Line)
	}
	if expected.Column != actual.Column {
		t.Errorf("Position column: expected %d, got %d", expected.Column, actual.Column)
	}
}

// BenchmarkHelper provides utilities for benchmark tests.
// Note: Benchmark functions should be implemented in individual test files
// to avoid import cycles.
type BenchmarkHelper struct {
	Input string
}

// NewBenchmarkHelper creates a new benchmark helper.
func NewBenchmarkHelper(input string) *BenchmarkHelper {
	return &BenchmarkHelper{Input: input}
}

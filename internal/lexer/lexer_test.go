package lexer

import (
	"testing"

	"github.com/inkbytefo/go-minus/internal/testutil"
	"github.com/inkbytefo/go-minus/internal/token"
)

func TestNextToken(t *testing.T) {
	tests := []testutil.LexerTestCase{
		{
			Name:  "Basic operators",
			Input: "=+(){},;",
			Expected: []token.Token{
				testutil.CreateTestToken(token.ASSIGN, "=", 1, 1),
				testutil.CreateTestToken(token.PLUS, "+", 1, 2),
				testutil.CreateTestToken(token.LPAREN, "(", 1, 3),
				testutil.CreateTestToken(token.RPAREN, ")", 1, 4),
				testutil.CreateTestToken(token.LBRACE, "{", 1, 5),
				testutil.CreateTestToken(token.RBRACE, "}", 1, 6),
				testutil.CreateTestToken(token.COMMA, ",", 1, 7),
				testutil.CreateTestToken(token.SEMICOLON, ";", 1, 8),
				testutil.CreateTestToken(token.EOF, "", 1, 9),
			},
		},
		{
			Name:  "Keywords",
			Input: "func var true false if else return class",
			Expected: []token.Token{
				testutil.CreateTestToken(token.FUNC, "func", 1, 1),
				testutil.CreateTestToken(token.VAR, "var", 1, 6),
				testutil.CreateTestToken(token.TRUE, "true", 1, 10),
				testutil.CreateTestToken(token.FALSE, "false", 1, 15),
				testutil.CreateTestToken(token.IF, "if", 1, 21),
				testutil.CreateTestToken(token.ELSE, "else", 1, 24),
				testutil.CreateTestToken(token.RETURN, "return", 1, 29),
				testutil.CreateTestToken(token.CLASS, "class", 1, 36),
				testutil.CreateTestToken(token.EOF, "", 1, 42),
			},
		},
		{
			Name:  "Identifiers and numbers",
			Input: "foobar 123 456.789",
			Expected: []token.Token{
				testutil.CreateTestToken(token.IDENT, "foobar", 1, 1),
				testutil.CreateTestToken(token.INT, "123", 1, 8),
				testutil.CreateTestToken(token.FLOAT, "456.789", 1, 12),
				testutil.CreateTestToken(token.EOF, "", 1, 20),
			},
		},
		{
			Name:  "String literals",
			Input: `"hello world" 'c'`,
			Expected: []token.Token{
				testutil.CreateTestToken(token.STRING, "hello world", 1, 1),
				testutil.CreateTestToken(token.CHAR, "c", 1, 15),
				testutil.CreateTestToken(token.EOF, "", 1, 18),
			},
		},
		{
			Name:  "Comparison operators",
			Input: "== != < > <= >=",
			Expected: []token.Token{
				testutil.CreateTestToken(token.EQ, "==", 1, 1),
				testutil.CreateTestToken(token.NOT_EQ, "!=", 1, 4),
				testutil.CreateTestToken(token.LT, "<", 1, 7),
				testutil.CreateTestToken(token.GT, ">", 1, 9),
				testutil.CreateTestToken(token.LTOEQ, "<=", 1, 11),
				testutil.CreateTestToken(token.GTOEQ, ">=", 1, 14),
				testutil.CreateTestToken(token.EOF, "", 1, 17),
			},
		},
		{
			Name:  "Logical operators",
			Input: "&& || !",
			Expected: []token.Token{
				testutil.CreateTestToken(token.LOGICAL_AND, "&&", 1, 1),
				testutil.CreateTestToken(token.LOGICAL_OR, "||", 1, 4),
				testutil.CreateTestToken(token.BANG, "!", 1, 7),
				testutil.CreateTestToken(token.EOF, "", 1, 8),
			},
		},
		{
			Name:  "Arithmetic operators",
			Input: "+ - * / %",
			Expected: []token.Token{
				testutil.CreateTestToken(token.PLUS, "+", 1, 1),
				testutil.CreateTestToken(token.MINUS, "-", 1, 3),
				testutil.CreateTestToken(token.ASTERISK, "*", 1, 5),
				testutil.CreateTestToken(token.SLASH, "/", 1, 7),
				testutil.CreateTestToken(token.MODULO, "%", 1, 9),
				testutil.CreateTestToken(token.EOF, "", 1, 10),
			},
		},
		{
			Name:  "Access modifiers",
			Input: "public private protected",
			Expected: []token.Token{
				testutil.CreateTestToken(token.PUBLIC, "public", 1, 1),
				testutil.CreateTestToken(token.PRIVATE, "private", 1, 8),
				testutil.CreateTestToken(token.PROTECTED, "protected", 1, 16),
				testutil.CreateTestToken(token.EOF, "", 1, 26),
			},
		},
		{
			Name:  "Exception handling",
			Input: "try catch finally throw",
			Expected: []token.Token{
				testutil.CreateTestToken(token.TRY, "try", 1, 1),
				testutil.CreateTestToken(token.CATCH, "catch", 1, 5),
				testutil.CreateTestToken(token.FINALLY, "finally", 1, 11),
				testutil.CreateTestToken(token.THROW, "throw", 1, 19),
				testutil.CreateTestToken(token.EOF, "", 1, 25),
			},
		},
		{
			Name:  "Template keywords",
			Input: "template this super",
			Expected: []token.Token{
				testutil.CreateTestToken(token.TEMPLATE, "template", 1, 1),
				testutil.CreateTestToken(token.THIS, "this", 1, 10),
				testutil.CreateTestToken(token.SUPER, "super", 1, 15),
				testutil.CreateTestToken(token.EOF, "", 1, 21),
			},
		},
	}

	// Run tests manually to avoid import cycle
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			l := New(tt.Input)

			for i, expectedToken := range tt.Expected {
				tok := l.NextToken()

				if tok.Type != expectedToken.Type {
					t.Errorf("Token %d: expected type %q, got %q", i, expectedToken.Type, tok.Type)
				}

				if tok.Literal != expectedToken.Literal {
					t.Errorf("Token %d: expected literal %q, got %q", i, expectedToken.Literal, tok.Literal)
				}
			}
		})
	}
}

func TestNextTokenWithComments(t *testing.T) {
	input := `
	// This is a single line comment
	var x = 5; // Another comment
	/* This is a
	   multi-line comment */
	var y = 10;
	`

	expected := []token.Token{
		testutil.CreateTestToken(token.VAR, "var", 3, 2),
		testutil.CreateTestToken(token.IDENT, "x", 3, 6),
		testutil.CreateTestToken(token.ASSIGN, "=", 3, 8),
		testutil.CreateTestToken(token.INT, "5", 3, 10),
		testutil.CreateTestToken(token.SEMICOLON, ";", 3, 11),
		testutil.CreateTestToken(token.VAR, "var", 6, 2),
		testutil.CreateTestToken(token.IDENT, "y", 6, 6),
		testutil.CreateTestToken(token.ASSIGN, "=", 6, 8),
		testutil.CreateTestToken(token.INT, "10", 6, 10),
		testutil.CreateTestToken(token.SEMICOLON, ";", 6, 12),
		testutil.CreateTestToken(token.EOF, "", 6, 13),
	}

	l := New(input)
	for i, expectedToken := range expected {
		tok := l.NextToken()

		if tok.Type != expectedToken.Type {
			t.Errorf("Token %d: expected type %q, got %q", i, expectedToken.Type, tok.Type)
		}

		if tok.Literal != expectedToken.Literal {
			t.Errorf("Token %d: expected literal %q, got %q", i, expectedToken.Literal, tok.Literal)
		}

		if tok.Line != expectedToken.Line {
			t.Errorf("Token %d: expected line %d, got %d", i, expectedToken.Line, tok.Line)
		}
	}
}

func TestNextTokenWithWhitespace(t *testing.T) {
	input := "   var    x   =   5   ;   "

	expected := []token.Token{
		testutil.CreateTestToken(token.VAR, "var", 1, 4),
		testutil.CreateTestToken(token.IDENT, "x", 1, 10),
		testutil.CreateTestToken(token.ASSIGN, "=", 1, 14),
		testutil.CreateTestToken(token.INT, "5", 1, 18),
		testutil.CreateTestToken(token.SEMICOLON, ";", 1, 22),
		testutil.CreateTestToken(token.EOF, "", 1, 26),
	}

	l := New(input)
	for i, expectedToken := range expected {
		tok := l.NextToken()

		if tok.Type != expectedToken.Type {
			t.Errorf("Token %d: expected type %q, got %q", i, expectedToken.Type, tok.Type)
		}

		if tok.Literal != expectedToken.Literal {
			t.Errorf("Token %d: expected literal %q, got %q", i, expectedToken.Literal, tok.Literal)
		}
	}
}

func TestIllegalTokens(t *testing.T) {
	input := "@#$"

	expected := []token.Token{
		testutil.CreateTestToken(token.ILLEGAL, "@", 1, 1),
		testutil.CreateTestToken(token.ILLEGAL, "#", 1, 2),
		testutil.CreateTestToken(token.ILLEGAL, "$", 1, 3),
		testutil.CreateTestToken(token.EOF, "", 1, 4),
	}

	l := New(input)
	for i, expectedToken := range expected {
		tok := l.NextToken()

		if tok.Type != expectedToken.Type {
			t.Errorf("Token %d: expected type %q, got %q", i, expectedToken.Type, tok.Type)
		}

		if tok.Literal != expectedToken.Literal {
			t.Errorf("Token %d: expected literal %q, got %q", i, expectedToken.Literal, tok.Literal)
		}
	}
}

func BenchmarkLexer(b *testing.B) {
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
		l := New(input)
		for {
			tok := l.NextToken()
			if tok.Type == token.EOF {
				break
			}
		}
	}
}

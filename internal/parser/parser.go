package parser

import (
	"github.com/inkbytefo/go-minus/internal/ast"
	"github.com/inkbytefo/go-minus/internal/lexer"
	"github.com/inkbytefo/go-minus/internal/token"
)

// Prefix ve infix ayrıştırma fonksiyonları için tip tanımları
type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

// Parser, token dizisini Soyut Sözdizimi Ağacı'na (AST) dönüştürür.
type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

// New, yeni bir Parser oluşturur.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Prefix ve infix ayrıştırma fonksiyonlarını kaydet
	p.registerPrefixFunctions()
	p.registerInfixFunctions()

	// İki token oku, böylece curToken ve peekToken ayarlanır.
	p.nextToken()
	p.nextToken()

	return p
}

// ParseProgram, GO-Minus programının tamamını ayrıştırır ve bir AST kök düğümü döndürür.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

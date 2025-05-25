package parser

import (
	"fmt"

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

	// Prefix ayrıştırma fonksiyonlarını kaydet
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.CHAR, p.parseCharLiteral)
	p.registerPrefix(token.TRUE, p.parseBooleanLiteral)
	p.registerPrefix(token.FALSE, p.parseBooleanLiteral)
	p.registerPrefix(token.NULL, p.parseNullLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNC, p.parseFunctionLiteral)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)
	p.registerPrefix(token.NEW, p.parseNewExpression)
	p.registerPrefix(token.TEMPLATE, p.parseTemplateExpression)
	p.registerPrefix(token.THIS, p.parseThisExpression)
	p.registerPrefix(token.SUPER, p.parseSuperExpression)

	// Infix ayrıştırma fonksiyonlarını kaydet
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.MODULO, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LTOEQ, p.parseInfixExpression)
	p.registerInfix(token.GTOEQ, p.parseInfixExpression)
	p.registerInfix(token.LOGICAL_AND, p.parseInfixExpression)
	p.registerInfix(token.LOGICAL_OR, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)
	p.registerInfix(token.DOT, p.parseMemberExpression)
	p.registerInfix(token.ARROW, p.parseMemberExpression)
	p.registerInfix(token.ASSIGN, p.parseAssignExpression)
	p.registerInfix(token.DEFINE, p.parseShortVarDeclExpression)
	p.registerInfix(token.INCREMENT, p.parsePostfixExpression)
	p.registerInfix(token.DECREMENT, p.parsePostfixExpression)

	// İki token oku, böylece curToken ve peekToken ayarlanır.
	p.nextToken()
	p.nextToken()

	return p
}

// registerPrefix, bir prefix ayrıştırma fonksiyonunu kaydeder.
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfix, bir infix ayrıştırma fonksiyonunu kaydeder.
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// nextToken bir sonraki token'ı alır.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// peekTokenIs, bir sonraki token'ın belirli bir türde olup olmadığını kontrol eder.
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// curTokenIs, mevcut token'ın belirli bir türde olup olmadığını kontrol eder.
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// expectPeek, bir sonraki token'ın belirli bir türde olmasını bekler.
// Eğer öyleyse, token'ı ilerletir ve true döndürür.
// Değilse, bir hata ekler ve false döndürür.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// peekError, beklenen token türü ile ilgili bir hata ekler.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Satır %d, Sütun %d: %s bekleniyordu, %s alındı",
		p.peekToken.Line, p.peekToken.Column, t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// Errors, ayrıştırma sırasında karşılaşılan hataları döndürür.
func (p *Parser) Errors() []string {
	return p.errors
}

// synchronize, hata durumunda parser'ı senkronize eder.
// Bu fonksiyon, hata durumunda parser'ın ilerlemesini sağlar ve
// bir sonraki ifadenin başlangıcına kadar token'ları atlar.
func (p *Parser) synchronize() {
	// Bir sonraki ifadenin başlangıcına kadar token'ları atla
	for !p.curTokenIs(token.EOF) {
		// Noktalı virgül, bir ifadenin sonunu belirtir
		if p.curTokenIs(token.SEMICOLON) {
			p.nextToken()
			return
		}

		// Aşağıdaki token'lar genellikle bir ifadenin başlangıcını belirtir
		switch p.peekToken.Type {
		case token.PACKAGE, token.IMPORT, token.FUNC, token.VAR, token.CONST,
			token.IF, token.FOR, token.WHILE, token.CLASS, token.RETURN,
			token.TRY, token.THROW, token.SCOPE:
			return
		}

		p.nextToken()
	}
}

// peekPrecedence, bir sonraki token'ın önceliğini döndürür.
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

// curPrecedence, mevcut token'ın önceliğini döndürür.
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
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

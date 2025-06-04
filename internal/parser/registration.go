package parser

import "github.com/inkbytefo/go-minus/internal/token"

// registerPrefixFunctions, tüm prefix ayrıştırma fonksiyonlarını kaydeder.
func (p *Parser) registerPrefixFunctions() {
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	
	// Literal parsing functions
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.CHAR, p.parseCharLiteral)
	p.registerPrefix(token.TRUE, p.parseBooleanLiteral)
	p.registerPrefix(token.FALSE, p.parseBooleanLiteral)
	p.registerPrefix(token.NULL, p.parseNullLiteral)
	
	// Prefix operators
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	
	// Grouping and collections
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)
	
	// Control flow and expressions
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNC, p.parseFunctionLiteral)
	
	// Object-oriented features
	p.registerPrefix(token.NEW, p.parseNewExpression)
	p.registerPrefix(token.THIS, p.parseThisExpression)
	p.registerPrefix(token.SUPER, p.parseSuperExpression)
	
	// Template system
	p.registerPrefix(token.TEMPLATE, p.parseTemplateExpression)
}

// registerInfixFunctions, tüm infix ayrıştırma fonksiyonlarını kaydeder.
func (p *Parser) registerInfixFunctions() {
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	
	// Arithmetic operators
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.MODULO, p.parseInfixExpression)
	
	// Comparison operators
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LTOEQ, p.parseInfixExpression)
	p.registerInfix(token.GTOEQ, p.parseInfixExpression)
	
	// Logical operators
	p.registerInfix(token.LOGICAL_AND, p.parseInfixExpression)
	p.registerInfix(token.LOGICAL_OR, p.parseInfixExpression)
	
	// Access and call operators
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)
	p.registerInfix(token.DOT, p.parseMemberExpression)
	p.registerInfix(token.ARROW, p.parseMemberExpression)
	
	// Assignment operators
	p.registerInfix(token.ASSIGN, p.parseAssignExpression)
	p.registerInfix(token.DEFINE, p.parseShortVarDeclExpression)
	
	// Postfix operators
	p.registerInfix(token.INCREMENT, p.parsePostfixExpression)
	p.registerInfix(token.DECREMENT, p.parsePostfixExpression)
}

// registerPrefix, bir prefix ayrıştırma fonksiyonunu kaydeder.
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfix, bir infix ayrıştırma fonksiyonunu kaydeder.
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

package parser

import (
	"fmt"
	"goplus/internal/ast"
	"goplus/internal/lexer"
	"goplus/internal/token"
	"strconv"
)

// Operatör öncelik seviyeleri
const (
	_ int = iota
	LOWEST
	ASSIGN      // =
	LOGICAL_OR  // ||
	LOGICAL_AND // &&
	EQUALS      // ==
	LESSGREATER // > veya <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X veya !X
	CALL        // myFunction(X)
	INDEX       // array[index]
	MEMBER      // foo.bar
)

// Operatör öncelik tablosu
var precedences = map[token.TokenType]int{
	token.ASSIGN:      ASSIGN,
	token.EQ:          EQUALS,
	token.NOT_EQ:      EQUALS,
	token.LT:          LESSGREATER,
	token.GT:          LESSGREATER,
	token.LTOEQ:       LESSGREATER,
	token.GTOEQ:       LESSGREATER,
	token.PLUS:        SUM,
	token.MINUS:       SUM,
	token.SLASH:       PRODUCT,
	token.ASTERISK:    PRODUCT,
	token.MODULO:      PRODUCT,
	token.LPAREN:      CALL,
	token.LBRACKET:    INDEX,
	token.DOT:         MEMBER,
	token.ARROW:       MEMBER,
	token.LOGICAL_AND: LOGICAL_AND,
	token.LOGICAL_OR:  LOGICAL_OR,
}

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

// ParseProgram, GO+ programının tamamını ayrıştırır ve bir AST kök düğümü döndürür.
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

// parseStatement, bir ifadeyi ayrıştırır.
func (p *Parser) parseStatement() ast.Statement {
	var stmt ast.Statement

	switch p.curToken.Type {
	case token.PACKAGE:
		stmt = p.parsePackageStatement()
	case token.IMPORT:
		stmt = p.parseImportStatement()
	case token.VAR:
		stmt = p.parseVarStatement()
	case token.CONST:
		stmt = p.parseConstStatement()
	case token.RETURN:
		stmt = p.parseReturnStatement()
	case token.IF:
		stmt = p.parseIfStatement()
	case token.FOR:
		stmt = p.parseForStatement()
	case token.WHILE:
		stmt = p.parseWhileStatement()
	case token.CLASS:
		stmt = p.parseClassStatement()
	case token.FUNC:
		if p.peekTokenIs(token.LPAREN) {
			stmt = p.parseMethodStatement()
		} else {
			stmt = p.parseFunctionStatement()
		}
	case token.TRY:
		stmt = p.parseTryCatchStatement()
	case token.THROW:
		stmt = p.parseThrowStatement()
	case token.SCOPE:
		stmt = p.parseScopeStatement()
	default:
		stmt = p.parseExpressionStatement()
	}

	// Hata durumunda senkronize et
	if len(p.errors) > 0 && stmt == nil {
		p.synchronize()
	}

	return stmt
}

// parseVarStatement, bir değişken tanımlama ifadesini ayrıştırır.
func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Opsiyonel tip
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		stmt.Type = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	} else if p.peekTokenIs(token.LBRACKET) {
		p.nextToken()
		// Dizi tipi
		// TODO: Dizi tipi ayrıştırma eklenecek
	} else if p.peekTokenIs(token.FUNC) {
		p.nextToken()
		// Fonksiyon tipi
		// TODO: Fonksiyon tipi ayrıştırma eklenecek
	} else if p.peekTokenIs(token.MAP) {
		p.nextToken()
		// Map tipi
		// TODO: Map tipi ayrıştırma eklenecek
	} else if p.peekTokenIs(token.CHAN) {
		p.nextToken()
		// Kanal tipi
		// TODO: Kanal tipi ayrıştırma eklenecek
	} else if p.peekTokenIs(token.INTERFACE) {
		p.nextToken()
		// Arayüz tipi
		// TODO: Arayüz tipi ayrıştırma eklenecek
	} else if p.peekTokenIs(token.STRUCT) {
		p.nextToken()
		// Struct tipi
		// TODO: Struct tipi ayrıştırma eklenecek
	}

	// Opsiyonel değer
	if p.peekTokenIs(token.ASSIGN) {
		p.nextToken()
		p.nextToken()
		stmt.Value = p.parseExpression(LOWEST)
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseConstStatement, bir sabit tanımlama ifadesini ayrıştırır.
func (p *Parser) parseConstStatement() *ast.ConstStatement {
	stmt := &ast.ConstStatement{Token: p.curToken}

	// Tek bir sabit
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

		// Opsiyonel tip
		if p.peekTokenIs(token.IDENT) {
			p.nextToken()
			stmt.Type = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		}

		if !p.expectPeek(token.ASSIGN) {
			return nil
		}

		p.nextToken()
		stmt.Value = p.parseExpression(LOWEST)
	} else if p.peekTokenIs(token.LPAREN) {
		// Çoklu sabit
		p.nextToken() // '(' token'ını atla

		// İlk sabit
		if p.peekTokenIs(token.IDENT) {
			p.nextToken()
			stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

			// Opsiyonel tip
			if p.peekTokenIs(token.IDENT) {
				p.nextToken()
				stmt.Type = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			}

			if !p.expectPeek(token.ASSIGN) {
				return nil
			}

			p.nextToken()
			stmt.Value = p.parseExpression(LOWEST)
		}

		// Diğer sabitler (şu anda tek bir sabit destekleniyor, çoklu sabit için genişletilebilir)
		for p.peekTokenIs(token.COMMA) {
			p.nextToken() // ',' token'ını atla

			if p.peekTokenIs(token.IDENT) {
				p.nextToken()
				// Çoklu sabit için burada bir dizi oluşturulabilir

				// Opsiyonel tip
				if p.peekTokenIs(token.IDENT) {
					p.nextToken()
				}

				if !p.expectPeek(token.ASSIGN) {
					return nil
				}

				p.nextToken()
				// Çoklu sabit için burada bir dizi oluşturulabilir
			}
		}

		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		p.peekError(token.IDENT)
		return nil
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement, bir dönüş ifadesini ayrıştırır.
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	if !p.curTokenIs(token.SEMICOLON) {
		stmt.ReturnValue = p.parseExpression(LOWEST)
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpressionStatement, bir ifade cümlesini ayrıştırır.
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpression, bir ifadeyi ayrıştırır.
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

// noPrefixParseFnError, bir prefix ayrıştırma fonksiyonu bulunamadığında bir hata ekler.
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("Satır %d, Sütun %d: %s için prefix ayrıştırma fonksiyonu bulunamadı",
		p.curToken.Line, p.curToken.Column, t)
	p.errors = append(p.errors, msg)
}

// parseIdentifier, bir tanımlayıcıyı ayrıştırır.
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parseIntegerLiteral, bir tamsayı değişmez değerini ayrıştırır.
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Satır %d, Sütun %d: %q bir tamsayıya dönüştürülemedi",
			p.curToken.Line, p.curToken.Column, p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

// parseFloatLiteral, bir ondalık sayı değişmez değerini ayrıştırır.
func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("Satır %d, Sütun %d: %q bir ondalık sayıya dönüştürülemedi",
			p.curToken.Line, p.curToken.Column, p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

// parseStringLiteral, bir string değişmez değerini ayrıştırır.
func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

// parseCharLiteral, bir karakter değişmez değerini ayrıştırır.
func (p *Parser) parseCharLiteral() ast.Expression {
	if len(p.curToken.Literal) != 1 {
		msg := fmt.Sprintf("Satır %d, Sütun %d: %q bir karakter değil",
			p.curToken.Line, p.curToken.Column, p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	return &ast.CharLiteral{Token: p.curToken, Value: rune(p.curToken.Literal[0])}
}

// parseBooleanLiteral, bir boolean değişmez değerini ayrıştırır.
func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &ast.BooleanLiteral{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

// parseNullLiteral, bir null değişmez değerini ayrıştırır.
func (p *Parser) parseNullLiteral() ast.Expression {
	return &ast.NullLiteral{Token: p.curToken}
}

// parsePrefixExpression, bir önek ifadesini ayrıştırır.
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

// parseInfixExpression, bir araek ifadesini ayrıştırır.
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

// parseGroupedExpression, bir gruplandırılmış ifadeyi ayrıştırır.
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

// parseIfExpression, bir if ifadesini ayrıştırır.
func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

// parseIfStatement, bir if ifadesini ayrıştırır ve bir ExpressionStatement olarak döndürür.
func (p *Parser) parseIfStatement() ast.Statement {
	return &ast.ExpressionStatement{
		Token:      p.curToken,
		Expression: p.parseIfExpression(),
	}
}

// parseBlockStatement, bir blok ifadesini ayrıştırır.
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

// parseFunctionLiteral, bir fonksiyon değişmez değerini ayrıştırır.
func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	// Fonksiyon adı (opsiyonel)
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		// Fonksiyon adı için bir alan eklenebilir
		// lit.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	// Opsiyonel dönüş tipi
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		lit.ReturnType = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	} else if p.peekTokenIs(token.LBRACKET) {
		p.nextToken()
		// Dizi dönüş tipi
		// TODO: Dizi tipi ayrıştırma eklenecek
	} else if p.peekTokenIs(token.FUNC) {
		p.nextToken()
		// Fonksiyon dönüş tipi
		// TODO: Fonksiyon tipi ayrıştırma eklenecek
	} else if p.peekTokenIs(token.MAP) {
		p.nextToken()
		// Map dönüş tipi
		// TODO: Map tipi ayrıştırma eklenecek
	} else if p.peekTokenIs(token.CHAN) {
		p.nextToken()
		// Kanal dönüş tipi
		// TODO: Kanal tipi ayrıştırma eklenecek
	} else if p.peekTokenIs(token.INTERFACE) {
		p.nextToken()
		// Arayüz dönüş tipi
		// TODO: Arayüz tipi ayrıştırma eklenecek
	} else if p.peekTokenIs(token.STRUCT) {
		p.nextToken()
		// Struct dönüş tipi
		// TODO: Struct tipi ayrıştırma eklenecek
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

// parseFunctionStatement, bir fonksiyon tanımını ayrıştırır.
func (p *Parser) parseFunctionStatement() ast.Statement {
	stmt := &ast.ExpressionStatement{
		Token: p.curToken,
	}

	// Fonksiyon adı
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()

		// Eğer bir sonraki token parantez ise, bu bir fonksiyon tanımıdır
		if p.peekTokenIs(token.LPAREN) {
			funcLit := &ast.FunctionLiteral{Token: stmt.Token}
			// Fonksiyon adı için bir alan eklenebilir
			// funcLit.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

			p.nextToken() // '(' token'ını al

			funcLit.Parameters = p.parseFunctionParameters()

			// Opsiyonel dönüş tipi
			if p.peekTokenIs(token.IDENT) {
				p.nextToken()
				funcLit.ReturnType = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			}

			if !p.expectPeek(token.LBRACE) {
				return nil
			}

			funcLit.Body = p.parseBlockStatement()

			stmt.Expression = funcLit
		} else {
			// Değişken tanımı veya başka bir ifade
			p.nextToken()
			stmt.Expression = p.parseExpression(LOWEST)
		}
	} else {
		// Anonim fonksiyon
		stmt.Expression = p.parseFunctionLiteral()
	}

	return stmt
}

// parseFunctionParameters, bir fonksiyonun parametrelerini ayrıştırır.
func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	// İlk parametre
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Parametre tipi (opsiyonel)
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		// Parametre tipi için bir alan eklenebilir
		// ident.Type = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	identifiers = append(identifiers, ident)

	// Diğer parametreler
	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // ',' token'ını atla
		p.nextToken() // Parametre adını al

		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

		// Parametre tipi (opsiyonel)
		if p.peekTokenIs(token.IDENT) {
			p.nextToken()
			// Parametre tipi için bir alan eklenebilir
			// ident.Type = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		}

		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

// parseCallExpression, bir fonksiyon çağrısını ayrıştırır.
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

// parseExpressionList, bir ifade listesini ayrıştırır.
func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

// parseArrayLiteral, bir dizi değişmez değerini ayrıştırır.
func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}

// parseIndexExpression, bir dizin ifadesini ayrıştırır.
func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.curToken, Left: left}

	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return exp
}

// parseHashLiteral, bir hash değişmez değerini ayrıştırır.
func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	if p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		return hash
	}

	p.nextToken()
	key := p.parseExpression(LOWEST)

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.nextToken()
	value := p.parseExpression(LOWEST)

	hash.Pairs[key] = value

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		key := p.parseExpression(LOWEST)

		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken()
		value := p.parseExpression(LOWEST)

		hash.Pairs[key] = value
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return hash
}

// parseMemberExpression, bir üye erişim ifadesini ayrıştırır.
func (p *Parser) parseMemberExpression(object ast.Expression) ast.Expression {
	exp := &ast.MemberExpression{
		Token:  p.curToken,
		Object: object,
	}

	p.nextToken()

	// Üye adı
	if p.curTokenIs(token.IDENT) {
		exp.Member = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	} else if p.curTokenIs(token.THIS) {
		exp.Member = p.parseThisExpression()
	} else if p.curTokenIs(token.SUPER) {
		exp.Member = p.parseSuperExpression()
	} else {
		// Diğer ifadeler (örneğin, metot çağrısı)
		exp.Member = p.parseExpression(MEMBER)
	}

	// Metot çağrısı
	if p.peekTokenIs(token.LPAREN) {
		p.nextToken()
		exp.Member = p.parseCallExpression(exp.Member)
	}

	return exp
}

// parseAssignExpression, bir atama ifadesini ayrıştırır.
func (p *Parser) parseAssignExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(precedence)

	return exp
}

// parseShortVarDeclExpression, bir kısa değişken tanımlama ifadesini ayrıştırır.
// Örnek: x := 5
func (p *Parser) parseShortVarDeclExpression(left ast.Expression) ast.Expression {
	// Kısa değişken tanımlama, bir atama ifadesi gibi ayrıştırılır
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(precedence)

	return exp
}

// parseForStatement, bir for döngüsünü ayrıştırır.
func (p *Parser) parseForStatement() *ast.ForStatement {
	stmt := &ast.ForStatement{Token: p.curToken}

	p.nextToken()

	// C tarzı for döngüsü: for i := 0; i < 10; i++ { ... }
	if !p.curTokenIs(token.LBRACE) {
		// Başlangıç ifadesi
		if !p.curTokenIs(token.SEMICOLON) {
			stmt.Init = p.parseStatement()
		}

		if !p.curTokenIs(token.SEMICOLON) {
			if !p.expectPeek(token.SEMICOLON) {
				return nil
			}
		}

		p.nextToken()

		// Koşul
		if !p.curTokenIs(token.SEMICOLON) {
			stmt.Condition = p.parseExpression(LOWEST)
		}

		if !p.curTokenIs(token.SEMICOLON) {
			if !p.expectPeek(token.SEMICOLON) {
				return nil
			}
		}

		p.nextToken()

		// Sonrası ifadesi
		if !p.curTokenIs(token.LBRACE) {
			stmt.Post = p.parseStatement()
		}
	}

	if !p.curTokenIs(token.LBRACE) {
		if !p.expectPeek(token.LBRACE) {
			return nil
		}
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// parseWhileStatement, bir while döngüsünü ayrıştırır.
func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curToken}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// parseClassStatement, bir sınıf tanımını ayrıştırır.
func (p *Parser) parseClassStatement() *ast.ClassStatement {
	stmt := &ast.ClassStatement{Token: p.curToken}

	// Sınıf adı
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Opsiyonel şablon parametreleri
	if p.peekTokenIs(token.LT) {
		p.nextToken() // '<' token'ını atla

		// İlk şablon parametresi
		if p.peekTokenIs(token.IDENT) {
			p.nextToken()
			// Şablon parametreleri için bir alan eklenebilir
			// param := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			// stmt.TemplateParameters = append(stmt.TemplateParameters, param)

			// Diğer şablon parametreleri
			for p.peekTokenIs(token.COMMA) {
				p.nextToken() // ',' token'ını atla

				if p.peekTokenIs(token.IDENT) {
					p.nextToken()
					// Şablon parametreleri için bir alan eklenebilir
					// param := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
					// stmt.TemplateParameters = append(stmt.TemplateParameters, param)
				}
			}
		}

		if !p.expectPeek(token.GT) {
			return nil
		}
	}

	// Opsiyonel kalıtım
	if p.peekTokenIs(token.EXTENDS) {
		p.nextToken()

		if !p.expectPeek(token.IDENT) {
			return nil
		}

		stmt.Extends = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

		// Opsiyonel şablon argümanları
		if p.peekTokenIs(token.LT) {
			p.nextToken() // '<' token'ını atla

			// İlk şablon argümanı
			if p.peekTokenIs(token.IDENT) {
				p.nextToken()
				// Şablon argümanları için bir alan eklenebilir
				// arg := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
				// stmt.ExtendsTemplateArguments = append(stmt.ExtendsTemplateArguments, arg)

				// Diğer şablon argümanları
				for p.peekTokenIs(token.COMMA) {
					p.nextToken() // ',' token'ını atla

					if p.peekTokenIs(token.IDENT) {
						p.nextToken()
						// Şablon argümanları için bir alan eklenebilir
						// arg := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
						// stmt.ExtendsTemplateArguments = append(stmt.ExtendsTemplateArguments, arg)
					}
				}
			}

			if !p.expectPeek(token.GT) {
				return nil
			}
		}
	}

	// Opsiyonel arayüz uygulamaları
	if p.peekTokenIs(token.IMPLEMENTS) {
		p.nextToken()

		if !p.expectPeek(token.IDENT) {
			return nil
		}

		impl := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		stmt.Implements = append(stmt.Implements, impl)

		// Opsiyonel şablon argümanları
		if p.peekTokenIs(token.LT) {
			p.nextToken() // '<' token'ını atla

			// İlk şablon argümanı
			if p.peekTokenIs(token.IDENT) {
				p.nextToken()
				// Şablon argümanları için bir alan eklenebilir
				// arg := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
				// impl.TemplateArguments = append(impl.TemplateArguments, arg)

				// Diğer şablon argümanları
				for p.peekTokenIs(token.COMMA) {
					p.nextToken() // ',' token'ını atla

					if p.peekTokenIs(token.IDENT) {
						p.nextToken()
						// Şablon argümanları için bir alan eklenebilir
						// arg := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
						// impl.TemplateArguments = append(impl.TemplateArguments, arg)
					}
				}
			}

			if !p.expectPeek(token.GT) {
				return nil
			}
		}

		// Diğer arayüz uygulamaları
		for p.peekTokenIs(token.COMMA) {
			p.nextToken()

			if !p.expectPeek(token.IDENT) {
				return nil
			}

			impl := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			stmt.Implements = append(stmt.Implements, impl)

			// Opsiyonel şablon argümanları
			if p.peekTokenIs(token.LT) {
				p.nextToken() // '<' token'ını atla

				// İlk şablon argümanı
				if p.peekTokenIs(token.IDENT) {
					p.nextToken()
					// Şablon argümanları için bir alan eklenebilir
					// arg := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
					// impl.TemplateArguments = append(impl.TemplateArguments, arg)

					// Diğer şablon argümanları
					for p.peekTokenIs(token.COMMA) {
						p.nextToken() // ',' token'ını atla

						if p.peekTokenIs(token.IDENT) {
							p.nextToken()
							// Şablon argümanları için bir alan eklenebilir
							// arg := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
							// impl.TemplateArguments = append(impl.TemplateArguments, arg)
						}
					}
				}

				if !p.expectPeek(token.GT) {
					return nil
				}
			}
		}
	}

	// Sınıf gövdesi
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	// Sınıf üyeleri
	stmt.Body = p.parseClassBody()

	return stmt
}

// parseClassBody, bir sınıf gövdesini ayrıştırır.
func (p *Parser) parseClassBody() *ast.BlockStatement {
	body := &ast.BlockStatement{Token: p.curToken}
	body.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		// Erişim belirleyicileri
		if p.curTokenIs(token.PUBLIC) || p.curTokenIs(token.PRIVATE) || p.curTokenIs(token.PROTECTED) {
			// Erişim belirleyicisi için bir alan eklenebilir
			// var accessModifier token.TokenType = p.curToken.Type
			p.nextToken()
		}

		// Üye değişkenler
		if p.curTokenIs(token.VAR) {
			stmt := p.parseVarStatement()
			// Erişim belirleyicisi için bir alan eklenebilir
			// stmt.AccessModifier = accessModifier
			body.Statements = append(body.Statements, stmt)
		} else if p.curTokenIs(token.CONST) {
			// Sabit üyeler
			stmt := p.parseConstStatement()
			// Erişim belirleyicisi için bir alan eklenebilir
			// stmt.AccessModifier = accessModifier
			body.Statements = append(body.Statements, stmt)
		} else if p.curTokenIs(token.FUNC) {
			// Metotlar
			stmt := p.parseMethodStatement()
			// Erişim belirleyicisi için bir alan eklenebilir
			// stmt.AccessModifier = accessModifier
			body.Statements = append(body.Statements, stmt)
		} else {
			// Diğer ifadeler
			stmt := p.parseStatement()
			if stmt != nil {
				body.Statements = append(body.Statements, stmt)
			}
		}

		p.nextToken()
	}

	return body
}

// parseMethodStatement, bir metot tanımını ayrıştırır.
func (p *Parser) parseMethodStatement() *ast.MethodStatement {
	stmt := &ast.MethodStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Receiver = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	stmt.Parameters = p.parseFunctionParameters()

	// Opsiyonel dönüş tipi
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		// TODO: Tip ayrıştırma eklenecek
		// stmt.ReturnType = p.parseType()
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// parseTryCatchStatement, bir try-catch ifadesini ayrıştırır.
func (p *Parser) parseTryCatchStatement() *ast.TryCatchStatement {
	stmt := &ast.TryCatchStatement{Token: p.curToken}

	// Try bloğu
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Try = p.parseBlockStatement()

	// En az bir catch bloğu olmalı
	if !p.peekTokenIs(token.CATCH) {
		msg := fmt.Sprintf("Satır %d, Sütun %d: try bloğundan sonra catch bloğu bekleniyor",
			p.curToken.Line, p.curToken.Column)
		p.errors = append(p.errors, msg)
		return nil
	}

	// Catch blokları
	for p.peekTokenIs(token.CATCH) {
		p.nextToken()
		catch := &ast.CatchClause{Token: p.curToken}

		// Opsiyonel parametre
		if p.peekTokenIs(token.LPAREN) {
			p.nextToken()

			if !p.expectPeek(token.IDENT) {
				return nil
			}

			catch.Parameter = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

			// Opsiyonel tip
			if p.peekTokenIs(token.IDENT) {
				p.nextToken()
				catch.Type = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			} else if p.peekTokenIs(token.LBRACKET) {
				p.nextToken()
				// Dizi tipi
				// TODO: Dizi tipi ayrıştırma eklenecek
			} else if p.peekTokenIs(token.FUNC) {
				p.nextToken()
				// Fonksiyon tipi
				// TODO: Fonksiyon tipi ayrıştırma eklenecek
			} else if p.peekTokenIs(token.MAP) {
				p.nextToken()
				// Map tipi
				// TODO: Map tipi ayrıştırma eklenecek
			} else if p.peekTokenIs(token.CHAN) {
				p.nextToken()
				// Kanal tipi
				// TODO: Kanal tipi ayrıştırma eklenecek
			} else if p.peekTokenIs(token.INTERFACE) {
				p.nextToken()
				// Arayüz tipi
				// TODO: Arayüz tipi ayrıştırma eklenecek
			} else if p.peekTokenIs(token.STRUCT) {
				p.nextToken()
				// Struct tipi
				// TODO: Struct tipi ayrıştırma eklenecek
			}

			if !p.expectPeek(token.RPAREN) {
				return nil
			}
		}

		// Catch bloğu gövdesi
		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		catch.Body = p.parseBlockStatement()
		stmt.Catches = append(stmt.Catches, catch)
	}

	// Opsiyonel finally bloğu
	if p.peekTokenIs(token.FINALLY) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		stmt.Finally = p.parseBlockStatement()
	}

	return stmt
}

// parseThrowStatement, bir throw ifadesini ayrıştırır.
func (p *Parser) parseThrowStatement() *ast.ThrowStatement {
	stmt := &ast.ThrowStatement{Token: p.curToken}

	p.nextToken()

	// Fırlatılacak ifade
	stmt.Value = p.parseExpression(LOWEST)

	// Opsiyonel noktalı virgül
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseThisExpression, bir this ifadesini ayrıştırır.
func (p *Parser) parseThisExpression() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parseSuperExpression, bir super ifadesini ayrıştırır.
func (p *Parser) parseSuperExpression() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parseScopeStatement, bir scope ifadesini ayrıştırır.
func (p *Parser) parseScopeStatement() *ast.ScopeStatement {
	stmt := &ast.ScopeStatement{Token: p.curToken}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// parseNewExpression, bir new ifadesini ayrıştırır.
func (p *Parser) parseNewExpression() ast.Expression {
	exp := &ast.NewExpression{Token: p.curToken}

	p.nextToken()
	exp.Class = p.parseExpression(CALL)

	if p.peekTokenIs(token.LPAREN) {
		p.nextToken()
		exp.Arguments = p.parseExpressionList(token.RPAREN)
	}

	return exp
}

// parseTemplateExpression, bir şablon ifadesini ayrıştırır.
func (p *Parser) parseTemplateExpression() ast.Expression {
	exp := &ast.TemplateExpression{Token: p.curToken}

	if !p.expectPeek(token.LT) {
		return nil
	}

	// Şablon parametreleri
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// İlk parametre
	param := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Parametre tipi (opsiyonel)
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		// Parametre tipi için bir alan eklenebilir
		// param.Type = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	exp.Parameters = append(exp.Parameters, param)

	// Diğer parametreler
	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // ',' token'ını atla

		if !p.expectPeek(token.IDENT) {
			return nil
		}

		param := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

		// Parametre tipi (opsiyonel)
		if p.peekTokenIs(token.IDENT) {
			p.nextToken()
			// Parametre tipi için bir alan eklenebilir
			// param.Type = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		}

		exp.Parameters = append(exp.Parameters, param)
	}

	if !p.expectPeek(token.GT) {
		return nil
	}

	// Şablonun gövdesi
	p.nextToken()

	// Şablon gövdesi bir fonksiyon, sınıf veya başka bir ifade olabilir
	if p.curTokenIs(token.FUNC) {
		exp.Body = p.parseFunctionLiteral()
	} else if p.curTokenIs(token.CLASS) {
		// Sınıf ifadesi için bir alan eklenebilir
		// exp.Body = p.parseClassExpression()
		exp.Body = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	} else {
		exp.Body = p.parseExpression(LOWEST)
	}

	return exp
}

// parsePackageStatement, bir paket bildirimini ayrıştırır.
func (p *Parser) parsePackageStatement() ast.Statement {
	stmt := &ast.PackageStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseImportStatement, bir import bildirimini ayrıştırır.
func (p *Parser) parseImportStatement() ast.Statement {
	stmt := &ast.ImportStatement{Token: p.curToken}

	// Tek bir import
	if p.peekTokenIs(token.STRING) {
		p.nextToken()
		stmt.Path = &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
	} else if p.peekTokenIs(token.LPAREN) {
		// Çoklu import
		p.nextToken() // '(' token'ını atla

		// İlk import
		if p.peekTokenIs(token.STRING) {
			p.nextToken()
			stmt.Path = &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
		}

		// Diğer importlar (şu anda tek bir import destekleniyor, çoklu import için genişletilebilir)
		for p.peekTokenIs(token.COMMA) {
			p.nextToken() // ',' token'ını atla

			if p.peekTokenIs(token.STRING) {
				p.nextToken()
				// Çoklu import için burada bir dizi oluşturulabilir
			}
		}

		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		p.peekError(token.STRING)
		return nil
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

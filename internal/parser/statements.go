package parser

import (
	"github.com/inkbytefo/go-minus/internal/ast"
	"github.com/inkbytefo/go-minus/internal/token"
)

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

// parsePackageStatement, bir package bildirimini ayrıştırır.
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

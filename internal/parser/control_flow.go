package parser

import (
	"fmt"

	"github.com/inkbytefo/go-minus/internal/ast"
	"github.com/inkbytefo/go-minus/internal/token"
)

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
			p.nextToken() // '(' token'ını atla

			if p.peekTokenIs(token.IDENT) {
				p.nextToken()
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

// parseScopeStatement, bir scope ifadesini ayrıştırır.
func (p *Parser) parseScopeStatement() *ast.ScopeStatement {
	stmt := &ast.ScopeStatement{Token: p.curToken}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

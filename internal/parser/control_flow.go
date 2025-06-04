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
func (p *Parser) parseForStatement() ast.Statement {
	stmt := &ast.ForStatement{Token: p.curToken}

	p.nextToken()

	// For loop türlerini belirle
	if !p.curTokenIs(token.LBRACE) {
		// İlk ifadeyi parse et
		firstExpr := p.parseExpressionUntil(LOWEST, token.SEMICOLON)

		if p.curTokenIs(token.SEMICOLON) {
			// C-style for loop: for init; condition; post { ... }
			stmt.Init = &ast.ExpressionStatement{
				Token:      p.curToken,
				Expression: firstExpr,
			}

			p.nextToken() // semicolon'u atla

			// Condition
			if !p.curTokenIs(token.SEMICOLON) {
				stmt.Condition = p.parseExpressionUntil(LOWEST, token.SEMICOLON)
			}

			if !p.expectPeek(token.SEMICOLON) {
				return nil
			}

			p.nextToken() // semicolon'u atla

			// Post
			if !p.curTokenIs(token.LBRACE) {
				postExpr := p.parseExpressionUntil(LOWEST, token.LBRACE)
				stmt.Post = &ast.ExpressionStatement{
					Token:      p.curToken,
					Expression: postExpr,
				}
			}
		} else {
			// While-style for loop: for condition { ... }
			stmt.Condition = firstExpr
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
func (p *Parser) parseWhileStatement() ast.Statement {
	stmt := &ast.WhileStatement{Token: p.curToken}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// parseSwitchStatement, bir switch ifadesini ayrıştırır.
func (p *Parser) parseSwitchStatement() ast.Statement {
	stmt := &ast.SwitchStatement{Token: p.curToken}

	p.nextToken()

	// Opsiyonel switch tag (ifade)
	if !p.curTokenIs(token.LBRACE) {
		stmt.Tag = p.parseExpression(LOWEST)
		p.nextToken()
	}

	if !p.curTokenIs(token.LBRACE) {
		p.reportUnexpectedTokenWithMessage("{ bekleniyor")
		return nil
	}

	// Case clause'ları parse et
	for p.peekTokenIs(token.CASE) || p.peekTokenIs(token.DEFAULT) {
		p.nextToken()

		caseClause := p.parseCaseClause()
		if caseClause != nil {
			stmt.Cases = append(stmt.Cases, caseClause)
		}
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return stmt
}

// parseCaseClause, bir case veya default clause'unu ayrıştırır.
func (p *Parser) parseCaseClause() *ast.CaseClause {
	clause := &ast.CaseClause{Token: p.curToken}

	if p.curTokenIs(token.CASE) {
		p.nextToken()

		// Case değerlerini parse et (virgülle ayrılmış)
		clause.Values = append(clause.Values, p.parseExpression(LOWEST))

		for p.peekTokenIs(token.COMMA) {
			p.nextToken() // comma'yı atla
			p.nextToken()
			clause.Values = append(clause.Values, p.parseExpression(LOWEST))
		}
	}
	// DEFAULT için Values nil kalır

	if !p.expectPeek(token.COLON) {
		return nil
	}

	// Case body'sini parse et - statement'ları topla
	for p.peekTokenIs(token.IDENT) || p.peekTokenIs(token.RETURN) ||
		p.peekTokenIs(token.IF) || p.peekTokenIs(token.FOR) ||
		p.peekTokenIs(token.WHILE) || p.peekTokenIs(token.VAR) ||
		p.peekTokenIs(token.CONST) || p.peekTokenIs(token.FUNC) {

		p.nextToken()
		stmt := p.parseStatement()
		if stmt != nil {
			clause.Body = append(clause.Body, stmt)
		}
	}

	return clause
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

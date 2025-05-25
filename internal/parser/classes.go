package parser

import (
	"github.com/inkbytefo/go-minus/internal/ast"
	"github.com/inkbytefo/go-minus/internal/token"
)

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

// parseThisExpression, bir this ifadesini ayrıştırır.
func (p *Parser) parseThisExpression() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parseSuperExpression, bir super ifadesini ayrıştırır.
func (p *Parser) parseSuperExpression() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
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

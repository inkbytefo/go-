package parser

import (
	"github.com/inkbytefo/go-minus/internal/ast"
	"github.com/inkbytefo/go-minus/internal/token"
)

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
	funcStmt := &ast.FunctionStatement{
		Token: p.curToken,
	}

	// Fonksiyon adı
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	funcStmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Parametreler
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	funcStmt.Parameters = p.parseFunctionParameters()

	// Opsiyonel dönüş tipi
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		funcStmt.ReturnType = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	// Fonksiyon gövdesi
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	funcStmt.Body = p.parseBlockStatement()

	return funcStmt
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
	exp.Arguments = p.parseCallArguments()
	return exp
}

// parseCallArguments, fonksiyon çağrısı argümanlarını ayrıştırır.
func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
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

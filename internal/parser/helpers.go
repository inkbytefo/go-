package parser

import (
	"fmt"

	"github.com/inkbytefo/go-minus/internal/token"
)

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

// skipSemicolon, opsiyonel semicolon'u atlar.
func (p *Parser) skipSemicolon() {
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
}

// expectToken, belirli bir token türünü bekler ve hata kontrolü yapar.
func (p *Parser) expectToken(t token.TokenType) bool {
	if p.curTokenIs(t) {
		return true
	}

	msg := fmt.Sprintf("Satır %d, Sütun %d: %s bekleniyordu, %s alındı",
		p.curToken.Line, p.curToken.Column, t, p.curToken.Type)
	p.errors = append(p.errors, msg)
	return false
}

// isAtEnd, parser'ın dosyanın sonuna gelip gelmediğini kontrol eder.
func (p *Parser) isAtEnd() bool {
	return p.curTokenIs(token.EOF)
}

// advance, parser'ı bir token ilerletir ve eski token'ı döndürür.
func (p *Parser) advance() token.Token {
	previous := p.curToken
	p.nextToken()
	return previous
}

// match, mevcut token'ın verilen türlerden biriyle eşleşip eşleşmediğini kontrol eder.
func (p *Parser) match(types ...token.TokenType) bool {
	for _, t := range types {
		if p.curTokenIs(t) {
			p.nextToken()
			return true
		}
	}
	return false
}

// check, mevcut token'ın belirli bir türde olup olmadığını kontrol eder (token'ı ilerletmez).
func (p *Parser) check(t token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.curTokenIs(t)
}

// previous, önceki token'ı döndürür (bu implementasyonda mevcut değil, gelecekte eklenebilir).
func (p *Parser) previous() token.Token {
	// Bu fonksiyon şu anda kullanılmıyor, ancak gelecekte token history için eklenebilir
	return p.curToken
}

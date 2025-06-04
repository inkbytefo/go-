package parser

import (
	"fmt"

	"github.com/inkbytefo/go-minus/internal/token"
)

// Errors, ayrıştırma sırasında karşılaşılan hataları döndürür.
func (p *Parser) Errors() []string {
	return p.errors
}

// addError, parser'a yeni bir hata ekler.
func (p *Parser) addError(msg string) {
	p.errors = append(p.errors, msg)
}

// addErrorf, formatted bir hata mesajı ekler.
func (p *Parser) addErrorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	p.errors = append(p.errors, msg)
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

// reportUnexpectedToken, beklenmeyen token hatası rapor eder.
func (p *Parser) reportUnexpectedToken(expected, actual token.TokenType) {
	msg := fmt.Sprintf("Satır %d, Sütun %d: %s bekleniyordu, %s alındı",
		p.curToken.Line, p.curToken.Column, expected, actual)
	p.addError(msg)
}

// reportUnexpectedTokenWithMessage, özel mesajla beklenmeyen token hatası rapor eder.
func (p *Parser) reportUnexpectedTokenWithMessage(message string) {
	msg := fmt.Sprintf("Satır %d, Sütun %d: %s",
		p.curToken.Line, p.curToken.Column, message)
	p.addError(msg)
}

// reportMissingToken, eksik token hatası rapor eder.
func (p *Parser) reportMissingToken(expected token.TokenType) {
	msg := fmt.Sprintf("Satır %d, Sütun %d: %s eksik",
		p.curToken.Line, p.curToken.Column, expected)
	p.addError(msg)
}

// reportInvalidSyntax, geçersiz syntax hatası rapor eder.
func (p *Parser) reportInvalidSyntax(context string) {
	msg := fmt.Sprintf("Satır %d, Sütun %d: %s içinde geçersiz syntax",
		p.curToken.Line, p.curToken.Column, context)
	p.addError(msg)
}

// reportSemanticError, semantik hata rapor eder.
func (p *Parser) reportSemanticError(message string) {
	msg := fmt.Sprintf("Satır %d, Sütun %d: Semantik hata: %s",
		p.curToken.Line, p.curToken.Column, message)
	p.addError(msg)
}

// hasErrors, parser'da hata olup olmadığını kontrol eder.
func (p *Parser) hasErrors() bool {
	return len(p.errors) > 0
}

// clearErrors, tüm hataları temizler.
func (p *Parser) clearErrors() {
	p.errors = []string{}
}

// getLastError, son hatayı döndürür.
func (p *Parser) getLastError() string {
	if len(p.errors) == 0 {
		return ""
	}
	return p.errors[len(p.errors)-1]
}

// errorCount, toplam hata sayısını döndürür.
func (p *Parser) errorCount() int {
	return len(p.errors)
}

// recoverFromPanic, panic durumunda parser'ı kurtarır.
func (p *Parser) recoverFromPanic() {
	if r := recover(); r != nil {
		msg := fmt.Sprintf("Satır %d, Sütun %d: Parser panic: %v",
			p.curToken.Line, p.curToken.Column, r)
		p.addError(msg)
		p.synchronize()
	}
}

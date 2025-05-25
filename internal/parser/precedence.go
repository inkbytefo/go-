package parser

import "github.com/inkbytefo/go-minus/internal/token"

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

package lexer

import (
	"github.com/inkbytefo/go-minus/internal/token"
)

// Lexer holds the state of the scanner.
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line number
	column       int  // current column number
	startPos     int  // token başlangıç pozisyonu
	startLine    int  // token başlangıç satırı
	startColumn  int  // token başlangıç sütunu
}

// New creates a new Lexer.
func New(input string) *Lexer {
	l := &Lexer{
		input:       input,
		line:        1,
		column:      1,
		startLine:   1,
		startColumn: 1,
	}
	l.readChar() // Initialize l.ch, l.position, and l.readPosition
	return l
}

// readChar gives us the next character and advances our position in the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" character, signifying EOF or not read anything yet
	} else {
		l.ch = l.input[l.readPosition]
	}

	// Satır ve sütun numaralarını güncelle
	if l.ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}

	l.position = l.readPosition
	l.readPosition++
}

// peekChar returns the next character without advancing the position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// isLetter checks if the character is a letter or underscore
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit checks if the character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// isWhitespace checks if the character is a whitespace
func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// readIdentifier reads an identifier from the input
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) { // Tanımlayıcılar harf veya alt çizgi ile başlar, sonrasında harf, rakam veya alt çizgi gelebilir
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a number from the input
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	// Ondalık sayılar için nokta kontrolü
	if l.ch == '.' && isDigit(l.peekChar()) {
		l.readChar() // Noktayı geç
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return l.input[position:l.position]
}

// readString reads a string literal from the input
func (l *Lexer) readString(delimiter byte) string {
	position := l.position + 1 // Başlangıç tırnak işaretini atla
	for {
		l.readChar()
		if l.ch == delimiter || l.ch == 0 {
			break
		}

		// Kaçış dizileri için
		if l.ch == '\\' {
			l.readChar() // Kaçış karakterini atla
		}
	}

	if l.ch == 0 {
		// Sonlandırılmamış string
		return l.input[position:l.position]
	}

	result := l.input[position:l.position]
	// Bitiş tırnak işaretini atlama - NextToken'da readChar() çağrılacak
	return result
}

// skipWhitespace skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

// skipComment skips a comment
func (l *Lexer) skipComment() {
	if l.ch == '/' && l.peekChar() == '/' {
		// Tek satırlık yorum
		for l.ch != '\n' && l.ch != 0 {
			l.readChar()
		}
	} else if l.ch == '/' && l.peekChar() == '*' {
		// Çok satırlı yorum
		l.readChar() // '/' karakterini atla
		l.readChar() // '*' karakterini atla

		for !(l.ch == '*' && l.peekChar() == '/') && l.ch != 0 {
			l.readChar()
		}

		if l.ch != 0 {
			l.readChar() // '*' karakterini atla
			l.readChar() // '/' karakterini atla
		}
	}
}

// newToken creates a new token
func (l *Lexer) newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: literal,
		Line:    l.startLine,
		Column:  l.startColumn,
		Pos:     l.startPos,
		End:     l.position,
	}
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	// Token başlangıç pozisyonunu kaydet
	l.startPos = l.position
	l.startLine = l.line
	l.startColumn = l.column

	// Yorum kontrolü
	if l.ch == '/' && (l.peekChar() == '/' || l.peekChar() == '*') {
		l.skipComment()
		return l.NextToken()
	}

	var tok token.Token

	switch l.ch {
	// Tek karakterli operatörler ve ayırıcılar
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.EQ, "==")
		} else {
			tok = l.newToken(token.ASSIGN, "=")
		}
	case '+':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.PLUS_ASSIGN, "+=")
		} else if l.peekChar() == '+' {
			l.readChar()
			tok = l.newToken(token.INCREMENT, "++") // Artırma operatörü
		} else {
			tok = l.newToken(token.PLUS, "+")
		}
	case '-':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.MINUS_ASSIGN, "-=")
		} else if l.peekChar() == '-' {
			l.readChar()
			tok = l.newToken(token.DECREMENT, "--") // Azaltma operatörü
		} else if l.peekChar() == '>' {
			l.readChar()
			tok = l.newToken(token.ARROW, "->") // Ok operatörü
		} else {
			tok = l.newToken(token.MINUS, "-")
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.NOT_EQ, "!=")
		} else {
			tok = l.newToken(token.BANG, "!")
		}
	case '*':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.MUL_ASSIGN, "*=")
		} else {
			tok = l.newToken(token.ASTERISK, "*")
		}
	case '/':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.DIV_ASSIGN, "/=")
		} else {
			tok = l.newToken(token.SLASH, "/")
		}
	case '%':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.MOD_ASSIGN, "%=")
		} else {
			tok = l.newToken(token.MODULO, "%")
		}
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.LTOEQ, "<=")
		} else if l.peekChar() == '<' {
			l.readChar()
			if l.peekChar() == '=' {
				l.readChar()
				tok = l.newToken(token.LEFT_SHIFT_ASSIGN, "<<=")
			} else {
				tok = l.newToken(token.LEFT_SHIFT, "<<")
			}
		} else {
			tok = l.newToken(token.LT, "<")
		}
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.GTOEQ, ">=")
		} else if l.peekChar() == '>' {
			l.readChar()
			if l.peekChar() == '=' {
				l.readChar()
				tok = l.newToken(token.RIGHT_SHIFT_ASSIGN, ">>=")
			} else {
				tok = l.newToken(token.RIGHT_SHIFT, ">>")
			}
		} else {
			tok = l.newToken(token.GT, ">")
		}
	case '&':
		if l.peekChar() == '&' {
			l.readChar()
			tok = l.newToken(token.LOGICAL_AND, "&&")
		} else if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.AND_ASSIGN, "&=")
		} else {
			tok = l.newToken(token.BIT_AND, "&")
		}
	case '|':
		if l.peekChar() == '|' {
			l.readChar()
			tok = l.newToken(token.LOGICAL_OR, "||")
		} else if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.OR_ASSIGN, "|=")
		} else {
			tok = l.newToken(token.BIT_OR, "|")
		}
	case '^':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.XOR_ASSIGN, "^=")
		} else {
			tok = l.newToken(token.BIT_XOR, "^")
		}
	case '~':
		tok = l.newToken(token.BIT_NOT, "~")
	case ':':
		if l.peekChar() == ':' {
			l.readChar()
			tok = l.newToken(token.SCOPE_RES, "::")
		} else if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(token.DEFINE, ":=") // Kısa değişken tanımlama
		} else {
			tok = l.newToken(token.COLON, ":")
		}

	// Ayırıcılar
	case ',':
		tok = l.newToken(token.COMMA, ",")
	case ';':
		tok = l.newToken(token.SEMICOLON, ";")
	case '.':
		tok = l.newToken(token.DOT, ".")
	case '(':
		tok = l.newToken(token.LPAREN, "(")
	case ')':
		tok = l.newToken(token.RPAREN, ")")
	case '{':
		tok = l.newToken(token.LBRACE, "{")
	case '}':
		tok = l.newToken(token.RBRACE, "}")
	case '[':
		tok = l.newToken(token.LBRACKET, "[")
	case ']':
		tok = l.newToken(token.RBRACKET, "]")

	// String ve karakter literalleri
	case '"':
		literal := l.readString('"')
		tok = l.newToken(token.STRING, literal)
	case '\'':
		literal := l.readString('\'')
		tok = l.newToken(token.CHAR, literal)
	case '`':
		literal := l.readString('`')
		tok = l.newToken(token.STRING, literal)

	case 0:
		tok = l.newToken(token.EOF, "")
	default:
		if isLetter(l.ch) {
			// Tanımlayıcı veya anahtar kelime
			literal := l.readIdentifier()
			tokenType := token.LookupIdent(literal)
			return l.newToken(tokenType, literal)
		} else if isDigit(l.ch) {
			// Sayı literali
			literal := l.readNumber()
			// Basit bir kontrol: eğer nokta içeriyorsa FLOAT, değilse INT
			if contains(literal, '.') {
				return l.newToken(token.FLOAT, literal)
			}
			return l.newToken(token.INT, literal)
		} else {
			tok = l.newToken(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

// contains checks if a string contains a specific character
func contains(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

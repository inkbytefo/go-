package token

// TokenType, bir token'ın türünü temsil eden bir string'dir.
type TokenType string

// Token, kaynak koddaki bir token'ı temsil eder.
type Token struct {
	Type     TokenType // Token türü (örn: IDENT, INT, LPAREN)
	Literal  string    // Token'ın değişmez değeri (örn: "x", "123", "(")
	Line     int       // Token'ın bulunduğu satır numarası
	Column   int       // Token'ın bulunduğu sütun numarası
	Pos      int       // Token'ın dosyadaki başlangıç pozisyonu (byte offset)
	End      int       // Token'ın dosyadaki bitiş pozisyonu (byte offset)
	Position Position  // Token'ın konumu
}

// Anahtar kelimeler ve token türleri
const (
	// Özel token türleri
	ILLEGAL TokenType = "ILLEGAL" // Tanınmayan token veya karakter
	EOF     TokenType = "EOF"     // Dosya sonu

	// Tanımlayıcılar + Değişmez Değerler (Literals)
	IDENT  TokenType = "IDENT"  // main, foobar, x, y, ...
	INT    TokenType = "INT"    // 1343456
	FLOAT  TokenType = "FLOAT"  // 3.14
	STRING TokenType = "STRING" // "hello world"
	CHAR   TokenType = "CHAR"   // 'a'

	// Operatörler
	ASSIGN    TokenType = "="
	PLUS      TokenType = "+"
	MINUS     TokenType = "-"
	BANG      TokenType = "!"
	ASTERISK  TokenType = "*"
	SLASH     TokenType = "/"
	MODULO    TokenType = "%"
	INCREMENT TokenType = "++"
	DECREMENT TokenType = "--"

	LT     TokenType = "<"
	GT     TokenType = ">"
	LTOEQ  TokenType = "<="
	GTOEQ  TokenType = ">="
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="

	// Bileşik atama operatörleri
	PLUS_ASSIGN  TokenType = "+="
	MINUS_ASSIGN TokenType = "-="
	MUL_ASSIGN   TokenType = "*="
	DIV_ASSIGN   TokenType = "/="
	MOD_ASSIGN   TokenType = "%="

	// Bit operatörleri
	BIT_AND     TokenType = "&"
	BIT_OR      TokenType = "|"
	BIT_XOR     TokenType = "^"
	BIT_NOT     TokenType = "~"
	LEFT_SHIFT  TokenType = "<<"
	RIGHT_SHIFT TokenType = ">>"

	// Bit bileşik atama operatörleri
	AND_ASSIGN         TokenType = "&="
	OR_ASSIGN          TokenType = "|="
	XOR_ASSIGN         TokenType = "^="
	LEFT_SHIFT_ASSIGN  TokenType = "<<="
	RIGHT_SHIFT_ASSIGN TokenType = ">>="

	// Mantıksal operatörler
	LOGICAL_AND TokenType = "&&"
	LOGICAL_OR  TokenType = "||"

	// C++ tarzı operatörler
	ARROW     TokenType = "->" // Pointer üye erişimi için
	SCOPE_RES TokenType = "::" // Kapsam çözümleme operatörü

	// Ayırıcılar (Delimiters)
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";" // İsteğe bağlı
	COLON     TokenType = ":"
	DOT       TokenType = "."
	DEFINE    TokenType = ":=" // Kısa değişken tanımlama

	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"

	// Anahtar Kelimeler (Keywords)
	PACKAGE     TokenType = "PACKAGE"
	IMPORT      TokenType = "IMPORT"
	FUNC        TokenType = "FUNC"
	VAR         TokenType = "VAR"
	CONST       TokenType = "CONST"
	RETURN      TokenType = "RETURN"
	IF          TokenType = "IF"
	ELSE        TokenType = "ELSE"
	FOR         TokenType = "FOR"
	WHILE       TokenType = "WHILE" // C++ tarzı döngü için eklenebilir
	BREAK       TokenType = "BREAK"
	CONTINUE    TokenType = "CONTINUE"
	STRUCT      TokenType = "STRUCT"
	INTERFACE   TokenType = "INTERFACE"
	MAP         TokenType = "MAP"
	CHAN        TokenType = "CHAN"
	GO          TokenType = "GO"
	DEFER       TokenType = "DEFER"
	SELECT      TokenType = "SELECT"
	SWITCH      TokenType = "SWITCH"
	CASE        TokenType = "CASE"
	DEFAULT     TokenType = "DEFAULT"
	TYPE        TokenType = "TYPE"
	FALLTHROUGH TokenType = "FALLTHROUGH"
	RANGE       TokenType = "RANGE"

	// GO+ Eklemeleri
	CLASS      TokenType = "CLASS"
	TEMPLATE   TokenType = "TEMPLATE" // template<T>
	THROW      TokenType = "THROW"
	TRY        TokenType = "TRY"
	CATCH      TokenType = "CATCH"
	FINALLY    TokenType = "FINALLY"
	SCOPE      TokenType = "SCOPE"     // scope {} (RAII için)
	PUBLIC     TokenType = "PUBLIC"    // Sınıf üyeleri için
	PRIVATE    TokenType = "PRIVATE"   // Sınıf üyeleri için
	PROTECTED  TokenType = "PROTECTED" // Sınıf üyeleri için
	NEW        TokenType = "NEW"
	DELETE     TokenType = "DELETE" // Manuel bellek yönetimi için (unsafe bloklarda)
	UNSAFE     TokenType = "UNSAFE"
	ALLOC      TokenType = "ALLOC" // Manuel bellek yönetimi için
	FREE       TokenType = "FREE"  // Manuel bellek yönetimi için
	THIS       TokenType = "THIS"
	SUPER      TokenType = "SUPER"
	NULL       TokenType = "NULL"
	TRUE       TokenType = "TRUE"
	FALSE      TokenType = "FALSE"
	EXTENDS    TokenType = "EXTENDS"    // Sınıf kalıtımı için
	IMPLEMENTS TokenType = "IMPLEMENTS" // Arayüz uygulaması için
	VIRTUAL    TokenType = "VIRTUAL"    // Sanal metotlar için
	OVERRIDE   TokenType = "OVERRIDE"   // Metot ezme için
	FINAL      TokenType = "FINAL"      // Son sınıf/metot için
	ABSTRACT   TokenType = "ABSTRACT"   // Soyut sınıf/metot için
	STATIC     TokenType = "STATIC"     // Statik üyeler için
	CONST_EXPR TokenType = "CONSTEXPR"  // Derleme zamanı sabit ifadeleri için
	NAMESPACE  TokenType = "NAMESPACE"  // İsim alanları için
	USING      TokenType = "USING"      // İsim alanı kullanımı için
	FRIEND     TokenType = "FRIEND"     // Arkadaş sınıf/fonksiyon için
	OPERATOR   TokenType = "OPERATOR"   // Operatör aşırı yükleme için
)

// keywords, anahtar kelimeleri token türleriyle eşler.
var keywords = map[string]TokenType{
	// Go anahtar kelimeleri
	"package":     PACKAGE,
	"import":      IMPORT,
	"func":        FUNC,
	"var":         VAR,
	"const":       CONST,
	"return":      RETURN,
	"if":          IF,
	"else":        ELSE,
	"for":         FOR,
	"while":       WHILE,
	"break":       BREAK,
	"continue":    CONTINUE,
	"struct":      STRUCT,
	"interface":   INTERFACE,
	"map":         MAP,
	"chan":        CHAN,
	"go":          GO,
	"defer":       DEFER,
	"select":      SELECT,
	"switch":      SWITCH,
	"case":        CASE,
	"default":     DEFAULT,
	"type":        TYPE,
	"fallthrough": FALLTHROUGH,
	"range":       RANGE,

	// GO+ anahtar kelimeleri
	"class":      CLASS,
	"template":   TEMPLATE,
	"throw":      THROW,
	"try":        TRY,
	"catch":      CATCH,
	"finally":    FINALLY,
	"scope":      SCOPE,
	"public":     PUBLIC,
	"private":    PRIVATE,
	"protected":  PROTECTED,
	"new":        NEW,
	"delete":     DELETE,
	"unsafe":     UNSAFE,
	"alloc":      ALLOC,
	"free":       FREE,
	"this":       THIS,
	"super":      SUPER,
	"nil":        NULL, // Go'daki nil'e karşılık
	"true":       TRUE,
	"false":      FALSE,
	"extends":    EXTENDS,
	"implements": IMPLEMENTS,
	"virtual":    VIRTUAL,
	"override":   OVERRIDE,
	"final":      FINAL,
	"abstract":   ABSTRACT,
	"static":     STATIC,
	"constexpr":  CONST_EXPR,
	"namespace":  NAMESPACE,
	"using":      USING,
	"friend":     FRIEND,
	"operator":   OPERATOR,
}

// LookupIdent, verilen tanımlayıcının bir anahtar kelime olup olmadığını kontrol eder.
// Eğer anahtar kelimeyse, ilgili TokenType'ı döndürür, değilse IDENT döndürür.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

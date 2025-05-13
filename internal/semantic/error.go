package semantic

import (
	"fmt"
	"gominus/internal/token"
	"strings"
)

// ErrorLevel, hata seviyesini temsil eder.
type ErrorLevel int

const (
	ERROR_LEVEL ErrorLevel = iota
	WARNING_LEVEL
	INFO_LEVEL
)

// String, hata seviyesinin string temsilini döndürür.
func (el ErrorLevel) String() string {
	switch el {
	case ERROR_LEVEL:
		return "error"
	case WARNING_LEVEL:
		return "warning"
	case INFO_LEVEL:
		return "info"
	default:
		return "unknown"
	}
}

// Color, hata seviyesinin rengini döndürür.
func (el ErrorLevel) Color() string {
	switch el {
	case ERROR_LEVEL:
		return "\033[1;31m" // Kırmızı
	case WARNING_LEVEL:
		return "\033[1;33m" // Sarı
	case INFO_LEVEL:
		return "\033[1;34m" // Mavi
	default:
		return "\033[0m" // Normal
	}
}

// SemanticError, semantik analiz sırasında oluşan bir hatayı temsil eder.
type SemanticError struct {
	Level   ErrorLevel
	Token   token.Token
	Message string
	Hints   []string
}

// String, hatanın string temsilini döndürür.
func (se *SemanticError) String() string {
	var builder strings.Builder

	// Dosya ve konum bilgisi
	builder.WriteString(fmt.Sprintf("Satır %d, Sütun %d: ", se.Token.Line, se.Token.Column))

	// Hata seviyesi
	builder.WriteString(fmt.Sprintf("%s%s\033[0m: ", se.Level.Color(), se.Level.String()))

	// Hata mesajı
	builder.WriteString(se.Message)

	// İpuçları
	if len(se.Hints) > 0 {
		builder.WriteString("\nİpuçları:\n")
		for _, hint := range se.Hints {
			builder.WriteString(fmt.Sprintf("  - %s\n", hint))
		}
	}

	return builder.String()
}

// NewError, yeni bir hata oluşturur.
func NewError(token token.Token, format string, args ...interface{}) *SemanticError {
	return &SemanticError{
		Level:   ERROR_LEVEL,
		Token:   token,
		Message: fmt.Sprintf(format, args...),
		Hints:   []string{},
	}
}

// NewWarning, yeni bir uyarı oluşturur.
func NewWarning(token token.Token, format string, args ...interface{}) *SemanticError {
	return &SemanticError{
		Level:   WARNING_LEVEL,
		Token:   token,
		Message: fmt.Sprintf(format, args...),
		Hints:   []string{},
	}
}

// NewInfo, yeni bir bilgi oluşturur.
func NewInfo(token token.Token, format string, args ...interface{}) *SemanticError {
	return &SemanticError{
		Level:   INFO_LEVEL,
		Token:   token,
		Message: fmt.Sprintf(format, args...),
		Hints:   []string{},
	}
}

// AddHint, hataya bir ipucu ekler.
func (se *SemanticError) AddHint(format string, args ...interface{}) *SemanticError {
	se.Hints = append(se.Hints, fmt.Sprintf(format, args...))
	return se
}

// ErrorReporter, hata raporlama işlemlerini gerçekleştirir.
type ErrorReporter struct {
	Errors   []*SemanticError
	Warnings []*SemanticError
	Infos    []*SemanticError
}

// NewErrorReporter, yeni bir hata raporlayıcı oluşturur.
func NewErrorReporter() *ErrorReporter {
	return &ErrorReporter{
		Errors:   []*SemanticError{},
		Warnings: []*SemanticError{},
		Infos:    []*SemanticError{},
	}
}

// ReportError, bir hata raporlar.
func (er *ErrorReporter) ReportError(token token.Token, format string, args ...interface{}) *SemanticError {
	error := NewError(token, format, args...)
	er.Errors = append(er.Errors, error)
	return error
}

// ReportWarning, bir uyarı raporlar.
func (er *ErrorReporter) ReportWarning(token token.Token, format string, args ...interface{}) *SemanticError {
	warning := NewWarning(token, format, args...)
	er.Warnings = append(er.Warnings, warning)
	return warning
}

// ReportInfo, bir bilgi raporlar.
func (er *ErrorReporter) ReportInfo(token token.Token, format string, args ...interface{}) *SemanticError {
	info := NewInfo(token, format, args...)
	er.Infos = append(er.Infos, info)
	return info
}

// HasErrors, hata olup olmadığını kontrol eder.
func (er *ErrorReporter) HasErrors() bool {
	return len(er.Errors) > 0
}

// HasWarnings, uyarı olup olmadığını kontrol eder.
func (er *ErrorReporter) HasWarnings() bool {
	return len(er.Warnings) > 0
}

// HasInfos, bilgi olup olmadığını kontrol eder.
func (er *ErrorReporter) HasInfos() bool {
	return len(er.Infos) > 0
}

// GetAllMessages, tüm mesajları döndürür.
func (er *ErrorReporter) GetAllMessages() []string {
	var messages []string

	for _, error := range er.Errors {
		messages = append(messages, error.String())
	}

	for _, warning := range er.Warnings {
		messages = append(messages, warning.String())
	}

	for _, info := range er.Infos {
		messages = append(messages, info.String())
	}

	return messages
}

// PrintAllMessages, tüm mesajları yazdırır.
func (er *ErrorReporter) PrintAllMessages() {
	for _, message := range er.GetAllMessages() {
		fmt.Println(message)
	}
}

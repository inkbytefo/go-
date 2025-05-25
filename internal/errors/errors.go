// Package errors provides standardized error types and utilities for the GO-Minus compiler.
package errors

import (
	"fmt"
	"strings"

	"github.com/inkbytefo/go-minus/internal/token"
)

// ErrorType represents the type of error.
type ErrorType int

const (
	// SyntaxError represents a syntax error in the source code.
	SyntaxError ErrorType = iota
	// SemanticError represents a semantic error (type mismatch, undefined variable, etc.).
	SemanticError
	// IRGenError represents an error during IR generation.
	IRGenError
	// CodeGenError represents an error during code generation.
	CodeGenError
	// IOError represents an I/O error (file not found, permission denied, etc.).
	IOError
	// InternalError represents an internal compiler error.
	InternalError
)

// String returns the string representation of the error type.
func (et ErrorType) String() string {
	switch et {
	case SyntaxError:
		return "Syntax Error"
	case SemanticError:
		return "Semantic Error"
	case IRGenError:
		return "IR Generation Error"
	case CodeGenError:
		return "Code Generation Error"
	case IOError:
		return "I/O Error"
	case InternalError:
		return "Internal Error"
	default:
		return "Unknown Error"
	}
}

// CompilerError represents a structured error with position information.
type CompilerError struct {
	Type     ErrorType
	Message  string
	Position token.Position
	File     string
	Hint     string
	Cause    error
}

// Error implements the error interface.
func (e *CompilerError) Error() string {
	var parts []string

	// Add error type
	parts = append(parts, e.Type.String())

	// Add position information if available
	if e.Position.IsValid() {
		if e.File != "" {
			parts = append(parts, fmt.Sprintf("at %s:%d:%d", e.File, e.Position.Line, e.Position.Column))
		} else {
			parts = append(parts, fmt.Sprintf("at line %d, column %d", e.Position.Line, e.Position.Column))
		}
	} else if e.File != "" {
		parts = append(parts, fmt.Sprintf("in %s", e.File))
	}

	// Add main message
	result := strings.Join(parts, " ") + ": " + e.Message

	// Add hint if available
	if e.Hint != "" {
		result += "\n  Hint: " + e.Hint
	}

	// Add cause if available
	if e.Cause != nil {
		result += "\n  Caused by: " + e.Cause.Error()
	}

	return result
}

// Unwrap returns the underlying cause error.
func (e *CompilerError) Unwrap() error {
	return e.Cause
}

// Is checks if the error matches the target error type.
func (e *CompilerError) Is(target error) bool {
	if t, ok := target.(*CompilerError); ok {
		return e.Type == t.Type
	}
	return false
}

// ErrorList represents a collection of errors.
type ErrorList []*CompilerError

// Add adds an error to the list.
func (el *ErrorList) Add(err *CompilerError) {
	*el = append(*el, err)
}

// Error implements the error interface.
func (el ErrorList) Error() string {
	if len(el) == 0 {
		return "no errors"
	}

	if len(el) == 1 {
		return el[0].Error()
	}

	var parts []string
	for i, err := range el {
		if i >= 10 { // Limit to first 10 errors
			parts = append(parts, fmt.Sprintf("... and %d more errors", len(el)-i))
			break
		}
		parts = append(parts, err.Error())
	}

	return strings.Join(parts, "\n")
}

// HasErrors returns true if the list contains any errors.
func (el ErrorList) HasErrors() bool {
	return len(el) > 0
}

// Count returns the number of errors in the list.
func (el ErrorList) Count() int {
	return len(el)
}

// Filter returns a new ErrorList containing only errors of the specified type.
func (el ErrorList) Filter(errorType ErrorType) ErrorList {
	var filtered ErrorList
	for _, err := range el {
		if err.Type == errorType {
			filtered = append(filtered, err)
		}
	}
	return filtered
}

// NewSyntaxError creates a new syntax error.
func NewSyntaxError(pos token.Position, file, message string, args ...interface{}) *CompilerError {
	return &CompilerError{
		Type:     SyntaxError,
		Message:  fmt.Sprintf(message, args...),
		Position: pos,
		File:     file,
	}
}

// NewSemanticError creates a new semantic error.
func NewSemanticError(pos token.Position, file, message string, args ...interface{}) *CompilerError {
	return &CompilerError{
		Type:     SemanticError,
		Message:  fmt.Sprintf(message, args...),
		Position: pos,
		File:     file,
	}
}

// NewIRGenError creates a new IR generation error.
func NewIRGenError(pos token.Position, file, message string, args ...interface{}) *CompilerError {
	return &CompilerError{
		Type:     IRGenError,
		Message:  fmt.Sprintf(message, args...),
		Position: pos,
		File:     file,
	}
}

// NewCodeGenError creates a new code generation error.
func NewCodeGenError(message string, args ...interface{}) *CompilerError {
	return &CompilerError{
		Type:    CodeGenError,
		Message: fmt.Sprintf(message, args...),
	}
}

// NewIOError creates a new I/O error.
func NewIOError(file, message string, cause error, args ...interface{}) *CompilerError {
	return &CompilerError{
		Type:    IOError,
		Message: fmt.Sprintf(message, args...),
		File:    file,
		Cause:   cause,
	}
}

// NewInternalError creates a new internal error.
func NewInternalError(message string, cause error, args ...interface{}) *CompilerError {
	return &CompilerError{
		Type:    InternalError,
		Message: fmt.Sprintf(message, args...),
		Cause:   cause,
	}
}

// WithHint adds a hint to an existing error.
func WithHint(err *CompilerError, hint string, args ...interface{}) *CompilerError {
	err.Hint = fmt.Sprintf(hint, args...)
	return err
}

// WithCause adds a cause to an existing error.
func WithCause(err *CompilerError, cause error) *CompilerError {
	err.Cause = cause
	return err
}

// WrapError wraps a standard error as a CompilerError.
func WrapError(err error, errorType ErrorType, message string, args ...interface{}) *CompilerError {
	return &CompilerError{
		Type:    errorType,
		Message: fmt.Sprintf(message, args...),
		Cause:   err,
	}
}

// ErrorReporter provides a centralized way to collect and report errors.
type ErrorReporter struct {
	errors ErrorList
	file   string
}

// NewErrorReporter creates a new error reporter.
func NewErrorReporter(file string) *ErrorReporter {
	return &ErrorReporter{
		file: file,
	}
}

// ReportSyntaxError reports a syntax error.
func (er *ErrorReporter) ReportSyntaxError(pos token.Position, message string, args ...interface{}) {
	er.errors.Add(NewSyntaxError(pos, er.file, message, args...))
}

// ReportSemanticError reports a semantic error.
func (er *ErrorReporter) ReportSemanticError(pos token.Position, message string, args ...interface{}) {
	er.errors.Add(NewSemanticError(pos, er.file, message, args...))
}

// ReportIRGenError reports an IR generation error.
func (er *ErrorReporter) ReportIRGenError(pos token.Position, message string, args ...interface{}) {
	er.errors.Add(NewIRGenError(pos, er.file, message, args...))
}

// ReportCodeGenError reports a code generation error.
func (er *ErrorReporter) ReportCodeGenError(message string, args ...interface{}) {
	er.errors.Add(NewCodeGenError(message, args...))
}

// ReportIOError reports an I/O error.
func (er *ErrorReporter) ReportIOError(message string, cause error, args ...interface{}) {
	er.errors.Add(NewIOError(er.file, message, cause, args...))
}

// ReportInternalError reports an internal error.
func (er *ErrorReporter) ReportInternalError(message string, cause error, args ...interface{}) {
	er.errors.Add(NewInternalError(message, cause, args...))
}

// AddError adds an existing error to the reporter.
func (er *ErrorReporter) AddError(err *CompilerError) {
	er.errors.Add(err)
}

// Errors returns all collected errors.
func (er *ErrorReporter) Errors() ErrorList {
	return er.errors
}

// HasErrors returns true if any errors have been reported.
func (er *ErrorReporter) HasErrors() bool {
	return er.errors.HasErrors()
}

// Clear clears all collected errors.
func (er *ErrorReporter) Clear() {
	er.errors = nil
}

// SetFile sets the current file being processed.
func (er *ErrorReporter) SetFile(file string) {
	er.file = file
}

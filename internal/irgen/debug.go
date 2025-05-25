// This file contains the implementation of debug information generation for the GO-Minus compiler.
package irgen

import (
	"fmt"
	"path/filepath"

	"github.com/llir/llvm/ir"
)

// DebugInfo manages debug information generation.
type DebugInfo struct {
	module      *ir.Module
	sourceFile  string
	sourceDir   string
	currentLine int
	currentCol  int
	enabled     bool
}

// NewDebugInfo creates a new DebugInfo instance.
func NewDebugInfo(module *ir.Module) *DebugInfo {
	return &DebugInfo{
		module:  module,
		enabled: true,
	}
}

// InitCompileUnit initializes the compile unit metadata.
func (d *DebugInfo) InitCompileUnit(filename, directory, producer string, isOptimized bool, flags string, runtimeVersion int) {
	d.sourceFile = filename
	d.sourceDir = directory

	// In a real implementation, we would create DWARF debug info here
	// For now, we'll just store the source file information
	fmt.Printf("Debug info initialized for %s/%s\n", directory, filename)
}

// CreateFunction creates function metadata.
func (d *DebugInfo) CreateFunction(fn *ir.Func, name string, linkageName string, file interface{}, line int, isLocal bool, isDefinition bool, scopeLine int, flags interface{}, isOptimized bool) interface{} {
	if !d.enabled {
		return nil
	}

	// In a real implementation, we would create DWARF debug info for the function
	// For now, we'll just add a comment to the function
	comment := fmt.Sprintf("; function %s, line %d, file %s", name, line, d.sourceFile)

	// We can't directly modify the function in the llir/llvm library
	// In a real implementation, we would use the LLVM C API to attach debug info

	// For now, just print the debug info
	if false { // Disable debug printing to avoid cluttering the output
		fmt.Printf("Debug info for function %s: %s\n", name, comment)
	}

	return nil
}

// SetLocation sets the current source location.
func (d *DebugInfo) SetLocation(line, col int, filename string) {
	d.currentLine = line
	d.currentCol = col

	if filename != "" && d.sourceFile != filename {
		d.sourceFile = filename
		d.sourceDir = filepath.Dir(filename)
	}
}

// AttachLocation attaches location metadata to an instruction.
func (d *DebugInfo) AttachLocation(inst ir.Instruction) {
	if !d.enabled || d.currentLine <= 0 {
		return
	}

	// In a real implementation, we would attach DWARF debug info to the instruction
	// For now, we'll just add a comment to the instruction
	comment := fmt.Sprintf("; line %d, col %d, file %s", d.currentLine, d.currentCol, d.sourceFile)

	// We can't directly modify the instruction in the llir/llvm library
	// In a real implementation, we would use the LLVM C API to attach debug info

	// For now, just print the debug info
	if false { // Disable debug printing to avoid cluttering the output
		fmt.Printf("Debug info for instruction %v: %s\n", inst, comment)
	}
}

// getOrCreateFileMetadata returns file metadata for the given file.
func (d *DebugInfo) getOrCreateFileMetadata(filename, directory string) interface{} {
	// In a real implementation, we would create DWARF debug info for the file
	// For now, we'll just return a dummy value
	return nil
}

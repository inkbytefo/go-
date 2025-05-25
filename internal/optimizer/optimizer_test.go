package optimizer

import (
	"strings"
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// TestOptimizationLevels tests different optimization levels.
func TestOptimizationLevels(t *testing.T) {
	// Create a simple LLVM IR module
	module := ir.NewModule()

	// Add a simple function that adds two integers
	f := module.NewFunc("add", types.I32, ir.NewParam("a", types.I32), ir.NewParam("b", types.I32))
	entry := f.NewBlock("entry")
	a := entry.NewLoad(types.I32, f.Params[0])
	b := entry.NewLoad(types.I32, f.Params[1])
	result := entry.NewAdd(a, b)
	entry.NewRet(result)

	// Test each optimization level
	levels := []OptimizationLevel{O0, O1, O2, O3}
	for _, level := range levels {
		t.Run(getLevelName(level), func(t *testing.T) {
			opt := New(level)
			optimizedModule, err := opt.OptimizeModule(module)
			if err != nil {
				t.Fatalf("Failed to optimize module: %v", err)
			}

			// Check that the optimized module is not nil
			if optimizedModule == nil {
				t.Fatal("Optimized module is nil")
			}

			// For O0, the module should be unchanged
			if level == O0 {
				if optimizedModule != module {
					t.Error("O0 optimization should return the original module")
				}
			}
		})
	}
}

// TestConstantFolding tests the constant folding optimization.
func TestConstantFolding(t *testing.T) {
	// Create a simple LLVM IR module
	module := ir.NewModule()

	// Add a function that adds two constants
	f := module.NewFunc("const_add", types.I32)
	entry := f.NewBlock("entry")
	a := constant.NewInt(types.I32, 2)
	b := constant.NewInt(types.I32, 3)
	result := entry.NewAdd(a, b)
	entry.NewRet(result)

	// Apply constant folding
	pass := &ConstantFoldingPass{}
	optimizedModule, err := pass.Apply(module)
	if err != nil {
		t.Fatalf("Failed to apply constant folding: %v", err)
	}

	// Check that the optimized module is not nil
	if optimizedModule == nil {
		t.Fatal("Optimized module is nil")
	}

	// In a real implementation, we would check that the add instruction
	// has been replaced with a constant 5
}

// TestDeadCodeElimination tests the dead code elimination optimization.
func TestDeadCodeElimination(t *testing.T) {
	// Create a simple LLVM IR module
	module := ir.NewModule()

	// Add a function with dead code
	f := module.NewFunc("dead_code", types.Void)
	entry := f.NewBlock("entry")
	
	// This variable is never used
	deadVar := entry.NewAlloca(types.I32)
	entry.NewStore(constant.NewInt(types.I32, 42), deadVar)
	
	// This variable is used
	liveVar := entry.NewAlloca(types.I32)
	entry.NewStore(constant.NewInt(types.I32, 123), liveVar)
	loadedValue := entry.NewLoad(types.I32, liveVar)
	
	// Use the loaded value
	entry.NewRet(loadedValue)

	// Apply dead code elimination
	pass := &DeadCodeEliminationPass{}
	optimizedModule, err := pass.Apply(module)
	if err != nil {
		t.Fatalf("Failed to apply dead code elimination: %v", err)
	}

	// Check that the optimized module is not nil
	if optimizedModule == nil {
		t.Fatal("Optimized module is nil")
	}

	// In a real implementation, we would check that the dead variable
	// and its store instruction have been eliminated
}

// TestGetOptimizedIRString tests the GetOptimizedIRString function.
func TestGetOptimizedIRString(t *testing.T) {
	// Create a simple LLVM IR string
	irString := `
define i32 @add(i32 %a, i32 %b) {
entry:
  %0 = add i32 %a, %b
  ret i32 %0
}
`

	// Test each optimization level
	levels := []OptimizationLevel{O0, O1, O2, O3}
	for _, level := range levels {
		t.Run(getLevelName(level), func(t *testing.T) {
			opt := New(level)
			optimizedIR, err := opt.GetOptimizedIRString(irString)
			
			// If LLVM tools are not available, the function should return the original IR
			if err != nil && !strings.Contains(err.Error(), "LLVM opt") {
				t.Fatalf("Failed to optimize IR string: %v", err)
			}

			// Check that the optimized IR is not empty
			if optimizedIR == "" {
				t.Fatal("Optimized IR is empty")
			}

			// For O0, the IR should be unchanged
			if level == O0 && optimizedIR != irString {
				t.Error("O0 optimization should return the original IR")
			}
		})
	}
}

// Helper function to get the name of an optimization level
func getLevelName(level OptimizationLevel) string {
	switch level {
	case O0:
		return "O0"
	case O1:
		return "O1"
	case O2:
		return "O2"
	case O3:
		return "O3"
	default:
		return "Unknown"
	}
}

// Package runtime provides a public API for the GO-Minus runtime.
package runtime

// RuntimeOptions represents options for the runtime.
type RuntimeOptions struct {
	DebugMode bool
	TraceMode bool
	GCMode    string
}

// DefaultRuntimeOptions returns the default runtime options.
func DefaultRuntimeOptions() RuntimeOptions {
	return RuntimeOptions{
		DebugMode: false,
		TraceMode: false,
		GCMode:    "default",
	}
}

// Runtime represents a GO-Minus runtime.
type Runtime struct {
	options RuntimeOptions
	errors  []string
}

// New creates a new GO-Minus runtime with the given options.
func New(options RuntimeOptions) *Runtime {
	return &Runtime{
		options: options,
		errors:  []string{},
	}
}

// Errors returns the runtime errors.
func (r *Runtime) Errors() []string {
	return r.errors
}

// Execute executes a GO-Minus program.
func (r *Runtime) Execute(programPath string, args []string) (int, error) {
	// TODO: Implement the execution process
	return 0, nil
}

// ExecuteIR executes a GO-Minus IR program.
func (r *Runtime) ExecuteIR(irPath string, args []string) (int, error) {
	// TODO: Implement the execution process
	return 0, nil
}

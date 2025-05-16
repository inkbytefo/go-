# GO-Minus Project Directory Structure

This document explains the directory structure of the GO-Minus project and the purpose of each directory.

## Main Directories

```
go-minus/
├── cmd/                      # Command-line applications
│   ├── gominus/              # GO-Minus compiler
│   ├── gompm/                # GO-Minus package manager
│   ├── gomfmt/               # GO-Minus code formatting tool
│   ├── gomdoc/               # GO-Minus documentation tool
│   ├── gomtest/              # GO-Minus testing tool
│   └── gomlsp/               # GO-Minus language server
├── docs/                     # Documentation
│   ├── tutorial/             # Tutorials
│   ├── reference/            # Reference documents
│   ├── images/               # Documentation images
│   └── ...                   # Other documentation files
├── examples/                 # Example code
│   ├── basic/                # Basic examples
│   ├── advanced/             # Advanced examples
│   ├── vulkan/               # Vulkan examples
│   └── memory/               # Memory management examples
├── internal/                 # Internal packages
│   ├── ast/                  # Abstract syntax tree
│   ├── codegen/              # Code generation
│   ├── irgen/                # IR generation
│   ├── lexer/                # Lexical analysis
│   ├── optimizer/            # Optimization
│   ├── parser/               # Syntax analysis
│   ├── semantic/             # Semantic analysis
│   └── token/                # Token definitions
├── pkg/                      # Public packages
│   ├── compiler/             # Compiler API
│   └── runtime/              # Runtime API
├── stdlib/                   # Standard library
│   ├── async/                # Asynchronous operations
│   ├── concurrent/           # Concurrency
│   ├── container/            # Data structures
│   ├── core/                 # Core functions
│   ├── fmt/                  # Formatting
│   ├── io/                   # Input/output operations
│   ├── math/                 # Mathematical operations
│   ├── memory/               # Memory management
│   ├── net/                  # Network operations
│   ├── regex/                # Regular expressions
│   ├── strings/              # String operations
│   ├── time/                 # Time operations
│   └── vulkan/               # Vulkan bindings
├── tests/                    # Tests
│   ├── compiler/             # Compiler tests
│   ├── basic/                # Basic language feature tests
│   └── ...                   # Other test files
├── website/                  # Website
│   ├── content/              # Content
│   ├── static/               # Static files
│   └── templates/            # Templates
├── .gitignore                # Git ignore list
├── go.mod                    # Go module definition
├── go.sum                    # Go dependency checksum
├── LICENSE                   # License file
└── README.md                 # Main README file
```

## Directory Descriptions

### cmd/

This directory contains the command-line applications of the GO-Minus project. Each subdirectory represents a separate application.

- **gominus**: GO-Minus compiler
- **gompm**: GO-Minus package manager
- **gomfmt**: GO-Minus code formatting tool
- **gomdoc**: GO-Minus documentation tool
- **gomtest**: GO-Minus testing tool
- **gomlsp**: GO-Minus language server

### docs/

This directory contains the documentation of the GO-Minus project.

- **tutorial/**: Tutorials and getting started guides
- **reference/**: Language reference and API documentation
- **images/**: Images used in documentation

### examples/

This directory contains example code written in the GO-Minus language.

- **basic/**: Simple examples showing basic language features
- **advanced/**: Complex examples showing advanced language features
- **vulkan/**: Examples using the Vulkan API
- **memory/**: Examples showing memory management features

### internal/

This directory contains the internal components of the GO-Minus compiler. These packages are not directly used from outside the project.

- **ast/**: Abstract syntax tree definitions and operations
- **codegen/**: Machine code generation
- **irgen/**: LLVM IR generation
- **lexer/**: Lexical analysis
- **optimizer/**: Code optimization
- **parser/**: Syntax analysis
- **semantic/**: Semantic analysis
- **token/**: Token definitions

### pkg/

This directory contains the public packages of the GO-Minus project. These packages can be used by other projects.

- **compiler/**: Compiler API
- **runtime/**: Runtime API

### stdlib/

This directory contains the standard library of the GO-Minus language.

- **async/**: Packages for asynchronous operations
- **concurrent/**: Packages for concurrency
- **container/**: Data structures
- **core/**: Core functions
- **fmt/**: Formatting operations
- **io/**: Input/output operations
- **math/**: Mathematical operations
- **memory/**: Memory management
- **net/**: Network operations
- **regex/**: Regular expressions
- **strings/**: String operations
- **time/**: Time operations
- **vulkan/**: Vulkan API bindings

### tests/

This directory contains the test files of the GO-Minus project.

- **compiler/**: Compiler tests
- **basic/**: Basic language feature tests

### website/

This directory contains the website files of the GO-Minus project.

- **content/**: Website content
- **static/**: Static files (CSS, JS, images)
- **templates/**: HTML templates

## Files

- **.gitignore**: Specifies files to be ignored by Git
- **go.mod**: Go module definition
- **go.sum**: Go dependency checksum
- **LICENSE**: License file
- **README.md**: Main README file
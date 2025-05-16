# GO-Minus Development Tools

This document describes the tools you can use when developing with the GO-Minus programming language. These tools help you perform operations such as compiling, testing, documenting, formatting, and package management of GO-Minus code.

## Contents

1. [Compiler (gominus)](#compiler-gominus)
2. [Package Manager (gompm)](#package-manager-gompm)
3. [Testing Tool (gomtest)](#testing-tool-gomtest)
4. [Documentation Tool (gomdoc)](#documentation-tool-gomdoc)
5. [Code Formatting Tool (gomfmt)](#code-formatting-tool-gomfmt)
6. [Language Server (gomlsp)](#language-server-gomlsp)
7. [Debugging Tool (gomdebug)](#debugging-tool-gomdebug)
8. [IDE Integration](#ide-integration)

## Compiler (gominus)

The GO-Minus compiler compiles GO-Minus source code to create executable files.

### Usage

```bash
# Compile a single file
gominus file.gom

# Compile multiple files
gominus file1.gom file2.gom

# Specify output file
gominus -o program file.gom

# Specify optimization level
gominus -O2 file.gom

# Add debugging information
gominus -g file.gom

# Enable warnings
gominus -Wall file.gom

# Compile for a specific target platform
gominus -target=x86_64-linux file.gom

# Compile as a library
gominus -lib file.gom

# Compile and run
gominus -run file.gom
```

### Output Formats

The GO-Minus compiler supports the following output formats:

- Executable file (default)
- Object file (`.o`)
- LLVM IR (`.ll`)
- Library (`.a`, `.so`, `.dll`)

## Package Manager (gompm)

The GO-Minus Package Manager (gompm) performs downloading, installing, removing, updating, and searching operations for GO-Minus packages. This tool facilitates the management of packages in the GO-Minus ecosystem and automates dependency resolution processes.

### Features

- Creating new projects
- Installing and removing packages
- Updating packages
- Listing installed packages
- Searching in the package repository
- Viewing package information
- Dependency management
- Development dependencies support
- Version constraints
- Package repository integration

### Usage

```bash
# Start a new project
gompm init [project-name]

# Install a package
gompm install package-name

# Install a specific version
gompm install package-name@1.0.0

# Install as a development dependency
gompm install package-name --dev

# Remove a package
gompm remove package-name

# Update all packages
gompm update

# List installed packages
gompm list

# Search in the package repository
gompm search query

# Show package information
gompm info package-name

# Show help message
gompm help

# Show version information
gompm version
```

### Package Configuration File

The package configuration file (`gompm.json`) contains the project's metadata and dependencies:

```json
{
  "name": "project-name",
  "version": "1.0.0",
  "description": "Project description",
  "author": "Author Name",
  "license": "MIT",
  "dependencies": {
    "package1": "1.0.0",
    "package2": "^2.0.0"
  },
  "devDependencies": {
    "test-package": "1.0.0"
  },
  "keywords": ["keyword", "term"],
  "homepage": "https://example.com",
  "repository": "https://github.com/user/project"
}
```

### Version Specification

Package versions are specified according to [Semantic Versioning](https://semver.org/) rules. Version specification formats:

- `1.0.0`: Exactly version 1.0.0
- `^1.0.0`: 1.0.0 or higher, but less than 2.0.0
- `~1.0.0`: 1.0.0 or higher, but less than 1.1.0
- `>=1.0.0`: 1.0.0 or higher
- `<=1.0.0`: 1.0.0 or lower
- `1.0.0 - 2.0.0`: Between 1.0.0 and 2.0.0 (inclusive)
- `latest`: Latest version

### Package Repository

GO-Minus packages are stored in a central package repository (https://repo.gominus.org). This repository contains the metadata and source code of packages. gompm downloads and installs packages from this repository.

### Dependency Resolution

gompm automatically resolves and installs package dependencies. The dependency resolution algorithm selects the most appropriate versions considering version constraints. This helps prevent version conflicts between packages.

### Security

gompm uses digital signatures to verify the integrity of packages. This helps prevent the installation of malicious packages. It also performs a security scan to check for vulnerabilities in packages.

### Example: Creating a New Project

```bash
# Create a new project
gompm init my-project

# Go to the project directory
cd my-project

# Install packages
gompm install logger@1.0.0
gompm install http-client@^2.0.0
gompm install test-framework@latest --dev

# List installed packages
gompm list
```

### Example: Viewing Package Information

```bash
# View package information
gompm info http-client

# Output:
# Package: http-client
# Version: 2.1.0
# Description: HTTP client library
# Author: GO-Minus Team
# License: MIT
# Dependencies:
#   - logger@1.0.0
#   - url-parser@^1.5.0
# Keywords: http, client, request, response
# Homepage: https://example.com/http-client
# Repository: https://github.com/gominus/http-client
```

For more information, see the [Package Manager Documentation](../cmd/gompm/README.md).

## Testing Tool (gomtest)

The GO-Minus Testing Tool (gomtest) is used to test GO-Minus code. It finds test files, runs tests, and reports results.

### Usage

```bash
# Run tests in the current directory
gomtest

# Run tests with detailed output
gomtest -v

# Run tests in subdirectories as well
gomtest -r

# Run tests matching a specific pattern
gomtest -pattern=TestAdd

# Run benchmark tests
gomtest -bench

# Perform code coverage analysis
gomtest -cover

# Run tests in specified directories
gomtest ./pkg ./internal
```

### Test Files

Test files are files ending with the `_test.gom` suffix. Test functions start with the `Test` prefix and take a `*testing.T` parameter:

```go
// math_test.gom
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; expected 5", result)
    }
}
```

## Documentation Tool (gomdoc)

The GO-Minus Documentation Tool (gomdoc) is used to document GO-Minus code. It analyzes comments and structures in the code to create documentation.

### Usage

```bash
# Document the package in the current directory
gomdoc

# Create documentation in HTML format
gomdoc -html

# Create documentation in Markdown format
gomdoc -markdown

# Save documentation to a specific directory
gomdoc -output=docs

# Start a documentation server
gomdoc -server

# Start a documentation server on a specific port
gomdoc -server -port=8080

# Document specified packages
gomdoc ./pkg ./internal
```

### Documentation Comments

GO-Minus documentation comments start with `//` or `/* */` and are placed immediately before the documented item:

```go
// Add adds two numbers and returns the result.
//
// Parameters:
//   - a: First number
//   - b: Second number
//
// Return value:
//   - The sum of the two numbers
func Add(a, b int) int {
    return a + b
}
```

## Code Formatting Tool (gomfmt)

The GO-Minus Code Formatting Tool (gomfmt) is used to format GO-Minus code in a standard way. It makes the code style consistent.

### Usage

```bash
# Format GO-Minus files in the current directory
gomfmt

# Write changes to files
gomfmt -w

# Show changes in diff format
gomfmt -d

# List files that need formatting
gomfmt -l

# Format files in subdirectories as well
gomfmt -r

# Simplify code
gomfmt -s

# Format specified files
gomfmt file1.gom file2.gom

# Format files in specified directories
gomfmt ./pkg ./internal
```

## Language Server (gomlsp)

The GO-Minus Language Server (gomlsp) provides a Language Server Protocol (LSP) implementation for GO-Minus. This is used for integration with IDEs and text editors.

### Features

- Code completion
- Error and warning display
- Go to definition
- Rename
- Code formatting
- Code folding
- Symbol search
- Find references
- Code actions

### Usage

```bash
# Start the language server
gomlsp

# Start on a specific port
gomlsp --port=8080

# Enable verbose logging
gomlsp --verbose
```

## Debugging Tool (gomdebug)

The GO-Minus Debugging Tool (gomdebug) is used to debug GO-Minus programs.

### Usage

```bash
# Start the program in debug mode
gomdebug program

# Start the program with specific arguments
gomdebug program arg1 arg2

# Add a breakpoint at a specific file and line
gomdebug --break=file.gom:10 program

# Start the debug server
gomdebug --server program

# Start the debug server on a specific port
gomdebug --server --port=8080 program
```

### Debugging Commands

During a debugging session, you can use the following commands:

- `break file:line`: Add a breakpoint
- `continue`: Continue execution
- `step`: Step one line (enter functions)
- `next`: Step one line (skip functions)
- `print expression`: Evaluate and print an expression
- `backtrace`: Show call stack
- `frame n`: Go to the nth call frame
- `list`: Show source code
- `quit`: Exit the debugger

## IDE Integration

GO-Minus can be integrated with various IDEs and text editors.

### Visual Studio Code

The GO-Minus extension for VS Code provides the following features:

- Syntax highlighting
- Code completion
- Error and warning display
- Go to definition
- Rename
- Code formatting
- Code folding
- Symbol search
- Find references
- Code actions

### JetBrains IDEs

The GO-Minus extension for JetBrains IDEs (IntelliJ IDEA, GoLand, etc.) provides the following features:

- Syntax highlighting
- Code completion
- Error and warning display
- Go to definition
- Rename
- Code formatting
- Code folding
- Symbol search
- Find references
- Code actions

### Vim/Neovim

The GO-Minus extension for Vim/Neovim provides the following features:

- Syntax highlighting
- Code completion (with coc.nvim or YouCompleteMe)
- Error and warning display
- Go to definition
- Code formatting

### Emacs

The GO-Minus mode for Emacs provides the following features:

- Syntax highlighting
- Code completion (with company-mode)
- Error and warning display (with flycheck)
- Go to definition
- Code formatting

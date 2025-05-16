# GO-Minus Development Tools

This directory contains development tools for the GO-Minus programming language. These tools include various utilities that can be used when developing with GO-Minus.

## Tools

### gompm - GO-Minus Package Manager

A tool used to manage GO-Minus packages. It performs package installation, removal, updating, and searching operations.

```bash
# Start a new GO-Minus project
gompm -init myproject

# Install packages
gompm -install fmt strings

# Remove a package
gompm -uninstall fmt

# Update packages
gompm -update

# List installed packages
gompm -list

# Search for a package
gompm -search json
```

### gomtest - GO-Minus Testing Tool

A tool used to test GO-Minus code. It finds test files, runs tests, and reports results.

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

### gomdoc - GO-Minus Documentation Tool

A tool used to document GO-Minus code. It analyzes comments and structures in the code to create documentation.

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

### gomfmt - GO-Minus Code Formatting Tool

A tool used to format GO-Minus code in a standard way. It makes the code style consistent.

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

## Installation

To install GO-Minus development tools, after installing the GO-Minus compiler, you can run the following commands:

```bash
# Install the GO-Minus package manager
go build -o gompm ./tools/gompm

# Install the GO-Minus testing tool
go build -o gomtest ./tools/gomtest

# Install the GO-Minus documentation tool
go build -o gomdoc ./tools/gomdoc

# Install the GO-Minus code formatting tool
go build -o gomfmt ./tools/gomfmt
```

You can copy the generated executable files to a directory in your PATH environment variable.
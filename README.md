# GO-Minus Programming Language

🎉 **BÜYÜK HABER: İlk Çalışan Versiyon Tamamlandı!** 🎉

GO-Minus is a programming language that includes all features of the Go programming language and extends it with C++-like features (classes, templates, exception handling, etc.). GO-Minus files use the `.gom` extension.

## 📊 Proje Durumu (Son Güncelleme)

**Tamamlanma Oranı**: %25-30 🚀
**Durum**: ✅ Temel Go syntax'ı başarıyla çalışıyor!
**Son Başarı**: `fmt.Println("Hello, World!")` tam olarak parse ediliyor ve LLVM IR üretiyor

### ✅ Çalışan Özellikler:
- ✅ **Parser**: Function calls, member access, string literals
- ✅ **Semantic Analysis**: Built-in functions, package resolution, type checking
- ✅ **IR Generation**: LLVM IR üretimi, function calls, string literals
- ✅ **Standard Library**: fmt, os, io, strings, math packages
- ✅ **Package System**: Import statements, package.function calls

### 🎯 Sonraki Hedefler:
- 🔄 Executable generation (LLVM toolchain setup)
- 🔄 Variables ve data types (int, float, bool)
- 🔄 Arithmetic operations (+, -, *, /)
- 🔄 Control flow (if/else, loops)

## 🎯 Purpose

The main purpose of the GO-Minus language is to combine Go's simplicity, fast compilation times, and powerful concurrency model with C++'s low-level system control, performance optimizations, template metaprogramming (simplified), and rich OOP capabilities. It aims to find the "sweet spot" for both high-performance system programming and rapid application development.

## ✨ Features

- **Go Compatibility**: Supports all features of Go
- **Classes and Inheritance**: C++-like class and inheritance support
- **Templates**: Template support for generic programming
- **Exception Handling**: Try-catch-finally structures
- **Access Modifiers**: Public, private, protected access control
- **LLVM Based**: Strong optimization and multi-platform support
- **Rich Standard Library**: Basic data structures, I/O operations, concurrency support
- **Development Tools**: Package manager, testing tool, documentation tool, code formatting tool
- **IDE Integration**: Plugins for VS Code, JetBrains IDEs, Vim/Neovim, Emacs

## 📂 Project Structure

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
└── README.md                 # This file
```

## 🚀 Getting Started

### Prerequisites

- Go 1.20 or later
- LLVM 14 or later (for code generation)
- Git

### Installation

```bash
# Clone the GO-Minus compiler
git clone https://github.com/inkbytefo/go-minus.git
cd go-minus

# Download dependencies
make deps

# Build the GO-Minus compiler and tools
make build-all

# Install binaries (optional)
make install
```

### Quick Build

```bash
# Build just the compiler
make build

# Build all development tools
make build-tools

# Run tests
make test

# Run linter and code quality checks
make check
```

### Hello World (✅ ÇALIŞIYOR!)

```go
// main.gom
package main

import "fmt"

func main() {
    fmt.Println("Hello, GO-Minus!")
}
```

```bash
# Şu anda çalışan komutlar:
gominus main.gom                    # ✅ LLVM IR üretir (main.ll)

# Yakında gelecek:
# gominus run main.gom              # 🔄 Executable generation (LLVM araçları gerekli)
```

**Mevcut Çıktı**: `main.ll` dosyası başarıyla oluşturuluyor ve geçerli LLVM IR içeriyor!

### Class Example

```go
// person.gom
package main

import "fmt"

class Person {
    private:
        string name
        int age

    public:
        Person(string name, int age) {
            this.name = name
            this.age = age
        }

        string getName() {
            return this.name
        }

        int getAge() {
            return this.age
        }

        void birthday() {
            this.age++
        }
}

func main() {
    person := Person("John", 30)
    fmt.Println("Name:", person.getName())
    fmt.Println("Age:", person.getAge())

    person.birthday()
    fmt.Println("New age:", person.getAge())
}
```

```bash
# Compile and run the program
gominus run person.gom
```

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/lexer -v
go test ./internal/parser -v
go test ./internal/semantic -v

# Run integration tests
go test ./test -v

# Run benchmarks
go test -bench=. ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run linter and code quality checks
make check
```

## 🛠️ Development Tools

The GO-Minus language provides the following development tools:

- **gompm**: GO-Minus Package Manager (`cmd/gompm`)
- **gomtest**: GO-Minus Testing Tool (`cmd/gomtest`)
- **gomdoc**: GO-Minus Documentation Tool (`cmd/gomdoc`)
- **gomfmt**: GO-Minus Code Formatting Tool (`cmd/gomfmt`)
- **gomlsp**: GO-Minus Language Server (`cmd/gomlsp`)
- **gomdebug**: GO-Minus Debugging Tool (`cmd/gomdebug`)

## 🔌 IDE Integration

The GO-Minus language provides plugins for the following IDEs:

- **VS Code**: [GO-Minus VS Code Plugin](ide/vscode/README.md)
- **JetBrains IDEs**: [GO-Minus JetBrains Plugin](ide/jetbrains/README.md)
- **Vim/Neovim**: [GO-Minus Vim Plugin](ide/vim/README.md)
- **Emacs**: [GO-Minus Emacs Plugin](ide/emacs/README.md)

## 📚 Documentation

- [Getting Started Guide](docs/tutorial/getting-started.md)
- [Language Reference](docs/reference/README.md)
- [Standard Library](stdlib/README.md)
- [Tutorials](docs/tutorial/README.md)
- [Best Practices](docs/best-practices.md)
- [FAQ](docs/faq.md)
- [Why GO-Minus?](docs/why-gominus.md)
- [Progress Report](docs/progress.md)
- [Development Plan](docs/development-plan.md)

## 👥 Community

- [GitHub](https://github.com/inkbytefo/go-minus)
- [Website](website/index.html)
- [Discord](https://discord.gg/gominus)
- [Forum](https://forum.gominus.org)

## 🤝 Contributing

To contribute to the GO-Minus project, please read the [contribution guide](CONTRIBUTING.md). All contributors are considered to have agreed to our [code of conduct](CODE_OF_CONDUCT.md).

## 📄 License

The GO-Minus language is licensed under the [MIT License](LICENSE).
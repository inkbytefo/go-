# GO-Minus Programming Language Makefile
# This Makefile provides common development tasks for the GO-Minus project

# Variables
BINARY_NAME=gominus
PACKAGE_NAME=github.com/inkbytefo/go-minus
VERSION?=0.1.0
BUILD_TIME=$(shell date +%Y-%m-%dT%H:%M:%S%z)
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.gitCommit=$(GIT_COMMIT)"

# Go related variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOLINT=golangci-lint

# Directories
BUILD_DIR=build
DIST_DIR=dist
TOOLS_DIR=tools

# Default target
.PHONY: all
all: clean deps test build

# Help target
.PHONY: help
help:
	@echo "GO-Minus Development Commands:"
	@echo ""
	@echo "Building:"
	@echo "  build          Build the main compiler"
	@echo "  build-all      Build all tools and binaries"
	@echo "  build-tools    Build development tools"
	@echo "  install        Install binaries to GOPATH/bin"
	@echo ""
	@echo "Testing:"
	@echo "  test           Run all tests"
	@echo "  test-verbose   Run tests with verbose output"
	@echo "  test-coverage  Run tests with coverage report"
	@echo "  bench          Run benchmark tests"
	@echo ""
	@echo "Code Quality:"
	@echo "  fmt            Format Go code"
	@echo "  lint           Run linter"
	@echo "  vet            Run go vet"
	@echo "  check          Run all code quality checks"
	@echo ""
	@echo "Dependencies:"
	@echo "  deps           Download dependencies"
	@echo "  deps-update    Update dependencies"
	@echo "  deps-tidy      Clean up dependencies"
	@echo ""
	@echo "Cleanup:"
	@echo "  clean          Clean build artifacts"
	@echo "  clean-all      Clean everything including dependencies"
	@echo ""
	@echo "Release:"
	@echo "  release        Build release binaries for all platforms"
	@echo "  package        Create distribution packages"

# Build targets
.PHONY: build
build: deps
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/gominus

.PHONY: build-all
build-all: build build-tools

.PHONY: build-tools
build-tools: deps
	@echo "Building development tools..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/gompm ./cmd/gompm
	$(GOBUILD) -o $(BUILD_DIR)/gomtest ./cmd/gomtest
	$(GOBUILD) -o $(BUILD_DIR)/gomdoc ./cmd/gomdoc
	$(GOBUILD) -o $(BUILD_DIR)/gomfmt ./cmd/gomfmt
	$(GOBUILD) -o $(BUILD_DIR)/gomlsp ./cmd/gomlsp
	$(GOBUILD) -o $(BUILD_DIR)/gomdebug ./cmd/gomdebug

.PHONY: install
install: build-all
	@echo "Installing binaries..."
	$(GOCMD) install $(LDFLAGS) ./cmd/gominus
	$(GOCMD) install ./cmd/gompm
	$(GOCMD) install ./cmd/gomtest
	$(GOCMD) install ./cmd/gomdoc
	$(GOCMD) install ./cmd/gomfmt
	$(GOCMD) install ./cmd/gomlsp
	$(GOCMD) install ./cmd/gomdebug

# Test targets
.PHONY: test
test: deps
	@echo "Running tests..."
	$(GOTEST) -race -short ./...

.PHONY: test-verbose
test-verbose: deps
	@echo "Running tests with verbose output..."
	$(GOTEST) -race -v ./...

.PHONY: test-coverage
test-coverage: deps
	@echo "Running tests with coverage..."
	@mkdir -p $(BUILD_DIR)
	$(GOTEST) -race -coverprofile=$(BUILD_DIR)/coverage.out ./...
	$(GOCMD) tool cover -html=$(BUILD_DIR)/coverage.out -o $(BUILD_DIR)/coverage.html
	@echo "Coverage report generated: $(BUILD_DIR)/coverage.html"

.PHONY: bench
bench: deps
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem ./...

# Code quality targets
.PHONY: fmt
fmt:
	@echo "Formatting Go code..."
	$(GOFMT) -s -w .

.PHONY: lint
lint:
	@echo "Running linter..."
	@if command -v $(GOLINT) >/dev/null 2>&1; then \
		$(GOLINT) run; \
	else \
		echo "golangci-lint not found. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

.PHONY: vet
vet:
	@echo "Running go vet..."
	$(GOCMD) vet ./...

.PHONY: check
check: fmt vet lint test

# Dependency targets
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download

.PHONY: deps-update
deps-update:
	@echo "Updating dependencies..."
	$(GOMOD) get -u ./...
	$(GOMOD) tidy

.PHONY: deps-tidy
deps-tidy:
	@echo "Tidying dependencies..."
	$(GOMOD) tidy

# Clean targets
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)

.PHONY: clean-all
clean-all: clean
	@echo "Cleaning dependencies..."
	$(GOMOD) clean -cache

# Release targets
.PHONY: release
release: clean deps test
	@echo "Building release binaries..."
	@mkdir -p $(DIST_DIR)
	
	# Linux AMD64
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/gominus
	
	# Linux ARM64
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm64 ./cmd/gominus
	
	# Windows AMD64
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd/gominus
	
	# macOS AMD64
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/gominus
	
	# macOS ARM64
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64 ./cmd/gominus

.PHONY: package
package: release
	@echo "Creating distribution packages..."
	@cd $(DIST_DIR) && \
	for binary in $(BINARY_NAME)-*; do \
		if [[ $$binary == *.exe ]]; then \
			zip $${binary%.exe}.zip $$binary; \
		else \
			tar -czf $$binary.tar.gz $$binary; \
		fi; \
	done

# Development helpers
.PHONY: dev-setup
dev-setup:
	@echo "Setting up development environment..."
	$(GOGET) github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GOMOD) download

.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: debug
debug: deps
	@echo "Building debug version..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -gcflags="all=-N -l" -o $(BUILD_DIR)/$(BINARY_NAME)-debug ./cmd/gominus

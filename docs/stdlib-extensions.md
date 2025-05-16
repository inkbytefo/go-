# GO-Minus Standard Library Extensions

This document describes the new modules and extensions added to the standard library of the GO-Minus programming language.

## Overview

The GO-Minus standard library is based on Go's standard library and has been extended using GO-Minus features such as classes, templates, and exception handling. With recent updates, the following new modules and extensions have been added:

1. **Trie Implementation**
2. **Buffered IO Implementation**
3. **Regex Package**

## 1. Trie Implementation

Trie (also known as a prefix tree or digital tree) is a tree data structure used to efficiently store and search for string keys. The Trie implementation added to GO-Minus's container package is ideal for operations such as prefix-based searching and autocomplete.

### Features

- **Generic Type Support**: Can be used for any type
- **Word Insertion, Search, and Deletion**: Basic Trie operations
- **Prefix Search**: Finding words that start with a specific prefix
- **Listing All Words**: Listing all words in the Trie
- **Empty Check and Size Calculation**: Checking if the Trie is empty and calculating its size

### Example Usage

```go
import "container/trie"

// Create a Trie for string values
t := trie.Trie.New<string>()

// Add words
t.Insert("apple", "apple")
t.Insert("banana", "banana")
t.Insert("application", "application")

// Search for a word
value, found := t.Search("apple")
if found {
    fmt.Println("apple:", value) // Prints "apple"
}

// Prefix check
if t.StartsWith("app") {
    fmt.Println("There are words starting with the prefix 'app'")
}

// Get all words starting with a specific prefix
appWords := t.GetWordsWithPrefix("app")
for word, value := range appWords {
    fmt.Printf("%s: %s\n", word, value)
}
```

### Performance

The Trie data structure has the following complexities for string keys:

- **Insertion**: O(m), m = key length
- **Search**: O(m), m = key length
- **Deletion**: O(m), m = key length
- **Prefix Search**: O(m), m = prefix length
- **Finding All Words Starting with a Specific Prefix**: O(n), n = number of matching words

### Implementation Details

The Trie implementation is implemented in the `stdlib/container/trie/trie.gom` file. The implementation consists of the following components:

- **TrieNode<T>**: Class representing each node in the Trie tree
- **Trie<T>**: Class representing the Trie data structure

```go
// TrieNode represents a trie node.
class TrieNode<T> {
    private:
        map[rune]TrieNode<T> children
        bool isEndOfWord
        T value
        bool hasValue

    public:
        // Create a new node
        static func New<T>() *TrieNode<T> {
            node := new TrieNode<T>()
            node.children = make(map[rune]TrieNode<T>)
            node.isEndOfWord = false
            node.hasValue = false
            return node
        }

        // ... other methods
}

// Trie represents a prefix tree.
class Trie<T> {
    private:
        TrieNode<T> root
        int size

    public:
        // Create a new Trie
        static func New<T>() *Trie<T> {
            t := new Trie<T>()
            t.root = *TrieNode.New<T>()
            return t
        }

        // ... other methods
}
```

## 2. Buffered IO Implementation

Buffered I/O is a technique used to improve performance when accessing slow I/O sources such as disk or network. The Buffered IO implementation added to GO-Minus's io package groups small read/write operations into larger blocks, reducing the number of system calls and improving overall performance.

### Features

- **Buffered Reading**: BufferedReader for efficient reading operations
- **Buffered Writing**: BufferedWriter for efficient writing operations
- **Customizable Buffer Size**: Adjusting buffer size according to needs
- **Line-by-Line Reading**: Reading text files line by line
- **Byte and String Writing**: Writing support for both byte arrays and strings

### Example Usage

```go
import (
    "io"
    "io/buffered"
)

// BufferedReader example
file := io.Open("input.txt")
defer file.Close()

reader := buffered.BufferedReader.New(file, 4096)

// Reading line by line
for {
    line, err := reader.ReadLine()
    if err == io.EOF {
        break
    }
    if err != nil {
        // Error handling
        break
    }

    fmt.Println(line)
}

// BufferedWriter example
outFile := io.Create("output.txt")
defer outFile.Close()

writer := buffered.BufferedWriter.New(outFile, 4096)

// Writing data
writer.WriteString("Hello, World!\n")
writer.WriteString("This is a test.\n")

// Flush the buffer
writer.Flush()
```

### Performance

Buffered I/O provides significant performance improvements, especially in situations where small read/write operations are performed frequently. For example, when reading a file line by line or writing in small chunks, using buffered I/O can be much faster than unbuffered I/O.

### Implementation Details

The Buffered IO implementation is implemented in the `stdlib/io/buffered/buffered.gom` file. The implementation consists of the following components:

- **BufferedReader**: Class for buffered reading operations
- **BufferedWriter**: Class for buffered writing operations

```go
// BufferedReader is used for buffered reading operations.
class BufferedReader {
    private:
        io.Reader reader
        []byte buffer
        int bufferSize
        int readPos
        int writePos
        bool eof

    public:
        // Create a new BufferedReader
        static func New(reader io.Reader, bufferSize int) *BufferedReader {
            // ... implementation
        }

        // ... other methods
}

// BufferedWriter is used for buffered writing operations.
class BufferedWriter {
    private:
        io.Writer writer
        []byte buffer
        int bufferSize
        int count

    public:
        // Create a new BufferedWriter
        static func New(writer io.Writer, bufferSize int) *BufferedWriter {
            // ... implementation
        }

        // ... other methods
}
```

## 3. Regex Package

Regular Expressions provide a special language and syntax used to search for, match, and replace specific patterns in text. The Regex package added to GO-Minus can be used in many scenarios such as text processing, form validation, data extraction, and transformation.

### Features

- **Regular Expression Pattern Compilation**: Compiling and reusing patterns
- **Text Matching**: Checking if a text matches a pattern
- **Finding All Matches**: Finding all matches in a text
- **Text Replacement**: Replacing matched text with other text
- **Text Splitting**: Splitting a text according to a specific pattern
- **Case Sensitive and Insensitive Modes**: Controlling case sensitivity
- **Multiline Mode**: Matching line beginnings and endings in multiline texts

### Example Usage

```go
import "regex"

// Compile a regular expression pattern
pattern := regex.Compile("hello")

// Text matching
if pattern.Match("hello world") {
    fmt.Println("Match found!")
}

// Case insensitive pattern
patternIgnoreCase := regex.CompileIgnoreCase("hello")

// Find all matches
matches := patternIgnoreCase.FindAll("Hello, hello, HELLO!")
fmt.Println("Number of matches:", len(matches))

// Text replacement
result := patternIgnoreCase.Replace("Hello, hello, HELLO!", "hi")
fmt.Println(result) // "hi, hi, hi!"

// Text splitting
parts := regex.Split(",", "apple,banana,orange")
for _, part := range parts {
    fmt.Println(part)
}
```

### Implementation Details

The Regex package is implemented in the `stdlib/regex/regex.gom` file. The package consists of the following components:

- **RegexPattern**: Class representing a compiled regular expression pattern
- **Compile, CompileIgnoreCase, CompileMultiline**: Pattern compilation functions
- **Match, Replace, Split**: Helper functions

```go
// RegexPattern represents a compiled regular expression pattern.
class RegexPattern {
    private:
        string pattern
        bool caseSensitive
        bool multiline
        bool compiled
        bool hasSpecialChars

    public:
        // Create a new RegexPattern
        static func New(pattern string, caseSensitive bool, multiline bool) *RegexPattern {
            // ... implementation
        }

        // ... other methods
}

// Compile compiles a regular expression pattern.
func Compile(pattern string) *RegexPattern {
    return RegexPattern.New(pattern, true, false)
}

// CompileIgnoreCase compiles a regular expression pattern with case insensitivity.
func CompileIgnoreCase(pattern string) *RegexPattern {
    return RegexPattern.New(pattern, false, false)
}

// ... other functions
```

## Conclusion

These new modules and extensions added to the GO-Minus standard library enhance the capabilities of the GO-Minus programming language and provide developers with more powerful tools. These extensions enable GO-Minus to be used in a wider range of applications.

For more information and examples, you can refer to the following documents:

- [Trie Example](examples/container/trie-example.md)
- [Buffered IO Example](examples/io/buffered-io-example.md)
- [Regex Example](examples/regex/regex-example.md)

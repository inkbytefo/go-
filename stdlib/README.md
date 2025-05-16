# GO-Minus Standard Library

This directory contains the standard library for the GO-Minus programming language. The GO-Minus standard library is based on Go's standard library and has been extended using GO-Minus features such as classes, templates, and exception handling.

## Packages

### core
Core package for basic data types and functions.

### container
Packages for data structures:
- **list**: Doubly linked list implementation
- **vector**: Dynamic array implementation
- **map**: Key-value mapping implementation
- **set**: Set implementation

### concurrent
Packages for concurrency:
- **channel**: Channels for communication between goroutines
- **mutex**: Mutexes for mutual exclusion
- **waitgroup**: WaitGroup for waiting for goroutines to complete

### fmt
Package for formatted input/output operations.

### io
Basic interfaces and functions for input/output operations.

### math
Mathematical functions and constants.

### strings
String processing functions.

## Usage

To use packages from the GO-Minus standard library, simply import the relevant package in your GO-Minus program:

```go
import "fmt"
import "container/list"
import "concurrent/channel"

func main() {
    // Use the fmt package
    fmt.Println("Hello, World!")

    // Use the list package
    l := list.New<int>()
    l.PushBack(10)
    l.PushBack(20)

    // Use the channel package
    ch := channel.New<string>(1)
    ch.Send("Hello")
    msg := ch.Receive()
    fmt.Println(msg)
}
```

## Development

The GO-Minus standard library is continuously being expanded along with the development of the GO-Minus language. New packages and functions will continue to be added.
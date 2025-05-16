# GO-Minus Getting Started Guide

This guide explains the steps needed to start using the GO-Minus programming language. GO-Minus is a programming language that includes all features of the Go programming language and extends it with C++-like features.

## Contents

1. [GO-Minus Installation](#go-minus-installation)
2. [First GO-Minus Program](#first-go-minus-program)
3. [Basic Syntax](#basic-syntax)
4. [Classes and Objects](#classes-and-objects)
5. [Templates](#templates)
6. [Exception Handling](#exception-handling)
7. [Packages and Modules](#packages-and-modules)
8. [Compiling and Running](#compiling-and-running)
9. [IDE Integration](#ide-integration)
10. [Next Steps](#next-steps)

## GO-Minus Installation

Follow these steps to install GO-Minus:

### Prerequisites

- Go 1.18 or higher
- LLVM 14.0 or higher
- Git

### Installation Steps

1. Clone the GO-Minus repository:

```bash
git clone https://github.com/gominus/gominus.git
cd gominus
```

2. Build the GO-Minus compiler:

```bash
go build -o gominus ./cmd/gominus
```

3. Add the compiler to PATH:

For Windows:
```bash
copy gominus.exe C:\Windows\System32\
```

For Linux/macOS:
```bash
sudo cp gominus /usr/local/bin/
```

4. Verify the installation:

```bash
gominus version
```

## First GO-Minus Program

Let's write your first program with GO-Minus. Open a text editor and save the following code in a file named `hello.gom`:

```go
// hello.gom
package main

import "fmt"

func main() {
    fmt.Println("Hello, GO-Minus!")
}
```

To compile and run the program:

```bash
gominus run hello.gom
```

Output:
```
Hello, GO-Minus!
```

## Basic Syntax

GO-Minus is based on Go's syntax and adds C++-like features.

### Variables and Constants

```go
// Variable declaration
var x int = 10
var y = 20 // Type inference
z := 30    // Short variable declaration

// Constant declaration
const pi = 3.14159
const (
    a = 1
    b = 2
)
```

### Control Structures

```go
// If-else
if x > 0 {
    fmt.Println("Positive")
} else if x < 0 {
    fmt.Println("Negative")
} else {
    fmt.Println("Zero")
}

// For loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Range loop
nums := []int{1, 2, 3, 4, 5}
for i, num := range nums {
    fmt.Printf("Index: %d, Value: %d\n", i, num)
}

// Switch
switch day {
case "Monday":
    fmt.Println("First day of the week")
case "Friday":
    fmt.Println("Last workday of the week")
default:
    fmt.Println("Another day")
}
```

### Functions

```go
// Basic function
func add(a int, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero error")
    }
    return a / b, nil
}

// Variadic parameters
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

## Classes and Objects

GO-Minus provides C++-like class and object support.

```go
// Class definition
class Person {
    private:
        string name
        int age

    public:
        // Constructor
        Person(string name, int age) {
            this.name = name
            this.age = age
        }

        // Getter methods
        string getName() {
            return this.name
        }

        int getAge() {
            return this.age
        }

        // Method
        void birthday() {
            this.age++
        }

        // Static method
        static Person createDefault() {
            return Person("John Doe", 30)
        }
}

// Class usage
func main() {
    // Object creation
    person := Person("John", 30)

    // Method calling
    fmt.Println("Name:", person.getName())
    fmt.Println("Age:", person.getAge())

    person.birthday()
    fmt.Println("New age:", person.getAge())

    // Static method calling
    defaultPerson := Person.createDefault()
    fmt.Println("Default name:", defaultPerson.getName())
}
```

### Inheritance

```go
// Base class
class Animal {
    protected:
        string name

    public:
        Animal(string name) {
            this.name = name
        }

        virtual string makeSound() {
            return "..."
        }
}

// Derived class
class Dog : Animal {
    private:
        string breed

    public:
        Dog(string name, string breed) : Animal(name) {
            this.breed = breed
        }

        override string makeSound() {
            return "Woof!"
        }

        string getBreed() {
            return this.breed
        }
}

// Inheritance usage
func main() {
    dog := Dog("Buddy", "Golden Retriever")
    fmt.Println("Sound:", dog.makeSound())
    fmt.Println("Breed:", dog.getBreed())

    // Polymorphism
    var animal Animal = dog
    fmt.Println("Polymorphic sound:", animal.makeSound())
}
```

## Templates

GO-Minus provides C++-like template support.

```go
// Template class
template<T>
class Stack {
    private:
        T[] items
        int size

    public:
        Stack() {
            this.items = T[]{}
            this.size = 0
        }

        void push(T item) {
            this.items = append(this.items, item)
            this.size++
        }

        T pop() {
            if this.size == 0 {
                throw new Exception("Stack is empty")
            }

            this.size--
            item := this.items[this.size]
            this.items = this.items[:this.size]
            return item
        }

        bool isEmpty() {
            return this.size == 0
        }
}

// Template function
template<T>
T max(T a, T b) {
    if a > b {
        return a
    }
    return b
}

// Template usage
func main() {
    // Template class usage
    intStack := Stack<int>()
    intStack.push(1)
    intStack.push(2)
    intStack.push(3)

    fmt.Println(intStack.pop())  // 3
    fmt.Println(intStack.pop())  // 2

    // Template function usage
    fmt.Println(max<int>(5, 10))      // 10
    fmt.Println(max<string>("a", "b")) // "b"
}
```

## Exception Handling

GO-Minus provides C++-like exception handling support.

```go
// Exception definition
class DivisionByZeroException : Exception {
    public:
        DivisionByZeroException() : Exception("Division by zero") {}
}

// Throwing an exception
func divide(a, b float64) float64 {
    if b == 0 {
        throw new DivisionByZeroException()
    }
    return a / b
}

// Catching an exception
func main() {
    try {
        result := divide(10, 0)
        fmt.Println("Result:", result)
    } catch (DivisionByZeroException e) {
        fmt.Println("Error:", e.message)
    } catch (Exception e) {
        fmt.Println("General error:", e.message)
    } finally {
        fmt.Println("Operation completed")
    }
}
```

## Packages and Modules

GO-Minus uses Go's package and module system.

### Creating a Package

```go
// math/calculator.gom
package math

class Calculator {
    public:
        static int add(int a, int b) {
            return a + b
        }

        static int subtract(int a, int b) {
            return a - b
        }
}

// Exported function
func Multiply(a, b int) int {
    return a * b
}
```

### Using a Package

```go
// main.gom
package main

import (
    "fmt"
    "myapp/math"
)

func main() {
    // Class usage
    sum := math.Calculator.add(5, 3)
    fmt.Println("Sum:", sum)

    // Function usage
    product := math.Multiply(4, 2)
    fmt.Println("Product:", product)
}
```

## Compiling and Running

There are various commands to compile and run GO-Minus programs.

### Single File Compilation and Execution

```bash
# Compile and run
gominus run hello.gom

# Compile only
gominus build hello.gom

# Run
./hello
```

### Project Compilation

```bash
# In the project directory
gominus build

# With a specific output name
gominus build -o myapp

# Run
./myapp
```

## IDE Integration

GO-Minus provides plugins for various IDEs.

### VS Code

1. Open VS Code
2. Go to the Extensions tab
3. Search for "GO-Minus"
4. Install the GO-Minus extension

### JetBrains IDEs

1. Open your JetBrains IDE (IntelliJ IDEA, GoLand, etc.)
2. Go to the Plugins tab
3. Search for "GO-Minus" in the Marketplace
4. Install the GO-Minus plugin

### Vim/Neovim

```bash
# With Vim-Plug
Plug 'gominus/vim-gominus'

# With Packer
use 'gominus/vim-gominus'
```

## Next Steps

To continue learning the GO-Minus language:

1. Explore the [Language Reference](../reference/README.md) document
2. Read the [Standard Library](../../stdlib/README.md) documentation
3. Examine the example projects in the [Examples](../examples) directory
4. Join our [Discord](https://discord.gg/gominus) server
5. Contribute to the project via [GitHub](https://github.com/gominus/gominus)

If you encounter any issues while programming with GO-Minus, you can check the [FAQ](../faq.md) document or ask for help through our community channels.

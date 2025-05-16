# GO-Minus Trie Example

This example demonstrates the use of the Trie data structure found in GO-Minus's standard library. Trie is a tree data structure used to efficiently store and search for string keys.

## What is a Trie?

Trie (also known as a prefix tree or digital tree) is a tree data structure used to store string keys. Each node represents a character, and the path from the root node to any node forms a string. Trie is ideal for operations such as prefix-based searching and autocomplete.

## Basic Trie Usage

```go
// trie_basic.gom
package main

import (
    "container/trie"
    "fmt"
)

func main() {
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
    } else {
        fmt.Println("apple not found")
    }

    // Search for a non-existent word
    value, found = t.Search("orange")
    if found {
        fmt.Println("orange:", value)
    } else {
        fmt.Println("orange not found") // This line executes
    }

    // Prefix check
    if t.StartsWith("app") {
        fmt.Println("There are words starting with the prefix 'app'")
    } else {
        fmt.Println("No words start with the prefix 'app'")
    }

    // Number of words in the Trie
    fmt.Println("Word count:", t.Size()) // Prints 3

    // Delete a word
    t.Delete("banana")
    fmt.Println("Word count after deleting 'banana':", t.Size()) // Prints 2
}
```

## Output

```
apple: apple
orange not found
There are words starting with the prefix 'app'
Word count: 3
Word count after deleting 'banana': 2
```

## Prefix Search

One of the most powerful features of Trie is its ability to quickly find all words that start with a specific prefix. This is ideal for features such as autocomplete.

```go
// trie_prefix.gom
package main

import (
    "container/trie"
    "fmt"
)

func main() {
    // Create a new Trie
    t := trie.Trie.New<string>()

    // Add words
    t.Insert("apple", "apple")
    t.Insert("application", "application")
    t.Insert("append", "append")
    t.Insert("banana", "banana")
    t.Insert("ball", "ball")
    t.Insert("cat", "cat")

    // Find all words starting with the prefix "app"
    appWords := t.GetWordsWithPrefix("app")
    fmt.Println("Words starting with the prefix 'app':")
    for word, value := range appWords {
        fmt.Printf("  %s: %s\n", word, value)
    }

    // Find all words starting with the prefix "ba"
    baWords := t.GetWordsWithPrefix("ba")
    fmt.Println("\nWords starting with the prefix 'ba':")
    for word, value := range baWords {
        fmt.Printf("  %s: %s\n", word, value)
    }

    // Words starting with the prefix "x" (empty result)
    xWords := t.GetWordsWithPrefix("x")
    fmt.Println("\nWords starting with the prefix 'x':")
    if len(xWords) == 0 {
        fmt.Println("  No results found")
    }
}
```

## Output

```
Words starting with the prefix 'app':
  apple: apple
  application: application
  append: append

Words starting with the prefix 'ba':
  banana: banana
  ball: ball

Words starting with the prefix 'x':
  No results found
```

## Autocomplete Application

As a practical application of Trie, let's create a simple autocomplete system:

```go
// trie_autocomplete.gom
package main

import (
    "container/trie"
    "fmt"
    "strings"
)

// AutocompleteSystem, autocomplete system
class AutocompleteSystem {
    private:
        trie.Trie<string> dictionary

    public:
        // Create a new autocomplete system
        AutocompleteSystem() {
            this.dictionary = trie.Trie.New<string>()
        }

        // Add a word to the dictionary
        void AddWord(string word, string meaning) {
            this.dictionary.Insert(strings.ToLower(word), meaning)
        }

        // Find words starting with the prefix
        map[string]string GetSuggestions(string prefix) {
            return this.dictionary.GetWordsWithPrefix(strings.ToLower(prefix))
        }

        // Search for a word
        string GetMeaning(string word) {
            meaning, found := this.dictionary.Search(strings.ToLower(word))
            if found {
                return meaning
            }
            return "Not found"
        }
}

func main() {
    // Create an autocomplete system
    ac := AutocompleteSystem()

    // Add words to the dictionary
    ac.AddWord("apple", "apple")
    ac.AddWord("application", "application")
    ac.AddWord("append", "append")
    ac.AddWord("banana", "banana")
    ac.AddWord("ball", "ball")
    ac.AddWord("cat", "cat")
    ac.AddWord("computer", "computer")
    ac.AddWord("calculate", "calculate")

    // User input simulation
    userInputs := []string{"a", "ap", "app", "c", "co", "x"}

    for _, input := range userInputs {
        fmt.Printf("User input: '%s'\n", input)
        suggestions := ac.GetSuggestions(input)

        if len(suggestions) == 0 {
            fmt.Println("  No suggestions found")
        } else {
            fmt.Println("  Suggestions:")
            for word, meaning := range suggestions {
                fmt.Printf("    %s: %s\n", word, meaning)
            }
        }
        fmt.Println()
    }
}
```

## Output

```
User input: 'a'
  Suggestions:
    apple: apple
    application: application
    append: append

User input: 'ap'
  Suggestions:
    apple: apple
    application: application
    append: append

User input: 'app'
  Suggestions:
    apple: apple
    application: application
    append: append

User input: 'c'
  Suggestions:
    cat: cat
    computer: computer
    calculate: calculate

User input: 'co'
  Suggestions:
    computer: computer

User input: 'x'
  No suggestions found
```

## Using Trie with Different Data Types

Since Trie is a generic data structure, it can be used for values other than strings:

```go
// trie_generic.gom
package main

import (
    "container/trie"
    "fmt"
)

func main() {
    // Create a Trie for int values
    intTrie := trie.Trie.New<int>()

    // Add words and values
    intTrie.Insert("one", 1)
    intTrie.Insert("two", 2)
    intTrie.Insert("three", 3)
    intTrie.Insert("ten", 10)
    intTrie.Insert("twenty", 20)

    // Search for values
    value, found := intTrie.Search("two")
    if found {
        fmt.Println("two:", value) // Prints 2
    }

    // Create a Trie for struct values
    type Person struct {
        Name string
        Age  int
    }

    personTrie := trie.Trie.New<Person>()

    // Add people
    personTrie.Insert("john", Person{Name: "John Smith", Age: 30})
    personTrie.Insert("jane", Person{Name: "Jane Doe", Age: 25})

    // Search for people
    person, found := personTrie.Search("jane")
    if found {
        fmt.Printf("jane: %+v\n", person) // Prints {Name:Jane Doe Age:25}
    }
}
```

## Output

```
two: 2
jane: {Name:Jane Doe Age:25}
```

These examples demonstrate how to use GO-Minus's Trie data structure. For more information, you can refer to the [Trie Documentation](../../stdlib/container/trie/README.md).

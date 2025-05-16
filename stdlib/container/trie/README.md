# GO-Minus Trie Package

This package provides a Trie (Prefix Tree) data structure implementation for the GO-Minus programming language. Trie is a tree data structure used to efficiently store and search for string keys.

## Features

- Generic type support (can be used for any type)
- Word insertion, search, and deletion operations
- Prefix search
- Finding all words starting with a specific prefix
- Listing all words in the Trie
- Empty check and size calculation

## Usage

```go
import "container/trie"

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
        fmt.Println("Value:", value) // Prints "apple"
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

    // Delete a word
    t.Delete("banana")

    // Get all words in the Trie
    allWords := t.GetAllWords()
    for word, value := range allWords {
        fmt.Printf("%s: %s\n", word, value)
    }

    // Clear the Trie
    t.Clear()
}
```

## Performance

The Trie data structure has the following complexities for string keys:

- Insertion: O(m), m = key length
- Search: O(m), m = key length
- Deletion: O(m), m = key length
- Prefix search: O(m), m = prefix length
- Finding all words starting with a specific prefix: O(n), n = number of matching words

## Implementation Details

The Trie implementation consists of two main classes:

1. `TrieNode<T>`: Represents each node in the Trie tree. Each node contains child nodes, information about whether it is the end of a word, and a value.

2. `Trie<T>`: Represents the Trie data structure. Contains the root node and methods for operating on the trie.

## Use Cases

The Trie data structure is useful in the following scenarios:

- Autocomplete
- Spell checking
- Dictionary implementation
- Prefix-based search
- IP routing tables
- Text compression algorithms

## Limitations

- The Trie data structure may take up more space in terms of memory usage compared to other data structures.
- Trie is optimized for string keys. Hash tables or binary search trees may be more suitable for other types of keys.

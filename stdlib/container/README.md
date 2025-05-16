# GO-Minus Container Package

This directory contains container data structures for the GO-Minus programming language. These data structures are used to store, organize, and process data.

## Packages

### list
Doubly linked list implementation. Used to store elements in a sequential manner and navigate in both directions.

### vector
Dynamic array implementation. Used to store elements in a sequential manner and provide fast access by index.

### deque
Double-ended queue implementation. Supports adding and removing elements from both the front and back.

### heap
Priority queue implementation. Used to store elements according to priority order and provide fast access to the highest/lowest priority element.

### trie
Prefix tree implementation. Used to efficiently store and search for string keys. Ideal for operations such as prefix-based searching and autocomplete.

## Usage

To use GO-Minus container packages, simply import the relevant package in your GO-Minus program:

```go
import "container/list"
import "container/vector"
import "container/deque"
import "container/heap"
import "container/trie"

func main() {
    // list package usage
    l := list.List.New<int>()
    l.PushBack(10)
    l.PushBack(20)

    // vector package usage
    v := vector.Vector.New<string>(10)
    v.Push("hello")
    v.Push("world")

    // deque package usage
    d := deque.Deque.New<float>(10)
    d.PushBack(3.14)
    d.PushFront(2.71)

    // heap package usage
    h := heap.MinHeap.New<int>()
    h.Push(5)
    h.Push(3)
    h.Push(7)

    // trie package usage
    t := trie.Trie.New<string>()
    t.Insert("apple", "apple")
    t.Insert("banana", "banana")
}
```

## Performance

Each data structure has different performance characteristics for different operations. To choose the most suitable data structure for your application's needs, refer to the documentation of each data structure.

## Documentation

For more information and example usage for each package, refer to the README.md file of the relevant package:

- [list/README.md](list/README.md)
- [vector/README.md](vector/README.md)
- [deque/README.md](deque/README.md)
- [heap/README.md](heap/README.md)
- [trie/README.md](trie/README.md)

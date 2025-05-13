# GO-Minus Trie (Önek Ağacı) Örneği

Bu örnek, GO-Minus'un standart kütüphanesinde bulunan Trie (Önek Ağacı) veri yapısının kullanımını göstermektedir. Trie, string anahtarları verimli bir şekilde depolamak ve aramak için kullanılan bir ağaç veri yapısıdır.

## Trie Nedir?

Trie (önek ağacı veya dijital ağaç olarak da bilinir), string anahtarları depolamak için kullanılan bir ağaç veri yapısıdır. Her düğüm, bir karakteri temsil eder ve kök düğümden herhangi bir düğüme giden yol, bir string oluşturur. Trie, özellikle önek tabanlı arama ve otomatik tamamlama gibi işlemler için idealdir.

## Temel Trie Kullanımı

```go
// trie_basic.gom
package main

import (
    "container/trie"
    "fmt"
)

func main() {
    // String değerler için bir Trie oluştur
    t := trie.Trie.New<string>()
    
    // Kelimeler ekle
    t.Insert("apple", "elma")
    t.Insert("banana", "muz")
    t.Insert("application", "uygulama")
    
    // Kelime ara
    value, found := t.Search("apple")
    if found {
        fmt.Println("apple:", value) // "elma" yazdırır
    } else {
        fmt.Println("apple bulunamadı")
    }
    
    // Var olmayan bir kelime ara
    value, found = t.Search("orange")
    if found {
        fmt.Println("orange:", value)
    } else {
        fmt.Println("orange bulunamadı") // Bu satır çalışır
    }
    
    // Önek kontrolü
    if t.StartsWith("app") {
        fmt.Println("'app' öneki ile başlayan kelimeler var")
    } else {
        fmt.Println("'app' öneki ile başlayan kelime yok")
    }
    
    // Trie'deki kelime sayısı
    fmt.Println("Kelime sayısı:", t.Size()) // 3 yazdırır
    
    // Kelime sil
    t.Delete("banana")
    fmt.Println("'banana' silindikten sonra kelime sayısı:", t.Size()) // 2 yazdırır
}
```

## Çıktı

```
apple: elma
orange bulunamadı
'app' öneki ile başlayan kelimeler var
Kelime sayısı: 3
'banana' silindikten sonra kelime sayısı: 2
```

## Önek Araması

Trie'nin en güçlü özelliklerinden biri, belirli bir önekle başlayan tüm kelimeleri hızlı bir şekilde bulabilmesidir. Bu, otomatik tamamlama gibi özellikler için idealdir.

```go
// trie_prefix.gom
package main

import (
    "container/trie"
    "fmt"
)

func main() {
    // Yeni bir Trie oluştur
    t := trie.Trie.New<string>()
    
    // Kelimeler ekle
    t.Insert("apple", "elma")
    t.Insert("application", "uygulama")
    t.Insert("append", "ekle")
    t.Insert("banana", "muz")
    t.Insert("ball", "top")
    t.Insert("cat", "kedi")
    
    // "app" öneki ile başlayan tüm kelimeleri bul
    appWords := t.GetWordsWithPrefix("app")
    fmt.Println("'app' öneki ile başlayan kelimeler:")
    for word, value := range appWords {
        fmt.Printf("  %s: %s\n", word, value)
    }
    
    // "ba" öneki ile başlayan tüm kelimeleri bul
    baWords := t.GetWordsWithPrefix("ba")
    fmt.Println("\n'ba' öneki ile başlayan kelimeler:")
    for word, value := range baWords {
        fmt.Printf("  %s: %s\n", word, value)
    }
    
    // "x" öneki ile başlayan kelimeler (boş sonuç)
    xWords := t.GetWordsWithPrefix("x")
    fmt.Println("\n'x' öneki ile başlayan kelimeler:")
    if len(xWords) == 0 {
        fmt.Println("  Sonuç bulunamadı")
    }
}
```

## Çıktı

```
'app' öneki ile başlayan kelimeler:
  apple: elma
  application: uygulama
  append: ekle

'ba' öneki ile başlayan kelimeler:
  banana: muz
  ball: top

'x' öneki ile başlayan kelimeler:
  Sonuç bulunamadı
```

## Otomatik Tamamlama Uygulaması

Trie'nin pratik bir uygulaması olarak, basit bir otomatik tamamlama sistemi oluşturalım:

```go
// trie_autocomplete.gom
package main

import (
    "container/trie"
    "fmt"
    "strings"
)

// AutocompleteSystem, otomatik tamamlama sistemi
class AutocompleteSystem {
    private:
        trie.Trie<string> dictionary
    
    public:
        // Yeni bir otomatik tamamlama sistemi oluştur
        AutocompleteSystem() {
            this.dictionary = trie.Trie.New<string>()
        }
        
        // Sözlüğe kelime ekle
        void AddWord(string word, string meaning) {
            this.dictionary.Insert(strings.ToLower(word), meaning)
        }
        
        // Önek ile başlayan kelimeleri bul
        map[string]string GetSuggestions(string prefix) {
            return this.dictionary.GetWordsWithPrefix(strings.ToLower(prefix))
        }
        
        // Kelime ara
        string GetMeaning(string word) {
            meaning, found := this.dictionary.Search(strings.ToLower(word))
            if found {
                return meaning
            }
            return "Bulunamadı"
        }
}

func main() {
    // Otomatik tamamlama sistemi oluştur
    ac := AutocompleteSystem()
    
    // Sözlüğe kelimeler ekle
    ac.AddWord("apple", "elma")
    ac.AddWord("application", "uygulama")
    ac.AddWord("append", "ekle")
    ac.AddWord("banana", "muz")
    ac.AddWord("ball", "top")
    ac.AddWord("cat", "kedi")
    ac.AddWord("computer", "bilgisayar")
    ac.AddWord("calculate", "hesapla")
    
    // Kullanıcı girişi simülasyonu
    userInputs := []string{"a", "ap", "app", "c", "co", "x"}
    
    for _, input := range userInputs {
        fmt.Printf("Kullanıcı girişi: '%s'\n", input)
        suggestions := ac.GetSuggestions(input)
        
        if len(suggestions) == 0 {
            fmt.Println("  Öneri bulunamadı")
        } else {
            fmt.Println("  Öneriler:")
            for word, meaning := range suggestions {
                fmt.Printf("    %s: %s\n", word, meaning)
            }
        }
        fmt.Println()
    }
}
```

## Çıktı

```
Kullanıcı girişi: 'a'
  Öneriler:
    apple: elma
    application: uygulama
    append: ekle

Kullanıcı girişi: 'ap'
  Öneriler:
    apple: elma
    application: uygulama
    append: ekle

Kullanıcı girişi: 'app'
  Öneriler:
    apple: elma
    application: uygulama
    append: ekle

Kullanıcı girişi: 'c'
  Öneriler:
    cat: kedi
    computer: bilgisayar
    calculate: hesapla

Kullanıcı girişi: 'co'
  Öneriler:
    computer: bilgisayar

Kullanıcı girişi: 'x'
  Öneri bulunamadı
```

## Farklı Veri Tipleri ile Trie Kullanımı

Trie, jenerik bir veri yapısı olduğu için, string dışındaki değerler için de kullanılabilir:

```go
// trie_generic.gom
package main

import (
    "container/trie"
    "fmt"
)

func main() {
    // Int değerler için bir Trie oluştur
    intTrie := trie.Trie.New<int>()
    
    // Kelimeler ve değerler ekle
    intTrie.Insert("one", 1)
    intTrie.Insert("two", 2)
    intTrie.Insert("three", 3)
    intTrie.Insert("ten", 10)
    intTrie.Insert("twenty", 20)
    
    // Değerleri ara
    value, found := intTrie.Search("two")
    if found {
        fmt.Println("two:", value) // 2 yazdırır
    }
    
    // Struct değerler için bir Trie oluştur
    type Person struct {
        Name string
        Age  int
    }
    
    personTrie := trie.Trie.New<Person>()
    
    // Kişileri ekle
    personTrie.Insert("john", Person{Name: "John Smith", Age: 30})
    personTrie.Insert("jane", Person{Name: "Jane Doe", Age: 25})
    
    // Kişileri ara
    person, found := personTrie.Search("jane")
    if found {
        fmt.Printf("jane: %+v\n", person) // {Name:Jane Doe Age:25} yazdırır
    }
}
```

## Çıktı

```
two: 2
jane: {Name:Jane Doe Age:25}
```

Bu örnekler, GO-Minus'un Trie veri yapısının nasıl kullanılacağını göstermektedir. Daha fazla bilgi için [Trie Belgelendirmesi](../../stdlib/container/trie/README.md) belgesine bakabilirsiniz.

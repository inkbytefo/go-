# GO-Minus Trie Paketi

Bu paket, GO-Minus programlama dili için Trie (Önek Ağacı) veri yapısı implementasyonu sağlar. Trie, string anahtarları verimli bir şekilde depolamak ve aramak için kullanılan bir ağaç veri yapısıdır.

## Özellikler

- Jenerik tip desteği (herhangi bir tip için kullanılabilir)
- Kelime ekleme, arama ve silme işlemleri
- Önek araması
- Belirli bir önekle başlayan tüm kelimeleri bulma
- Trie'deki tüm kelimeleri listeleme
- Boş kontrol ve boyut hesaplama

## Kullanım

```go
import "container/trie"

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
        fmt.Println("Değer:", value) // "elma" yazdırır
    }
    
    // Önek kontrolü
    if t.StartsWith("app") {
        fmt.Println("'app' öneki ile başlayan kelimeler var")
    }
    
    // Belirli önekle başlayan tüm kelimeleri al
    appWords := t.GetWordsWithPrefix("app")
    for word, value := range appWords {
        fmt.Printf("%s: %s\n", word, value)
    }
    
    // Kelime sil
    t.Delete("banana")
    
    // Trie'deki tüm kelimeleri al
    allWords := t.GetAllWords()
    for word, value := range allWords {
        fmt.Printf("%s: %s\n", word, value)
    }
    
    // Trie'yi temizle
    t.Clear()
}
```

## Performans

Trie veri yapısı, string anahtarları için aşağıdaki karmaşıklıklara sahiptir:

- Ekleme: O(m), m = anahtar uzunluğu
- Arama: O(m), m = anahtar uzunluğu
- Silme: O(m), m = anahtar uzunluğu
- Önek araması: O(m), m = önek uzunluğu
- Belirli önekle başlayan tüm kelimeleri bulma: O(n), n = eşleşen kelime sayısı

## Uygulama Detayları

Trie implementasyonu, iki ana sınıftan oluşur:

1. `TrieNode<T>`: Trie ağacındaki her bir düğümü temsil eder. Her düğüm, çocuk düğümleri, bir kelimenin sonu olup olmadığı bilgisini ve bir değer içerir.

2. `Trie<T>`: Trie veri yapısını temsil eder. Kök düğümü ve trie üzerinde işlem yapmak için metotları içerir.

## Kullanım Senaryoları

Trie veri yapısı aşağıdaki senaryolarda kullanışlıdır:

- Otomatik tamamlama
- Yazım denetimi
- Sözlük implementasyonu
- Önek tabanlı arama
- IP yönlendirme tabloları
- Metin sıkıştırma algoritmaları

## Sınırlamalar

- Trie veri yapısı, bellek kullanımı açısından diğer veri yapılarına göre daha fazla yer kaplayabilir.
- Trie, string anahtarlar için optimize edilmiştir. Diğer tip anahtarlar için hash tabloları veya ikili arama ağaçları daha uygun olabilir.

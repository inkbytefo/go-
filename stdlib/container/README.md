# GO-Minus Container Paketi

Bu dizin, GO-Minus programlama dili için container (kapsayıcı) veri yapılarını içerir. Bu veri yapıları, verileri depolamak, düzenlemek ve işlemek için kullanılır.

## Paketler

### list
Çift bağlı liste implementasyonu. Elemanları sıralı bir şekilde depolamak ve her iki yönde de gezinmek için kullanılır.

### vector
Dinamik dizi implementasyonu. Elemanları sıralı bir şekilde depolamak ve indeks ile hızlı erişim sağlamak için kullanılır.

### deque
Çift uçlu kuyruk implementasyonu. Hem baştan hem de sondan eleman ekleme ve çıkarma işlemlerini destekler.

### heap
Öncelik kuyruğu implementasyonu. Elemanları öncelik sırasına göre depolamak ve en yüksek/düşük öncelikli elemana hızlı erişim sağlamak için kullanılır.

### trie
Önek ağacı implementasyonu. String anahtarları verimli bir şekilde depolamak ve aramak için kullanılır. Önek tabanlı arama ve otomatik tamamlama gibi işlemler için idealdir.

## Kullanım

GO-Minus container paketlerini kullanmak için, GO-Minus programınızda ilgili paketi import etmeniz yeterlidir:

```go
import "container/list"
import "container/vector"
import "container/deque"
import "container/heap"
import "container/trie"

func main() {
    // list paketi kullanımı
    l := list.List.New<int>()
    l.PushBack(10)
    l.PushBack(20)
    
    // vector paketi kullanımı
    v := vector.Vector.New<string>(10)
    v.Push("hello")
    v.Push("world")
    
    // deque paketi kullanımı
    d := deque.Deque.New<float>(10)
    d.PushBack(3.14)
    d.PushFront(2.71)
    
    // heap paketi kullanımı
    h := heap.MinHeap.New<int>()
    h.Push(5)
    h.Push(3)
    h.Push(7)
    
    // trie paketi kullanımı
    t := trie.Trie.New<string>()
    t.Insert("apple", "elma")
    t.Insert("banana", "muz")
}
```

## Performans

Her veri yapısı, farklı operasyonlar için farklı performans karakteristiklerine sahiptir. Uygulamanızın ihtiyaçlarına en uygun veri yapısını seçmek için, her bir veri yapısının belgelendirmesine bakın.

## Belgelendirme

Her bir paket için daha fazla bilgi ve örnek kullanım için, ilgili paketin README.md dosyasına bakın:

- [list/README.md](list/README.md)
- [vector/README.md](vector/README.md)
- [deque/README.md](deque/README.md)
- [heap/README.md](heap/README.md)
- [trie/README.md](trie/README.md)

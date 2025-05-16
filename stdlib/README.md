# GO-Minus Standart Kütüphanesi

Bu dizin, GO-Minus programlama dili için standart kütüphaneyi içerir. GO-Minus standart kütüphanesi, Go'nun standart kütüphanesini temel alır ve GO-Minus'ın sınıf, şablon ve istisna işleme gibi özelliklerini kullanarak genişletilmiştir.

## Paketler

### core
Temel veri tipleri ve fonksiyonlar için çekirdek paket.

### container
Veri yapıları için paketler:
- **list**: Çift bağlı liste implementasyonu
- **vector**: Dinamik dizi implementasyonu
- **map**: Anahtar-değer eşleştirme implementasyonu
- **set**: Küme implementasyonu

### concurrent
Eşzamanlılık için paketler:
- **channel**: Goroutine'ler arası iletişim için kanallar
- **mutex**: Karşılıklı dışlama için mutex'ler
- **waitgroup**: Goroutine'lerin tamamlanmasını beklemek için WaitGroup

### fmt
Biçimlendirilmiş giriş/çıkış işlemleri için paket.

### io
Giriş/çıkış işlemleri için temel arayüzler ve fonksiyonlar.

### math
Matematiksel fonksiyonlar ve sabitler.

### strings
Dize işleme fonksiyonları.

## Kullanım

GO-Minus standart kütüphanesindeki paketleri kullanmak için, GO-Minus programınızda ilgili paketi import etmeniz yeterlidir:

```go
import "fmt"
import "container/list"
import "concurrent/channel"

func main() {
    // fmt paketini kullan
    fmt.Println("Merhaba, Dünya!")

    // list paketini kullan
    l := list.New<int>()
    l.PushBack(10)
    l.PushBack(20)

    // channel paketini kullan
    ch := channel.New<string>(1)
    ch.Send("Merhaba")
    msg := ch.Receive()
    fmt.Println(msg)
}
```

## Geliştirme

GO-Minus standart kütüphanesi, GO-Minus dilinin gelişimiyle birlikte sürekli olarak genişletilmektedir. Yeni paketler ve fonksiyonlar eklenmeye devam edilecektir.
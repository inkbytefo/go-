# GO-Minus Async Paketi

Bu paket, GO-Minus programlama dili için asenkron giriş/çıkış (async I/O) işlemlerini sağlar. Asenkron I/O, uygulamaların I/O işlemleri için beklemeden diğer işlemleri gerçekleştirmesine olanak tanır, bu da performansı ve ölçeklenebilirliği artırır.

## Özellikler

- Platform bağımsız asenkron I/O API'si
- Yüksek performanslı olay döngüsü (event loop) mekanizması
- Promise/Future pattern entegrasyonu
- Farklı işletim sistemleri için I/O multiplexing desteği (epoll, kqueue, IOCP)
- Asenkron dosya, ağ ve diğer I/O işlemleri için tutarlı bir API
- Hata yönetimi ve iptal mekanizmaları

## Kullanım

### Olay Döngüsü Oluşturma

```go
import (
    "async"
    "fmt"
)

func main() {
    // Olay döngüsü oluştur
    loop, err := async.EventLoop.New()
    if err != nil {
        fmt.Println("Olay döngüsü oluşturma hatası:", err)
        return
    }
    
    // Olay döngüsünü başlat
    err = loop.Start()
    if err != nil {
        fmt.Println("Olay döngüsü başlatma hatası:", err)
        return
    }
    
    // Uygulama işlemleri...
    
    // Olay döngüsünü durdur
    err = loop.Stop()
    if err != nil {
        fmt.Println("Olay döngüsü durdurma hatası:", err)
        return
    }
}
```

### Asenkron Dosya İşlemleri

```go
import (
    "async"
    "fmt"
)

func main() {
    // Olay döngüsü oluştur
    loop, _ := async.EventLoop.New()
    loop.Start()
    
    // Asenkron dosya aç
    file, err := async.OpenFile("example.txt", async.O_RDONLY, 0)
    if err != nil {
        fmt.Println("Dosya açma hatası:", err)
        return
    }
    
    // Asenkron okuma
    buffer := make([]byte, 1024)
    future := file.Read(buffer)
    
    // Sonucu bekle
    n, err := future.Get()
    if err != nil {
        fmt.Println("Okuma hatası:", err)
        return
    }
    
    fmt.Printf("Okunan veri: %s\n", string(buffer[:n]))
    
    // Dosyayı kapat
    closeErr := file.Close().Get()
    if closeErr != nil {
        fmt.Println("Dosya kapatma hatası:", closeErr)
    }
    
    // Olay döngüsünü durdur
    loop.Stop()
}
```

### Asenkron Ağ İşlemleri

```go
import (
    "async"
    "fmt"
)

func main() {
    // Olay döngüsü oluştur
    loop, _ := async.EventLoop.New()
    loop.Start()
    
    // Asenkron TCP bağlantısı kur
    socket, err := async.Dial("tcp", "example.com:80")
    if err != nil {
        fmt.Println("Bağlantı hatası:", err)
        return
    }
    
    // Asenkron yazma
    request := "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"
    writeFuture := socket.Write([]byte(request))
    
    // Yazma sonucunu bekle
    n, err := writeFuture.Get()
    if err != nil {
        fmt.Println("Yazma hatası:", err)
        return
    }
    
    fmt.Printf("Yazılan bayt sayısı: %d\n", n)
    
    // Asenkron okuma
    buffer := make([]byte, 1024)
    readFuture := socket.Read(buffer)
    
    // Okuma sonucunu bekle
    n, err = readFuture.Get()
    if err != nil {
        fmt.Println("Okuma hatası:", err)
        return
    }
    
    fmt.Printf("Yanıt: %s\n", string(buffer[:n]))
    
    // Bağlantıyı kapat
    socket.Close().Get()
    
    // Olay döngüsünü durdur
    loop.Stop()
}
```

## Performans İpuçları

1. **Tek Olay Döngüsü**: Uygulama başına tek bir olay döngüsü kullanın.
2. **Callback Kullanımı**: Uzun süren işlemler için Get() yerine callback'leri tercih edin.
3. **Tamponlama**: Küçük okuma/yazma işlemleri yerine daha büyük tamponlar kullanın.
4. **Havuzlama**: Tampon ve diğer nesneler için nesne havuzları kullanın.
5. **Zaman Aşımı**: Uzun süren işlemler için zaman aşımı ayarlayın.

## Sınırlamalar

1. **Bloke Edici İşlemler**: Olay döngüsü içinde bloke edici işlemler kullanmayın.
2. **Çoklu İş Parçacığı**: Olay döngüsü tek iş parçacığı üzerinde çalışır, çoklu iş parçacığı desteği için ek önlemler gereklidir.
3. **Bellek Kullanımı**: Büyük tamponlar bellek kullanımını artırabilir.
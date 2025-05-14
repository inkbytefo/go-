# GO-Minus Asenkron IO Implementasyonu

Bu belge, GO-Minus programlama dilinin asenkron IO implementasyonunu açıklamaktadır.

## Genel Bakış

GO-Minus'un asenkron IO implementasyonu, yüksek performanslı, ölçeklenebilir ve platform bağımsız bir asenkron IO altyapısı sağlar. Bu altyapı, aşağıdaki bileşenlerden oluşur:

1. **Asenkron IO Arayüzleri**: Asenkron okuma, yazma ve diğer IO işlemleri için arayüzler
2. **Event Loop**: Asenkron olayları işleyen bir olay döngüsü
3. **Platform Bağımlı IO Multiplexing**: Linux (epoll), macOS/BSD (kqueue) ve Windows (IOCP) için implementasyonlar
4. **Future/Promise Pattern**: Asenkron işlemlerin sonuçlarını temsil eden ve işleyen bir pattern
5. **Asenkron Dosya ve Soket İşlemleri**: Asenkron dosya ve ağ işlemleri için implementasyonlar
6. **Performans Optimizasyonları**: CPU kullanımı, lock-free veri yapıları, sistem çağrıları ve buffer havuzu optimizasyonları

## Asenkron IO Arayüzleri

GO-Minus'un asenkron IO arayüzleri, asenkron okuma, yazma ve diğer IO işlemleri için bir dizi arayüz sağlar:

```go
// AsyncCloser, asenkron kapatma işlemleri için bir arayüzdür.
interface AsyncCloser {
    func Close() Future<e>
}

// AsyncReader, asenkron okuma işlemleri için bir arayüzdür.
interface AsyncReader {
    func Read(p []byte) Future<int>
}

// AsyncWriter, asenkron yazma işlemleri için bir arayüzdür.
interface AsyncWriter {
    func Write(p []byte) Future<int>
}

// AsyncReadWriter, hem asenkron okuma hem de yazma işlemleri için bir arayüzdür.
interface AsyncReadWriter {
    AsyncReader
    AsyncWriter
}

// AsyncReadCloser, asenkron okuma ve kapatma işlemleri için bir arayüzdür.
interface AsyncReadCloser {
    AsyncReader
    AsyncCloser
}

// AsyncWriteCloser, asenkron yazma ve kapatma işlemleri için bir arayüzdür.
interface AsyncWriteCloser {
    AsyncWriter
    AsyncCloser
}

// AsyncReadWriteCloser, asenkron okuma, yazma ve kapatma işlemleri için bir arayüzdür.
interface AsyncReadWriteCloser {
    AsyncReader
    AsyncWriter
    AsyncCloser
}

// AsyncSeeker, asenkron konumlandırma işlemleri için bir arayüzdür.
interface AsyncSeeker {
    func Seek(offset int64, whence int) Future<int64>
}

// AsyncReadWriteSeeker, asenkron okuma, yazma ve konumlandırma işlemleri için bir arayüzdür.
interface AsyncReadWriteSeeker {
    AsyncReader
    AsyncWriter
    AsyncSeeker
}
```

## Event Loop

Event Loop, asenkron olayları işleyen bir olay döngüsüdür. GO-Minus'un event loop implementasyonu, aşağıdaki özelliklere sahiptir:

- Olay tabanlı asenkron IO işlemleri
- Görev kuyruğu yönetimi
- Olay işleme mekanizması
- Platform bağımsız poller entegrasyonu

```go
// EventLoop, asenkron işlemleri yöneten bir olay döngüsüdür.
class EventLoop {
    // New, yeni bir EventLoop oluşturur.
    static func New() (*EventLoop, error)
    
    // Start, olay döngüsünü başlatır.
    func (loop *EventLoop) Start() error
    
    // Stop, olay döngüsünü durdurur.
    func (loop *EventLoop) Stop() error
    
    // Register, bir dosya tanımlayıcısını olay döngüsüne kaydeder.
    func (loop *EventLoop) Register(fd int, events int, handler AsyncHandler) (int, error)
    
    // Unregister, bir dosya tanımlayıcısını olay döngüsünden kaldırır.
    func (loop *EventLoop) Unregister(id int) error
    
    // Wakeup, olay döngüsünü uyandırır.
    func (loop *EventLoop) Wakeup()
    
    // Post, bir olayı olay döngüsüne gönderir.
    func (loop *EventLoop) Post(event Event, handler AsyncHandler)
}
```

## Platform Bağımlı IO Multiplexing

GO-Minus, farklı platformlar için IO multiplexing implementasyonları sağlar:

### Linux için epoll

```go
// EpollPoller, Linux epoll API'sini kullanarak bir poller implementasyonu sağlar.
class EpollPoller {
    // Add, bir dosya tanımlayıcısını epoll örneğine ekler.
    func (ep *EpollPoller) Add(fd int, events int) error
    
    // Remove, bir dosya tanımlayıcısını epoll örneğinden kaldırır.
    func (ep *EpollPoller) Remove(fd int) error
    
    // Modify, bir dosya tanımlayıcısının epoll olaylarını değiştirir.
    func (ep *EpollPoller) Modify(fd int, events int) error
    
    // Wait, epoll olaylarını bekler.
    func (ep *EpollPoller) Wait(timeout int) ([]Event, error)
    
    // Close, epoll örneğini kapatır.
    func (ep *EpollPoller) Close() error
}
```

### macOS/BSD için kqueue

```go
// KqueuePoller, macOS/BSD kqueue API'sini kullanarak bir poller implementasyonu sağlar.
class KqueuePoller {
    // Add, bir dosya tanımlayıcısını kqueue örneğine ekler.
    func (kp *KqueuePoller) Add(fd int, events int) error
    
    // Remove, bir dosya tanımlayıcısını kqueue örneğinden kaldırır.
    func (kp *KqueuePoller) Remove(fd int) error
    
    // Modify, bir dosya tanımlayıcısının kqueue olaylarını değiştirir.
    func (kp *KqueuePoller) Modify(fd int, events int) error
    
    // Wait, kqueue olaylarını bekler.
    func (kp *KqueuePoller) Wait(timeout int) ([]Event, error)
    
    // Close, kqueue örneğini kapatır.
    func (kp *KqueuePoller) Close() error
}
```

### Windows için IOCP

```go
// IOCPPoller, Windows IOCP API'sini kullanarak bir poller implementasyonu sağlar.
class IOCPPoller {
    // Add, bir handle'ı IOCP'ye ekler.
    func (ip *IOCPPoller) Add(fd int, events int) error
    
    // Remove, bir handle'ı IOCP'den kaldırır.
    func (ip *IOCPPoller) Remove(fd int) error
    
    // Modify, bir handle'ın IOCP olaylarını değiştirir.
    func (ip *IOCPPoller) Modify(fd int, events int) error
    
    // Wait, IOCP olaylarını bekler.
    func (ip *IOCPPoller) Wait(timeout int) ([]Event, error)
    
    // Close, IOCP örneğini kapatır.
    func (ip *IOCPPoller) Close() error
}
```

## Future/Promise Pattern

GO-Minus'un Future/Promise pattern implementasyonu, asenkron işlemlerin sonuçlarını temsil eden ve işleyen bir pattern sağlar:

```go
// AsyncFuture, asenkron bir işlemin sonucunu temsil eder.
class AsyncFuture<T> {
    // Get, Future'ın sonucunu döndürür.
    func (af *AsyncFuture<T>) Get() (T, error)
    
    // GetWithTimeout, belirtilen süre içinde Future'ın sonucunu döndürür.
    func (af *AsyncFuture<T>) GetWithTimeout(timeout time.Duration) (T, error, bool)
    
    // IsDone, Future'ın tamamlanıp tamamlanmadığını kontrol eder.
    func (af *AsyncFuture<T>) IsDone() bool
    
    // IsCancelled, Future'ın iptal edilip edilmediğini kontrol eder.
    func (af *AsyncFuture<T>) IsCancelled() bool
    
    // Cancel, Future'ı iptal eder.
    func (af *AsyncFuture<T>) Cancel() bool
    
    // Then, Future tamamlandığında çağrılacak bir callback ekler.
    func (af *AsyncFuture<T>) Then(callback func(T)) *AsyncFuture<T>
    
    // Catch, Future bir hata ile tamamlandığında çağrılacak bir callback ekler.
    func (af *AsyncFuture<T>) Catch(callback func(error)) *AsyncFuture<T>
    
    // Finally, Future tamamlandığında veya iptal edildiğinde çağrılacak bir callback ekler.
    func (af *AsyncFuture<T>) Finally(callback func()) *AsyncFuture<T>
    
    // Map, Future'ın sonucunu dönüştürür.
    template<U> func (af *AsyncFuture<T>) Map(mapper func(T) U) *AsyncFuture<U>
    
    // FlatMap, Future'ın sonucunu başka bir Future'a dönüştürür.
    template<U> func (af *AsyncFuture<T>) FlatMap(mapper func(T) *AsyncFuture<U>) *AsyncFuture<U>
}

// AsyncPromise, bir AsyncFuture'ın sonucunu ayarlamak için kullanılır.
class AsyncPromise<T> {
    // Complete, Promise'i tamamlar ve sonucu ayarlar.
    func (p *AsyncPromise<T>) Complete(result T)
    
    // CompleteWithError, Promise'i bir hata ile tamamlar.
    func (p *AsyncPromise<T>) CompleteWithError(err interface{})
    
    // GetFuture, Promise ile ilişkili Future'ı döndürür.
    func (p *AsyncPromise<T>) GetFuture() *AsyncFuture<T>
}
```

## Performans Optimizasyonları

GO-Minus'un asenkron IO implementasyonu, aşağıdaki performans optimizasyonlarını içerir:

### CPU Kullanımı Optimizasyonu

- Optimize edilmiş event loop
- İş parçacığı havuzu optimizasyonu
- CPU çekirdeklerine bağlama (CPU affinity)

### Lock-Free Veri Yapıları

- Lock-free kuyruk
- İş çalma algoritması (work stealing)
- Atomik sayaçlar

### Sistem Çağrıları Optimizasyonu (Devam Ediyor)

- Sistem çağrı sayısını azaltma
- Batch işleme
- Syscall overhead azaltma

### Buffer Havuzu (Devam Ediyor)

- Önceden ayrılmış buffer havuzu
- Buffer yeniden kullanımı
- Buffer boyutu optimizasyonu

## Örnek Kullanım

```go
import (
    "async"
    "fmt"
)

func main() {
    // Event loop oluştur
    loop, err := async.EventLoop.New()
    if err != nil {
        fmt.Println("Event loop oluşturma hatası:", err)
        return
    }
    
    // Event loop başlat
    err = loop.Start()
    if err != nil {
        fmt.Println("Event loop başlatma hatası:", err)
        return
    }
    defer loop.Stop()
    
    // Asenkron bir işlem başlat
    future := async.SupplyAsync<string>(loop, func() string {
        // Asenkron işlem simülasyonu
        time.Sleep(1 * time.Second)
        return "Merhaba, Dünya!"
    })
    
    // Callback ekle
    future.Then(func(result string) {
        fmt.Println("Sonuç:", result)
    }).Catch(func(err error) {
        fmt.Println("Hata:", err)
    })
    
    // Sonucu bekle
    result, err := future.Get()
    if err != nil {
        fmt.Println("Hata:", err)
        return
    }
    
    fmt.Println("Sonuç:", result)
}
```

## Sonuç

GO-Minus'un asenkron IO implementasyonu, yüksek performanslı, ölçeklenebilir ve platform bağımsız bir asenkron IO altyapısı sağlar. Bu altyapı, modern uygulamaların ihtiyaç duyduğu asenkron IO işlemlerini gerçekleştirmek için gerekli tüm bileşenleri içerir.

Asenkron IO implementasyonu, GO-Minus'un standart kütüphanesinin bir parçası olarak sunulur ve `async` paketi altında bulunur.

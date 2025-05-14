# GO-Minus Bellek Yönetimi Paketi

Bu paket, GO-Minus programlama dili için Hibrit Akıllı Bellek Yönetimi Sistemi'ni sağlar. Bu sistem, programcılara maksimum esneklik ve performans sağlayarak, hem yüksek performanslı sistem programlama hem de hızlı uygulama geliştirme için ideal bir ortam sunar.

## Özellikler

- **Otomatik Bellek Yönetimi (Garbage Collection)**: Go'nun garbage collector'ünü kullanarak otomatik bellek yönetimi sağlar.
- **Manuel Bellek Yönetimi**: Performans kritik bölümler için manuel bellek yönetimi seçeneği sunar.
- **Bölgesel Bellek Yönetimi (Region-Based Memory Management)**: Belirli kod bloklarını "bellek bölgeleri" olarak işaretleme ve bu bölgelerdeki tüm bellek ayırma işlemlerini bölge sonunda otomatik olarak serbest bırakma olanağı sunar.
- **Yaşam Süresi Analizi (Lifetime Analysis)**: Değişkenlerin yaşam sürelerini analiz eder ve potansiyel bellek sızıntılarını veya dangling pointer'ları tespit eder.
- **Profil Tabanlı Otomatik Optimizasyon**: Uygulama çalışırken bellek kullanım desenlerini analiz eder ve gelecekteki çalıştırmalarda bellek yönetimini otomatik olarak optimize eder.
- **Bellek Havuzu Şablonları**: Belirli veri yapıları için özelleştirilmiş bellek havuzları oluşturmayı kolaylaştıran şablonlar sunar.

## Kullanım

### Otomatik Bellek Yönetimi

```go
func processData() {
    // Otomatik bellek yönetimi
    data := createLargeData()
    processData(data)
    // data otomatik olarak temizlenir
}
```

### Manuel Bellek Yönetimi

```go
func processImage(Image image) {
    unsafe {
        // Manuel bellek ayırma
        buffer := allocate<byte>(image.width * image.height * 4)
        defer free(buffer) // İşlev sonunda belleği serbest bırak
        
        // Performans kritik işlemler
        // ...
    }
}
```

### Bölgesel Bellek Yönetimi

```go
import "memory"

func processLargeData(data []byte) {
    // Bellek bölgesi tanımlama
    region := memory.NewRegion()
    defer region.Free()
    
    // Bu bölgedeki tüm bellek ayırmaları bölge ile birlikte serbest bırakılır
    buffer := region.Allocate[byte](1024 * 1024)
    
    // Performans kritik işlemler...
}
```

### Bellek Havuzu Şablonları

```go
import "memory"

// Özelleştirilmiş bellek havuzu şablonu
pool := memory.NewPool[MyStruct](1000)

// Havuzdan nesne alma
obj := pool.Get()

// İşlemler...

// Havuza geri döndürme
pool.Return(obj)
```

### Profil Tabanlı Otomatik Optimizasyon

```go
import "memory"
import "time"

func main() {
    // Profil tabanlı otomatik optimizasyon
    memory.EnableProfiling(time.Minute, "memory_profile.json")
    
    // Uygulama kodu
    // ...
    
    // Profil verilerini kaydet
    memory.SaveProfile("memory_profile.json")
}
```

## Sınıflar ve Arayüzler

### MemoryManager

Bellek yönetimi için temel arayüz.

```go
interface MemoryManager {
    func Allocate(size uint64) unsafe.Pointer
    func Free(ptr unsafe.Pointer)
    func GetStats() MemoryStats
}
```

### Region

Bölgesel bellek yönetimi için kullanılan sınıf.

```go
class Region {
    static func New() *Region
    static func NewWithBlockSize(blockSize uint64) *Region
    func Allocate(size uint64) unsafe.Pointer
    func Allocate<T>(count uint64) *T
    func Free()
    func GetStats() RegionStats
}
```

### Pool<T>

Bellek havuzu şablonu.

```go
class Pool<T> {
    static func New(capacity uint64) *Pool<T>
    func Get() *T
    func Return(ptr *T)
    func GetStats() PoolStats
    func Clear()
    func Resize(newCapacity uint64)
}
```

### AdvancedMemoryProfiler

Gelişmiş bellek profilleme için kullanılan sınıf.

```go
class AdvancedMemoryProfiler {
    static func New(saveInterval time.Duration, filePath string) *AdvancedMemoryProfiler
    func RecordAllocation(ptr unsafe.Pointer, size uint64)
    func RecordFree(ptr unsafe.Pointer)
    func GetAllocationSize(ptr unsafe.Pointer) uint64
    func SaveProfile() error
    func AnalyzeAndOptimize() []OptimizationRule
    func Stop()
}
```

## Global Fonksiyonlar

```go
func EnableProfiling(saveInterval time.Duration, filePath string)
func SaveProfile(filePath string) error
func GetStats() MemoryStats
func Allocate(size uint64) unsafe.Pointer
func Free(ptr unsafe.Pointer)
func NewRegion() *Region
func NewRegionWithBlockSize(blockSize uint64) *Region
func NewPool<T>(capacity uint64) *Pool<T>
```

## Örnekler

Daha fazla örnek için [örnekler dizinine](../../examples/memory) bakın.

## Lisans

GO-Minus Bellek Yönetimi Paketi, GO-Minus projesi ile aynı lisans altında dağıtılmaktadır.

# GO-Minus Standart Kütüphane Referansı

Bu belge, GO-Minus programlama dilinin standart kütüphanesinin kapsamlı bir referansını sağlamaktadır. GO-Minus standart kütüphanesi, Go'nun standart kütüphanesini temel alır ve GO-Minus'un sınıf, şablon ve istisna işleme gibi özelliklerini kullanarak genişletilmiştir.

## İçindekiler

1. [Temel Paketler](#temel-paketler)
2. [IO Paketi](#io-paketi)
3. [Container Paketi](#container-paketi)
4. [Net Paketi](#net-paketi)
5. [Time Paketi](#time-paketi)
6. [Regex Paketi](#regex-paketi)
7. [Concurrent Paketi](#concurrent-paketi)
8. [Math Paketi](#math-paketi)
9. [Encoding Paketi](#encoding-paketi)
10. [Crypto Paketi](#crypto-paketi)

## Temel Paketler

GO-Minus, Go'nun temel paketlerini içerir ve genişletir:

- `fmt`: Biçimlendirilmiş I/O işlemleri
- `io`: Temel I/O arayüzleri
- `os`: İşletim sistemi fonksiyonları
- `strings`: Dize işleme
- `errors`: Hata işleme

## IO Paketi

IO paketi, giriş/çıkış işlemleri için temel arayüzler ve fonksiyonlar sağlar.

### Temel IO Arayüzleri

```go
// Reader, okuma işlemleri için bir arayüzdür.
interface Reader {
    func Read(p []byte) (n int, err error)
}

// Writer, yazma işlemleri için bir arayüzdür.
interface Writer {
    func Write(p []byte) (n int, err error)
}

// Closer, kapatma işlemleri için bir arayüzdür.
interface Closer {
    func Close() error
}

// ReadWriter, hem okuma hem de yazma işlemleri için bir arayüzdür.
interface ReadWriter {
    Reader
    Writer
}
```

### Buffered IO

Buffered IO paketi, performansı artırmak için veri okuma ve yazma işlemlerini gruplar halinde gerçekleştirir.

```go
// BufferedReader, bir Reader'ı tamponlayan bir okuyucudur.
class BufferedReader {
    public:
        // Yeni bir BufferedReader oluşturur.
        static func New(reader io.Reader, bufferSize int) *BufferedReader

        // Tampondan veri okur.
        func Read(p []byte) (n int, err error)

        // Bir satır okur.
        func ReadLine() (string, error)

        // Tamponun boyutunu döndürür.
        func BufferSize() int
}

// BufferedWriter, bir Writer'ı tamponlayan bir yazıcıdır.
class BufferedWriter {
    public:
        // Yeni bir BufferedWriter oluşturur.
        static func New(writer io.Writer, bufferSize int) *BufferedWriter

        // Tampona veri yazar.
        func Write(p []byte) (n int, err error)

        // Bir dize yazar.
        func WriteString(s string) (n int, err error)

        // Tamponu boşaltır.
        func Flush() error

        // Tamponun boyutunu döndürür.
        func BufferSize() int
}
```

Daha fazla bilgi için [Buffered IO Örneği](examples/io/buffered-io-example.md) belgesine bakın.

### Memory-mapped IO

Memory-mapped IO paketi (`io/mmap`), dosyaları doğrudan belleğe eşleyerek, normal dosya I/O işlemlerine göre daha hızlı erişim sağlar. Bu paket, özellikle büyük dosyalarla çalışırken performans avantajları sunar.

#### Temel Sınıflar ve Fonksiyonlar

```go
// MMapFile, belleğe eşlenmiş bir dosyayı temsil eder.
class MMapFile {
    private:
        void* data
        int64 size
        int prot
        int flags
        bool isMapped

    public:
        // Yeni bir MMapFile oluşturur.
        static func New(file *os.File, prot int, flags int) (*MMapFile, error)

        // Belleğe eşlenmiş dosyayı kapatır.
        func Close() error

        // Belleğe eşlenmiş dosyanın boyutunu döndürür.
        func Len() int64

        // Belleğe eşlenmiş dosyanın veri işaretçisini döndürür.
        func Data() unsafe.Pointer

        // Belleğe eşlenmiş dosyanın içeriğini byte dizisi olarak döndürür.
        func Bytes() []byte

        // Belleğe eşlenmiş dosyadan belirtilen konumdan itibaren veri okur.
        func ReadAt(p []byte, off int64) (n int, err error)

        // Belleğe eşlenmiş dosyaya belirtilen konumdan itibaren veri yazar.
        func WriteAt(p []byte, off int64) (n int, err error)

        // Belleğe eşlenmiş dosyadaki değişiklikleri diske yazar.
        func Flush() error
}

// MMapError, memory-mapped IO işlemlerinde oluşan hataları temsil eder.
class MMapError extends error {
    private:
        string message
        int code

    public:
        MMapError(string message, int code)
        string Error()
        int Code()
}
```

#### Koruma ve Eşleme Bayrakları

```go
// Koruma bayrakları
const (
    PROT_READ  = 0x1 // Okuma izni
    PROT_WRITE = 0x2 // Yazma izni
    PROT_EXEC  = 0x4 // Çalıştırma izni
)

// Eşleme bayrakları
const (
    MAP_SHARED  = 0x01 // Değişiklikler diğer süreçlerle paylaşılır
    MAP_PRIVATE = 0x02 // Değişiklikler özeldir
)
```

#### Yardımcı Fonksiyonlar

```go
// Bir dosyayı belleğe eşler.
func Map(file *os.File, prot int, flags int) (*MMapFile, error)

// Bir dosyanın belirli bir bölgesini belleğe eşler.
func MapRegion(file *os.File, length int64, prot int, flags int, offset int64) (*MMapFile, error)
```

#### Örnek Kullanım

```go
// Dosyayı aç
file, err := os.OpenFile("data.bin", os.O_RDWR, 0)
if err != nil {
    panic(err)
}
defer file.Close()

// Dosyayı belleğe eşle (okuma ve yazma izni ile)
mmapFile, err := mmap.Map(file, mmap.PROT_READ|mmap.PROT_WRITE, mmap.MAP_SHARED)
if err != nil {
    panic(err)
}
defer mmapFile.Close()

// Belleğe eşlenmiş dosyadan oku
data := make([]byte, 10)
n, err := mmapFile.ReadAt(data, 0)
if err != nil {
    panic(err)
}
fmt.Printf("Okunan: %d bayt, İçerik: %v\n", n, data)

// Belleğe eşlenmiş dosyaya yaz
writeData := []byte{1, 2, 3, 4, 5}
n, err = mmapFile.WriteAt(writeData, 0)
if err != nil {
    panic(err)
}
fmt.Printf("Yazılan: %d bayt\n", n)

// Değişiklikleri diske yaz
err = mmapFile.Flush()
if err != nil {
    panic(err)
}
```

#### Performans Avantajları

Memory-mapped IO, özellikle büyük dosyalarla çalışırken, normal dosya I/O işlemlerine göre önemli performans avantajları sağlar:

1. **Daha Az Sistem Çağrısı**: Dosya belleğe eşlendikten sonra, okuma ve yazma işlemleri için sistem çağrısı gerekmez.
2. **Sayfa Önbelleği**: İşletim sistemi, belleğe eşlenmiş dosyaları sayfa önbelleğinde tutar, bu da tekrarlanan erişimleri hızlandırır.
3. **Sıfır Kopyalama**: Veriler, kullanıcı alanı ve çekirdek alanı arasında kopyalanmadan doğrudan erişilebilir.
4. **Talep Üzerine Sayfalama**: İşletim sistemi, sadece erişilen sayfaları belleğe yükler, bu da büyük dosyalarla çalışırken bellek kullanımını azaltır.

Daha fazla bilgi için [Memory-mapped IO Belgelendirmesi](../stdlib/io/mmap/README.md) belgesine bakın.

## Container Paketi

Container paketi, çeşitli veri yapıları sağlar.

### Trie (Önek Ağacı)

Trie, string anahtarları verimli bir şekilde depolamak ve aramak için kullanılan bir ağaç veri yapısıdır.

```go
// Trie, bir önek ağacını temsil eder.
class Trie<T> {
    public:
        // Yeni bir Trie oluşturur.
        static func New<T>() *Trie<T>

        // Bir anahtar-değer çifti ekler.
        func Insert(key string, value T)

        // Bir anahtarı arar.
        func Search(key string) (T, bool)

        // Bir anahtarı siler.
        func Delete(key string) bool

        // Belirli bir önekle başlayan tüm anahtarları döndürür.
        func GetWordsWithPrefix(prefix string) map[string]T

        // Trie'deki tüm anahtarları döndürür.
        func GetAllWords() map[string]T

        // Trie'nin boş olup olmadığını kontrol eder.
        func IsEmpty() bool

        // Trie'deki anahtar sayısını döndürür.
        func Size() int

        // Trie'yi temizler.
        func Clear()
}
```

Daha fazla bilgi için [Trie Örneği](examples/container/trie-example.md) belgesine bakın.

## Net Paketi

Net paketi (`net`), ağ programlama için sınıflar ve fonksiyonlar sağlar. TCP ve UDP protokolleri üzerinden ağ bağlantıları oluşturma, veri gönderme ve alma işlemleri için gerekli araçları içerir.

### Temel Arayüzler

```go
// Addr, bir ağ adresi arayüzünü temsil eder.
interface Addr {
    func Network() string
    func String() string
}

// Conn, bir ağ bağlantısı arayüzünü temsil eder.
interface Conn {
    func Read(b []byte) (n int, err error)
    func Write(b []byte) (n int, err error)
    func Close() error
    func LocalAddr() Addr
    func RemoteAddr() Addr
    func SetDeadline(t time.Time) error
    func SetReadDeadline(t time.Time) error
    func SetWriteDeadline(t time.Time) error
}

// Listener, bir ağ dinleyicisi arayüzünü temsil eder.
interface Listener {
    func Accept() (Conn, error)
    func Close() error
    func Addr() Addr
}
```

### IP Adresleri

```go
// IPAddr, bir IP adresini temsil eder.
class IPAddr {
    private:
        net.IPAddr goIPAddr

    public:
        // Yeni bir IPAddr oluşturur.
        static func New(ip []byte) *IPAddr

        // Bir IP adresi dizesini ayrıştırır.
        static func Parse(s string) (*IPAddr, error)

        // Ağ adı döndürür.
        func Network() string

        // IP adresini bir dize olarak döndürür.
        func String() string

        // IP adresinin loopback olup olmadığını kontrol eder.
        func IsLoopback() bool

        // IP adresinin global unicast olup olmadığını kontrol eder.
        func IsGlobalUnicast() bool

        // IP adresinin multicast olup olmadığını kontrol eder.
        func IsMulticast() bool

        // IP adresinin belirtilmemiş olup olmadığını kontrol eder.
        func IsUnspecified() bool

        // IPv4 adresini döndürür veya IPv4 değilse nil döndürür.
        func To4() []byte

        // IPv6 adresini döndürür.
        func To16() []byte
}
```

### TCP ve UDP Adresleri

```go
// TCPAddr, bir TCP adresini temsil eder.
class TCPAddr {
    private:
        net.TCPAddr goTCPAddr

    public:
        // Yeni bir TCPAddr oluşturur.
        static func New(ip []byte, port int) *TCPAddr

        // Bir TCP adresi dizesini ayrıştırır.
        static func Parse(s string) (*TCPAddr, error)

        // Ağ adı döndürür.
        func Network() string

        // TCP adresini bir dize olarak döndürür.
        func String() string

        // IP adresini döndürür.
        func IP() []byte

        // Port numarasını döndürür.
        func Port() int
}

// UDPAddr, bir UDP adresini temsil eder.
class UDPAddr {
    private:
        net.UDPAddr goUDPAddr

    public:
        // Yeni bir UDPAddr oluşturur.
        static func New(ip []byte, port int) *UDPAddr

        // Bir UDP adresi dizesini ayrıştırır.
        static func Parse(s string) (*UDPAddr, error)

        // Ağ adı döndürür.
        func Network() string

        // UDP adresini bir dize olarak döndürür.
        func String() string

        // IP adresini döndürür.
        func IP() []byte

        // Port numarasını döndürür.
        func Port() int
}
```

### TCP ve UDP Bağlantıları

```go
// TCPConn, bir TCP bağlantısını temsil eder.
class TCPConn {
    private:
        net.TCPConn goTCPConn

    public:
        // Bağlantıdan veri okur.
        func Read(b []byte) (n int, err error)

        // Bağlantıya veri yazar.
        func Write(b []byte) (n int, err error)

        // Bağlantıyı kapatır.
        func Close() error

        // Yerel adresi döndürür.
        func LocalAddr() Addr

        // Uzak adresi döndürür.
        func RemoteAddr() Addr

        // Bağlantı için son tarih ayarlar.
        func SetDeadline(t time.Time) error

        // Okuma işlemi için son tarih ayarlar.
        func SetReadDeadline(t time.Time) error

        // Yazma işlemi için son tarih ayarlar.
        func SetWriteDeadline(t time.Time) error

        // Nagle algoritmasını devre dışı bırakır.
        func SetNoDelay(noDelay bool) error

        // Keep-alive özelliğini etkinleştirir.
        func SetKeepAlive(keepalive bool) error

        // Keep-alive periyodunu ayarlar.
        func SetKeepAlivePeriod(d time.Duration) error
}

// UDPConn, bir UDP bağlantısını temsil eder.
class UDPConn {
    private:
        net.UDPConn goUDPConn

    public:
        // Bağlantıdan veri okur.
        func Read(b []byte) (n int, err error)

        // Bağlantıya veri yazar.
        func Write(b []byte) (n int, err error)

        // Bağlantıyı kapatır.
        func Close() error

        // Yerel adresi döndürür.
        func LocalAddr() Addr

        // Uzak adresi döndürür.
        func RemoteAddr() Addr

        // Bağlantı için son tarih ayarlar.
        func SetDeadline(t time.Time) error

        // Okuma işlemi için son tarih ayarlar.
        func SetReadDeadline(t time.Time) error

        // Yazma işlemi için son tarih ayarlar.
        func SetWriteDeadline(t time.Time) error

        // Belirli bir UDP adresinden veri okur.
        func ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)

        // Belirli bir UDP adresine veri yazar.
        func WriteToUDP(b []byte, addr *UDPAddr) (n int, err error)
}
```

### TCP Dinleyici

```go
// TCPListener, bir TCP dinleyicisini temsil eder.
class TCPListener {
    private:
        net.TCPListener goTCPListener

    public:
        // Bir bağlantı kabul eder.
        func Accept() (Conn, error)

        // Dinleyiciyi kapatır.
        func Close() error

        // Dinleyici adresini döndürür.
        func Addr() Addr

        // Kabul işlemi için son tarih ayarlar.
        func SetDeadline(t time.Time) error
}
```

### Yardımcı Fonksiyonlar

```go
// Belirtilen ağ ve adresle bir bağlantı kurar.
func Dial(network string, address string) (Conn, error)

// Belirtilen ağ ve adresle bir bağlantı kurar ve zaman aşımı ayarlar.
func DialTimeout(network string, address string, timeout time.Duration) (Conn, error)

// Belirtilen ağ ve adresle bir dinleyici oluşturur.
func Listen(network string, address string) (Listener, error)

// Bir ana bilgisayar adını IP adreslerine çözer.
func LookupHost(host string) ([]string, error)

// Bir ana bilgisayar adını IP adreslerine çözer.
func LookupIP(host string) ([]*IPAddr, error)

// Bir servis adını port numarasına çözer.
func LookupPort(network string, service string) (int, error)
```

### Örnek Kullanım: TCP İstemci

```go
// TCP bağlantısı kur
conn, err := net.Dial("tcp", "example.com:80")
if err != nil {
    fmt.Println("Bağlantı hatası:", err)
    return
}
defer conn.Close()

// Veri gönder
request := "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"
_, err = conn.Write([]byte(request))
if err != nil {
    fmt.Println("Yazma hatası:", err)
    return
}

// Yanıt al
buffer := make([]byte, 1024)
n, err := conn.Read(buffer)
if err != nil {
    fmt.Println("Okuma hatası:", err)
    return
}

fmt.Println("Yanıt:", string(buffer[:n]))
```

### Örnek Kullanım: TCP Sunucu

```go
// TCP dinleyici oluştur
listener, err := net.Listen("tcp", ":8080")
if err != nil {
    fmt.Println("Dinleyici oluşturma hatası:", err)
    return
}
defer listener.Close()

fmt.Println("Sunucu başlatıldı: 8080 portu dinleniyor...")

for {
    // Bağlantı kabul et
    conn, err := listener.Accept()
    if err != nil {
        fmt.Println("Bağlantı kabul hatası:", err)
        continue
    }

    // Her bağlantı için yeni bir goroutine başlat
    go func(c net.Conn) {
        defer c.Close()

        // İstemci bilgilerini göster
        fmt.Printf("İstemci bağlandı: %s\n", c.RemoteAddr())

        // Veri oku
        buffer := make([]byte, 1024)
        n, err := c.Read(buffer)
        if err != nil {
            fmt.Println("Okuma hatası:", err)
            return
        }

        fmt.Printf("Alınan veri: %s\n", string(buffer[:n]))

        // Yanıt gönder
        response := "Merhaba, istemci!"
        _, err = c.Write([]byte(response))
        if err != nil {
            fmt.Println("Yazma hatası:", err)
            return
        }
    }(conn)
}
```

Daha fazla bilgi için [Net Paketi Belgelendirmesi](../stdlib/net/README.md) belgesine bakın.

## Time Paketi

Time paketi (`time`), zaman noktaları, süreler, zaman dilimleri ve zaman biçimlendirme işlemleri için sınıflar ve fonksiyonlar sağlar.

### Zaman Noktaları

```go
// Time, bir zaman noktasını temsil eder.
class Time {
    private:
        time.Time goTime

    public:
        // Belirtilen zaman değerleriyle yeni bir Time oluşturur.
        static func New(year int, month int, day int, hour int, min int, sec int, nsec int, loc *Location) Time

        // Şu anki zamanı döndürür.
        static func Now() Time

        // Unix zaman damgasından bir Time oluşturur.
        static func Unix(sec int64, nsec int64) Time

        // Belirtilen düzende bir zaman dizesini ayrıştırır.
        static func Parse(layout string, value string) (Time, error)

        // Yılı döndürür.
        func Year() int

        // Ayı döndürür (1-12).
        func Month() int

        // Ayın gününü döndürür.
        func Day() int

        // Saati döndürür (0-23).
        func Hour() int

        // Dakikayı döndürür (0-59).
        func Minute() int

        // Saniyeyi döndürür (0-59).
        func Second() int

        // Nanosaniyeyi döndürür (0-999999999).
        func Nanosecond() int

        // Haftanın gününü döndürür (0-6, 0 = Pazar).
        func Weekday() int

        // Yılın gününü döndürür (1-365/366).
        func YearDay() int

        // Zaman dilimini döndürür.
        func Location() *Location

        // UTC zaman dilimindeki zamanı döndürür.
        func UTC() Time

        // Yerel zaman dilimindeki zamanı döndürür.
        func Local() Time

        // Belirtilen zaman dilimindeki zamanı döndürür.
        func In(loc *Location) Time

        // Unix zaman damgasını döndürür.
        func Unix() int64

        // Unix zaman damgasını nanosaniye cinsinden döndürür.
        func UnixNano() int64

        // Zamanı belirtilen düzende biçimlendirir.
        func Format(layout string) string

        // Zamanı RFC3339 formatında bir dize olarak döndürür.
        func String() string

        // Belirtilen süreyi ekler.
        func Add(d Duration) Time

        // İki zaman arasındaki farkı döndürür.
        func Sub(u Time) Duration

        // Belirtilen yıl, ay ve gün sayısını ekler.
        func AddDate(years int, months int, days int) Time

        // t'nin u'dan önce olup olmadığını kontrol eder.
        func Before(u Time) bool

        // t'nin u'dan sonra olup olmadığını kontrol eder.
        func After(u Time) bool

        // t'nin u'ya eşit olup olmadığını kontrol eder.
        func Equal(u Time) bool
}
```

### Süreler

```go
// Duration, iki zaman noktası arasındaki süreyi temsil eder.
class Duration {
    private:
        time.Duration goDuration

    public:
        // Nanosaniye cinsinden bir süre oluşturur.
        static func New(nanoseconds int64) Duration

        // Süreyi nanosaniye cinsinden döndürür.
        func Nanoseconds() int64

        // Süreyi mikrosaniye cinsinden döndürür.
        func Microseconds() int64

        // Süreyi milisaniye cinsinden döndürür.
        func Milliseconds() int64

        // Süreyi saniye cinsinden döndürür.
        func Seconds() float64

        // Süreyi dakika cinsinden döndürür.
        func Minutes() float64

        // Süreyi saat cinsinden döndürür.
        func Hours() float64

        // Süreyi bir dize olarak döndürür.
        func String() string
}
```

### Zaman Dilimleri

```go
// Location, bir zaman dilimini temsil eder.
class Location {
    private:
        time.Location goLoc

    public:
        // Belirtilen isimle bir zaman dilimi yükler.
        static func LoadLocation(name string) (*Location, error)

        // Belirtilen isim ve ofsetle sabit bir zaman dilimi oluşturur.
        static func FixedZone(name string, offset int) *Location

        // Zaman diliminin adını döndürür.
        func String() string
}
```

### Zaman Sabitleri

```go
// Süre sabitleri
const (
    Nanosecond  = 1
    Microsecond = 1000 * Nanosecond
    Millisecond = 1000 * Microsecond
    Second      = 1000 * Millisecond
    Minute      = 60 * Second
    Hour        = 60 * Minute
)

// Tarih/saat biçimlendirme düzenleri
const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700"
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    Stamp       = "Jan _2 15:04:05"
    StampMilli  = "Jan _2 15:04:05.000"
    StampMicro  = "Jan _2 15:04:05.000000"
    StampNano   = "Jan _2 15:04:05.000000000"
    DateTime    = "2006-01-02 15:04:05"
    DateOnly    = "2006-01-02"
    TimeOnly    = "15:04:05"
)
```

### Yardımcı Fonksiyonlar

```go
// Belirtilen süre kadar bekler.
func Sleep(d Duration)

// Belirtilen süre sonra bir değer gönderen bir kanal döndürür.
func After(d Duration) <-chan Time

// Belirtilen aralıklarla bir değer gönderen bir kanal döndürür.
func Tick(d Duration) <-chan Time

// Belirtilen zamandan bu yana geçen süreyi döndürür.
func Since(t Time) Duration

// Belirtilen zamana kadar kalan süreyi döndürür.
func Until(t Time) Duration

// Bir süre dizesini ayrıştırır.
func ParseDuration(s string) (Duration, error)
```

### Örnek Kullanım

```go
// Şu anki zamanı al
now := time.Now()
fmt.Println("Şu an:", now)

// Belirli bir zaman oluştur
t := time.New(2023, 5, 15, 14, 30, 0, 0, time.UTC)
fmt.Println("Belirli zaman:", t)

// Zaman bileşenlerini al
fmt.Printf("Yıl: %d, Ay: %d, Gün: %d\n", t.Year(), t.Month(), t.Day())
fmt.Printf("Saat: %d, Dakika: %d, Saniye: %d\n", t.Hour(), t.Minute(), t.Second())

// Zaman biçimlendirme
fmt.Println("RFC3339:", t.Format(time.RFC3339))
fmt.Println("Tarih:", t.Format(time.DateOnly))
fmt.Println("Saat:", t.Format(time.TimeOnly))

// Zaman ayrıştırma
timeStr := "2023-05-15T14:30:00Z"
parsedTime, err := time.Parse(time.RFC3339, timeStr)
if err != nil {
    fmt.Println("Ayrıştırma hatası:", err)
} else {
    fmt.Println("Ayrıştırılan zaman:", parsedTime)
}

// Süre oluştur
d := time.New(5 * time.Second)
fmt.Println("Süre:", d)

// Süre bileşenlerini al
fmt.Printf("Saniye: %.2f\n", d.Seconds())
fmt.Printf("Milisaniye: %d\n", d.Milliseconds())

// Süre ayrıştırma
durationStr := "1h30m15s"
parsedDuration, err := time.ParseDuration(durationStr)
if err != nil {
    fmt.Println("Ayrıştırma hatası:", err)
} else {
    fmt.Println("Ayrıştırılan süre:", parsedDuration)
    fmt.Printf("Saat: %.2f\n", parsedDuration.Hours())
}

// Bekleme
fmt.Println("1 saniye bekleniyor...")
time.Sleep(time.Second)
fmt.Println("Bekleme tamamlandı.")

// Zamanlayıcı
fmt.Println("2 saniye sonra bir mesaj alacaksınız...")
<-time.After(2 * time.Second)
fmt.Println("2 saniye geçti!")
```

Daha fazla bilgi için [Time Paketi Belgelendirmesi](../stdlib/time/README.md) belgesine bakın.

## Regex Paketi

Regex paketi, düzenli ifadeler için sınıflar ve fonksiyonlar sağlar.

```go
// RegexPattern, bir düzenli ifade desenini temsil eder.
class RegexPattern {
    public:
        // Yeni bir RegexPattern oluşturur.
        static func New(pattern string, caseSensitive bool, multiline bool) *RegexPattern

        // Bir metni düzenli ifade deseniyle eşleştirir.
        func Match(text string) bool

        // Bir metindeki tüm eşleşmeleri bulur.
        func FindAll(text string) []string

        // Bir metindeki tüm eşleşmeleri belirtilen metinle değiştirir.
        func Replace(text string, replacement string) string

        // Bir metindeki ilk eşleşmeyi belirtilen metinle değiştirir.
        func ReplaceFirst(text string, replacement string) string
}

// Bir düzenli ifade desenini derler.
func Compile(pattern string) *RegexPattern

// Bir düzenli ifade desenini büyük/küçük harf duyarsız olarak derler.
func CompileIgnoreCase(pattern string) *RegexPattern

// Bir düzenli ifade desenini çok satırlı olarak derler.
func CompileMultiline(pattern string) *RegexPattern

// Bir metni düzenli ifade deseniyle eşleştirir.
func Match(pattern string, text string) bool

// Bir metindeki tüm eşleşmeleri belirtilen metinle değiştirir.
func Replace(pattern string, text string, replacement string) string

// Bir metni belirtilen desene göre böler.
func Split(pattern string, text string) []string
```

Daha fazla bilgi için [Regex Örneği](examples/regex/regex-example.md) belgesine bakın.

## Concurrent Paketi

Concurrent paketi, eşzamanlılık işlemleri için sınıflar ve fonksiyonlar sağlar.

```go
// Semaphore, bir semafor veri yapısını temsil eder.
class Semaphore {
    public:
        // Yeni bir Semaphore oluşturur.
        static func New(count int) *Semaphore

        // Semafor değerini azaltır.
        func Acquire()

        // Semafor değerini artırır.
        func Release()

        // Semafor değerini döndürür.
        func Count() int
}

// Barrier, bir bariyer veri yapısını temsil eder.
class Barrier {
    public:
        // Yeni bir Barrier oluşturur.
        static func New(count int) *Barrier

        // Bariyerde bekler.
        func Wait()

        // Bariyer sayısını döndürür.
        func Count() int
}

// ThreadPool, bir iş parçacığı havuzunu temsil eder.
class ThreadPool {
    public:
        // Yeni bir ThreadPool oluşturur.
        static func New(size int) *ThreadPool

        // Bir görevi havuza ekler.
        func Submit(task func()) Future<void>

        // Havuzu kapatır.
        func Shutdown()
}

// Future, gelecekte tamamlanacak bir işlemi temsil eder.
class Future<T> {
    public:
        // İşlemin tamamlanmasını bekler ve sonucu döndürür.
        func Get() T

        // İşlemin tamamlanıp tamamlanmadığını kontrol eder.
        func IsDone() bool

        // İşlemi iptal eder.
        func Cancel() bool
}
```

## Math Paketi

Math paketi, matematiksel işlemler için fonksiyonlar ve sabitler sağlar.

```go
// Matematiksel sabitler
const (
    E   = 2.71828182845904523536028747135266249775724709369995957496696763
    Pi  = 3.14159265358979323846264338327950288419716939937510582097494459
    Phi = 1.61803398874989484820458683436563811772030917980576286213544862
)

// Mutlak değer
func Abs(x float64) float64

// Karekök
func Sqrt(x float64) float64

// Üs alma
func Pow(x, y float64) float64

// Logaritma
func Log(x float64) float64
func Log10(x float64) float64

// Trigonometrik fonksiyonlar
func Sin(x float64) float64
func Cos(x float64) float64
func Tan(x float64) float64
```

## Encoding Paketi

Encoding paketi, veri kodlama ve çözme işlemleri için sınıflar ve fonksiyonlar sağlar.

```go
// JSON kodlama ve çözme
func MarshalJSON(v interface{}) ([]byte, error)
func UnmarshalJSON(data []byte, v interface{}) error

// XML kodlama ve çözme
func MarshalXML(v interface{}) ([]byte, error)
func UnmarshalXML(data []byte, v interface{}) error

// Base64 kodlama ve çözme
func EncodeBase64(src []byte) string
func DecodeBase64(s string) ([]byte, error)
```

## Crypto Paketi

Crypto paketi, kriptografik işlemler için sınıflar ve fonksiyonlar sağlar.

```go
// Hash fonksiyonları
func MD5(data []byte) []byte
func SHA1(data []byte) []byte
func SHA256(data []byte) []byte

// HMAC
func HMAC(key, data []byte, hash func([]byte) []byte) []byte

// AES şifreleme
func AESEncrypt(key, plaintext []byte) ([]byte, error)
func AESDecrypt(key, ciphertext []byte) ([]byte, error)

// RSA şifreleme
func RSAEncrypt(publicKey, plaintext []byte) ([]byte, error)
func RSADecrypt(privateKey, ciphertext []byte) ([]byte, error)
```

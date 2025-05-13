# GO-Minus Net Paketi

Bu paket, GO-Minus programlama dili için ağ işlemlerini sağlar. TCP ve UDP protokolleri üzerinden ağ bağlantıları oluşturma, veri gönderme ve alma işlemleri için fonksiyonlar ve sınıflar içerir.

## Özellikler

- TCP ve UDP bağlantıları oluşturma
- Ağ sunucuları oluşturma
- IP adresleri ve port numaraları işleme
- Ana bilgisayar adı çözümleme
- Zaman aşımı ve son tarih ayarlama
- Keep-alive ve diğer soket seçenekleri

## Kullanım

### TCP İstemci

```go
import (
    "fmt"
    "net"
)

func main() {
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
}
```

### TCP Sunucu

```go
import (
    "fmt"
    "net"
)

func main() {
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
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    
    // İstemci bilgilerini göster
    fmt.Printf("İstemci bağlandı: %s\n", conn.RemoteAddr())
    
    // Veri oku
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Okuma hatası:", err)
        return
    }
    
    fmt.Printf("Alınan veri: %s\n", string(buffer[:n]))
    
    // Yanıt gönder
    response := "Merhaba, istemci!"
    _, err = conn.Write([]byte(response))
    if err != nil {
        fmt.Println("Yazma hatası:", err)
        return
    }
}
```

### UDP İstemci

```go
import (
    "fmt"
    "net"
)

func main() {
    // UDP bağlantısı kur
    conn, err := net.Dial("udp", "example.com:53")
    if err != nil {
        fmt.Println("Bağlantı hatası:", err)
        return
    }
    defer conn.Close()
    
    // Veri gönder
    message := "Merhaba, UDP sunucu!"
    _, err = conn.Write([]byte(message))
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
}
```

### UDP Sunucu

```go
import (
    "fmt"
    "net"
)

func main() {
    // UDP adresi oluştur
    addr, err := net.UDPAddr.Parse(":8053")
    if err != nil {
        fmt.Println("Adres ayrıştırma hatası:", err)
        return
    }
    
    // UDP dinleyici oluştur
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Dinleyici oluşturma hatası:", err)
        return
    }
    defer conn.Close()
    
    fmt.Println("UDP sunucu başlatıldı: 8053 portu dinleniyor...")
    
    buffer := make([]byte, 1024)
    for {
        // Veri al
        n, clientAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("Okuma hatası:", err)
            continue
        }
        
        fmt.Printf("İstemci: %s, Veri: %s\n", clientAddr, string(buffer[:n]))
        
        // Yanıt gönder
        response := "Merhaba, UDP istemci!"
        _, err = conn.WriteToUDP([]byte(response), clientAddr)
        if err != nil {
            fmt.Println("Yazma hatası:", err)
            continue
        }
    }
}
```

### IP Adresleri

```go
import (
    "fmt"
    "net"
)

func main() {
    // IP adresi ayrıştır
    ipAddr, err := net.IPAddr.Parse("192.168.1.1")
    if err != nil {
        fmt.Println("IP adresi ayrıştırma hatası:", err)
        return
    }
    
    fmt.Println("IP adresi:", ipAddr)
    
    // IP özelliklerini kontrol et
    fmt.Println("Loopback mu:", ipAddr.IsLoopback())
    fmt.Println("Global unicast mı:", ipAddr.IsGlobalUnicast())
    fmt.Println("Multicast mı:", ipAddr.IsMulticast())
    
    // IPv4 ve IPv6 dönüşümleri
    ip4 := ipAddr.To4()
    if ip4 != nil {
        fmt.Println("IPv4:", ip4)
    }
    
    ip6 := ipAddr.To16()
    fmt.Println("IPv6:", ip6)
    
    // Ana bilgisayar adı çözümleme
    ips, err := net.LookupIP("example.com")
    if err != nil {
        fmt.Println("Ana bilgisayar adı çözümleme hatası:", err)
        return
    }
    
    fmt.Println("example.com IP adresleri:")
    for _, ip := range ips {
        fmt.Println("  ", ip)
    }
    
    // Port numarası çözümleme
    port, err := net.LookupPort("tcp", "http")
    if err != nil {
        fmt.Println("Port çözümleme hatası:", err)
        return
    }
    
    fmt.Println("HTTP port numarası:", port)
}
```

### Zaman Aşımı ve Son Tarih

```go
import (
    "fmt"
    "net"
    "time"
)

func main() {
    // Zaman aşımı ile bağlantı kur
    conn, err := net.DialTimeout("tcp", "example.com:80", 5*time.Second)
    if err != nil {
        fmt.Println("Bağlantı hatası:", err)
        return
    }
    defer conn.Close()
    
    // Okuma için son tarih ayarla
    err = conn.SetReadDeadline(time.Now().Add(10*time.Second))
    if err != nil {
        fmt.Println("Son tarih ayarlama hatası:", err)
        return
    }
    
    // Yazma için son tarih ayarla
    err = conn.SetWriteDeadline(time.Now().Add(5*time.Second))
    if err != nil {
        fmt.Println("Son tarih ayarlama hatası:", err)
        return
    }
    
    // Veri gönder ve al
    // ...
}
```

## TCP Seçenekleri

TCP bağlantıları için çeşitli seçenekler ayarlanabilir:

```go
import (
    "fmt"
    "net"
    "time"
)

func main() {
    // TCP bağlantısı kur
    conn, err := net.Dial("tcp", "example.com:80")
    if err != nil {
        fmt.Println("Bağlantı hatası:", err)
        return
    }
    defer conn.Close()
    
    tcpConn := conn.(*net.TCPConn)
    
    // Nagle algoritmasını devre dışı bırak
    err = tcpConn.SetNoDelay(true)
    if err != nil {
        fmt.Println("SetNoDelay hatası:", err)
        return
    }
    
    // Keep-alive özelliğini etkinleştir
    err = tcpConn.SetKeepAlive(true)
    if err != nil {
        fmt.Println("SetKeepAlive hatası:", err)
        return
    }
    
    // Keep-alive periyodunu ayarla
    err = tcpConn.SetKeepAlivePeriod(30*time.Second)
    if err != nil {
        fmt.Println("SetKeepAlivePeriod hatası:", err)
        return
    }
    
    // Veri gönder ve al
    // ...
}
```

## Performans İpuçları

1. **Bağlantı Havuzu**: Çok sayıda kısa süreli bağlantı yerine, bağlantı havuzu kullanın.
2. **Tamponlama**: Küçük paketler yerine, daha büyük tamponlar kullanın.
3. **Goroutine'ler**: Her bağlantı için ayrı bir goroutine kullanın.
4. **Keep-alive**: Uzun süreli bağlantılar için keep-alive özelliğini etkinleştirin.
5. **NoDelay**: Düşük gecikme gerektiren uygulamalar için Nagle algoritmasını devre dışı bırakın.

## Sınırlamalar

1. **Güvenlik**: Bu paket, TLS/SSL şifreleme sağlamaz. Güvenli bağlantılar için `crypto/tls` paketini kullanın.
2. **Protokoller**: Sadece TCP ve UDP protokollerini destekler. Diğer protokoller için ek paketler gerekebilir.
3. **IPv6**: IPv6 desteği sağlanmıştır, ancak bazı sistemlerde sınırlamalar olabilir.

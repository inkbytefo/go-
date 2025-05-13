# GO-Minus Time Paketi

Bu paket, GO-Minus programlama dili için zaman işlemlerini sağlar. Zaman noktaları, süreler, zaman dilimleri ve zaman biçimlendirme işlemleri için fonksiyonlar ve sınıflar içerir.

## Özellikler

- Zaman noktaları oluşturma ve işleme
- Süreler oluşturma ve işleme
- Zaman dilimleri yönetimi
- Zaman biçimlendirme ve ayrıştırma
- Zamanlayıcılar ve kronometre
- Tarih aritmetiği

## Kullanım

### Temel Zaman İşlemleri

```go
import (
    "fmt"
    "time"
)

func main() {
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
    fmt.Println("Özel format:", t.Format("2006-01-02 15:04:05 MST"))
    
    // Zaman ayrıştırma
    timeStr := "2023-05-15T14:30:00Z"
    parsedTime, err := time.Parse(time.RFC3339, timeStr)
    if err != nil {
        fmt.Println("Ayrıştırma hatası:", err)
    } else {
        fmt.Println("Ayrıştırılan zaman:", parsedTime)
    }
}
```

### Süreler

```go
import (
    "fmt"
    "time"
)

func main() {
    // Süre oluştur
    d1 := time.New(5 * time.Second)
    d2 := time.New(2 * time.Minute)
    
    fmt.Println("d1:", d1)
    fmt.Println("d2:", d2)
    
    // Süre aritmetiği
    sum := time.New(d1.Nanoseconds() + d2.Nanoseconds())
    fmt.Println("d1 + d2:", sum)
    
    // Süre bileşenlerini al
    fmt.Printf("Saniye: %.2f\n", d2.Seconds())
    fmt.Printf("Dakika: %.2f\n", d2.Minutes())
    
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
}
```

### Zaman Dilimleri

```go
import (
    "fmt"
    "time"
)

func main() {
    // Zaman dilimi yükle
    istanbul, err := time.LoadLocation("Europe/Istanbul")
    if err != nil {
        fmt.Println("Zaman dilimi yükleme hatası:", err)
        return
    }
    
    // Belirli bir zaman diliminde zaman oluştur
    t := time.New(2023, 5, 15, 14, 30, 0, 0, istanbul)
    fmt.Println("İstanbul zamanı:", t)
    
    // UTC'ye dönüştür
    utcTime := t.UTC()
    fmt.Println("UTC zamanı:", utcTime)
    
    // Yerel zamana dönüştür
    localTime := t.Local()
    fmt.Println("Yerel zaman:", localTime)
    
    // Sabit zaman dilimi oluştur
    fixedZone := time.FixedZone("GMT+3", 3*60*60)
    fixedTime := t.In(fixedZone)
    fmt.Println("Sabit zaman dilimi:", fixedTime)
}
```

### Tarih Aritmetiği

```go
import (
    "fmt"
    "time"
)

func main() {
    // Başlangıç zamanı
    t := time.New(2023, 5, 15, 14, 30, 0, 0, time.UTC)
    fmt.Println("Başlangıç zamanı:", t)
    
    // Süre ekle
    t1 := t.Add(time.Hour)
    fmt.Println("1 saat sonra:", t1)
    
    // Tarih ekle
    t2 := t.AddDate(0, 1, 0)
    fmt.Println("1 ay sonra:", t2)
    
    // İki zaman arasındaki fark
    diff := t2.Sub(t)
    fmt.Printf("Fark: %v (%.2f gün)\n", diff, diff.Hours()/24)
    
    // Karşılaştırma
    fmt.Println("t1 < t2:", t1.Before(t2))
    fmt.Println("t2 > t1:", t2.After(t1))
    
    // Geçen süre
    start := time.Now()
    // ... işlem ...
    elapsed := time.Since(start)
    fmt.Println("Geçen süre:", elapsed)
}
```

### Zamanlayıcılar

```go
import (
    "fmt"
    "time"
)

func main() {
    // Tek seferlik zamanlayıcı
    fmt.Println("2 saniye sonra bir mesaj alacaksınız...")
    <-time.After(2 * time.Second)
    fmt.Println("2 saniye geçti!")
    
    // Periyodik zamanlayıcı
    ticker := time.Tick(1 * time.Second)
    count := 0
    
    fmt.Println("Her saniye bir mesaj alacaksınız...")
    for t := range ticker {
        fmt.Println("Tik:", t)
        count++
        if count >= 5 {
            break
        }
    }
}
```

## Zaman Biçimlendirme Düzenleri

GO-Minus, zaman biçimlendirme için aşağıdaki önceden tanımlanmış düzenleri sağlar:

- `ANSIC`: "Mon Jan _2 15:04:05 2006"
- `UnixDate`: "Mon Jan _2 15:04:05 MST 2006"
- `RubyDate`: "Mon Jan 02 15:04:05 -0700 2006"
- `RFC822`: "02 Jan 06 15:04 MST"
- `RFC822Z`: "02 Jan 06 15:04 -0700"
- `RFC850`: "Monday, 02-Jan-06 15:04:05 MST"
- `RFC1123`: "Mon, 02 Jan 2006 15:04:05 MST"
- `RFC1123Z`: "Mon, 02 Jan 2006 15:04:05 -0700"
- `RFC3339`: "2006-01-02T15:04:05Z07:00"
- `RFC3339Nano`: "2006-01-02T15:04:05.999999999Z07:00"
- `Kitchen`: "3:04PM"
- `Stamp`: "Jan _2 15:04:05"
- `StampMilli`: "Jan _2 15:04:05.000"
- `StampMicro`: "Jan _2 15:04:05.000000"
- `StampNano`: "Jan _2 15:04:05.000000000"
- `DateTime`: "2006-01-02 15:04:05"
- `DateOnly`: "2006-01-02"
- `TimeOnly`: "15:04:05"

## Süre Sabitleri

GO-Minus, süre hesaplamaları için aşağıdaki sabitleri sağlar:

- `Nanosecond`: 1 nanosaniye
- `Microsecond`: 1 mikrosaniye (1000 nanosaniye)
- `Millisecond`: 1 milisaniye (1000 mikrosaniye)
- `Second`: 1 saniye (1000 milisaniye)
- `Minute`: 1 dakika (60 saniye)
- `Hour`: 1 saat (60 dakika)

## Performans

Time paketi, zaman işlemleri için yüksek performanslı bir implementasyon sağlar. Zaman noktaları, monoton saat kullanılarak oluşturulur, bu da sistem saati değişikliklerinden etkilenmez.

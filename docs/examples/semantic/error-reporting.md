# GO-Minus Gelişmiş Hata Raporlama Örneği

Bu örnek, GO-Minus'un gelişmiş hata raporlama sistemini göstermektedir. GO-Minus, semantik analiz sırasında karşılaşılan hataları, uyarıları ve bilgileri ayrıntılı bir şekilde raporlar.

## Hata Raporlama Özellikleri

- Farklı hata seviyeleri (hata, uyarı, bilgi)
- Dosya ve konum bilgisi
- Hata mesajı
- İpuçları ve düzeltme önerileri
- Renkli çıktı

## Örnek Kod

```go
// error_reporting.gom
package main

import "fmt"

func main() {
    // Tip uyuşmazlığı hatası
    var x int = 10
    var y string = "20"
    
    // Hata: int ve string toplanamaz
    z := x + y
    
    fmt.Println("Sonuç:", z)
}
```

## Hata Çıktısı

```
Satır 10, Sütun 9: HATA: Tip uyuşmazlığı: 'int' ve 'string' tipleri '+' operatörü ile kullanılamaz
İpuçları:
  - 'y' değişkenini int'e dönüştürmeyi deneyin: 'strconv.Atoi(y)'
  - 'x' değişkenini string'e dönüştürmeyi deneyin: 'fmt.Sprint(x)'
  - Veya 'fmt.Sprintf("%d%s", x, y)' kullanarak birleştirmeyi deneyin
```

## Düzeltilmiş Kod

```go
// error_reporting_fixed.gom
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Tip uyuşmazlığı hatası
    var x int = 10
    var y string = "20"
    
    // Çözüm 1: y'yi int'e dönüştür
    yInt, err := strconv.Atoi(y)
    if err != nil {
        fmt.Println("Dönüştürme hatası:", err)
        return
    }
    z1 := x + yInt
    
    // Çözüm 2: x'i string'e dönüştür
    z2 := fmt.Sprint(x) + y
    
    // Çözüm 3: fmt.Sprintf kullan
    z3 := fmt.Sprintf("%d%s", x, y)
    
    fmt.Println("Sonuç 1:", z1) // 30
    fmt.Println("Sonuç 2:", z2) // "1020"
    fmt.Println("Sonuç 3:", z3) // "1020"
}
```

## Hata Seviyeleri

GO-Minus, üç farklı hata seviyesi sunar:

1. **Hata (Error)**: Programın derlenmesini engelleyen ciddi sorunlar
2. **Uyarı (Warning)**: Potansiyel sorunlar, program derlenebilir ancak beklenmeyen davranışlara neden olabilir
3. **Bilgi (Info)**: Bilgilendirme amaçlı mesajlar, genellikle iyileştirme önerileri içerir

## Örnek Uyarı

```go
// warning_example.gom
package main

func main() {
    // Kullanılmayan değişken uyarısı
    x := 10
    
    // x değişkeni kullanılmıyor
}
```

## Uyarı Çıktısı

```
Satır 5, Sütun 5: UYARI: Değişken 'x' tanımlanmış ancak kullanılmamış
İpuçları:
  - Değişkeni kullanın
  - Değişkeni '_' ile değiştirin
  - Değişkeni kaldırın
```

## Hata Raporlama API'si

GO-Minus, hata raporlama için aşağıdaki API'yi sunar:

```go
// Hata raporlama örneği
func reportError(token Token, format string, args ...interface{}) *SemanticError
func reportWarning(token Token, format string, args ...interface{}) *SemanticError
func reportInfo(token Token, format string, args ...interface{}) *SemanticError

// İpucu ekleme
func (se *SemanticError) AddHint(format string, args ...interface{}) *SemanticError
```

## Hata Kurtarma

GO-Minus, semantik analiz sırasında hatalarla karşılaşıldığında analizi durdurmak yerine, mümkün olduğunca devam etmeye çalışır. Bu, tek bir çalıştırmada birden fazla hatanın tespit edilmesini sağlar.

```go
// error_recovery.gom
package main

import "fmt"

func main() {
    // Hata 1: Tanımlanmamış değişken
    fmt.Println(undefinedVar)
    
    // Hata 2: Tip uyuşmazlığı
    var x int = "string"
    
    // Hata 3: Bilinmeyen fonksiyon
    unknownFunction()
    
    // Bu satıra kadar analiz devam eder
    fmt.Println("Bu satır analiz edilir")
}
```

## Hata Kurtarma Çıktısı

```
Satır 7, Sütun 17: HATA: Tanımlanmamış tanımlayıcı: 'undefinedVar'
İpuçları:
  - Değişkeni kullanmadan önce tanımlayın
  - Yazım hatası olabilir, benzer değişkenler: []

Satır 10, Sütun 14: HATA: Tip uyuşmazlığı: 'string' tipindeki değer 'int' tipindeki değişkene atanamaz
İpuçları:
  - String'i int'e dönüştürmek için 'strconv.Atoi()' kullanın
  - Değişken tipini 'string' olarak değiştirin

Satır 13, Sütun 5: HATA: Tanımlanmamış fonksiyon: 'unknownFunction'
İpuçları:
  - Fonksiyonu kullanmadan önce tanımlayın
  - Yazım hatası olabilir, benzer fonksiyonlar: []
```

Bu örnekler, GO-Minus'un gelişmiş hata raporlama sisteminin nasıl kullanılacağını göstermektedir. Daha fazla bilgi için [Semantik Analiz Belgelendirmesi](../../docs/reference/semantic-analysis.md) belgesine bakabilirsiniz.

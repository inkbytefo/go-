# GO-Minus Semantik Analiz İyileştirmeleri

Bu belge, GO-Minus programlama dilinin semantik analiz bileşeninde yapılan iyileştirmeleri açıklamaktadır.

## Genel Bakış

Semantik analiz, bir programın anlamsal doğruluğunu kontrol eden derleyici aşamasıdır. Sözdizimi analizi (parsing) aşamasından sonra gerçekleştirilir ve programın semantik kurallarına uygunluğunu denetler. GO-Minus'un semantik analiz bileşeni, aşağıdaki iyileştirmelerle güçlendirilmiştir:

1. **Gelişmiş Hata Raporlama Sistemi**
2. **Tip Çıkarımı Modülü**
3. **Hata Kurtarma Mekanizmaları**

## 1. Gelişmiş Hata Raporlama Sistemi

GO-Minus'un gelişmiş hata raporlama sistemi, programcılara daha açıklayıcı ve yardımcı hata mesajları sunmak için tasarlanmıştır.

### Özellikler

- **Farklı Hata Seviyeleri**: Hatalar, uyarılar ve bilgi mesajları için farklı seviyeler
- **Renkli Çıktı**: Hata seviyelerine göre renklendirilmiş çıktı
- **Dosya ve Konum Bilgisi**: Hatanın tam olarak nerede oluştuğunu gösteren dosya, satır ve sütun bilgisi
- **İpuçları ve Düzeltme Önerileri**: Hatanın nasıl düzeltilebileceğine dair ipuçları
- **Benzer Tanımlayıcı Önerileri**: Yazım hatası olabilecek tanımlayıcılar için benzer isim önerileri

### Örnek Hata Mesajı

```
Satır 10, Sütun 9: HATA: Tip uyuşmazlığı: 'int' ve 'string' tipleri '+' operatörü ile kullanılamaz
İpuçları:
  - 'y' değişkenini int'e dönüştürmeyi deneyin: 'strconv.Atoi(y)'
  - 'x' değişkenini string'e dönüştürmeyi deneyin: 'fmt.Sprint(x)'
  - Veya 'fmt.Sprintf("%d%s", x, y)' kullanarak birleştirmeyi deneyin
```

### Uygulama Detayları

Gelişmiş hata raporlama sistemi, `internal/semantic/error.go` dosyasında uygulanmıştır. Sistem, aşağıdaki bileşenlerden oluşur:

- **ErrorLevel**: Hata seviyelerini tanımlayan enum (ERROR, WARNING, INFO)
- **SemanticError**: Hata bilgilerini içeren yapı
- **ErrorReporter**: Hataları toplayan ve raporlayan sınıf

```go
// ErrorLevel, hata seviyesini temsil eder.
type ErrorLevel int

const (
    ERROR ErrorLevel = iota
    WARNING
    INFO
)

// SemanticError, semantik analiz sırasında oluşan bir hatayı temsil eder.
type SemanticError struct {
    Level   ErrorLevel
    Token   token.Token
    Message string
    Hints   []string
}

// ErrorReporter, hata raporlama işlemlerini gerçekleştirir.
type ErrorReporter struct {
    Errors   []*SemanticError
    Warnings []*SemanticError
    Infos    []*SemanticError
}
```

## 2. Tip Çıkarımı Modülü

Tip çıkarımı modülü, değişken tiplerinin açıkça belirtilmediği durumlarda, derleyicinin değişken tiplerini otomatik olarak belirlemesini sağlar.

### Özellikler

- **Değişken Tanımlamalarında Tip Çıkarımı**: `:=` operatörü ile tanımlanan değişkenlerin tiplerini çıkarma
- **Fonksiyon Dönüş Tiplerinde Tip Çıkarımı**: Dönüş tipi belirtilmeyen fonksiyonların dönüş tiplerini çıkarma
- **Karmaşık İfadelerde Tip Çıkarımı**: Karmaşık ifadelerin sonuç tiplerini çıkarma
- **Jenerik Fonksiyonlarda Tip Çıkarımı**: Jenerik fonksiyonların tip parametrelerini çıkarma
- **Şablon Sınıflarda Tip Çıkarımı**: Şablon sınıfların tip parametrelerini çıkarma

### Örnek Kod

```go
// Değişken tanımlamalarında tip çıkarımı
x := 10          // int olarak çıkarılır
y := 3.14        // float olarak çıkarılır
z := "Merhaba"   // string olarak çıkarılır
b := true        // bool olarak çıkarılır

// Fonksiyon dönüş tiplerinde tip çıkarımı
func add(a, b int) {
    return a + b  // int olarak çıkarılır
}

// Jenerik fonksiyonlarda tip çıkarımı
func first<T>(items []T) T {
    return items[0]
}

// Tip parametreleri çıkarılır
intResult := first([3]int{1, 2, 3})  // T = int
```

### Uygulama Detayları

Tip çıkarımı modülü, `internal/semantic/inference.go` dosyasında uygulanmıştır. Modül, aşağıdaki bileşenlerden oluşur:

- **TypeInference**: Tip çıkarımı işlemlerini gerçekleştiren sınıf
- **InferType**: Bir ifadenin tipini çıkaran ana fonksiyon
- **Özel çıkarım fonksiyonları**: Farklı ifade türleri için özel çıkarım fonksiyonları

```go
// TypeInference, tip çıkarımı işlemlerini gerçekleştirir.
type TypeInference struct {
    analyzer *Analyzer
}

// InferType, bir ifadenin tipini çıkarır.
func (ti *TypeInference) InferType(expr ast.Expression) Type {
    switch e := expr.(type) {
    case *ast.Identifier:
        return ti.inferIdentifierType(e)
    case *ast.IntegerLiteral:
        return &BasicType{Name: "int", Kind: INTEGER_TYPE}
    // ... diğer ifade türleri
    }
}
```

## 3. Hata Kurtarma Mekanizmaları

Hata kurtarma mekanizmaları, semantik analiz sırasında hatalarla karşılaşıldığında analizi durdurmak yerine, mümkün olduğunca devam etmeye çalışır. Bu, tek bir çalıştırmada birden fazla hatanın tespit edilmesini sağlar.

### Özellikler

- **Analiz Devam Etme**: Bir hata bulunduğunda analizi durdurmak yerine devam etme
- **Eksik Sembol Kurtarma**: Tanımlanmamış sembollerle karşılaşıldığında varsayılan tipler atama
- **Tip Uyuşmazlığı Kurtarma**: Tip uyuşmazlıklarında en iyi eşleşmeyi bulma
- **Eksik Üye Kurtarma**: Eksik sınıf üyeleriyle karşılaşıldığında varsayılan değerler atama
- **Sözdizimi Hatası Kurtarma**: Sözdizimi hatalarından sonra analizi sürdürme

### Örnek Kod

```go
// Hata kurtarma örneği
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

### Uygulama Detayları

Hata kurtarma mekanizmaları, `internal/semantic/semantic.go` dosyasında uygulanmıştır. Mekanizmalar, aşağıdaki bileşenlerden oluşur:

- **Hata bayrakları**: Hata durumlarını izleyen bayraklar
- **Varsayılan tipler**: Eksik semboller için varsayılan tipler
- **Kurtarma stratejileri**: Farklı hata türleri için kurtarma stratejileri

```go
// analyzeStatement, bir ifadeyi analiz eder.
func (a *Analyzer) analyzeStatement(stmt ast.Statement) Type {
    // Hata kurtarma modunu etkinleştir
    defer func() {
        if r := recover(); r != nil {
            // Hatayı raporla
            a.reportError(token.Token{}, "Analiz sırasında panik: %v", r)
            // Varsayılan tip döndür
            return &BasicType{Name: "unknown", Kind: UNKNOWN_TYPE}
        }
    }()
    
    // Normal analiz
    switch s := stmt.(type) {
    case *ast.VarStatement:
        return a.analyzeVarStatement(s)
    // ... diğer ifade türleri
    }
}
```

## Sonuç

GO-Minus'un semantik analiz bileşeninde yapılan iyileştirmeler, daha güçlü bir derleme deneyimi ve daha iyi hata mesajları sağlar. Bu iyileştirmeler, programcıların hatalarını daha hızlı bulmasına ve düzeltmesine yardımcı olur.

Daha fazla bilgi ve örnekler için aşağıdaki belgelere bakabilirsiniz:

- [Gelişmiş Hata Raporlama Örneği](examples/semantic/error-reporting.md)
- [Tip Çıkarımı Örneği](examples/semantic/type-inference.md)
- [Semantik Analiz Belgelendirmesi](reference/semantic-analysis.md)

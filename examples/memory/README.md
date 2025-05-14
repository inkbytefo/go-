# GO-Minus Bellek Yönetimi Örnekleri

Bu dizin, GO-Minus programlama dili için Hibrit Akıllı Bellek Yönetimi Sistemi'nin kullanımını gösteren örnekleri içerir.

## Örnekler

### 1. Hibrit Bellek Yönetimi

[hybrid_memory_management.gom](hybrid_memory_management.gom) dosyası, GO-Minus'un hibrit bellek yönetimi özelliklerinin kullanımını gösterir:

- Otomatik bellek yönetimi (garbage collection)
- Manuel bellek yönetimi
- Bölgesel bellek yönetimi (region-based memory management)
- Bellek havuzu şablonları (memory pool templates)
- Profil tabanlı otomatik optimizasyon

```bash
gominus run hybrid_memory_management.gom
```

### 2. Yaşam Süresi Analizi

[lifetime_analysis.gom](lifetime_analysis.gom) dosyası, GO-Minus'un yaşam süresi analizi özelliğinin kullanımını gösterir:

- Bellek sızıntısı tespiti
- Dangling pointer tespiti
- Güvenli referans yönetimi
- Yaşam süresi kapsamları

```bash
gominus run lifetime_analysis.gom
```

## Çıktılar

### Hibrit Bellek Yönetimi Örneği

```
GO-Minus Hibrit Akıllı Bellek Yönetimi Sistemi Örneği
====================================================
Otomatik Bellek Yönetimi Örneği
Toplam: 523776

Manuel Bellek Yönetimi Örneği
Toplam: 523776

Bölgesel Bellek Yönetimi Örneği
Toplam: 523776
Bölge İstatistikleri:
  Toplam Boyut: 1048576 bayt
  Kullanılan Boyut: 4096 bayt
  Blok Sayısı: 1
  Ayırma İşlemi Sayısı: 1

Bellek Havuzu Örneği
Object 0 : 0
Object 1 : 1
...
Object 99 : 99
Toplam: 4950
Havuz İstatistikleri:
  Kapasite: 1000
  Kullanılabilir: 100
  Alınan Nesne Sayısı: 100
  Geri Döndürülen Nesne Sayısı: 100

Profil Tabanlı Otomatik Optimizasyon Örneği
İterasyon 0 Toplam: 523776
İterasyon 10 Toplam: 534016
...
İterasyon 90 Toplam: 624896
Bellek İstatistikleri:
  Toplam Ayrılan: 419430400 bayt
  Toplam Serbest Bırakılan: 419430400 bayt
  Şu Anda Kullanılan: 0 bayt
  En Yüksek Kullanım: 4096 bayt
  Ayırma İşlemi Sayısı: 100
  Serbest Bırakma İşlemi Sayısı: 100

Hibrit Bellek Yönetimi Örneği
Manuel Bellek Toplamı: 523776
Otomatik Bellek: Otomatik Bellek : 42
Bölgesel Bellek Toplamı: 523776
Havuz Nesneleri Toplamı: 4950
```

### Yaşam Süresi Analizi Örneği

```
GO-Minus Yaşam Süresi Analizi Örneği
====================================
Bellek Sızıntısı Örneği
Bellek Sızıntısı Sayısı: 1
Sızıntı 1: person (Person)

Dangling Pointer Örneği
Dangling Pointer Sayısı: 1
Dangling Pointer 1: outerPointer (unsafe.Pointer)

Güvenli Referans Örneği
Bellek Sızıntısı Sayısı: 0
Dangling Pointer Sayısı: 0

Yaşam Süresi Analizi Örneği
Bellek Sızıntısı Sayısı: 1
Sızıntı 1: person1 (Person)
Dangling Pointer Sayısı: 0
```

## Bellek Yönetimi Stratejileri

GO-Minus, aşağıdaki bellek yönetimi stratejilerini destekler:

### 1. Otomatik Bellek Yönetimi (Garbage Collection)

GO-Minus, varsayılan olarak Go'nun garbage collector'ünü kullanır. Bu, bellek yönetimini otomatikleştirerek programcıların bellek sızıntıları ve dangling pointer'lar gibi yaygın sorunlarla uğraşmak zorunda kalmadan kod yazmasına olanak tanır.

### 2. Manuel Bellek Yönetimi

Performans kritik bölümler için, GO-Minus manuel bellek yönetimi seçeneği sunar. Bu, `unsafe` bloğu içinde bellek ayırma ve serbest bırakma işlemlerini doğrudan kontrol etmenize olanak tanır.

### 3. Bölgesel Bellek Yönetimi (Region-Based Memory Management)

GO-Minus'un yeni özelliği olan bölgesel bellek yönetimi, programcılara belirli kod bloklarını "bellek bölgeleri" olarak işaretleme ve bu bölgelerdeki tüm bellek ayırma işlemlerini bölge sonunda otomatik olarak serbest bırakma olanağı sunar.

### 4. Yaşam Süresi Analizi (Lifetime Analysis)

GO-Minus, Rust'tan esinlenen, ancak daha az katı bir yaşam süresi analizi sistemi sunar. Derleyici, değişkenlerin yaşam sürelerini analiz eder ve potansiyel bellek sızıntılarını veya dangling pointer'ları tespit eder.

### 5. Profil Tabanlı Otomatik Optimizasyon

GO-Minus, uygulama çalışırken bellek kullanım desenlerini analiz eden ve gelecekteki çalıştırmalarda bellek yönetimini otomatik olarak optimize eden bir sistem sunar.

### 6. Bellek Havuzu Şablonları

GO-Minus, belirli veri yapıları için özelleştirilmiş bellek havuzları oluşturmayı kolaylaştıran şablonlar sunar.

## Lisans

GO-Minus Bellek Yönetimi Örnekleri, GO-Minus projesi ile aynı lisans altında dağıtılmaktadır.

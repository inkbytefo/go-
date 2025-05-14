# GO-Minus Geliştirme Araçları

Bu belge, GO-Minus programlama dili ile geliştirme yaparken kullanabileceğiniz araçları açıklamaktadır. Bu araçlar, GO-Minus kodunu derleme, test etme, belgelendirme, biçimlendirme ve paket yönetimi gibi işlemleri gerçekleştirmenize yardımcı olur.

## İçindekiler

1. [Derleyici (gominus)](#derleyici-gominus)
2. [Paket Yöneticisi (gompm)](#paket-yöneticisi-gompm)
3. [Test Aracı (gomtest)](#test-aracı-gomtest)
4. [Belgelendirme Aracı (gomdoc)](#belgelendirme-aracı-gomdoc)
5. [Kod Biçimlendirme Aracı (gomfmt)](#kod-biçimlendirme-aracı-gomfmt)
6. [Dil Sunucusu (gomlsp)](#dil-sunucusu-gomlsp)
7. [Hata Ayıklama Aracı (gomdebug)](#hata-ayıklama-aracı-gomdebug)
8. [IDE Entegrasyonu](#ide-entegrasyonu)

## Derleyici (gominus)

GO-Minus derleyicisi, GO-Minus kaynak kodunu derleyerek çalıştırılabilir dosyalar oluşturur.

### Kullanım

```bash
# Tek bir dosyayı derle
gominus dosya.gom

# Birden fazla dosyayı derle
gominus dosya1.gom dosya2.gom

# Çıktı dosyasını belirt
gominus -o program dosya.gom

# Optimizasyon seviyesini belirt
gominus -O2 dosya.gom

# Hata ayıklama bilgisi ekle
gominus -g dosya.gom

# Uyarıları etkinleştir
gominus -Wall dosya.gom

# Belirli bir hedef platform için derle
gominus -target=x86_64-linux dosya.gom

# Kütüphane olarak derle
gominus -lib dosya.gom

# Derleme ve çalıştırma
gominus -run dosya.gom
```

### Çıktı Formatları

GO-Minus derleyicisi, aşağıdaki çıktı formatlarını destekler:

- Çalıştırılabilir dosya (varsayılan)
- Nesne dosyası (`.o`)
- LLVM IR (`.ll`)
- Kütüphane (`.a`, `.so`, `.dll`)

## Paket Yöneticisi (gompm)

GO-Minus Paket Yöneticisi (gompm), GO-Minus paketlerini indirme, kurma, kaldırma, güncelleme ve arama işlemlerini gerçekleştirir. Bu araç, GO-Minus ekosistemindeki paketlerin yönetimini kolaylaştırır ve bağımlılık çözümleme işlemlerini otomatikleştirir.

### Özellikler

- Yeni projeler oluşturma
- Paketleri kurma ve kaldırma
- Paketleri güncelleme
- Kurulu paketleri listeleme
- Paket deposunda arama yapma
- Paket bilgilerini görüntüleme
- Bağımlılık yönetimi
- Geliştirme bağımlılıkları desteği
- Sürüm kısıtlamaları
- Paket deposu entegrasyonu

### Kullanım

```bash
# Yeni bir proje başlat
gompm init [proje-adı]

# Paket kur
gompm install paket-adı

# Belirli bir sürümü kur
gompm install paket-adı@1.0.0

# Geliştirme bağımlılığı olarak kur
gompm install paket-adı --dev

# Paket kaldır
gompm remove paket-adı

# Tüm paketleri güncelle
gompm update

# Kurulu paketleri listele
gompm list

# Paket deposunda ara
gompm search sorgu

# Paket bilgisi göster
gompm info paket-adı

# Yardım mesajını göster
gompm help

# Sürüm bilgisini göster
gompm version
```

### Paket Yapılandırma Dosyası

Paket yapılandırma dosyası (`gompm.json`), projenin meta verilerini ve bağımlılıklarını içerir:

```json
{
  "name": "proje-adı",
  "version": "1.0.0",
  "description": "Proje açıklaması",
  "author": "Yazar Adı",
  "license": "MIT",
  "dependencies": {
    "paket1": "1.0.0",
    "paket2": "^2.0.0"
  },
  "devDependencies": {
    "test-paketi": "1.0.0"
  },
  "keywords": ["anahtar", "kelime"],
  "homepage": "https://example.com",
  "repository": "https://github.com/kullanici/proje"
}
```

### Sürüm Belirtme

Paket sürümleri, [Semantic Versioning](https://semver.org/) kurallarına göre belirtilir. Sürüm belirtme formatları:

- `1.0.0`: Tam olarak 1.0.0 sürümü
- `^1.0.0`: 1.0.0 veya daha yüksek, ancak 2.0.0'dan düşük
- `~1.0.0`: 1.0.0 veya daha yüksek, ancak 1.1.0'dan düşük
- `>=1.0.0`: 1.0.0 veya daha yüksek
- `<=1.0.0`: 1.0.0 veya daha düşük
- `1.0.0 - 2.0.0`: 1.0.0 ile 2.0.0 arasında (her ikisi de dahil)
- `latest`: En son sürüm

### Paket Deposu

GO-Minus paketleri, merkezi bir paket deposunda (https://repo.gominus.org) saklanır. Bu depo, paketlerin meta verilerini ve kaynak kodlarını içerir. gompm, bu depodan paketleri indirir ve kurar.

### Bağımlılık Çözümleme

gompm, paketlerin bağımlılıklarını otomatik olarak çözümler ve kurar. Bağımlılık çözümleme algoritması, sürüm kısıtlamalarını dikkate alarak en uygun sürümleri seçer. Bu, paketler arasındaki sürüm çakışmalarını önlemeye yardımcı olur.

### Güvenlik

gompm, paketlerin bütünlüğünü doğrulamak için dijital imzalar kullanır. Bu, kötü amaçlı paketlerin kurulmasını önlemeye yardımcı olur. Ayrıca, paketlerin güvenlik açıklarını kontrol etmek için bir güvenlik taraması da gerçekleştirir.

### Örnek: Yeni Bir Proje Oluşturma

```bash
# Yeni bir proje oluştur
gompm init my-project

# Proje dizinine git
cd my-project

# Paketleri kur
gompm install logger@1.0.0
gompm install http-client@^2.0.0
gompm install test-framework@latest --dev

# Kurulu paketleri listele
gompm list
```

### Örnek: Paket Bilgisi Görüntüleme

```bash
# Paket bilgisini görüntüle
gompm info http-client

# Çıktı:
# Paket: http-client
# Sürüm: 2.1.0
# Açıklama: HTTP istemci kütüphanesi
# Yazar: GO-Minus Ekibi
# Lisans: MIT
# Bağımlılıklar:
#   - logger@1.0.0
#   - url-parser@^1.5.0
# Anahtar Kelimeler: http, client, request, response
# Ana Sayfa: https://example.com/http-client
# Depo: https://github.com/gominus/http-client
```

Daha fazla bilgi için [Paket Yöneticisi Belgelendirmesi](../cmd/gompm/README.md) belgesine bakın.

## Test Aracı (gomtest)

GO-Minus Test Aracı (gomtest), GO-Minus kodunu test etmek için kullanılır. Test dosyalarını bulur, testleri çalıştırır ve sonuçları raporlar.

### Kullanım

```bash
# Mevcut dizindeki testleri çalıştır
gomtest

# Ayrıntılı çıktı ile testleri çalıştır
gomtest -v

# Alt dizinlerdeki testleri de çalıştır
gomtest -r

# Belirli bir test desenine uyan testleri çalıştır
gomtest -pattern=TestAdd

# Benchmark testlerini çalıştır
gomtest -bench

# Kod kapsama analizi yap
gomtest -cover

# Belirtilen dizinlerdeki testleri çalıştır
gomtest ./pkg ./internal
```

### Test Dosyaları

Test dosyaları, `_test.gom` son ekiyle biten dosyalardır. Test fonksiyonları, `Test` önekiyle başlar ve `*testing.T` parametresi alır:

```go
// math_test.gom
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; beklenen 5", result)
    }
}
```

## Belgelendirme Aracı (gomdoc)

GO-Minus Belgelendirme Aracı (gomdoc), GO-Minus kodunu belgelendirmek için kullanılır. Kod içindeki yorumları ve yapıları analiz ederek belgelendirme oluşturur.

### Kullanım

```bash
# Mevcut dizindeki paketi belgele
gomdoc

# HTML formatında belgelendirme oluştur
gomdoc -html

# Markdown formatında belgelendirme oluştur
gomdoc -markdown

# Belgelendirmeyi belirli bir dizine kaydet
gomdoc -output=docs

# Belgelendirme sunucusu başlat
gomdoc -server

# Belgelendirme sunucusunu belirli bir portta başlat
gomdoc -server -port=8080

# Belirtilen paketleri belgele
gomdoc ./pkg ./internal
```

### Belgelendirme Yorumları

GO-Minus belgelendirme yorumları, `//` veya `/* */` ile başlar ve belgelendirilen öğeden hemen önce yer alır:

```go
// Add, iki sayıyı toplar ve sonucu döndürür.
//
// Parametreler:
//   - a: İlk sayı
//   - b: İkinci sayı
//
// Dönüş değeri:
//   - İki sayının toplamı
func Add(a, b int) int {
    return a + b
}
```

## Kod Biçimlendirme Aracı (gomfmt)

GO-Minus Kod Biçimlendirme Aracı (gomfmt), GO-Minus kodunu standart bir biçimde düzenlemek için kullanılır. Kod stilini tutarlı hale getirir.

### Kullanım

```bash
# Mevcut dizindeki GO-Minus dosyalarını biçimlendir
gomfmt

# Değişiklikleri dosyalara yaz
gomfmt -w

# Değişiklikleri diff formatında göster
gomfmt -d

# Biçimlendirilmesi gereken dosyaları listele
gomfmt -l

# Alt dizinlerdeki dosyaları da biçimlendir
gomfmt -r

# Kodu basitleştir
gomfmt -s

# Belirtilen dosyaları biçimlendir
gomfmt dosya1.gom dosya2.gom

# Belirtilen dizinlerdeki dosyaları biçimlendir
gomfmt ./pkg ./internal
```

## Dil Sunucusu (gomlsp)

GO-Minus Dil Sunucusu (gomlsp), GO-Minus için Language Server Protocol (LSP) implementasyonu sağlar. Bu, IDE ve metin düzenleyicileri ile entegrasyon için kullanılır.

### Özellikler

- Kod tamamlama
- Hata ve uyarı gösterimi
- Tanıma gitme
- Yeniden adlandırma
- Kod biçimlendirme
- Kod katlama
- Sembol arama
- Referans bulma
- Kod eylemleri

### Kullanım

```bash
# Dil sunucusunu başlat
gomlsp

# Belirli bir portta başlat
gomlsp --port=8080

# Ayrıntılı günlük kaydı etkinleştir
gomlsp --verbose
```

## Hata Ayıklama Aracı (gomdebug)

GO-Minus Hata Ayıklama Aracı (gomdebug), GO-Minus programlarında hata ayıklamak için kullanılır.

### Kullanım

```bash
# Programı hata ayıklama modunda başlat
gomdebug program

# Belirli argümanlarla programı başlat
gomdebug program arg1 arg2

# Belirli bir dosya ve satırda kesme noktası ekle
gomdebug --break=dosya.gom:10 program

# Hata ayıklama sunucusunu başlat
gomdebug --server program

# Belirli bir portta hata ayıklama sunucusunu başlat
gomdebug --server --port=8080 program
```

### Hata Ayıklama Komutları

Hata ayıklama oturumu sırasında aşağıdaki komutları kullanabilirsiniz:

- `break dosya:satır`: Kesme noktası ekle
- `continue`: Çalıştırmaya devam et
- `step`: Bir satır ilerle (fonksiyonlara gir)
- `next`: Bir satır ilerle (fonksiyonları atla)
- `print ifade`: İfadeyi değerlendir ve yazdır
- `backtrace`: Çağrı yığınını göster
- `frame n`: n. çağrı çerçevesine geç
- `list`: Kaynak kodu göster
- `quit`: Hata ayıklayıcıdan çık

## IDE Entegrasyonu

GO-Minus, çeşitli IDE ve metin düzenleyicileri ile entegre edilebilir.

### Visual Studio Code

VS Code için GO-Minus eklentisi, aşağıdaki özellikleri sağlar:

- Sözdizimi vurgulama
- Kod tamamlama
- Hata ve uyarı gösterimi
- Tanıma gitme
- Yeniden adlandırma
- Kod biçimlendirme
- Kod katlama
- Sembol arama
- Referans bulma
- Kod eylemleri

### JetBrains IDE'leri

JetBrains IDE'leri (IntelliJ IDEA, GoLand, vb.) için GO-Minus eklentisi, aşağıdaki özellikleri sağlar:

- Sözdizimi vurgulama
- Kod tamamlama
- Hata ve uyarı gösterimi
- Tanıma gitme
- Yeniden adlandırma
- Kod biçimlendirme
- Kod katlama
- Sembol arama
- Referans bulma
- Kod eylemleri

### Vim/Neovim

Vim/Neovim için GO-Minus eklentisi, aşağıdaki özellikleri sağlar:

- Sözdizimi vurgulama
- Kod tamamlama (coc.nvim veya YouCompleteMe ile)
- Hata ve uyarı gösterimi
- Tanıma gitme
- Kod biçimlendirme

### Emacs

Emacs için GO-Minus modu, aşağıdaki özellikleri sağlar:

- Sözdizimi vurgulama
- Kod tamamlama (company-mode ile)
- Hata ve uyarı gösterimi (flycheck ile)
- Tanıma gitme
- Kod biçimlendirme

# GO-Minus Paket Yöneticisi (gompm)

GO-Minus Paket Yöneticisi (gompm), GO-Minus programlama dili için paket yönetimi sağlayan bir araçtır. Bu araç, paketleri indirme, kurma, kaldırma, güncelleme ve arama işlemlerini gerçekleştirir.

## Özellikler

- Yeni projeler oluşturma
- Paketleri kurma ve kaldırma
- Paketleri güncelleme
- Kurulu paketleri listeleme
- Paket deposunda arama yapma
- Paket bilgilerini görüntüleme
- Bağımlılık yönetimi
- Geliştirme bağımlılıkları desteği

## Kurulum

```bash
# GO-Minus derleyicisini kullanarak gompm'yi derle
gominus build -o gompm cmd/gompm/main.gom

# Derlenmiş dosyayı PATH içindeki bir dizine kopyala
cp gompm /usr/local/bin/
```

## Kullanım

### Yeni Proje Oluşturma

```bash
# Mevcut dizinde yeni bir proje oluştur
gompm init

# Belirli bir adla yeni bir proje oluştur
gompm init proje-adi
```

Bu komut, bir `gompm.json` dosyası oluşturur. Bu dosya, projenin adı, sürümü, açıklaması, bağımlılıkları ve diğer meta verileri içerir.

### Paket Kurma

```bash
# Paket kur
gompm install paket-adi

# Belirli bir sürümü kur
gompm install paket-adi@1.0.0

# Geliştirme bağımlılığı olarak kur
gompm install paket-adi --dev
```

Bu komut, belirtilen paketi ve bağımlılıklarını indirir ve kurar. Paket, `gompm.json` dosyasındaki bağımlılıklar listesine eklenir.

### Paket Kaldırma

```bash
# Paket kaldır
gompm remove paket-adi
```

Bu komut, belirtilen paketi kaldırır ve `gompm.json` dosyasındaki bağımlılıklar listesinden çıkarır.

### Paketleri Güncelleme

```bash
# Tüm paketleri güncelle
gompm update
```

Bu komut, tüm paketleri en son sürümlerine günceller.

### Kurulu Paketleri Listeleme

```bash
# Kurulu paketleri listele
gompm list
```

Bu komut, projede kurulu olan tüm paketleri ve sürümlerini listeler.

### Paket Arama

```bash
# Paket deposunda arama yap
gompm search sorgu
```

Bu komut, paket deposunda belirtilen sorguya göre arama yapar ve sonuçları listeler.

### Paket Bilgisi Görüntüleme

```bash
# Paket bilgisini görüntüle
gompm info paket-adi
```

Bu komut, belirtilen paketin adı, sürümü, açıklaması, yazarı, lisansı, bağımlılıkları ve diğer meta verilerini görüntüler.

### Yardım ve Sürüm Bilgisi

```bash
# Yardım mesajını görüntüle
gompm help

# Sürüm bilgisini görüntüle
gompm version
```

## Paket Yapılandırma Dosyası (gompm.json)

Paket yapılandırma dosyası, projenin meta verilerini ve bağımlılıklarını içerir. Örnek bir `gompm.json` dosyası:

```json
{
  "name": "proje-adi",
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

## Sürüm Belirtme

Paket sürümleri, [Semantic Versioning](https://semver.org/) kurallarına göre belirtilir. Sürüm belirtme formatları:

- `1.0.0`: Tam olarak 1.0.0 sürümü
- `^1.0.0`: 1.0.0 veya daha yüksek, ancak 2.0.0'dan düşük
- `~1.0.0`: 1.0.0 veya daha yüksek, ancak 1.1.0'dan düşük
- `>=1.0.0`: 1.0.0 veya daha yüksek
- `<=1.0.0`: 1.0.0 veya daha düşük
- `1.0.0 - 2.0.0`: 1.0.0 ile 2.0.0 arasında (her ikisi de dahil)
- `latest`: En son sürüm

## Paket Deposu

GO-Minus paketleri, merkezi bir paket deposunda (https://repo.gominus.org) saklanır. Bu depo, paketlerin meta verilerini ve kaynak kodlarını içerir.

## Bağımlılık Çözümleme

gompm, paketlerin bağımlılıklarını otomatik olarak çözümler ve kurar. Bağımlılık çözümleme algoritması, sürüm kısıtlamalarını dikkate alarak en uygun sürümleri seçer.

## Güvenlik

gompm, paketlerin bütünlüğünü doğrulamak için dijital imzalar kullanır. Bu, kötü amaçlı paketlerin kurulmasını önlemeye yardımcı olur.

## Katkıda Bulunma

Katkıda bulunmak için, GitHub deposunu (https://github.com/gominus/gompm) çatallayın, değişikliklerinizi yapın ve bir pull isteği gönderin.

## Lisans

GO-Minus Paket Yöneticisi, MIT lisansı altında lisanslanmıştır.

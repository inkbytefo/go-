# GO-Minus Geliştirme Araçları

Bu dizin, GO-Minus programlama dili için geliştirme araçlarını içerir. Bu araçlar, GO-Minus ile geliştirme yaparken kullanılabilecek çeşitli yardımcı programları içerir.

## Araçlar

### gompm - GO-Minus Paket Yöneticisi

GO-Minus paketlerini yönetmek için kullanılan bir araçtır. Paket yükleme, kaldırma, güncelleme ve arama işlemlerini gerçekleştirir.

```bash
# Yeni bir GO-Minus projesi başlat
gompm -init myproject

# Paket yükle
gompm -install fmt strings

# Paket kaldır
gompm -uninstall fmt

# Paketleri güncelle
gompm -update

# Yüklü paketleri listele
gompm -list

# Paket ara
gompm -search json
```

### gomtest - GO-Minus Test Aracı

GO-Minus kodunu test etmek için kullanılan bir araçtır. Test dosyalarını bulur, testleri çalıştırır ve sonuçları raporlar.

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

### gomdoc - GO-Minus Belgelendirme Aracı

GO-Minus kodunu belgelendirmek için kullanılan bir araçtır. Kod içindeki yorumları ve yapıları analiz ederek belgelendirme oluşturur.

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

### gomfmt - GO-Minus Kod Biçimlendirme Aracı

GO-Minus kodunu standart bir biçimde düzenlemek için kullanılan bir araçtır. Kod stilini tutarlı hale getirir.

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
gomfmt file1.gom file2.gom

# Belirtilen dizinlerdeki dosyaları biçimlendir
gomfmt ./pkg ./internal
```

## Kurulum

GO-Minus geliştirme araçlarını kurmak için, GO-Minus derleyicisini kurduktan sonra aşağıdaki komutları çalıştırabilirsiniz:

```bash
# GO-Minus paket yöneticisini kur
go build -o gompm ./tools/gompm

# GO-Minus test aracını kur
go build -o gomtest ./tools/gomtest

# GO-Minus belgelendirme aracını kur
go build -o gomdoc ./tools/gomdoc

# GO-Minus kod biçimlendirme aracını kur
go build -o gomfmt ./tools/gomfmt
```

Oluşturulan çalıştırılabilir dosyaları PATH ortam değişkeninizin bulunduğu bir dizine kopyalayabilirsiniz.
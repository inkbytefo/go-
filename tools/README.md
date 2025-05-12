# GO+ Geliştirme Araçları

Bu dizin, GO+ programlama dili için geliştirme araçlarını içerir. Bu araçlar, GO+ ile geliştirme yaparken kullanılabilecek çeşitli yardımcı programları içerir.

## Araçlar

### goppm - GO+ Paket Yöneticisi

GO+ paketlerini yönetmek için kullanılan bir araçtır. Paket yükleme, kaldırma, güncelleme ve arama işlemlerini gerçekleştirir.

```bash
# Yeni bir GO+ projesi başlat
goppm -init myproject

# Paket yükle
goppm -install fmt strings

# Paket kaldır
goppm -uninstall fmt

# Paketleri güncelle
goppm -update

# Yüklü paketleri listele
goppm -list

# Paket ara
goppm -search json
```

### goptest - GO+ Test Aracı

GO+ kodunu test etmek için kullanılan bir araçtır. Test dosyalarını bulur, testleri çalıştırır ve sonuçları raporlar.

```bash
# Mevcut dizindeki testleri çalıştır
goptest

# Ayrıntılı çıktı ile testleri çalıştır
goptest -v

# Alt dizinlerdeki testleri de çalıştır
goptest -r

# Belirli bir test desenine uyan testleri çalıştır
goptest -pattern=TestAdd

# Benchmark testlerini çalıştır
goptest -bench

# Kod kapsama analizi yap
goptest -cover

# Belirtilen dizinlerdeki testleri çalıştır
goptest ./pkg ./internal
```

### gopdoc - GO+ Belgelendirme Aracı

GO+ kodunu belgelendirmek için kullanılan bir araçtır. Kod içindeki yorumları ve yapıları analiz ederek belgelendirme oluşturur.

```bash
# Mevcut dizindeki paketi belgele
gopdoc

# HTML formatında belgelendirme oluştur
gopdoc -html

# Markdown formatında belgelendirme oluştur
gopdoc -markdown

# Belgelendirmeyi belirli bir dizine kaydet
gopdoc -output=docs

# Belgelendirme sunucusu başlat
gopdoc -server

# Belgelendirme sunucusunu belirli bir portta başlat
gopdoc -server -port=8080

# Belirtilen paketleri belgele
gopdoc ./pkg ./internal
```

### gopfmt - GO+ Kod Biçimlendirme Aracı

GO+ kodunu standart bir biçimde düzenlemek için kullanılan bir araçtır. Kod stilini tutarlı hale getirir.

```bash
# Mevcut dizindeki GO+ dosyalarını biçimlendir
gopfmt

# Değişiklikleri dosyalara yaz
gopfmt -w

# Değişiklikleri diff formatında göster
gopfmt -d

# Biçimlendirilmesi gereken dosyaları listele
gopfmt -l

# Alt dizinlerdeki dosyaları da biçimlendir
gopfmt -r

# Kodu basitleştir
gopfmt -s

# Belirtilen dosyaları biçimlendir
gopfmt file1.gop file2.gop

# Belirtilen dizinlerdeki dosyaları biçimlendir
gopfmt ./pkg ./internal
```

## Kurulum

GO+ geliştirme araçlarını kurmak için, GO+ derleyicisini kurduktan sonra aşağıdaki komutları çalıştırabilirsiniz:

```bash
# GO+ paket yöneticisini kur
go build -o goppm ./tools/goppm

# GO+ test aracını kur
go build -o goptest ./tools/goptest

# GO+ belgelendirme aracını kur
go build -o gopdoc ./tools/gopdoc

# GO+ kod biçimlendirme aracını kur
go build -o gopfmt ./tools/gopfmt
```

Oluşturulan çalıştırılabilir dosyaları PATH ortam değişkeninizin bulunduğu bir dizine kopyalayabilirsiniz.
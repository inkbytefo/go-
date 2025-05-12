# GO-Minus JetBrains Eklentisi

GO-Minus JetBrains Eklentisi, GO-Minus programlama dili için IntelliJ IDEA, GoLand, PyCharm, WebStorm, PhpStorm, Rider ve diğer JetBrains IDE'leri için destek sağlar. Bu eklenti, sözdizimi vurgulama, kod tamamlama, hata işaretleme, tanıma gitme, hata ayıklama gibi özellikler sunar.

## Özellikler

- Sözdizimi vurgulama
- Kod tamamlama
- Hata işaretleme
- Tanıma gitme
- Referansları bulma
- Kod biçimlendirme
- Hata ayıklama
- Testleri çalıştırma
- Kod kapsama analizi
- Yeniden düzenleme araçları
- Canlı şablonlar
- Kod analizi

## Kurulum

### Ön Koşullar

- GO-Minus derleyicisi ve araçları yüklü olmalıdır.
- `gomlsp` (GO-Minus Dil Sunucusu) ve `gomdebug` (GO-Minus Hata Ayıklama Aracı) PATH ortam değişkeninizde bulunmalıdır.

### JetBrains Marketplace'den Kurulum

1. JetBrains IDE'nizi açın
2. "File" > "Settings" > "Plugins" menüsünü açın
3. "Marketplace" sekmesine tıklayın
4. "GO+" araması yapın
5. "GO+ Language Support" eklentisini bulun ve "Install" düğmesine tıklayın
6. IDE'yi yeniden başlatın

### ZIP Dosyasından Kurulum

1. Eklentiyi derleyin veya indirin
2. JetBrains IDE'nizi açın
3. "File" > "Settings" > "Plugins" menüsünü açın
4. Dişli simgesine tıklayın ve "Install Plugin from Disk..." seçeneğini seçin
5. ZIP dosyasını seçin ve "OK" düğmesine tıklayın
6. IDE'yi yeniden başlatın

## Kullanım

### Sözdizimi Vurgulama

GO+ dosyaları (`.gop` uzantılı) otomatik olarak sözdizimi vurgulaması ile açılır.

### Kod Tamamlama

Kod yazarken, otomatik tamamlama önerileri görünecektir. Önerileri kabul etmek için `Tab` veya `Enter` tuşuna basın.

### Hata İşaretleme

Kod yazarken, sözdizimi ve semantik hatalar otomatik olarak işaretlenir.

### Tanıma Gitme

Bir sembolün tanımına gitmek için, sembolün üzerine `Ctrl+Click` (veya `Cmd+Click`) yapın veya sembolün üzerindeyken `Ctrl+B` (veya `Cmd+B`) tuşlarına basın.

### Kod Biçimlendirme

Bir belgeyi biçimlendirmek için, `Ctrl+Alt+L` (veya `Cmd+Alt+L`) tuşlarına basın.

### Hata Ayıklama

1. Hata ayıklamak istediğiniz GO+ dosyasını açın
2. Çalıştırma yapılandırması oluşturun ("Run" > "Edit Configurations...")
3. "+" düğmesine tıklayın ve "GO+" seçeneğini seçin
4. Yapılandırmayı adlandırın ve "OK" düğmesine tıklayın
5. Kesme noktaları ayarlamak için, satır numarasının solundaki boşluğa tıklayın
6. "Run" > "Debug" menüsünü seçin veya `Shift+F9` tuşlarına basın
7. Hata ayıklama araç çubuğunu kullanarak programı adım adım çalıştırın

### Testleri Çalıştırma

1. Test dosyasını açın (`*_test.gop`)
2. Test fonksiyonunun yanındaki yeşil "Run" düğmesine tıklayın veya sağ tıklayıp "Run" seçeneğini seçin

## Yapılandırma

Eklentiyi yapılandırmak için, "File" > "Settings" > "Languages & Frameworks" > "GO+" menüsünü açın.

### Dil Sunucusu Yolu

Dil sunucusu yolunu belirtin:

```
goplsp
```

### Hata Ayıklama Aracı Yolu

Hata ayıklama aracı yolunu belirtin:

```
gopdebug
```

### Kod Stili

GO+ kod stilini özelleştirin:

- Girinti boyutu
- Sekme kullanımı
- Satır sonu stili
- İçe aktarma düzeni
- Boşluk kullanımı

### Denetleyici Seçenekleri

Denetleyici seçeneklerini özelleştirin:

- Denetleyici seviyesi
- Özel denetleyici kuralları
- Denetleyici yapılandırma dosyası

## Sorun Giderme

### Dil Sunucusu Başlatılamıyor

1. `goplsp` komutunun PATH ortam değişkeninizde olduğunu kontrol edin
2. IDE'yi yeniden başlatın
3. Dil sunucusu yapılandırmasını kontrol edin

### Hata Ayıklama Aracı Başlatılamıyor

1. `gopdebug` komutunun PATH ortam değişkeninizde olduğunu kontrol edin
2. IDE'yi yeniden başlatın
3. Hata ayıklama yapılandırmasını kontrol edin

## Katkıda Bulunma

GO+ JetBrains Eklentisi, açık kaynaklı bir projedir. Katkıda bulunmak için, lütfen [katkı sağlama rehberini](../../CONTRIBUTING.md) okuyun.

## Lisans

GO+ JetBrains Eklentisi, GO+ projesi ile aynı lisans altında dağıtılmaktadır.
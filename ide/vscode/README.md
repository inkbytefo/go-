# GO-Minus VS Code Eklentisi

GO-Minus VS Code Eklentisi, GO-Minus programlama dili için Visual Studio Code desteği sağlar. Bu eklenti, sözdizimi vurgulama, kod tamamlama, hata işaretleme, tanıma gitme, hata ayıklama gibi özellikler sunar.

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

## Kurulum

### Ön Koşullar

- GO-Minus derleyicisi ve araçları yüklü olmalıdır.
- `gomlsp` (GO-Minus Dil Sunucusu) ve `gomdebug` (GO-Minus Hata Ayıklama Aracı) PATH ortam değişkeninizde bulunmalıdır.

### VS Code Marketplace'den Kurulum

1. VS Code'u açın
2. Uzantılar görünümünü açın (`Ctrl+Shift+X` veya `Cmd+Shift+X`)
3. "GO-Minus" araması yapın
4. "GO-Minus Language Support" eklentisini bulun ve "Install" düğmesine tıklayın

### VSIX Dosyasından Kurulum

1. Eklentiyi derleyin veya indirin
2. VS Code'u açın
3. Uzantılar görünümünü açın (`Ctrl+Shift+X` veya `Cmd+Shift+X`)
4. "..." menüsüne tıklayın ve "Install from VSIX..." seçeneğini seçin
5. VSIX dosyasını seçin ve "Install" düğmesine tıklayın

## Kullanım

### Sözdizimi Vurgulama

GO-Minus dosyaları (`.gom` uzantılı) otomatik olarak sözdizimi vurgulaması ile açılır.

### Kod Tamamlama

Kod yazarken, otomatik tamamlama önerileri görünecektir. Önerileri kabul etmek için `Tab` veya `Enter` tuşuna basın.

### Hata İşaretleme

Kod yazarken, sözdizimi ve semantik hatalar otomatik olarak işaretlenir.

### Tanıma Gitme

Bir sembolün tanımına gitmek için, sembolün üzerine `Ctrl+Click` (veya `Cmd+Click`) yapın veya sembolün üzerindeyken `F12` tuşuna basın.

### Kod Biçimlendirme

Bir belgeyi biçimlendirmek için, `Shift+Alt+F` tuşlarına basın veya sağ tıklayıp "Format Document" seçeneğini seçin.

### Hata Ayıklama

1. Hata ayıklamak istediğiniz GO-Minus dosyasını açın
2. `F5` tuşuna basın veya "Run" menüsünden "Start Debugging" seçeneğini seçin
3. Hata ayıklama yapılandırmasını seçin (ilk kez çalıştırıyorsanız, "GO-Minus Programını Çalıştır" seçeneğini seçin)
4. Kesme noktaları ayarlamak için, satır numarasının solundaki boşluğa tıklayın
5. Hata ayıklama araç çubuğunu kullanarak programı adım adım çalıştırın

### Testleri Çalıştırma

1. Test dosyasını açın (`*_test.gom`)
2. Sağ tıklayın ve "GO-Minus: Testleri Çalıştır" seçeneğini seçin

## Yapılandırma

Eklentiyi yapılandırmak için, VS Code ayarlarını açın (`Ctrl+,` veya `Cmd+,`) ve "GO-Minus" araması yapın.

### Dil Sunucusu Yolu

```json
"gominus.languageServerPath": "gomlsp"
```

### Hata Ayıklama Aracı Yolu

```json
"gominus.debuggerPath": "gomdebug"
```

### Kaydetme Sırasında Biçimlendirme

```json
"gominus.formatOnSave": true
```

### Kaydetme Sırasında Denetleme

```json
"gominus.lintOnSave": true
```

### Kaydetme Sırasında Test Çalıştırma

```json
"gominus.testOnSave": false
```

## Sorun Giderme

### Dil Sunucusu Başlatılamıyor

1. `gomlsp` komutunun PATH ortam değişkeninizde olduğunu kontrol edin
2. VS Code'u yeniden başlatın
3. Dil sunucusunu manuel olarak başlatmak için, komut paletini açın (`Ctrl+Shift+P` veya `Cmd+Shift+P`) ve "GO-Minus: Dil Sunucusunu Başlat" komutunu çalıştırın

### Hata Ayıklama Aracı Başlatılamıyor

1. `gomdebug` komutunun PATH ortam değişkeninizde olduğunu kontrol edin
2. VS Code'u yeniden başlatın
3. Hata ayıklama yapılandırmasını kontrol edin

## Katkıda Bulunma

GO-Minus VS Code Eklentisi, açık kaynaklı bir projedir. Katkıda bulunmak için, lütfen [katkı sağlama rehberini](../../CONTRIBUTING.md) okuyun.

## Lisans

GO-Minus VS Code Eklentisi, GO-Minus projesi ile aynı lisans altında dağıtılmaktadır.
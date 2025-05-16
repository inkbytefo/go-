# GO-Minus Dil Sunucusu (gomlsp)

GO-Minus Dil Sunucusu, GO-Minus programlama dili için Language Server Protocol (LSP) implementasyonudur. Bu sunucu, çeşitli IDE ve metin düzenleyicileri ile entegrasyon sağlayarak kod tamamlama, hata işaretleme, tanıma gitme, yeniden adlandırma gibi akıllı düzenleme özelliklerini destekler.

## Özellikler

- Kod tamamlama
- Sözdizimi hata işaretleme
- Semantik hata işaretleme
- Tanıma gitme
- Referansları bulma
- Belge sembolleri
- Çalışma alanı sembolleri
- Kod biçimlendirme
- Fare üzerinde bilgi gösterme (hover)
- Yeniden adlandırma

## Kurulum

GO-Minus Dil Sunucusu'nu kurmak için, GO-Minus derleyicisini kurduktan sonra aşağıdaki komutu çalıştırabilirsiniz:

```bash
go build -o gomlsp ./tools/gomlsp
```

Oluşturulan çalıştırılabilir dosyayı PATH ortam değişkeninizin bulunduğu bir dizine kopyalayabilirsiniz.

## Kullanım

GO+ Dil Sunucusu, standart giriş/çıkış veya TCP üzerinden çalışabilir:

```bash
# Standart giriş/çıkış üzerinden çalıştır
goplsp

# TCP sunucu olarak çalıştır (varsayılan port: 8080)
goplsp -mode=tcp

# TCP sunucu olarak belirtilen portta çalıştır
goplsp -mode=tcp -addr=:9090

# Log dosyasına yaz
goplsp -log=goplsp.log
```

## IDE Entegrasyonu

GO+ Dil Sunucusu, Language Server Protocol'ü destekleyen herhangi bir IDE veya metin düzenleyicisi ile kullanılabilir. Aşağıdaki IDE'ler için özel eklentiler mevcuttur:

- Visual Studio Code: [GO+ VS Code Eklentisi](../ide/vscode/README.md)
- JetBrains IDEs: [GO+ JetBrains Eklentisi](../ide/jetbrains/README.md)
- Vim/Neovim: [GO+ Vim Eklentisi](../ide/vim/README.md)
- Emacs: [GO+ Emacs Eklentisi](../ide/emacs/README.md)

## Geliştirme

GO+ Dil Sunucusu, GO+ dilinin gelişimiyle birlikte sürekli olarak geliştirilmektedir. Katkıda bulunmak için, lütfen [katkı sağlama rehberini](../../CONTRIBUTING.md) okuyun.

## Lisans

GO+ Dil Sunucusu, GO+ projesi ile aynı lisans altında dağıtılmaktadır.
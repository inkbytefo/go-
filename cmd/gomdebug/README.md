# GO-Minus Hata Ayıklama Aracı (gomdebug)

GO-Minus Hata Ayıklama Aracı, GO-Minus programlama dili için Debug Adapter Protocol (DAP) implementasyonudur. Bu araç, çeşitli IDE ve metin düzenleyicileri ile entegrasyon sağlayarak kesme noktaları, adım adım çalıştırma, değişken inceleme gibi hata ayıklama özelliklerini destekler.

## Özellikler

- Kesme noktaları ayarlama ve kaldırma
- Adım adım çalıştırma (adım, satır, fonksiyondan çıkma)
- Değişken inceleme
- Yığın izi görüntüleme
- İfade değerlendirme
- Koşullu kesme noktaları
- Çalışma zamanı değişken değiştirme

## Kurulum

GO-Minus Hata Ayıklama Aracı'nı kurmak için, GO-Minus derleyicisini kurduktan sonra aşağıdaki komutu çalıştırabilirsiniz:

```bash
go build -o gomdebug ./tools/gomdebug
```

Oluşturulan çalıştırılabilir dosyayı PATH ortam değişkeninizin bulunduğu bir dizine kopyalayabilirsiniz.

## Kullanım

GO+ Hata Ayıklama Aracı, standart giriş/çıkış veya TCP üzerinden çalışabilir:

```bash
# program.gop dosyasını hata ayıkla
gopdebug program.gop

# program.gop dosyasını hata ayıkla (alternatif sözdizimi)
gopdebug -program=program.gop

# TCP sunucu olarak çalıştır (varsayılan port: 8081)
gopdebug -mode=tcp program.gop

# TCP sunucu olarak belirtilen portta çalıştır
gopdebug -mode=tcp -addr=:9091 program.gop

# Log dosyasına yaz
gopdebug -log=debug.log program.gop
```

## IDE Entegrasyonu

GO-Minus Hata Ayıklama Aracı, Debug Adapter Protocol'ü destekleyen herhangi bir IDE veya metin düzenleyicisi ile kullanılabilir. Aşağıdaki IDE'ler için özel eklentiler mevcuttur:

- Visual Studio Code: [GO-Minus VS Code Eklentisi](../ide/vscode/README.md)
- JetBrains IDEs: [GO-Minus JetBrains Eklentisi](../ide/jetbrains/README.md)
- Vim/Neovim: [GO-Minus Vim Eklentisi](../ide/vim/README.md)
- Emacs: [GO-Minus Emacs Eklentisi](../ide/emacs/README.md)

## Geliştirme

GO-Minus Hata Ayıklama Aracı, GO-Minus dilinin gelişimiyle birlikte sürekli olarak geliştirilmektedir. Katkıda bulunmak için, lütfen [katkı sağlama rehberini](../../CONTRIBUTING.md) okuyun.

## Lisans

GO-Minus Hata Ayıklama Aracı, GO-Minus projesi ile aynı lisans altında dağıtılmaktadır.
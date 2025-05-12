# GO-Minus Programlama Dili

GO-Minus, Go programlama dilinin tüm özelliklerini içeren ve C++ benzeri özelliklerle (sınıflar, şablonlar, istisna işleme vb.) genişletilmiş bir programlama dilidir. GO-Minus dosyaları `.gom` uzantısını kullanır.

## 🎯 Amaç

GO-Minus dilinin temel amacı, Go'nun sadeliği, hızlı derleme süreleri ve güçlü eşzamanlılık modelini, C++'ın düşük seviyeli sistem kontrolü, performans optimizasyonları, şablon metaprogramlama (basitleştirilmiş) ve zengin OOP yetenekleriyle birleştirmektir. Hem yüksek performanslı sistem programlama hem de hızlı uygulama geliştirme için "tatlı noktayı" bulmayı hedefler.

## ✨ Özellikler

- **Go Uyumluluğu**: Go'nun tüm özelliklerini destekler
- **Sınıflar ve Kalıtım**: C++ benzeri sınıf ve kalıtım desteği
- **Şablonlar**: Jenerik programlama için şablon desteği
- **İstisna İşleme**: Try-catch-finally yapıları
- **Erişim Belirleyicileri**: Public, private, protected erişim kontrolü
- **LLVM Tabanlı**: Güçlü optimizasyon ve çoklu platform desteği
- **Zengin Standart Kütüphane**: Temel veri yapıları, I/O işlemleri, eşzamanlılık desteği
- **Geliştirme Araçları**: Paket yöneticisi, test aracı, belgelendirme aracı, kod biçimlendirme aracı
- **IDE Entegrasyonu**: VS Code, JetBrains IDE'leri, Vim/Neovim, Emacs için eklentiler

## 🚀 Başlangıç

### Kurulum

```bash
# GO-Minus derleyicisini klonlayın
git clone https://github.com/gominus/gominus.git
cd gominus

# GO-Minus derleyicisini derleyin
go build -o gominus ./cmd/gominus

# Derleyiciyi PATH'e ekleyin
# Windows için:
# copy gominus.exe C:\Windows\System32\
# Linux/macOS için:
# sudo cp gominus /usr/local/bin/
```

### Merhaba Dünya

```go
// main.gom
package main

import "fmt"

func main() {
    fmt.Println("Merhaba, GO-Minus!")
}
```

```bash
# Programı derleyin ve çalıştırın
gominus run main.gom
```

### Sınıf Örneği

```go
// person.gom
package main

import "fmt"

class Person {
    private:
        string name
        int age

    public:
        Person(string name, int age) {
            this.name = name
            this.age = age
        }

        string getName() {
            return this.name
        }

        int getAge() {
            return this.age
        }

        void birthday() {
            this.age++
        }
}

func main() {
    person := Person("Ahmet", 30)
    fmt.Println("İsim:", person.getName())
    fmt.Println("Yaş:", person.getAge())

    person.birthday()
    fmt.Println("Yeni yaş:", person.getAge())
}
```

```bash
# Programı derleyin ve çalıştırın
gominus run person.gom
```

## 🛠️ Geliştirme Araçları

GO-Minus dili, aşağıdaki geliştirme araçlarını sağlar:

- **gompm**: GO-Minus Paket Yöneticisi
- **gomtest**: GO-Minus Test Aracı
- **gomdoc**: GO-Minus Belgelendirme Aracı
- **gomfmt**: GO-Minus Kod Biçimlendirme Aracı
- **gomlsp**: GO-Minus Dil Sunucusu
- **gomdebug**: GO-Minus Hata Ayıklama Aracı

## 🔌 IDE Entegrasyonu

GO-Minus dili, aşağıdaki IDE'ler için eklentiler sağlar:

- **VS Code**: [GO-Minus VS Code Eklentisi](ide/vscode/README.md)
- **JetBrains IDEs**: [GO-Minus JetBrains Eklentisi](ide/jetbrains/README.md)
- **Vim/Neovim**: [GO-Minus Vim Eklentisi](ide/vim/README.md)
- **Emacs**: [GO-Minus Emacs Eklentisi](ide/emacs/README.md)

## 📚 Belgelendirme

- [Başlangıç Rehberi](docs/tutorial/getting-started.md)
- [Dil Referansı](docs/reference/README.md)
- [Standart Kütüphane](stdlib/README.md)
- [Öğreticiler](docs/tutorial/README.md)
- [En İyi Uygulamalar](docs/best-practices.md)
- [SSS](docs/faq.md)

## 👥 Topluluk

- [GitHub](https://github.com/gominus/gominus)
- [Web Sitesi](website/index.html)
- [Discord](https://discord.gg/gominus)
- [Forum](https://forum.gominus.org)

## 🤝 Katkıda Bulunma

GO-Minus projesine katkıda bulunmak için, lütfen [katkı sağlama rehberini](CONTRIBUTING.md) okuyun. Tüm katkıda bulunanlar, [davranış kurallarımıza](CODE_OF_CONDUCT.md) uymayı kabul etmiş sayılır.

## 📋 İlerleme

GO-Minus dilinin geliştirme sürecini takip etmek için [ilerleme raporunu](progress.md) inceleyebilirsiniz.

## 📄 Lisans

GO-Minus dili, [MIT Lisansı](LICENSE) altında lisanslanmıştır.
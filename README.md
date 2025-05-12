# GO+ Programlama Dili

GO+ (GO Plus), Go programlama dilinin tüm özelliklerini içeren ve C++ benzeri özelliklerle (sınıflar, şablonlar, istisna işleme vb.) genişletilmiş bir programlama dilidir. GO+ dosyaları `.gop` uzantısını kullanır.

## 🎯 Amaç

GO+ dilinin temel amacı, Go'nun sadeliği, hızlı derleme süreleri ve güçlü eşzamanlılık modelini, C++'ın düşük seviyeli sistem kontrolü, performans optimizasyonları, şablon metaprogramlama (basitleştirilmiş) ve zengin OOP yetenekleriyle birleştirmektir. Hem yüksek performanslı sistem programlama hem de hızlı uygulama geliştirme için "tatlı noktayı" bulmayı hedefler.

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
# GO+ derleyicisini klonlayın
git clone https://github.com/goplus/goplus.git
cd goplus

# GO+ derleyicisini derleyin
go build -o goplus ./cmd/goplus

# Derleyiciyi PATH'e ekleyin
# Windows için:
# copy goplus.exe C:\Windows\System32\
# Linux/macOS için:
# sudo cp goplus /usr/local/bin/
```

### Merhaba Dünya

```go
// main.gop
package main

import "fmt"

func main() {
    fmt.Println("Merhaba, GO+!")
}
```

```bash
# Programı derleyin ve çalıştırın
goplus run main.gop
```

### Sınıf Örneği

```go
// person.gop
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
goplus run person.gop
```

## 🛠️ Geliştirme Araçları

GO+ dili, aşağıdaki geliştirme araçlarını sağlar:

- **goppm**: GO+ Paket Yöneticisi
- **goptest**: GO+ Test Aracı
- **gopdoc**: GO+ Belgelendirme Aracı
- **gopfmt**: GO+ Kod Biçimlendirme Aracı
- **goplsp**: GO+ Dil Sunucusu
- **gopdebug**: GO+ Hata Ayıklama Aracı

## 🔌 IDE Entegrasyonu

GO+ dili, aşağıdaki IDE'ler için eklentiler sağlar:

- **VS Code**: [GO+ VS Code Eklentisi](ide/vscode/README.md)
- **JetBrains IDEs**: [GO+ JetBrains Eklentisi](ide/jetbrains/README.md)
- **Vim/Neovim**: [GO+ Vim Eklentisi](ide/vim/README.md)
- **Emacs**: [GO+ Emacs Eklentisi](ide/emacs/README.md)

## 📚 Belgelendirme

- [Başlangıç Rehberi](docs/tutorial/getting-started.md)
- [Dil Referansı](docs/reference/README.md)
- [Standart Kütüphane](stdlib/README.md)
- [Öğreticiler](docs/tutorial/README.md)
- [En İyi Uygulamalar](docs/best-practices.md)
- [SSS](docs/faq.md)

## 👥 Topluluk

- [GitHub](https://github.com/goplus/goplus)
- [Web Sitesi](website/index.html)
- [Discord](https://discord.gg/goplus)
- [Forum](https://forum.goplus.org)

## 🤝 Katkıda Bulunma

GO+ projesine katkıda bulunmak için, lütfen [katkı sağlama rehberini](CONTRIBUTING.md) okuyun. Tüm katkıda bulunanlar, [davranış kurallarımıza](CODE_OF_CONDUCT.md) uymayı kabul etmiş sayılır.

## 📋 İlerleme

GO+ dilinin geliştirme sürecini takip etmek için [ilerleme raporunu](progress.md) inceleyebilirsiniz.

## 📄 Lisans

GO+ dili, [MIT Lisansı](LICENSE) altında lisanslanmıştır.
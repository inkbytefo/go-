# GO-Minus Sık Sorulan Sorular (SSS)

Bu belge, GO-Minus programlama dili hakkında sık sorulan soruları ve cevaplarını içermektedir.

## Genel Sorular

### GO-Minus nedir?

GO-Minus, Go programlama dilinin tüm özelliklerini içeren ve C++ benzeri özelliklerle (sınıflar, şablonlar, istisna işleme vb.) genişletilmiş bir programlama dilidir. GO-Minus dosyaları `.gom` uzantısını kullanır.

### GO-Minus'un amacı nedir?

GO-Minus dilinin temel amacı, Go'nun sadeliği, hızlı derleme süreleri ve güçlü eşzamanlılık modelini, C++'ın düşük seviyeli sistem kontrolü, performans optimizasyonları, şablon metaprogramlama (basitleştirilmiş) ve zengin OOP yetenekleriyle birleştirmektir. Hem yüksek performanslı sistem programlama hem de hızlı uygulama geliştirme için "tatlı noktayı" bulmayı hedefler.

### GO-Minus'u neden kullanmalıyım?

GO-Minus'u şu durumlarda kullanmayı düşünebilirsiniz:

1. Go'nun sadeliğini ve verimliliğini seviyorsanız, ancak nesne yönelimli programlama ve şablonlar gibi C++ özelliklerine ihtiyaç duyuyorsanız
2. Yüksek performanslı sistem programlama yapıyorsanız, ancak C++'ın karmaşıklığından kaçınmak istiyorsanız
3. Gerçek zamanlı uygulamalar, oyunlar veya düşük gecikme gerektiren sistemler geliştiriyorsanız
4. Büyük ölçekli uygulamalar için daha güçlü bir tip sistemi ve nesne modeli arıyorsanız

### GO-Minus hangi platformları destekliyor?

GO-Minus, LLVM tabanlı olduğu için çeşitli platformları destekler:

- Windows (x86, x64)
- Linux (x86, x64, ARM)
- macOS (x64, ARM)
- FreeBSD
- Android
- iOS
- WebAssembly

### GO-Minus açık kaynaklı mı?

Evet, GO-Minus MIT Lisansı altında lisanslanmıştır. Kaynak kodu [GitHub](https://github.com/gominus/gominus) üzerinden erişilebilir.

## Teknik Sorular

### GO-Minus, Go ile uyumlu mu?

GO-Minus, Go'nun tüm özelliklerini destekler, ancak bazı ek sözdizimi ve semantik farklılıklar içerir. Çoğu Go kodu, GO-Minus derleyicisi tarafından derlenebilir, ancak GO-Minus'a özgü özellikleri kullanan kodlar Go derleyicisi tarafından derlenemez.

### GO-Minus'un performansı nasıl?

GO-Minus, LLVM optimizasyon geçişlerini kullanarak yüksek performanslı kod üretir. Performans testleri, GO-Minus'un Go ile karşılaştırılabilir ve bazı durumlarda daha iyi performans sunduğunu göstermektedir. Özellikle manuel bellek yönetimi seçeneği, performans kritik uygulamalarda avantaj sağlar.

### GO-Minus'ta garbage collection var mı?

Evet, GO-Minus, Go'nun garbage collector'ünü korur, ancak performans kritik bölümler için manuel bellek yönetimi seçeneği de sunar. Bu, gerçek zamanlı uygulamalar ve düşük gecikme gerektiren sistemler için idealdir.

### GO-Minus'ta sınıflar ve kalıtım nasıl çalışır?

GO-Minus, C++ benzeri sınıf ve kalıtım desteği sağlar:

```go
class Animal {
    protected:
        string name
    
    public:
        Animal(string name) {
            this.name = name
        }
        
        virtual string makeSound() {
            return "..."
        }
}

class Dog : Animal {
    public:
        Dog(string name) : Animal(name) {}
        
        override string makeSound() {
            return "Woof!"
        }
}
```

### GO-Minus'ta şablonlar nasıl çalışır?

GO-Minus, C++ benzeri şablon desteği sağlar:

```go
template<T>
class Stack {
    private:
        T[] items
        int size
    
    public:
        Stack() {
            this.items = T[]{}
            this.size = 0
        }
        
        void push(T item) {
            this.items = append(this.items, item)
            this.size++
        }
        
        T pop() {
            if this.size == 0 {
                throw new Exception("Stack is empty")
            }
            
            this.size--
            item := this.items[this.size]
            this.items = this.items[:this.size]
            return item
        }
}
```

### GO-Minus'ta istisna işleme nasıl çalışır?

GO-Minus, C++ benzeri istisna işleme desteği sağlar:

```go
try {
    result := divide(10, 0)
    fmt.Println("Sonuç:", result)
} catch (DivisionByZeroException e) {
    fmt.Println("Hata:", e.message)
} catch (Exception e) {
    fmt.Println("Genel hata:", e.message)
} finally {
    fmt.Println("İşlem tamamlandı")
}
```

### GO-Minus'ta Go'nun goroutine ve channel yapıları destekleniyor mu?

Evet, GO-Minus, Go'nun goroutine ve channel tabanlı eşzamanlılık modelini korur ve genişletir:

```go
func main() {
    messages := make(chan string)
    
    go func() {
        messages <- "Merhaba, Dünya!"
    }()
    
    msg := <-messages
    fmt.Println(msg)
}
```

### GO-Minus'ta C/C++ kütüphaneleri kullanabilir miyim?

Evet, GO-Minus, C ve C++ kütüphaneleriyle entegrasyon için çeşitli yöntemler sağlar:

1. C kütüphaneleri için Go'nun cgo mekanizması
2. C++ kütüphaneleri için doğrudan bağlama desteği
3. LLVM IR seviyesinde entegrasyon

### GO-Minus'ta modül sistemi nasıl çalışır?

GO-Minus, Go'nun modül sistemini kullanır, ancak `.gom` uzantılı dosyaları destekler:

```
module myapp

go 1.18

require (
    github.com/example/package v1.0.0
)
```

## Kurulum ve Kullanım

### GO-Minus'u nasıl kurarım?

GO-Minus'u kurmak için aşağıdaki adımları izleyin:

1. GO-Minus deposunu klonlayın:
```bash
git clone https://github.com/gominus/gominus.git
cd gominus
```

2. GO-Minus derleyicisini derleyin:
```bash
go build -o gominus ./cmd/gominus
```

3. Derleyiciyi PATH'e ekleyin:
```bash
# Windows için:
copy gominus.exe C:\Windows\System32\
# Linux/macOS için:
sudo cp gominus /usr/local/bin/
```

### GO-Minus programlarını nasıl derlerim ve çalıştırırım?

GO-Minus programlarını derlemek ve çalıştırmak için:

```bash
# Derleme ve çalıştırma
gominus run main.gom

# Sadece derleme
gominus build main.gom

# Çalıştırma
./main
```

### GO-Minus için hangi IDE'ler ve editörler destekleniyor?

GO-Minus, aşağıdaki IDE'ler ve editörler için eklentiler sağlar:

- VS Code: [GO-Minus VS Code Eklentisi](../ide/vscode/README.md)
- JetBrains IDEs: [GO-Minus JetBrains Eklentisi](../ide/jetbrains/README.md)
- Vim/Neovim: [GO-Minus Vim Eklentisi](../ide/vim/README.md)
- Emacs: [GO-Minus Emacs Eklentisi](../ide/emacs/README.md)

### GO-Minus paket yöneticisi nasıl kullanılır?

GO-Minus Paket Yöneticisi (gompm), GO-Minus paketlerini yönetmek için kullanılır:

```bash
# Paket kurma
gompm get github.com/example/package

# Paket güncelleme
gompm update github.com/example/package

# Paket kaldırma
gompm remove github.com/example/package
```

## Topluluk ve Destek

### GO-Minus topluluğuna nasıl katılabilirim?

GO-Minus topluluğuna katılmak için:

1. [GitHub](https://github.com/gominus/gominus) üzerinden projeyi takip edin
2. [Discord](https://discord.gg/gominus) sunucusuna katılın
3. [Forum](https://forum.gominus.org) üzerinden tartışmalara katılın
4. [Twitter](https://twitter.com/gominuslang) üzerinden duyuruları takip edin

### GO-Minus'a nasıl katkıda bulunabilirim?

GO-Minus'a katkıda bulunmak için:

1. [Katkı Sağlama Rehberi](../CONTRIBUTING.md) belgesini okuyun
2. [GitHub Issues](https://github.com/gominus/gominus/issues) üzerinden açık sorunları inceleyin
3. Çözüm için bir pull request gönderin
4. Belgelendirme, örnekler ve öğreticiler oluşturun
5. Hataları bildirin ve özellik isteklerinde bulunun

### GO-Minus için destek nereden alabilirim?

GO-Minus için destek almak için:

1. [Belgelendirme](../docs/reference/README.md) sayfalarını inceleyin
2. [SSS](../docs/faq.md) belgesini okuyun
3. [Discord](https://discord.gg/gominus) sunucusunda soru sorun
4. [Forum](https://forum.gominus.org) üzerinden yardım isteyin
5. [GitHub Issues](https://github.com/gominus/gominus/issues) üzerinden hata bildirin

### GO-Minus için ticari destek mevcut mu?

Evet, GO-Minus için ticari destek mevcuttur. Ticari destek hakkında bilgi almak için [contact@gominus.org](mailto:contact@gominus.org) adresine e-posta gönderin.

## Gelecek Planları

### GO-Minus'un gelecek planları nelerdir?

GO-Minus'un gelecek planları şunları içermektedir:

1. Vulkan ve OpenGL için resmi bağlayıcılar geliştirme
2. Çöp toplama optimizasyonu ve manuel bellek yönetimi seçeneklerini iyileştirme
3. Ekosistemi genişletme (daha fazla örnek proje, eğitim içeriği ve kütüphane)
4. IDE ve hata ayıklama araçlarını iyileştirme
5. Performans karşılaştırmaları ve benchmark'lar yayınlama

### GO-Minus ne zaman 1.0 sürümüne ulaşacak?

GO-Minus'un 1.0 sürümü, temel dil özellikleri, standart kütüphane, araçlar ve belgelendirme tamamlandığında yayınlanacaktır. Şu anki tahminlere göre, 1.0 sürümü 2024 yılının sonlarında yayınlanacaktır.

### GO-Minus'ta hangi yeni özellikler planlanıyor?

GO-Minus için planlanan yeni özellikler şunları içermektedir:

1. Daha güçlü şablon metaprogramlama yetenekleri
2. Daha iyi C++ entegrasyonu
3. Daha fazla standart kütüphane paketi
4. WebAssembly desteğinin iyileştirilmesi
5. Daha iyi IDE entegrasyonu ve hata ayıklama desteği

## Diğer Sorular

### GO-Minus, Go veya C++'ın yerini alacak mı?

GO-Minus, Go veya C++'ın yerini almak için tasarlanmamıştır. Bunun yerine, her iki dilin de güçlü yönlerini birleştirerek belirli kullanım senaryoları için alternatif bir seçenek sunmayı amaçlamaktadır.

### GO-Minus, Go veya C++ ile karşılaştırıldığında ne kadar olgun?

GO-Minus, Go ve C++ kadar olgun değildir, çünkü daha yeni bir dildir. Ancak, GO-Minus ekibi, dili hızla olgunlaştırmak ve kararlı bir ekosistem oluşturmak için çalışmaktadır.

### GO-Minus'u üretim ortamında kullanabilir miyim?

GO-Minus hala geliştirme aşamasındadır, bu nedenle üretim ortamında kullanmadan önce dikkatli olmanız önerilir. Ancak, bazı erken benimseyenler, GO-Minus'u belirli projelerde başarıyla kullanmaktadır.

### GO-Minus'un lisansı nedir?

GO-Minus, [MIT Lisansı](../LICENSE) altında lisanslanmıştır. Bu, GO-Minus'u hem açık kaynaklı hem de ticari projelerde özgürce kullanabileceğiniz anlamına gelir.

### GO-Minus adı nereden geliyor?

GO-Minus adı, Go programlama dilinin adından ve "minus" (eksi) kelimesinden gelmektedir. "Minus" eki, dilin Go'nun bazı karmaşıklıklarını azaltırken C++ benzeri özellikler eklediğini ifade eder.

### GO-Minus'un logosu nedir?

GO-Minus'un logosu, Go'nun maskotu olan Gopher'ın stilize edilmiş bir versiyonudur. Logo, GO-Minus'un Go ile olan bağlantısını ve C++ benzeri özellikleri temsil eden bir "−" (eksi) işareti içerir.

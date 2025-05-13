# GO-Minus Standart Kütüphane Genişletmeleri

Bu belge, GO-Minus programlama dilinin standart kütüphanesine eklenen yeni modülleri ve genişletmeleri açıklamaktadır.

## Genel Bakış

GO-Minus standart kütüphanesi, Go'nun standart kütüphanesini temel alır ve GO-Minus'un sınıf, şablon ve istisna işleme gibi özelliklerini kullanarak genişletilmiştir. Son güncellemelerle birlikte, aşağıdaki yeni modüller ve genişletmeler eklenmiştir:

1. **Trie (Önek Ağacı) Implementasyonu**
2. **Buffered IO Implementasyonu**
3. **Regex Paketi**

## 1. Trie (Önek Ağacı) Implementasyonu

Trie (önek ağacı veya dijital ağaç olarak da bilinir), string anahtarları verimli bir şekilde depolamak ve aramak için kullanılan bir ağaç veri yapısıdır. GO-Minus'un container paketine eklenen Trie implementasyonu, özellikle önek tabanlı arama ve otomatik tamamlama gibi işlemler için idealdir.

### Özellikler

- **Jenerik Tip Desteği**: Herhangi bir tip için kullanılabilir
- **Kelime Ekleme, Arama ve Silme**: Temel Trie işlemleri
- **Önek Araması**: Belirli bir önekle başlayan kelimeleri bulma
- **Tüm Kelimeleri Listeleme**: Trie'deki tüm kelimeleri listeleme
- **Boş Kontrol ve Boyut Hesaplama**: Trie'nin boş olup olmadığını kontrol etme ve boyutunu hesaplama

### Örnek Kullanım

```go
import "container/trie"

// String değerler için bir Trie oluştur
t := trie.Trie.New<string>()

// Kelimeler ekle
t.Insert("apple", "elma")
t.Insert("banana", "muz")
t.Insert("application", "uygulama")

// Kelime ara
value, found := t.Search("apple")
if found {
    fmt.Println("apple:", value) // "elma" yazdırır
}

// Önek kontrolü
if t.StartsWith("app") {
    fmt.Println("'app' öneki ile başlayan kelimeler var")
}

// Belirli önekle başlayan tüm kelimeleri al
appWords := t.GetWordsWithPrefix("app")
for word, value := range appWords {
    fmt.Printf("%s: %s\n", word, value)
}
```

### Performans

Trie veri yapısı, string anahtarları için aşağıdaki karmaşıklıklara sahiptir:

- **Ekleme**: O(m), m = anahtar uzunluğu
- **Arama**: O(m), m = anahtar uzunluğu
- **Silme**: O(m), m = anahtar uzunluğu
- **Önek Araması**: O(m), m = önek uzunluğu
- **Belirli Önekle Başlayan Tüm Kelimeleri Bulma**: O(n), n = eşleşen kelime sayısı

### Uygulama Detayları

Trie implementasyonu, `stdlib/container/trie/trie.gom` dosyasında uygulanmıştır. İmplementasyon, aşağıdaki bileşenlerden oluşur:

- **TrieNode<T>**: Trie ağacındaki her bir düğümü temsil eden sınıf
- **Trie<T>**: Trie veri yapısını temsil eden sınıf

```go
// TrieNode, bir trie düğümünü temsil eder.
class TrieNode<T> {
    private:
        map[rune]TrieNode<T> children
        bool isEndOfWord
        T value
        bool hasValue
    
    public:
        // Yeni bir düğüm oluştur
        static func New<T>() *TrieNode<T> {
            node := new TrieNode<T>()
            node.children = make(map[rune]TrieNode<T>)
            node.isEndOfWord = false
            node.hasValue = false
            return node
        }
        
        // ... diğer metotlar
}

// Trie, bir önek ağacını temsil eder.
class Trie<T> {
    private:
        TrieNode<T> root
        int size
    
    public:
        // Yeni bir Trie oluştur
        static func New<T>() *Trie<T> {
            t := new Trie<T>()
            t.root = *TrieNode.New<T>()
            return t
        }
        
        // ... diğer metotlar
}
```

## 2. Buffered IO Implementasyonu

Tamponlanmış I/O, disk veya ağ gibi yavaş I/O kaynaklarına erişim sırasında performansı artırmak için kullanılan bir tekniktir. GO-Minus'un io paketine eklenen Buffered IO implementasyonu, küçük okuma/yazma işlemlerini daha büyük bloklarda gruplandırarak, sistem çağrılarının sayısını azaltır ve genel performansı artırır.

### Özellikler

- **Tamponlanmış Okuma**: Verimli okuma işlemleri için BufferedReader
- **Tamponlanmış Yazma**: Verimli yazma işlemleri için BufferedWriter
- **Özelleştirilebilir Tampon Boyutu**: İhtiyaca göre tampon boyutunu ayarlama
- **Satır Satır Okuma**: Metin dosyalarını satır satır okuma
- **Byte ve String Yazma**: Hem byte dizileri hem de string'ler için yazma desteği

### Örnek Kullanım

```go
import (
    "io"
    "io/buffered"
)

// BufferedReader örneği
file := io.Open("input.txt")
defer file.Close()

reader := buffered.BufferedReader.New(file, 4096)

// Satır satır okuma
for {
    line, err := reader.ReadLine()
    if err == io.EOF {
        break
    }
    if err != nil {
        // Hata işleme
        break
    }
    
    fmt.Println(line)
}

// BufferedWriter örneği
outFile := io.Create("output.txt")
defer outFile.Close()

writer := buffered.BufferedWriter.New(outFile, 4096)

// Veri yazma
writer.WriteString("Hello, World!\n")
writer.WriteString("This is a test.\n")

// Tamponu boşalt
writer.Flush()
```

### Performans

Tamponlanmış I/O, özellikle küçük okuma/yazma işlemlerinin sık yapıldığı durumlarda önemli performans artışı sağlar. Örneğin, bir dosyayı satır satır okurken veya küçük parçalar halinde yazarken, tamponlanmış I/O kullanmak, tamponsuz I/O'ya göre çok daha hızlı olabilir.

### Uygulama Detayları

Buffered IO implementasyonu, `stdlib/io/buffered/buffered.gom` dosyasında uygulanmıştır. İmplementasyon, aşağıdaki bileşenlerden oluşur:

- **BufferedReader**: Tamponlanmış okuma işlemleri için sınıf
- **BufferedWriter**: Tamponlanmış yazma işlemleri için sınıf

```go
// BufferedReader, tamponlanmış okuma işlemleri için kullanılır.
class BufferedReader {
    private:
        io.Reader reader
        []byte buffer
        int bufferSize
        int readPos
        int writePos
        bool eof
    
    public:
        // Yeni bir BufferedReader oluştur
        static func New(reader io.Reader, bufferSize int) *BufferedReader {
            // ... implementasyon
        }
        
        // ... diğer metotlar
}

// BufferedWriter, tamponlanmış yazma işlemleri için kullanılır.
class BufferedWriter {
    private:
        io.Writer writer
        []byte buffer
        int bufferSize
        int count
    
    public:
        // Yeni bir BufferedWriter oluştur
        static func New(writer io.Writer, bufferSize int) *BufferedWriter {
            // ... implementasyon
        }
        
        // ... diğer metotlar
}
```

## 3. Regex Paketi

Düzenli ifadeler (Regular Expressions), metin içinde belirli desenleri aramak, eşleştirmek ve değiştirmek için kullanılan özel bir dil ve sözdizimi sunar. GO-Minus'a eklenen Regex paketi, metin işleme, form doğrulama, veri çıkarma ve dönüştürme gibi birçok senaryoda kullanılabilir.

### Özellikler

- **Düzenli İfade Deseni Derleme**: Desenleri derleme ve yeniden kullanma
- **Metin Eşleştirme**: Bir metnin bir desene uyup uymadığını kontrol etme
- **Tüm Eşleşmeleri Bulma**: Bir metindeki tüm eşleşmeleri bulma
- **Metin Değiştirme**: Eşleşen metinleri başka metinlerle değiştirme
- **Metin Bölme**: Bir metni belirli bir desene göre bölme
- **Büyük/Küçük Harf Duyarlı ve Duyarsız Modlar**: Büyük/küçük harf duyarlılığını kontrol etme
- **Çok Satırlı Mod**: Çok satırlı metinlerde satır başı ve sonu eşleştirme

### Örnek Kullanım

```go
import "regex"

// Düzenli ifade deseni derleme
pattern := regex.Compile("hello")

// Metin eşleştirme
if pattern.Match("hello world") {
    fmt.Println("Eşleşme bulundu!")
}

// Büyük/küçük harf duyarsız desen
patternIgnoreCase := regex.CompileIgnoreCase("hello")

// Tüm eşleşmeleri bulma
matches := patternIgnoreCase.FindAll("Hello, hello, HELLO!")
fmt.Println("Eşleşme sayısı:", len(matches))

// Metin değiştirme
result := patternIgnoreCase.Replace("Hello, hello, HELLO!", "hi")
fmt.Println(result) // "hi, hi, hi!"

// Metin bölme
parts := regex.Split(",", "apple,banana,orange")
for _, part := range parts {
    fmt.Println(part)
}
```

### Uygulama Detayları

Regex paketi, `stdlib/regex/regex.gom` dosyasında uygulanmıştır. Paket, aşağıdaki bileşenlerden oluşur:

- **RegexPattern**: Derlenmiş bir düzenli ifade desenini temsil eden sınıf
- **Compile, CompileIgnoreCase, CompileMultiline**: Desen derleme fonksiyonları
- **Match, Replace, Split**: Yardımcı fonksiyonlar

```go
// RegexPattern, derlenmiş bir düzenli ifade desenini temsil eder.
class RegexPattern {
    private:
        string pattern
        bool caseSensitive
        bool multiline
        bool compiled
        bool hasSpecialChars
    
    public:
        // Yeni bir RegexPattern oluştur
        static func New(pattern string, caseSensitive bool, multiline bool) *RegexPattern {
            // ... implementasyon
        }
        
        // ... diğer metotlar
}

// Compile, bir düzenli ifade desenini derler.
func Compile(pattern string) *RegexPattern {
    return RegexPattern.New(pattern, true, false)
}

// CompileIgnoreCase, bir düzenli ifade desenini büyük/küçük harf duyarsız olarak derler.
func CompileIgnoreCase(pattern string) *RegexPattern {
    return RegexPattern.New(pattern, false, false)
}

// ... diğer fonksiyonlar
```

## Sonuç

GO-Minus standart kütüphanesine eklenen bu yeni modüller ve genişletmeler, GO-Minus programlama dilinin yeteneklerini artırır ve programcılara daha güçlü araçlar sunar. Bu genişletmeler, GO-Minus'un daha geniş bir uygulama yelpazesinde kullanılmasını sağlar.

Daha fazla bilgi ve örnekler için aşağıdaki belgelere bakabilirsiniz:

- [Trie Örneği](examples/container/trie-example.md)
- [Buffered IO Örneği](examples/io/buffered-io-example.md)
- [Regex Örneği](examples/regex/regex-example.md)

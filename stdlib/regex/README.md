# GO-Minus Regex Paketi

Bu paket, GO-Minus programlama dili için düzenli ifade (regular expression) işlemlerini sağlar. Düzenli ifadeler, metin arama, eşleştirme, değiştirme ve bölme işlemleri için güçlü bir araçtır.

## Özellikler

- Düzenli ifade deseni derleme
- Metin eşleştirme
- Tüm eşleşmeleri bulma
- Metin değiştirme
- Metin bölme
- Büyük/küçük harf duyarlı ve duyarsız modlar
- Çok satırlı mod desteği

## Kullanım

```go
import "regex"

func main() {
    // Düzenli ifade deseni derleme
    pattern := regex.Compile("hello")
    
    // Metin eşleştirme
    if pattern.Match("hello world") {
        fmt.Println("Eşleşme bulundu!")
    }
    
    // Büyük/küçük harf duyarsız desen
    pattern = regex.CompileIgnoreCase("hello")
    
    // Büyük/küçük harf duyarsız eşleştirme
    if pattern.Match("Hello World") {
        fmt.Println("Eşleşme bulundu!")
    }
    
    // Tüm eşleşmeleri bulma
    matches := pattern.FindAll("Hello, hello, HELLO!")
    fmt.Println("Eşleşme sayısı:", len(matches))
    
    // Metin değiştirme
    result := pattern.Replace("Hello, hello, HELLO!", "hi")
    fmt.Println(result) // "hi, hi, hi!"
    
    // Metin bölme
    parts := regex.Split(",", "apple,banana,orange")
    for _, part := range parts {
        fmt.Println(part)
    }
}
```

## RegexPattern Sınıfı

`RegexPattern` sınıfı, bir düzenli ifade desenini temsil eder ve aşağıdaki metotları sağlar:

### Yapıcı Metotlar

- `New(pattern string, caseSensitive bool, multiline bool) *RegexPattern`: Yeni bir RegexPattern oluşturur.

### Eşleştirme Metotları

- `Match(text string) bool`: Bir metni düzenli ifade deseniyle eşleştirir.
- `FindAll(text string) []string`: Bir metindeki tüm eşleşmeleri bulur.
- `Replace(text string, replacement string) string`: Bir metindeki tüm eşleşmeleri belirtilen metinle değiştirir.
- `Split(text string) []string`: Bir metni düzenli ifade deseniyle böler.

### Bilgi Metotları

- `GetPattern() string`: Düzenli ifade desenini döndürür.
- `IsCaseSensitive() bool`: Düzenli ifade deseninin büyük/küçük harf duyarlı olup olmadığını döndürür.
- `IsMultiline() bool`: Düzenli ifade deseninin çok satırlı olup olmadığını döndürür.

## Yardımcı Fonksiyonlar

Paket, aşağıdaki yardımcı fonksiyonları sağlar:

- `Compile(pattern string) *RegexPattern`: Bir düzenli ifade desenini derler.
- `CompileIgnoreCase(pattern string) *RegexPattern`: Bir düzenli ifade desenini büyük/küçük harf duyarsız olarak derler.
- `CompileMultiline(pattern string) *RegexPattern`: Bir düzenli ifade desenini çok satırlı olarak derler.
- `Match(pattern string, text string) bool`: Bir metni düzenli ifade deseniyle eşleştirir.
- `MatchIgnoreCase(pattern string, text string) bool`: Bir metni büyük/küçük harf duyarsız olarak düzenli ifade deseniyle eşleştirir.
- `Replace(pattern string, text string, replacement string) string`: Bir metindeki tüm eşleşmeleri belirtilen metinle değiştirir.
- `Split(pattern string, text string) []string`: Bir metni düzenli ifade deseniyle böler.

## Desteklenen Düzenli İfade Sözdizimi

Bu paket, aşağıdaki düzenli ifade özelliklerini destekler:

- Basit metin eşleştirme
- Büyük/küçük harf duyarlı ve duyarsız modlar
- Çok satırlı mod

Not: Bu paket, şu anda karmaşık düzenli ifade desenlerini (özel karakterler, yakalama grupları, vb.) desteklememektedir. Gelecek sürümlerde daha fazla özellik eklenecektir.

## Performans

Düzenli ifade işlemleri, özellikle karmaşık desenler ve büyük metinler için yoğun işlem gerektirebilir. Performans için aşağıdaki ipuçlarını göz önünde bulundurun:

- Düzenli ifade desenlerini önceden derleyin ve yeniden kullanın.
- Mümkün olduğunda basit metin eşleştirme kullanın.
- Büyük metinlerde tüm eşleşmeleri bulmak yerine, ilk eşleşmeyi bulun.

## Sınırlamalar

- Bu paket, şu anda karmaşık düzenli ifade desenlerini desteklememektedir.
- Büyük/küçük harf duyarsız mod, sadece ASCII karakterleri için doğru çalışır.
- Çok satırlı mod, şu anda tam olarak desteklenmemektedir.

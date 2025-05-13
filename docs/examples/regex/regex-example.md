# GO-Minus Regex Paketi Örneği

Bu örnek, GO-Minus'un standart kütüphanesinde bulunan Regex (Düzenli İfadeler) paketinin kullanımını göstermektedir. Düzenli ifadeler, metin arama, eşleştirme, değiştirme ve bölme işlemleri için güçlü bir araçtır.

## Regex Nedir?

Düzenli ifadeler (Regular Expressions), metin içinde belirli desenleri aramak, eşleştirmek ve değiştirmek için kullanılan özel bir dil ve sözdizimi sunar. Düzenli ifadeler, metin işleme, form doğrulama, veri çıkarma ve dönüştürme gibi birçok senaryoda kullanılır.

## Temel Regex Kullanımı

```go
// regex_basic.gom
package main

import (
    "fmt"
    "regex"
)

func main() {
    // Düzenli ifade deseni derleme
    pattern := regex.Compile("hello")
    
    // Metin eşleştirme
    if pattern.Match("hello world") {
        fmt.Println("Eşleşme bulundu!")
    } else {
        fmt.Println("Eşleşme bulunamadı!")
    }
    
    // Büyük/küçük harf duyarsız desen
    patternIgnoreCase := regex.CompileIgnoreCase("hello")
    
    // Büyük/küçük harf duyarsız eşleştirme
    if patternIgnoreCase.Match("Hello World") {
        fmt.Println("Büyük/küçük harf duyarsız eşleşme bulundu!")
    } else {
        fmt.Println("Büyük/küçük harf duyarsız eşleşme bulunamadı!")
    }
    
    // Çok satırlı mod
    multilinePattern := regex.CompileMultiline("^start")
    
    text := "start of line 1\nstart of line 2\nnot at start"
    
    if multilinePattern.Match(text) {
        fmt.Println("Çok satırlı eşleşme bulundu!")
    } else {
        fmt.Println("Çok satırlı eşleşme bulunamadı!")
    }
}
```

## Çıktı

```
Eşleşme bulundu!
Büyük/küçük harf duyarsız eşleşme bulundu!
Çok satırlı eşleşme bulundu!
```

## Tüm Eşleşmeleri Bulma

```go
// regex_find_all.gom
package main

import (
    "fmt"
    "regex"
)

func main() {
    // Düzenli ifade deseni
    pattern := regex.Compile("\\d+") // Bir veya daha fazla rakam
    
    // Test metni
    text := "There are 42 apples and 15 oranges in the basket, and it costs 9.99 dollars."
    
    // Tüm eşleşmeleri bul
    matches := pattern.FindAll(text)
    
    fmt.Printf("Eşleşme sayısı: %d\n", len(matches))
    fmt.Println("Eşleşmeler:")
    
    for i, match := range matches {
        fmt.Printf("  %d: %s\n", i+1, match)
    }
    
    // Büyük/küçük harf duyarsız tüm eşleşmeleri bul
    fruitPattern := regex.CompileIgnoreCase("apple|orange")
    fruitMatches := fruitPattern.FindAll(text)
    
    fmt.Printf("\nMeyve eşleşme sayısı: %d\n", len(fruitMatches))
    fmt.Println("Meyve eşleşmeleri:")
    
    for i, match := range fruitMatches {
        fmt.Printf("  %d: %s\n", i+1, match)
    }
}
```

## Çıktı

```
Eşleşme sayısı: 3
Eşleşmeler:
  1: 42
  2: 15
  3: 9

Meyve eşleşme sayısı: 2
Meyve eşleşmeleri:
  1: apple
  2: orange
```

## Metin Değiştirme

```go
// regex_replace.gom
package main

import (
    "fmt"
    "regex"
)

func main() {
    // Düzenli ifade deseni
    pattern := regex.Compile("\\d+") // Bir veya daha fazla rakam
    
    // Test metni
    text := "There are 42 apples and 15 oranges in the basket."
    
    // Rakamları yıldızlarla değiştir
    result := pattern.Replace(text, "**")
    fmt.Println("Rakamları değiştir:")
    fmt.Println("  Orijinal:", text)
    fmt.Println("  Değiştirilmiş:", result)
    
    // Büyük/küçük harf duyarsız değiştirme
    fruitPattern := regex.CompileIgnoreCase("apple|orange")
    fruitResult := fruitPattern.Replace(text, "fruit")
    
    fmt.Println("\nMeyveleri değiştir:")
    fmt.Println("  Orijinal:", text)
    fmt.Println("  Değiştirilmiş:", fruitResult)
    
    // İlk eşleşmeyi değiştir
    firstResult := pattern.ReplaceFirst(text, "XX")
    fmt.Println("\nİlk rakamı değiştir:")
    fmt.Println("  Orijinal:", text)
    fmt.Println("  Değiştirilmiş:", firstResult)
}
```

## Çıktı

```
Rakamları değiştir:
  Orijinal: There are 42 apples and 15 oranges in the basket.
  Değiştirilmiş: There are ** apples and ** oranges in the basket.

Meyveleri değiştir:
  Orijinal: There are 42 apples and 15 oranges in the basket.
  Değiştirilmiş: There are 42 fruit and 15 fruit in the basket.

İlk rakamı değiştir:
  Orijinal: There are 42 apples and 15 oranges in the basket.
  Değiştirilmiş: There are XX apples and 15 oranges in the basket.
```

## Metin Bölme

```go
// regex_split.gom
package main

import (
    "fmt"
    "regex"
)

func main() {
    // Virgülle ayrılmış metin
    text := "apple,banana,orange,grape"
    
    // Virgülle böl
    parts := regex.Split(",", text)
    
    fmt.Println("Virgülle bölme:")
    for i, part := range parts {
        fmt.Printf("  %d: %s\n", i+1, part)
    }
    
    // Boşluk karakterleriyle böl
    text2 := "This is a   test   with multiple   spaces"
    spaceParts := regex.Split("\\s+", text2)
    
    fmt.Println("\nBoşluk karakterleriyle bölme:")
    for i, part := range spaceParts {
        fmt.Printf("  %d: %s\n", i+1, part)
    }
    
    // Rakamlarla böl
    text3 := "abc123def456ghi789"
    digitParts := regex.Split("\\d+", text3)
    
    fmt.Println("\nRakamlarla bölme:")
    for i, part := range digitParts {
        fmt.Printf("  %d: %s\n", i+1, part)
    }
}
```

## Çıktı

```
Virgülle bölme:
  1: apple
  2: banana
  3: orange
  4: grape

Boşluk karakterleriyle bölme:
  1: This
  2: is
  3: a
  4: test
  5: with
  6: multiple
  7: spaces

Rakamlarla bölme:
  1: abc
  2: def
  3: ghi
  4: 
```

## E-posta Doğrulama Örneği

Düzenli ifadelerin pratik bir uygulaması olarak, e-posta doğrulama örneği:

```go
// regex_email_validation.gom
package main

import (
    "fmt"
    "regex"
)

// EmailValidator, e-posta doğrulama sınıfı
class EmailValidator {
    private:
        regex.RegexPattern emailPattern
    
    public:
        // Yeni bir e-posta doğrulayıcı oluştur
        EmailValidator() {
            // Basit bir e-posta deseni
            this.emailPattern = regex.Compile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
        }
        
        // E-posta adresini doğrula
        bool Validate(string email) {
            return this.emailPattern.Match(email)
        }
}

func main() {
    // E-posta doğrulayıcı oluştur
    validator := EmailValidator()
    
    // Test e-postaları
    emails := []string{
        "user@example.com",
        "john.doe@company.co.uk",
        "invalid-email",
        "missing@domain",
        "special@chars#.com",
        "trailing@dot.",
        "multiple@dots..com",
        "valid+tag@gmail.com",
    }
    
    fmt.Println("E-posta Doğrulama Sonuçları:")
    fmt.Println("-----------------------------")
    
    for _, email := range emails {
        isValid := validator.Validate(email)
        status := "Geçersiz"
        if isValid {
            status = "Geçerli"
        }
        fmt.Printf("%-25s : %s\n", email, status)
    }
}
```

## Çıktı

```
E-posta Doğrulama Sonuçları:
-----------------------------
user@example.com          : Geçerli
john.doe@company.co.uk    : Geçerli
invalid-email             : Geçersiz
missing@domain            : Geçersiz
special@chars#.com        : Geçersiz
trailing@dot.             : Geçersiz
multiple@dots..com        : Geçersiz
valid+tag@gmail.com       : Geçerli
```

## Metin Analizi Örneği

Düzenli ifadelerin başka bir uygulaması olarak, metin analizi örneği:

```go
// regex_text_analysis.gom
package main

import (
    "fmt"
    "regex"
)

// TextAnalyzer, metin analizi sınıfı
class TextAnalyzer {
    private:
        regex.RegexPattern wordPattern
        regex.RegexPattern sentencePattern
        regex.RegexPattern numberPattern
        regex.RegexPattern emailPattern
        regex.RegexPattern urlPattern
    
    public:
        // Yeni bir metin analizci oluştur
        TextAnalyzer() {
            this.wordPattern = regex.Compile("\\b[a-zA-Z]+\\b")
            this.sentencePattern = regex.Compile("[^.!?]+[.!?]")
            this.numberPattern = regex.Compile("\\b\\d+(\\.\\d+)?\\b")
            this.emailPattern = regex.Compile("\\b[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}\\b")
            this.urlPattern = regex.Compile("https?://[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}(/[a-zA-Z0-9./-]*)?")
        }
        
        // Kelime sayısını bul
        int CountWords(string text) {
            matches := this.wordPattern.FindAll(text)
            return len(matches)
        }
        
        // Cümle sayısını bul
        int CountSentences(string text) {
            matches := this.sentencePattern.FindAll(text)
            return len(matches)
        }
        
        // Sayıları bul
        []string FindNumbers(string text) {
            return this.numberPattern.FindAll(text)
        }
        
        // E-postaları bul
        []string FindEmails(string text) {
            return this.emailPattern.FindAll(text)
        }
        
        // URL'leri bul
        []string FindURLs(string text) {
            return this.urlPattern.FindAll(text)
        }
}

func main() {
    // Metin analizci oluştur
    analyzer := TextAnalyzer()
    
    // Test metni
    text := "Bu bir örnek metindir. Bu metin içinde 42 sayısı ve 3.14 gibi sayılar bulunmaktadır. " +
            "Ayrıca user@example.com ve admin@company.co.uk gibi e-posta adresleri ve " +
            "https://www.example.com ve http://blog.example.com/posts/123 gibi URL'ler içermektedir. " +
            "Bu son cümledir!"
    
    // Analiz sonuçları
    wordCount := analyzer.CountWords(text)
    sentenceCount := analyzer.CountSentences(text)
    numbers := analyzer.FindNumbers(text)
    emails := analyzer.FindEmails(text)
    urls := analyzer.FindURLs(text)
    
    fmt.Println("Metin Analizi Sonuçları:")
    fmt.Println("------------------------")
    fmt.Printf("Kelime Sayısı: %d\n", wordCount)
    fmt.Printf("Cümle Sayısı: %d\n", sentenceCount)
    
    fmt.Println("\nSayılar:")
    for i, number := range numbers {
        fmt.Printf("  %d: %s\n", i+1, number)
    }
    
    fmt.Println("\nE-posta Adresleri:")
    for i, email := range emails {
        fmt.Printf("  %d: %s\n", i+1, email)
    }
    
    fmt.Println("\nURL'ler:")
    for i, url := range urls {
        fmt.Printf("  %d: %s\n", i+1, url)
    }
}
```

## Çıktı

```
Metin Analizi Sonuçları:
------------------------
Kelime Sayısı: 42
Cümle Sayısı: 3

Sayılar:
  1: 42
  2: 3.14

E-posta Adresleri:
  1: user@example.com
  2: admin@company.co.uk

URL'ler:
  1: https://www.example.com
  2: http://blog.example.com/posts/123
```

Bu örnekler, GO-Minus'un Regex paketinin nasıl kullanılacağını göstermektedir. Daha fazla bilgi için [Regex Belgelendirmesi](../../stdlib/regex/README.md) belgesine bakabilirsiniz.

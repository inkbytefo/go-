# GO-Minus Dil Referansı

Bu belge, GO-Minus programlama dilinin kapsamlı bir referansını sağlamaktadır. GO-Minus, Go programlama dilinin tüm özelliklerini içeren ve C++ benzeri özelliklerle genişletilmiş bir programlama dilidir.

## İçindekiler

1. [Sözdizimi](#sözdizimi)
2. [Veri Tipleri](#veri-tipleri)
3. [Değişkenler ve Sabitler](#değişkenler-ve-sabitler)
4. [Operatörler](#operatörler)
5. [Kontrol Yapıları](#kontrol-yapıları)
6. [Fonksiyonlar](#fonksiyonlar)
7. [Sınıflar ve Nesneler](#sınıflar-ve-nesneler)
8. [Şablonlar](#şablonlar)
9. [İstisna İşleme](#istisna-işleme)
10. [Paketler ve Modüller](#paketler-ve-modüller)
11. [Eşzamanlılık](#eşzamanlılık)
12. [Bellek Yönetimi](#bellek-yönetimi)
13. [Standart Kütüphane](#standart-kütüphane)

## Sözdizimi

GO-Minus, Go'nun sözdizimini temel alır ve C++ benzeri özellikler ekler. GO-Minus dosyaları `.gom` uzantısını kullanır.

### Temel Sözdizimi

```go
// Paket bildirimi
package main

// İçe aktarma bildirimleri
import (
    "fmt"
    "math"
)

// Ana fonksiyon
func main() {
    fmt.Println("Merhaba, GO-Minus!")
}
```

### Yorum Satırları

```go
// Tek satır yorum

/*
Çok
satırlı
yorum
*/
```

## Veri Tipleri

GO-Minus, Go'nun tüm temel veri tiplerini destekler ve bazı ek tipler ekler.

### Temel Tipler

- **Tamsayılar**: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
- **Kayan Noktalı Sayılar**: `float32`, `float64`
- **Karmaşık Sayılar**: `complex64`, `complex128`
- **Boolean**: `bool`
- **Karakter**: `char` (GO-Minus'a özgü, C++ benzeri)
- **Dizeler**: `string`
- **Byte**: `byte` (`uint8` için takma ad)
- **Rune**: `rune` (`int32` için takma ad)

### Bileşik Tipler

- **Diziler**: `[n]T`
- **Dilimler**: `[]T`
- **Haritalar**: `map[K]V`
- **Yapılar**: `struct`
- **Kanallar**: `chan T`
- **İşaretçiler**: `*T`
- **Fonksiyonlar**: `func(T1, T2) R`
- **Arayüzler**: `interface`
- **Sınıflar**: `class` (GO-Minus'a özgü, C++ benzeri)
- **Şablonlar**: `template<T>` (GO-Minus'a özgü, C++ benzeri)

## Değişkenler ve Sabitler

### Değişken Tanımlama

```go
// Tam tanımlama
var x int = 10

// Tip çıkarımı ile tanımlama
var y = 20

// Kısa değişken tanımlama
z := 30

// Çoklu değişken tanımlama
var a, b, c int = 1, 2, 3
i, j := 0, "sıfır"
```

### Sabit Tanımlama

```go
// Tek sabit tanımlama
const pi = 3.14159

// Çoklu sabit tanımlama
const (
    a = 1
    b = 2
    c = 3
)

// iota kullanımı
const (
    zero = iota  // 0
    one          // 1
    two          // 2
)
```

## Operatörler

### Aritmetik Operatörler

- `+`: Toplama
- `-`: Çıkarma
- `*`: Çarpma
- `/`: Bölme
- `%`: Mod alma
- `++`: Artırma
- `--`: Azaltma

### Karşılaştırma Operatörleri

- `==`: Eşittir
- `!=`: Eşit değildir
- `<`: Küçüktür
- `>`: Büyüktür
- `<=`: Küçük eşittir
- `>=`: Büyük eşittir

### Mantıksal Operatörler

- `&&`: VE
- `||`: VEYA
- `!`: DEĞİL

### Bitsel Operatörler

- `&`: Bitsel VE
- `|`: Bitsel VEYA
- `^`: Bitsel XOR
- `<<`: Sola kaydırma
- `>>`: Sağa kaydırma
- `&^`: Bit temizleme

### Atama Operatörleri

- `=`: Atama
- `+=`: Toplama ve atama
- `-=`: Çıkarma ve atama
- `*=`: Çarpma ve atama
- `/=`: Bölme ve atama
- `%=`: Mod alma ve atama
- `&=`: Bitsel VE ve atama
- `|=`: Bitsel VEYA ve atama
- `^=`: Bitsel XOR ve atama
- `<<=`: Sola kaydırma ve atama
- `>>=`: Sağa kaydırma ve atama

### Diğer Operatörler

- `&`: Adres alma
- `*`: İşaretçi dereference
- `<-`: Kanal operatörü
- `->`: Üye erişimi (GO-Minus'a özgü, C++ benzeri)
- `::`: Kapsam çözümleme (GO-Minus'a özgü, C++ benzeri)

## Kontrol Yapıları

### If-Else

```go
if x > 0 {
    fmt.Println("Pozitif")
} else if x < 0 {
    fmt.Println("Negatif")
} else {
    fmt.Println("Sıfır")
}

// Kısa ifade ile
if n := rand.Intn(10); n > 5 {
    fmt.Println("Büyük sayı:", n)
} else {
    fmt.Println("Küçük sayı:", n)
}
```

### For Döngüsü

```go
// Klasik for döngüsü
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While benzeri for döngüsü
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}

// Sonsuz döngü
for {
    fmt.Println("Sonsuz döngü")
    break
}

// Range ile döngü
nums := []int{1, 2, 3, 4, 5}
for i, num := range nums {
    fmt.Printf("Index: %d, Value: %d\n", i, num)
}
```

### Switch

```go
switch day {
case "Pazartesi":
    fmt.Println("Haftanın ilk günü")
case "Salı", "Çarşamba", "Perşembe":
    fmt.Println("Hafta ortası")
case "Cuma":
    fmt.Println("Haftanın son iş günü")
default:
    fmt.Println("Hafta sonu")
}

// Koşullu switch
switch {
case hour < 12:
    fmt.Println("Günaydın")
case hour < 18:
    fmt.Println("İyi günler")
default:
    fmt.Println("İyi akşamlar")
}
```

## Fonksiyonlar

### Temel Fonksiyon Tanımlama

```go
func add(a int, b int) int {
    return a + b
}

// Kısa parametre listesi
func subtract(a, b int) int {
    return a - b
}
```

### Çoklu Dönüş Değeri

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("sıfıra bölme hatası")
    }
    return a / b, nil
}

// İsimlendirilen dönüş değerleri
func calculate(a, b int) (sum, diff int) {
    sum = a + b
    diff = a - b
    return // Çıplak return
}
```

### Değişken Sayıda Parametre

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Kullanım
fmt.Println(sum(1, 2, 3, 4, 5))
```

### Fonksiyon Değişkenleri ve Kapanışlar

```go
// Fonksiyon değişkeni
var compute func(int, int) int

// Kapanış (closure)
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
```

## Sınıflar ve Nesneler

GO-Minus, C++ benzeri sınıf ve nesne desteği sağlar.

### Sınıf Tanımlama

```go
class Person {
    private:
        string name
        int age
    
    public:
        // Yapıcı metot
        Person(string name, int age) {
            this.name = name
            this.age = age
        }
        
        // Yıkıcı metot
        ~Person() {
            fmt.Println("Person nesnesi silindi")
        }
        
        // Getter metotları
        string getName() {
            return this.name
        }
        
        int getAge() {
            return this.age
        }
        
        // Setter metotları
        void setName(string name) {
            this.name = name
        }
        
        void setAge(int age) {
            this.age = age
        }
        
        // Metot
        void birthday() {
            this.age++
        }
        
        // Statik metot
        static Person createDefault() {
            return Person("John Doe", 30)
        }
}
```

### Kalıtım

```go
// Temel sınıf
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

// Türetilmiş sınıf
class Dog : Animal {
    private:
        string breed
    
    public:
        Dog(string name, string breed) : Animal(name) {
            this.breed = breed
        }
        
        override string makeSound() {
            return "Woof!"
        }
        
        string getBreed() {
            return this.breed
        }
}
```

### Çoklu Kalıtım

```go
class A {
    public:
        void methodA() {
            fmt.Println("Method A")
        }
}

class B {
    public:
        void methodB() {
            fmt.Println("Method B")
        }
}

class C : A, B {
    public:
        void methodC() {
            methodA()
            methodB()
            fmt.Println("Method C")
        }
}
```

### Arayüz Uygulaması

```go
interface Drawable {
    void draw()
}

class Circle : Drawable {
    private:
        float radius
    
    public:
        Circle(float radius) {
            this.radius = radius
        }
        
        void draw() {
            fmt.Println("Drawing a circle with radius", this.radius)
        }
}
```

## Şablonlar

GO-Minus, C++ benzeri şablon desteği sağlar.

### Şablon Sınıf

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
        
        bool isEmpty() {
            return this.size == 0
        }
}
```

### Şablon Fonksiyon

```go
template<T>
T max(T a, T b) {
    if a > b {
        return a
    }
    return b
}
```

### Şablon Kullanımı

```go
// Şablon sınıf kullanımı
intStack := Stack<int>()
intStack.push(1)
intStack.push(2)
intStack.push(3)

// Şablon fonksiyon kullanımı
fmt.Println(max<int>(5, 10))      // 10
fmt.Println(max<string>("a", "b")) // "b"
```

## İstisna İşleme

GO-Minus, C++ benzeri istisna işleme desteği sağlar.

### İstisna Tanımlama

```go
// Temel istisna sınıfı
class Exception {
    public:
        string message
        
        Exception(string message) {
            this.message = message
        }
}

// Özel istisna sınıfı
class DivisionByZeroException : Exception {
    public:
        DivisionByZeroException() : Exception("Division by zero") {}
}
```

### İstisna Fırlatma

```go
func divide(a, b float64) float64 {
    if b == 0 {
        throw new DivisionByZeroException()
    }
    return a / b
}
```

### İstisna Yakalama

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

## Paketler ve Modüller

GO-Minus, Go'nun paket ve modül sistemini kullanır.

### Paket Oluşturma

```go
// math/calculator.gom
package math

class Calculator {
    public:
        static int add(int a, int b) {
            return a + b
        }
        
        static int subtract(int a, int b) {
            return a - b
        }
}

// Dışa aktarılan fonksiyon
func Multiply(a, b int) int {
    return a * b
}
```

### Paket Kullanımı

```go
// main.gom
package main

import (
    "fmt"
    "myapp/math"
)

func main() {
    // Sınıf kullanımı
    sum := math.Calculator.add(5, 3)
    fmt.Println("Toplam:", sum)
    
    // Fonksiyon kullanımı
    product := math.Multiply(4, 2)
    fmt.Println("Çarpım:", product)
}
```

## Eşzamanlılık

GO-Minus, Go'nun goroutine ve channel tabanlı eşzamanlılık modelini korur ve genişletir.

### Goroutine

```go
func main() {
    go sayHello()
    time.Sleep(time.Second)
}

func sayHello() {
    fmt.Println("Merhaba, Dünya!")
}
```

### Kanallar

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

### Select

```go
func main() {
    c1 := make(chan string)
    c2 := make(chan string)
    
    go func() {
        time.Sleep(time.Second)
        c1 <- "bir"
    }()
    
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "iki"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("Alındı:", msg1)
        case msg2 := <-c2:
            fmt.Println("Alındı:", msg2)
        }
    }
}
```

## Bellek Yönetimi

GO-Minus, Go'nun garbage collector'ünü korur ve manuel bellek yönetimi seçeneği ekler.

### Garbage Collection

```go
func main() {
    // Otomatik bellek yönetimi
    data := createLargeData()
    processData(data)
    // data otomatik olarak temizlenir
}
```

### Manuel Bellek Yönetimi

```go
func main() {
    // Manuel bellek yönetimi modu
    unsafe {
        buffer := allocate<byte>(1024 * 1024)
        defer free(buffer)
        
        // Performans kritik işlemler
        // ...
    }
}
```

## Standart Kütüphane

GO-Minus, Go'nun standart kütüphanesini içerir ve ek kütüphaneler ekler.

### Temel Paketler

- `fmt`: Biçimlendirilmiş I/O işlemleri
- `io`: Temel I/O arayüzleri
- `os`: İşletim sistemi fonksiyonları
- `strings`: Dize işleme
- `math`: Matematiksel fonksiyonlar
- `time`: Zaman işlemleri
- `errors`: Hata işleme

### GO-Minus Özel Paketleri

- `container`: Veri yapıları (list, vector, deque, heap, trie, graph)
- `concurrent`: Eşzamanlılık araçları (semaphore, barrier, threadpool, future/promise)
- `core`: Temel GO-Minus tipleri ve fonksiyonları

Daha fazla bilgi için [Standart Kütüphane](../../stdlib/README.md) belgesine bakın.

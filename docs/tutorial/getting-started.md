# GO-Minus Başlangıç Rehberi

Bu rehber, GO-Minus programlama dilini kullanmaya başlamak için gereken adımları açıklamaktadır. GO-Minus, Go programlama dilinin tüm özelliklerini içeren ve C++ benzeri özelliklerle genişletilmiş bir programlama dilidir.

## İçindekiler

1. [GO-Minus Kurulumu](#go-minus-kurulumu)
2. [İlk GO-Minus Programı](#ilk-go-minus-programı)
3. [Temel Sözdizimi](#temel-sözdizimi)
4. [Sınıflar ve Nesneler](#sınıflar-ve-nesneler)
5. [Şablonlar](#şablonlar)
6. [İstisna İşleme](#istisna-işleme)
7. [Paketler ve Modüller](#paketler-ve-modüller)
8. [Derleme ve Çalıştırma](#derleme-ve-çalıştırma)
9. [IDE Entegrasyonu](#ide-entegrasyonu)
10. [Sonraki Adımlar](#sonraki-adımlar)

## GO-Minus Kurulumu

GO-Minus'u kurmak için aşağıdaki adımları izleyin:

### Ön Koşullar

- Go 1.18 veya üzeri
- LLVM 14.0 veya üzeri
- Git

### Kurulum Adımları

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

Windows için:
```bash
copy gominus.exe C:\Windows\System32\
```

Linux/macOS için:
```bash
sudo cp gominus /usr/local/bin/
```

4. Kurulumu doğrulayın:

```bash
gominus version
```

## İlk GO-Minus Programı

GO-Minus ile ilk programınızı yazalım. Bir metin editörü açın ve aşağıdaki kodu `hello.gom` adlı bir dosyaya kaydedin:

```go
// hello.gom
package main

import "fmt"

func main() {
    fmt.Println("Merhaba, GO-Minus!")
}
```

Programı derlemek ve çalıştırmak için:

```bash
gominus run hello.gom
```

Çıktı:
```
Merhaba, GO-Minus!
```

## Temel Sözdizimi

GO-Minus, Go'nun sözdizimini temel alır ve C++ benzeri özellikler ekler.

### Değişkenler ve Sabitler

```go
// Değişken tanımlama
var x int = 10
var y = 20 // Tip çıkarımı
z := 30    // Kısa değişken tanımlama

// Sabit tanımlama
const pi = 3.14159
const (
    a = 1
    b = 2
)
```

### Kontrol Yapıları

```go
// If-else
if x > 0 {
    fmt.Println("Pozitif")
} else if x < 0 {
    fmt.Println("Negatif")
} else {
    fmt.Println("Sıfır")
}

// For döngüsü
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Range döngüsü
nums := []int{1, 2, 3, 4, 5}
for i, num := range nums {
    fmt.Printf("Index: %d, Value: %d\n", i, num)
}

// Switch
switch day {
case "Pazartesi":
    fmt.Println("Haftanın ilk günü")
case "Cuma":
    fmt.Println("Haftanın son iş günü")
default:
    fmt.Println("Başka bir gün")
}
```

### Fonksiyonlar

```go
// Temel fonksiyon
func add(a int, b int) int {
    return a + b
}

// Çoklu dönüş değeri
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("sıfıra bölme hatası")
    }
    return a / b, nil
}

// Değişken sayıda parametre
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

## Sınıflar ve Nesneler

GO-Minus, C++ benzeri sınıf ve nesne desteği sağlar.

```go
// Sınıf tanımlama
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
        
        // Getter metotları
        string getName() {
            return this.name
        }
        
        int getAge() {
            return this.age
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

// Sınıf kullanımı
func main() {
    // Nesne oluşturma
    person := Person("Ahmet", 30)
    
    // Metot çağırma
    fmt.Println("İsim:", person.getName())
    fmt.Println("Yaş:", person.getAge())
    
    person.birthday()
    fmt.Println("Yeni yaş:", person.getAge())
    
    // Statik metot çağırma
    defaultPerson := Person.createDefault()
    fmt.Println("Varsayılan isim:", defaultPerson.getName())
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

// Kalıtım kullanımı
func main() {
    dog := Dog("Buddy", "Golden Retriever")
    fmt.Println("Ses:", dog.makeSound())
    fmt.Println("Irk:", dog.getBreed())
    
    // Polimorfizm
    var animal Animal = dog
    fmt.Println("Polimorfik ses:", animal.makeSound())
}
```

## Şablonlar

GO-Minus, C++ benzeri şablon desteği sağlar.

```go
// Şablon sınıf
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

// Şablon fonksiyon
template<T>
T max(T a, T b) {
    if a > b {
        return a
    }
    return b
}

// Şablon kullanımı
func main() {
    // Şablon sınıf kullanımı
    intStack := Stack<int>()
    intStack.push(1)
    intStack.push(2)
    intStack.push(3)
    
    fmt.Println(intStack.pop())  // 3
    fmt.Println(intStack.pop())  // 2
    
    // Şablon fonksiyon kullanımı
    fmt.Println(max<int>(5, 10))      // 10
    fmt.Println(max<string>("a", "b")) // "b"
}
```

## İstisna İşleme

GO-Minus, C++ benzeri istisna işleme desteği sağlar.

```go
// İstisna tanımlama
class DivisionByZeroException : Exception {
    public:
        DivisionByZeroException() : Exception("Division by zero") {}
}

// İstisna fırlatma
func divide(a, b float64) float64 {
    if b == 0 {
        throw new DivisionByZeroException()
    }
    return a / b
}

// İstisna yakalama
func main() {
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

## Derleme ve Çalıştırma

GO-Minus programlarını derlemek ve çalıştırmak için çeşitli komutlar bulunmaktadır.

### Tek Dosya Derleme ve Çalıştırma

```bash
# Derleme ve çalıştırma
gominus run hello.gom

# Sadece derleme
gominus build hello.gom

# Çalıştırma
./hello
```

### Proje Derleme

```bash
# Proje dizininde
gominus build

# Belirli bir çıktı adı ile
gominus build -o myapp

# Çalıştırma
./myapp
```

## IDE Entegrasyonu

GO-Minus, çeşitli IDE'ler için eklentiler sağlar.

### VS Code

1. VS Code'u açın
2. Eklentiler sekmesine gidin
3. "GO-Minus" araması yapın
4. GO-Minus eklentisini yükleyin

### JetBrains IDEs

1. JetBrains IDE'nizi açın (IntelliJ IDEA, GoLand, vb.)
2. Eklentiler sekmesine gidin
3. Marketplace'te "GO-Minus" araması yapın
4. GO-Minus eklentisini yükleyin

### Vim/Neovim

```bash
# Vim-Plug ile
Plug 'gominus/vim-gominus'

# Packer ile
use 'gominus/vim-gominus'
```

## Sonraki Adımlar

GO-Minus dilini öğrenmeye devam etmek için:

1. [Dil Referansı](../reference/README.md) belgesini inceleyin
2. [Standart Kütüphane](../../stdlib/README.md) belgelerini okuyun
3. [Örnekler](../examples) klasöründeki örnek projeleri inceleyin
4. [Discord](https://discord.gg/gominus) sunucumuza katılın
5. [GitHub](https://github.com/gominus/gominus) üzerinden projeye katkıda bulunun

GO-Minus ile programlama yaparken herhangi bir sorunla karşılaşırsanız, [SSS](../faq.md) belgesine göz atabilir veya topluluk kanallarımızdan yardım isteyebilirsiniz.

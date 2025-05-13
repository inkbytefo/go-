# GO-Minus Tip Çıkarımı Örneği

Bu örnek, GO-Minus'un tip çıkarımı (type inference) özelliğini göstermektedir. Tip çıkarımı, değişken tiplerinin açıkça belirtilmediği durumlarda, derleyicinin değişken tiplerini otomatik olarak belirlemesini sağlar.

## Tip Çıkarımı Özellikleri

- Değişken tanımlamalarında tip çıkarımı
- Fonksiyon dönüş tiplerinde tip çıkarımı
- Karmaşık ifadelerde tip çıkarımı
- Jenerik fonksiyonlarda tip çıkarımı
- Şablon sınıflarda tip çıkarımı

## Temel Tip Çıkarımı

```go
// type_inference_basic.gom
package main

import "fmt"

func main() {
    // Tip çıkarımı ile değişken tanımlama
    x := 10          // int olarak çıkarılır
    y := 3.14        // float olarak çıkarılır
    z := "Merhaba"   // string olarak çıkarılır
    b := true        // bool olarak çıkarılır
    
    // Tipleri yazdır
    fmt.Printf("x tipi: %T\n", x)
    fmt.Printf("y tipi: %T\n", y)
    fmt.Printf("z tipi: %T\n", z)
    fmt.Printf("b tipi: %T\n", b)
}
```

## Çıktı

```
x tipi: int
y tipi: float64
z tipi: string
b tipi: bool
```

## Karmaşık İfadelerde Tip Çıkarımı

```go
// type_inference_complex.gom
package main

import "fmt"

func main() {
    // Karmaşık ifadelerde tip çıkarımı
    a := 10
    b := 20
    
    // İfade sonucu int olarak çıkarılır
    c := a + b
    
    // İfade sonucu float64 olarak çıkarılır
    d := a + b / 2.0
    
    // İfade sonucu bool olarak çıkarılır
    e := a > b
    
    // İfade sonucu string olarak çıkarılır
    f := fmt.Sprintf("%d ve %d", a, b)
    
    fmt.Printf("c tipi: %T, değeri: %v\n", c, c)
    fmt.Printf("d tipi: %T, değeri: %v\n", d, d)
    fmt.Printf("e tipi: %T, değeri: %v\n", e, e)
    fmt.Printf("f tipi: %T, değeri: %v\n", f, f)
}
```

## Çıktı

```
c tipi: int, değeri: 30
d tipi: float64, değeri: 20
e tipi: bool, değeri: false
f tipi: string, değeri: 10 ve 20
```

## Fonksiyon Dönüş Tiplerinde Tip Çıkarımı

```go
// type_inference_functions.gom
package main

import "fmt"

// Dönüş tipi belirtilmemiş, int olarak çıkarılır
func add(a, b int) {
    return a + b
}

// Dönüş tipi belirtilmemiş, string olarak çıkarılır
func greeting(name string) {
    return "Merhaba, " + name
}

// Dönüş tipi belirtilmemiş, bool olarak çıkarılır
func isEven(n int) {
    return n % 2 == 0
}

func main() {
    result1 := add(10, 20)
    result2 := greeting("Ahmet")
    result3 := isEven(10)
    
    fmt.Printf("result1 tipi: %T, değeri: %v\n", result1, result1)
    fmt.Printf("result2 tipi: %T, değeri: %v\n", result2, result2)
    fmt.Printf("result3 tipi: %T, değeri: %v\n", result3, result3)
}
```

## Çıktı

```
result1 tipi: int, değeri: 30
result2 tipi: string, değeri: Merhaba, Ahmet
result3 tipi: bool, değeri: true
```

## Jenerik Fonksiyonlarda Tip Çıkarımı

```go
// type_inference_generics.gom
package main

import "fmt"

// Jenerik fonksiyon
func first<T>(items []T) T {
    return items[0]
}

// Jenerik fonksiyon, dönüş tipi çıkarılır
func pair<T, U>(first T, second U) {
    return struct {
        First  T
        Second U
    }{first, second}
}

func main() {
    // Tip parametreleri çıkarılır
    intResult := first([3]int{1, 2, 3})
    strResult := first([2]string{"a", "b"})
    
    // Tip parametreleri çıkarılır
    p := pair(10, "hello")
    
    fmt.Printf("intResult tipi: %T, değeri: %v\n", intResult, intResult)
    fmt.Printf("strResult tipi: %T, değeri: %v\n", strResult, strResult)
    fmt.Printf("p tipi: %T, değeri: %v\n", p, p)
}
```

## Çıktı

```
intResult tipi: int, değeri: 1
strResult tipi: string, değeri: a
p tipi: struct{First int; Second string}, değeri: {10 hello}
```

## Şablon Sınıflarda Tip Çıkarımı

```go
// type_inference_templates.gom
package main

import "fmt"

// Şablon sınıf
class Pair<T, U> {
    private:
        T first
        U second
    
    public:
        Pair(T first, U second) {
            this.first = first
            this.second = second
        }
        
        T getFirst() {
            return this.first
        }
        
        U getSecond() {
            return this.second
        }
        
        string toString() {
            return fmt.Sprintf("(%v, %v)", this.first, this.second)
        }
}

func main() {
    // Tip parametreleri çıkarılır
    p1 := Pair(10, "hello")
    p2 := Pair(3.14, true)
    
    fmt.Printf("p1 tipi: Pair<int, string>, değeri: %s\n", p1.toString())
    fmt.Printf("p2 tipi: Pair<float64, bool>, değeri: %s\n", p2.toString())
    
    // Tip parametreleri açıkça belirtilir
    p3 := Pair<string, int>("count", 42)
    fmt.Printf("p3 tipi: Pair<string, int>, değeri: %s\n", p3.toString())
}
```

## Çıktı

```
p1 tipi: Pair<int, string>, değeri: (10, hello)
p2 tipi: Pair<float64, bool>, değeri: (3.14, true)
p3 tipi: Pair<string, int>, değeri: (count, 42)
```

## Tip Çıkarımı Sınırlamaları

GO-Minus'un tip çıkarımı güçlü olsa da, bazı durumlarda tiplerin açıkça belirtilmesi gerekebilir:

1. Belirsiz tip çıkarımı durumlarında
2. Boş koleksiyonlar oluşturulduğunda
3. Fonksiyon çağrılarında belirsizlik olduğunda
4. Şablon sınıf örneklemelerinde belirsizlik olduğunda

```go
// type_inference_limitations.gom
package main

import "fmt"

func main() {
    // Belirsiz tip çıkarımı - açıkça belirtilmeli
    var emptySlice []int = []int{}  // Boş dilim
    var emptyMap map[string]int = map[string]int{}  // Boş map
    
    // Fonksiyon çağrılarında belirsizlik
    result := process<int>(10)  // Tip açıkça belirtilmeli
    
    fmt.Printf("emptySlice tipi: %T\n", emptySlice)
    fmt.Printf("emptyMap tipi: %T\n", emptyMap)
    fmt.Printf("result tipi: %T\n", result)
}

// Belirsiz jenerik fonksiyon
func process<T>(value T) T {
    return value
}
```

## Çıktı

```
emptySlice tipi: []int
emptyMap tipi: map[string]int
result tipi: int
```

Bu örnekler, GO-Minus'un tip çıkarımı özelliğinin nasıl kullanılacağını göstermektedir. Daha fazla bilgi için [Tip Çıkarımı Belgelendirmesi](../../docs/reference/type-inference.md) belgesine bakabilirsiniz.

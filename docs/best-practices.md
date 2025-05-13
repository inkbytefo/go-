# GO-Minus En İyi Uygulamalar

Bu belge, GO-Minus programlama dili ile geliştirme yaparken izlenmesi gereken en iyi uygulamaları içermektedir. Bu uygulamalar, kodunuzu daha okunabilir, bakımı daha kolay ve daha verimli hale getirmenize yardımcı olacaktır.

## İçindekiler

1. [Kod Stili](#kod-stili)
2. [Sınıf ve Nesne Tasarımı](#sınıf-ve-nesne-tasarımı)
3. [Şablonlar](#şablonlar)
4. [İstisna İşleme](#istisna-işleme)
5. [Bellek Yönetimi](#bellek-yönetimi)
6. [Eşzamanlılık](#eşzamanlılık)
7. [Performans Optimizasyonları](#performans-optimizasyonları)
8. [Paket ve Modül Organizasyonu](#paket-ve-modül-organizasyonu)
9. [Test ve Hata Ayıklama](#test-ve-hata-ayıklama)
10. [Dokümantasyon](#dokümantasyon)

## Kod Stili

### İsimlendirme Kuralları

- **Paketler**: Küçük harflerle, tek kelime, kısaltma yok
  ```go
  package math
  package httpserver
  ```

- **Değişkenler ve Fonksiyonlar**: camelCase
  ```go
  var userName string
  func calculateTotal() int
  ```

- **Sınıflar ve Yapılar**: PascalCase
  ```go
  class Person {}
  struct Point {}
  ```

- **Arayüzler**: PascalCase, genellikle bir eylem + er
  ```go
  interface Reader {}
  interface EventHandler {}
  ```

- **Sabitler**: UPPER_SNAKE_CASE veya PascalCase
  ```go
  const MAX_CONNECTIONS = 100
  const Pi = 3.14159
  ```

### Girintileme ve Boşluklar

- Tab yerine 4 boşluk kullanın
- Operatörler etrafında boşluk bırakın
- Fonksiyon çağrılarında parantezden önce boşluk bırakmayın
- Süslü parantezleri aynı satırda başlatın

```go
if x > 0 {
    // Kod
} else {
    // Kod
}

func add(a, b int) int {
    return a + b
}
```

### Yorum Yazma

- Paketler, fonksiyonlar ve karmaşık kod blokları için açıklayıcı yorumlar ekleyin
- Paket yorumları, paket bildirimi öncesinde olmalıdır
- Fonksiyon yorumları, fonksiyon bildirimi öncesinde olmalıdır
- Yorumlar tam cümleler olmalı ve nokta ile bitmelidir

```go
// Package math provides basic mathematical functions.
package math

// Calculate returns the sum of all numbers from 1 to n.
// It returns an error if n is negative.
func Calculate(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("negative number")
    }
    
    // Use the formula: sum = n * (n + 1) / 2
    return n * (n + 1) / 2, nil
}
```

### Kod Düzeni

- İçe aktarmaları gruplandırın: standart kütüphane, üçüncü taraf, yerel
- Değişken tanımlamalarını gruplandırın
- İlgili fonksiyonları bir arada tutun
- Dosya uzunluğunu 500 satırın altında tutmaya çalışın

```go
package main

import (
    // Standart kütüphane
    "fmt"
    "os"
    
    // Üçüncü taraf
    "github.com/example/package"
    
    // Yerel
    "myapp/utils"
)

// Kod
```

## Sınıf ve Nesne Tasarımı

### Sınıf Tasarımı

- Sınıfları tek bir sorumluluğa sahip olacak şekilde tasarlayın (Tek Sorumluluk İlkesi)
- Sınıf hiyerarşilerini basit tutun, derin kalıtım zincirlerinden kaçının
- Arayüzleri küçük ve odaklanmış tutun
- Erişim belirleyicilerini doğru kullanın (public, private, protected)

```go
// Kötü: Çok fazla sorumluluk
class UserManager {
    public:
        void register(User user)
        void login(string username, string password)
        void sendEmail(User user, string subject, string body)
        void generateReport()
}

// İyi: Tek sorumluluk
class UserManager {
    public:
        void register(User user)
        void login(string username, string password)
}

class EmailService {
    public:
        void sendEmail(User user, string subject, string body)
}

class ReportGenerator {
    public:
        void generateReport()
}
```

### Yapıcı ve Yıkıcı Metotlar

- Yapıcı metotlarda tüm üye değişkenlerini başlatın
- Yıkıcı metotlarda tüm kaynakları temizleyin
- Birden fazla yapıcı metot sağlamak için statik fabrika metotları kullanın

```go
class Resource {
    private:
        File* file
        
    public:
        // Yapıcı metot
        Resource(string filename) {
            this.file = openFile(filename)
            if this.file == null {
                throw new Exception("File not found")
            }
        }
        
        // Statik fabrika metodu
        static Resource fromString(string content) {
            tempFile := createTempFile()
            writeToFile(tempFile, content)
            return Resource(tempFile.getName())
        }
        
        // Yıkıcı metot
        ~Resource() {
            if this.file != null {
                this.file.close()
            }
        }
}
```

### Kalıtım ve Kompozisyon

- Kalıtım yerine kompozisyonu tercih edin ("is-a" yerine "has-a" ilişkisi)
- Kalıtımı yalnızca gerçek bir "is-a" ilişkisi olduğunda kullanın
- Çoklu kalıtımı dikkatli kullanın, elmas probleminden kaçının

```go
// Kötü: Gereksiz kalıtım
class Rectangle {
    protected:
        int width
        int height
}

class Square : Rectangle {
    // Kare, dikdörtgenin özel bir durumu olabilir,
    // ancak bu tasarım Liskov Substitution Principle'ı ihlal eder
}

// İyi: Kompozisyon
class Shape {
    public:
        virtual int area() = 0
}

class Rectangle : Shape {
    private:
        int width
        int height
        
    public:
        int area() {
            return width * height
        }
}

class Square : Shape {
    private:
        int side
        
    public:
        int area() {
            return side * side
        }
}
```

## Şablonlar

### Şablon Kullanımı

- Şablonları yalnızca gerektiğinde kullanın
- Şablon parametrelerini anlamlı isimlerle adlandırın
- Şablon kısıtlamalarını kullanarak tip güvenliğini sağlayın
- Şablon özelleştirmelerini kullanarak belirli tipler için optimize edilmiş implementasyonlar sağlayın

```go
// Genel şablon
template<T>
class Stack {
    private:
        T[] items
        
    public:
        void push(T item) {
            // Kod
        }
        
        T pop() {
            // Kod
        }
}

// Belirli bir tip için özelleştirme
template<>
class Stack<bool> {
    private:
        uint64[] bits
        
    public:
        void push(bool item) {
            // Bit manipülasyonu ile optimize edilmiş kod
        }
        
        bool pop() {
            // Bit manipülasyonu ile optimize edilmiş kod
        }
}
```

### Şablon Metaprogramlama

- Karmaşık şablon metaprogramlamadan kaçının
- Şablon metaprogramlamayı yalnızca gerçekten gerektiğinde kullanın
- Şablon metaprogramlama kodunu açıklayıcı yorumlarla belgelendirin

```go
// Derleme zamanında faktöriyel hesaplama
template<int N>
struct Factorial {
    static const int value = N * Factorial<N-1>::value
}

// Temel durum
template<>
struct Factorial<0> {
    static const int value = 1
}

// Kullanım
const int fact5 = Factorial<5>::value // 120
```

## İstisna İşleme

### İstisna Kullanımı

- İstisnaları yalnızca gerçekten istisnai durumlar için kullanın
- Akış kontrolü için istisnaları kullanmaktan kaçının
- Özel istisna sınıfları oluşturarak hata türlerini ayırın
- İstisna mesajlarını açıklayıcı ve yardımcı olacak şekilde yazın

```go
// Özel istisna sınıfları
class DatabaseException : Exception {
    public:
        DatabaseException(string message) : Exception(message) {}
}

class ConnectionException : DatabaseException {
    public:
        ConnectionException(string message) : DatabaseException(message) {}
}

// İstisna fırlatma
func connectToDatabase(string connectionString) {
    if connectionString == "" {
        throw new IllegalArgumentException("Connection string cannot be empty")
    }
    
    if !isServerReachable(connectionString) {
        throw new ConnectionException("Cannot reach database server")
    }
    
    // Bağlantı kodu
}
```

### İstisna Yakalama

- İstisnaları mümkün olduğunca spesifik olarak yakalayın
- Boş catch bloklarından kaçının
- Kaynakları temizlemek için finally bloklarını kullanın
- İstisnaları yeniden fırlatmadan önce ek bilgi ekleyin

```go
func processFile(string filename) {
    file := null
    
    try {
        file = openFile(filename)
        processData(file)
    } catch (FileNotFoundException e) {
        log.error("File not found: " + filename)
        throw e
    } catch (IOException e) {
        log.error("Error processing file: " + filename)
        throw new ApplicationException("Data processing failed", e)
    } finally {
        if file != null {
            file.close()
        }
    }
}
```

## Bellek Yönetimi

### Otomatik Bellek Yönetimi

- Mümkün olduğunca otomatik bellek yönetimini (garbage collection) kullanın
- Döngüsel referanslardan kaçının
- Büyük nesneleri kullandıktan sonra null'a ayarlayın
- Bellek sızıntılarını önlemek için kaynakları doğru şekilde kapatın

```go
func processLargeData() {
    data := loadLargeData()
    result := processData(data)
    
    // İşlem tamamlandıktan sonra büyük veriyi serbest bırak
    data = null
    
    return result
}
```

### Manuel Bellek Yönetimi

- Manuel bellek yönetimini yalnızca performans kritik bölümlerde kullanın
- Bellek ayırma ve serbest bırakma işlemlerini eşleştirin
- Bellek sızıntılarını önlemek için defer kullanın
- Dangling pointer'lardan kaçının

```go
func processImage(Image image) {
    unsafe {
        // Manuel bellek ayırma
        buffer := allocate<byte>(image.width * image.height * 4)
        defer free(buffer) // İşlev sonunda belleği serbest bırak
        
        // Performans kritik işlemler
        // ...
    }
}
```

## Eşzamanlılık

### Goroutine Kullanımı

- Goroutine'leri hafif tutun
- Çok sayıda goroutine oluşturmaktan çekinmeyin, ancak kontrolsüz çoğalmalarına izin vermeyin
- Goroutine'lerin tamamlanmasını beklemek için WaitGroup kullanın
- Goroutine sızıntılarından kaçının

```go
func processItems(items []Item) []Result {
    results := make([]Result, len(items))
    var wg sync.WaitGroup
    
    for i, item := range items {
        wg.Add(1)
        go func(i int, item Item) {
            defer wg.Done()
            results[i] = processItem(item)
        }(i, item)
    }
    
    wg.Wait()
    return results
}
```

### Kanal Kullanımı

- Veri paylaşımı için kanalları kullanın
- Kanal boyutunu dikkatli seçin (buffered vs unbuffered)
- Kanal kapanmasını doğru şekilde yönetin
- Deadlock'lardan kaçınmak için timeout kullanın

```go
func processItemsWithChannel(items []Item) []Result {
    numWorkers := runtime.NumCPU()
    jobs := make(chan Item, len(items))
    results := make(chan Result, len(items))
    
    // Worker'ları başlat
    for w := 0; w < numWorkers; w++ {
        go worker(jobs, results)
    }
    
    // İşleri gönder
    for _, item := range items {
        jobs <- item
    }
    close(jobs)
    
    // Sonuçları topla
    allResults := []Result{}
    for i := 0; i < len(items); i++ {
        allResults = append(allResults, <-results)
    }
    
    return allResults
}

func worker(jobs <-chan Item, results chan<- Result) {
    for item := range jobs {
        results <- processItem(item)
    }
}
```

### Mutex ve Senkronizasyon

- Paylaşılan verileri korumak için mutex kullanın
- Mutex'leri mümkün olduğunca kısa süre kilitleyin
- Deadlock'lardan kaçınmak için mutex'leri her zaman aynı sırayla kilitleyin
- RWMutex kullanarak okuma/yazma işlemlerini optimize edin

```go
class SafeCounter {
    private:
        int value
        sync.Mutex mutex
        
    public:
        void increment() {
            this.mutex.Lock()
            defer this.mutex.Unlock()
            this.value++
        }
        
        int getValue() {
            this.mutex.Lock()
            defer this.mutex.Unlock()
            return this.value
        }
}
```

## Performans Optimizasyonları

### Genel Optimizasyonlar

- Erken optimizasyondan kaçının
- Optimizasyon yapmadan önce profilleme yapın
- Algoritma ve veri yapısı seçimini optimize edin
- Bellek ayırmalarını minimize edin

```go
// Kötü: Gereksiz bellek ayırma
func concatenateStrings(strings []string) string {
    result := ""
    for _, s := range strings {
        result += s // Her döngüde yeni bir string oluşturur
    }
    return result
}

// İyi: StringBuilder kullanımı
func concatenateStrings(strings []string) string {
    var builder strings.Builder
    for _, s := range strings {
        builder.WriteString(s)
    }
    return builder.String()
}
```

### Döngü Optimizasyonları

- Döngü içinde bellek ayırmaktan kaçının
- Döngü değişkenlerini döngü dışında tanımlayın
- Döngü koşullarını basit tutun
- Döngü içinde fonksiyon çağrılarını minimize edin

```go
// Kötü: Döngü içinde bellek ayırma
func processItems(items []Item) {
    for i := 0; i < len(items); i++ {
        result := make([]byte, 1024) // Her döngüde yeni bellek ayırma
        processItem(items[i], result)
    }
}

// İyi: Döngü dışında bellek ayırma
func processItems(items []Item) {
    result := make([]byte, 1024) // Bir kez bellek ayırma
    for i := 0; i < len(items); i++ {
        processItem(items[i], result)
    }
}
```

### Bellek Optimizasyonları

- Veri yapılarının boyutunu ve düzenini optimize edin
- Gereksiz kopyalamalardan kaçının
- Büyük nesneler için pointer kullanın
- Bellek havuzları kullanarak bellek ayırma maliyetini azaltın

```go
// Kötü: Gereksiz kopyalama
func processLargeStruct(data LargeStruct) LargeStruct {
    // data'nın bir kopyası oluşturulur
    data.field = newValue
    return data // Dönüş değeri olarak bir kopya daha oluşturulur
}

// İyi: Pointer kullanımı
func processLargeStruct(data *LargeStruct) {
    // Doğrudan orijinal veri üzerinde çalışır
    data.field = newValue
}
```

## Paket ve Modül Organizasyonu

### Paket Yapısı

- Paketleri tek bir sorumluluğa sahip olacak şekilde tasarlayın
- Paket isimlerini anlamlı ve kısa tutun
- Döngüsel bağımlılıklardan kaçının
- İç implementasyon detaylarını gizleyin

```
myapp/
├── main.gom
├── config/
│   ├── config.gom
│   └── loader.gom
├── models/
│   ├── user.gom
│   └── product.gom
├── services/
│   ├── auth.gom
│   └── payment.gom
└── utils/
    ├── logger.gom
    └── helpers.gom
```

### İçe Aktarma Yönetimi

- Kullanılmayan içe aktarmaları kaldırın
- İçe aktarma takma adlarını yalnızca gerektiğinde kullanın
- Döngüsel içe aktarmalardan kaçının
- İçe aktarmaları gruplandırın ve sıralayın

```go
package main

import (
    // Standart kütüphane
    "fmt"
    "os"
    
    // Üçüncü taraf
    "github.com/example/package"
    
    // Yerel
    "myapp/config"
    "myapp/models"
    "myapp/services"
)
```

### API Tasarımı

- Temiz ve tutarlı API'ler tasarlayın
- Gereksiz karmaşıklıktan kaçının
- Kullanıcı dostu hata mesajları sağlayın
- API'nizi kapsamlı bir şekilde belgelendirin

```go
// Kötü: Karmaşık API
func processData(data []byte, options map[string]interface{}, callback func([]byte) error) ([]byte, error)

// İyi: Temiz API
class DataProcessor {
    public:
        // Yapılandırma için bir yapı kullanın
        struct Options {
            bool validateInput
            int timeout
            string format
        }
        
        // Temiz ve anlaşılır API
        Result process(Data data, Options options)
}
```

## Test ve Hata Ayıklama

### Birim Testleri

- Her paket için birim testleri yazın
- Testleri anlamlı şekilde adlandırın
- Test senaryolarını kapsamlı tutun
- Test verilerini test kodundan ayırın

```go
// math/calculator_test.gom
package math

import "testing"

func TestAdd_PositiveNumbers_ReturnsSum(t *testing.T) {
    // Arrange
    calculator := Calculator()
    
    // Act
    result := calculator.add(2, 3)
    
    // Assert
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}

func TestAdd_NegativeNumbers_ReturnsSum(t *testing.T) {
    // Arrange
    calculator := Calculator()
    
    // Act
    result := calculator.add(-2, -3)
    
    // Assert
    if result != -5 {
        t.Errorf("Expected -5, got %d", result)
    }
}
```

### Entegrasyon Testleri

- Bileşenler arasındaki etkileşimleri test edin
- Test ortamını gerçek ortama benzer şekilde kurun
- Dış bağımlılıkları mock'layın
- Test sonrası temizlik yapın

```go
// integration_test.gom
package integration

import "testing"

func TestUserService_RegisterUser_StoresInDatabase(t *testing.T) {
    // Test veritabanı kurulumu
    db := setupTestDatabase()
    defer cleanupTestDatabase(db)
    
    // Servis oluşturma
    userService := UserService(db)
    
    // Test
    user := User{name: "Test User", email: "test@example.com"}
    err := userService.registerUser(user)
    
    // Doğrulama
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    // Veritabanından kullanıcıyı kontrol et
    storedUser, err := db.getUserByEmail("test@example.com")
    if err != nil {
        t.Errorf("Failed to retrieve user: %v", err)
    }
    
    if storedUser.name != user.name {
        t.Errorf("Expected name %s, got %s", user.name, storedUser.name)
    }
}
```

### Hata Ayıklama

- Anlamlı log mesajları kullanın
- Hata ayıklama sembollerini dahil edin
- Hata ayıklama araçlarını (gomdebug) etkin bir şekilde kullanın
- Karmaşık kod için adım adım hata ayıklama yapın

```go
func processTransaction(transaction Transaction) {
    log.debug("Processing transaction: %v", transaction.id)
    
    if transaction.amount <= 0 {
        log.error("Invalid transaction amount: %v", transaction.amount)
        return
    }
    
    log.debug("Validating transaction...")
    if !validateTransaction(transaction) {
        log.error("Transaction validation failed")
        return
    }
    
    log.debug("Executing transaction...")
    result := executeTransaction(transaction)
    
    log.info("Transaction completed: %v, result: %v", transaction.id, result)
}
```

## Dokümantasyon

### Kod Dokümantasyonu

- Tüm paketleri, fonksiyonları ve karmaşık kod bloklarını belgelendirin
- Dokümantasyonu güncel tutun
- Örnekler ve kullanım senaryoları ekleyin
- Parametreleri ve dönüş değerlerini açıklayın

```go
// Package math provides basic mathematical functions.
package math

// Calculator performs basic arithmetic operations.
class Calculator {
    public:
        // Add returns the sum of two integers.
        // Example:
        //   calculator := Calculator()
        //   sum := calculator.add(2, 3) // Returns 5
        int add(int a, int b) {
            return a + b
        }
        
        // Divide returns the quotient of two numbers.
        // It returns an error if the divisor is zero.
        // Example:
        //   calculator := Calculator()
        //   result, err := calculator.divide(10, 2) // Returns 5, nil
        //   result, err := calculator.divide(10, 0) // Returns 0, error
        (float64, error) divide(float64 a, float64 b) {
            if b == 0 {
                return 0, errors.New("division by zero")
            }
            return a / b, nil
        }
}
```

### Proje Dokümantasyonu

- README dosyası oluşturun
- Kurulum ve kullanım talimatları ekleyin
- Katkıda bulunma rehberi ekleyin
- Lisans bilgisi ekleyin

```markdown
# MyApp

MyApp, GO-Minus ile yazılmış bir web uygulamasıdır.

## Kurulum

```bash
git clone https://github.com/example/myapp.git
cd myapp
gominus build
```

## Kullanım

```bash
./myapp --config=config.json
```

## Katkıda Bulunma

Katkıda bulunmak için lütfen [katkı rehberini](CONTRIBUTING.md) okuyun.

## Lisans

MIT
```

### API Dokümantasyonu

- API'nizi kapsamlı bir şekilde belgelendirin
- Endpoint'leri, parametreleri ve dönüş değerlerini açıklayın
- Hata kodlarını ve mesajlarını belgelendirin
- Örnek istekler ve yanıtlar ekleyin

```go
// UserService provides user management functionality.
class UserService {
    public:
        // RegisterUser registers a new user in the system.
        // Parameters:
        //   - user: The user to register
        // Returns:
        //   - error: nil if successful, otherwise an error
        // Errors:
        //   - ErrInvalidEmail: If the email is invalid
        //   - ErrUserExists: If a user with the same email already exists
        error registerUser(User user) {
            // Implementation
        }
        
        // GetUserByID retrieves a user by their ID.
        // Parameters:
        //   - id: The user ID
        // Returns:
        //   - user: The user if found
        //   - error: nil if successful, otherwise an error
        // Errors:
        //   - ErrUserNotFound: If no user with the given ID exists
        (User, error) getUserByID(string id) {
            // Implementation
        }
}
```

Bu en iyi uygulamaları izleyerek, GO-Minus ile daha temiz, daha verimli ve daha bakımı kolay kod yazabilirsiniz. Her proje farklıdır, bu nedenle bu rehberi projenizin özel ihtiyaçlarına göre uyarlayın.

# GO-Minus Buffered IO Örneği

Bu örnek, GO-Minus'un standart kütüphanesinde bulunan Buffered IO (Tamponlanmış Giriş/Çıkış) paketinin kullanımını göstermektedir. Tamponlanmış I/O, performansı artırmak için veri okuma ve yazma işlemlerini gruplar halinde gerçekleştirir.

## Buffered IO Nedir?

Tamponlanmış I/O, disk veya ağ gibi yavaş I/O kaynaklarına erişim sırasında performansı artırmak için kullanılan bir tekniktir. Küçük okuma/yazma işlemlerini daha büyük bloklarda gruplandırarak, sistem çağrılarının sayısını azaltır ve genel performansı artırır.

## Temel BufferedReader Kullanımı

```go
// buffered_reader_basic.gom
package main

import (
    "fmt"
    "io"
    "io/buffered"
    "os"
)

func main() {
    // Dosyayı aç
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Dosya açma hatası:", err)
        return
    }
    defer file.Close()
    
    // BufferedReader oluştur (4KB tampon boyutu)
    reader := buffered.BufferedReader.New(file, 4096)
    
    // Tampondan veri oku
    buffer := make([]byte, 1024)
    totalRead := 0
    
    for {
        n, err := reader.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Okuma hatası:", err)
            break
        }
        
        totalRead += n
        fmt.Printf("Okunan: %d bayt, İçerik: %s\n", n, buffer[:n])
    }
    
    fmt.Printf("Toplam okunan: %d bayt\n", totalRead)
}
```

## Satır Satır Okuma

BufferedReader, satır satır okuma işlemleri için de kullanılabilir:

```go
// buffered_reader_lines.gom
package main

import (
    "fmt"
    "io"
    "io/buffered"
    "os"
)

func main() {
    // Dosyayı aç
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Dosya açma hatası:", err)
        return
    }
    defer file.Close()
    
    // BufferedReader oluştur
    reader := buffered.BufferedReader.New(file, 4096)
    
    // Satır satır oku
    lineCount := 0
    
    for {
        line, err := reader.ReadLine()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Okuma hatası:", err)
            break
        }
        
        lineCount++
        fmt.Printf("Satır %d: %s\n", lineCount, line)
    }
    
    fmt.Printf("Toplam satır sayısı: %d\n", lineCount)
}
```

## Temel BufferedWriter Kullanımı

```go
// buffered_writer_basic.gom
package main

import (
    "fmt"
    "io/buffered"
    "os"
)

func main() {
    // Dosyayı oluştur
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Dosya oluşturma hatası:", err)
        return
    }
    defer file.Close()
    
    // BufferedWriter oluştur (4KB tampon boyutu)
    writer := buffered.BufferedWriter.New(file, 4096)
    
    // Veri yaz
    data := "Bu bir test metnidir.\nBu ikinci satırdır.\nBu üçüncü satırdır.\n"
    n, err := writer.Write([]byte(data))
    if err != nil {
        fmt.Println("Yazma hatası:", err)
        return
    }
    
    fmt.Printf("Yazılan: %d bayt\n", n)
    
    // Tamponu boşalt
    err = writer.Flush()
    if err != nil {
        fmt.Println("Flush hatası:", err)
        return
    }
    
    fmt.Println("Veriler başarıyla yazıldı ve tampon boşaltıldı.")
}
```

## String Yazma

BufferedWriter, string yazma işlemleri için de kullanılabilir:

```go
// buffered_writer_string.gom
package main

import (
    "fmt"
    "io/buffered"
    "os"
)

func main() {
    // Dosyayı oluştur
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Dosya oluşturma hatası:", err)
        return
    }
    defer file.Close()
    
    // BufferedWriter oluştur
    writer := buffered.BufferedWriter.New(file, 4096)
    
    // String yaz
    lines := []string{
        "Bu birinci satırdır.",
        "Bu ikinci satırdır.",
        "Bu üçüncü satırdır.",
    }
    
    for _, line := range lines {
        n, err := writer.WriteString(line + "\n")
        if err != nil {
            fmt.Println("Yazma hatası:", err)
            return
        }
        fmt.Printf("Yazılan: %d bayt\n", n)
    }
    
    // Tamponu boşalt
    err = writer.Flush()
    if err != nil {
        fmt.Println("Flush hatası:", err)
        return
    }
    
    fmt.Println("Veriler başarıyla yazıldı ve tampon boşaltıldı.")
}
```

## Dosya Kopyalama Örneği

Buffered IO'nun pratik bir uygulaması olarak, bir dosyayı başka bir dosyaya kopyalama örneği:

```go
// buffered_copy.gom
package main

import (
    "fmt"
    "io"
    "io/buffered"
    "os"
    "time"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Kullanım: program kaynak_dosya hedef_dosya")
        return
    }
    
    sourcePath := os.Args[1]
    destPath := os.Args[2]
    
    // Kaynak dosyayı aç
    sourceFile, err := os.Open(sourcePath)
    if err != nil {
        fmt.Println("Kaynak dosya açma hatası:", err)
        return
    }
    defer sourceFile.Close()
    
    // Hedef dosyayı oluştur
    destFile, err := os.Create(destPath)
    if err != nil {
        fmt.Println("Hedef dosya oluşturma hatası:", err)
        return
    }
    defer destFile.Close()
    
    // BufferedReader ve BufferedWriter oluştur
    reader := buffered.BufferedReader.New(sourceFile, 8192)
    writer := buffered.BufferedWriter.New(destFile, 8192)
    
    // Kopyalama işlemini başlat
    startTime := time.Now()
    
    buffer := make([]byte, 4096)
    totalBytes := 0
    
    for {
        n, err := reader.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Okuma hatası:", err)
            return
        }
        
        _, err = writer.Write(buffer[:n])
        if err != nil {
            fmt.Println("Yazma hatası:", err)
            return
        }
        
        totalBytes += n
    }
    
    // Tamponu boşalt
    err = writer.Flush()
    if err != nil {
        fmt.Println("Flush hatası:", err)
        return
    }
    
    duration := time.Since(startTime)
    
    fmt.Printf("Kopyalama tamamlandı: %s -> %s\n", sourcePath, destPath)
    fmt.Printf("Toplam: %.2f MB, Süre: %v, Hız: %.2f MB/s\n",
        float64(totalBytes)/1024/1024,
        duration,
        float64(totalBytes)/1024/1024/duration.Seconds())
}
```

## Performans Karşılaştırması

Tamponlanmış ve tamponsuz I/O arasındaki performans farkını gösteren bir örnek:

```go
// buffered_performance.gom
package main

import (
    "fmt"
    "io"
    "io/buffered"
    "os"
    "time"
)

func main() {
    // Test dosyası oluştur
    createTestFile("test.dat", 10*1024*1024) // 10MB
    
    // Tamponsuz kopyalama
    fmt.Println("Tamponsuz kopyalama:")
    unbufferedCopy("test.dat", "unbuffered.dat")
    
    // Tamponlu kopyalama
    fmt.Println("\nTamponlu kopyalama:")
    bufferedCopy("test.dat", "buffered.dat")
    
    // Temizlik
    os.Remove("test.dat")
    os.Remove("unbuffered.dat")
    os.Remove("buffered.dat")
}

// Test dosyası oluştur
func createTestFile(path string, size int) {
    file, err := os.Create(path)
    if err != nil {
        fmt.Println("Dosya oluşturma hatası:", err)
        return
    }
    defer file.Close()
    
    // Rastgele veri yaz
    data := make([]byte, 4096)
    for i := 0; i < len(data); i++ {
        data[i] = byte(i % 256)
    }
    
    remaining := size
    for remaining > 0 {
        chunk := remaining
        if chunk > len(data) {
            chunk = len(data)
        }
        
        n, err := file.Write(data[:chunk])
        if err != nil {
            fmt.Println("Yazma hatası:", err)
            return
        }
        
        remaining -= n
    }
    
    fmt.Printf("Test dosyası oluşturuldu: %s (%.2f MB)\n", path, float64(size)/1024/1024)
}

// Tamponsuz kopyalama
func unbufferedCopy(src, dst string) {
    // Kaynak dosyayı aç
    sourceFile, err := os.Open(src)
    if err != nil {
        fmt.Println("Kaynak dosya açma hatası:", err)
        return
    }
    defer sourceFile.Close()
    
    // Hedef dosyayı oluştur
    destFile, err := os.Create(dst)
    if err != nil {
        fmt.Println("Hedef dosya oluşturma hatası:", err)
        return
    }
    defer destFile.Close()
    
    // Kopyalama işlemini başlat
    startTime := time.Now()
    
    buffer := make([]byte, 4096)
    totalBytes := 0
    
    for {
        n, err := sourceFile.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Okuma hatası:", err)
            return
        }
        
        _, err = destFile.Write(buffer[:n])
        if err != nil {
            fmt.Println("Yazma hatası:", err)
            return
        }
        
        totalBytes += n
    }
    
    duration := time.Since(startTime)
    
    fmt.Printf("Toplam: %.2f MB, Süre: %v, Hız: %.2f MB/s\n",
        float64(totalBytes)/1024/1024,
        duration,
        float64(totalBytes)/1024/1024/duration.Seconds())
}

// Tamponlu kopyalama
func bufferedCopy(src, dst string) {
    // Kaynak dosyayı aç
    sourceFile, err := os.Open(src)
    if err != nil {
        fmt.Println("Kaynak dosya açma hatası:", err)
        return
    }
    defer sourceFile.Close()
    
    // Hedef dosyayı oluştur
    destFile, err := os.Create(dst)
    if err != nil {
        fmt.Println("Hedef dosya oluşturma hatası:", err)
        return
    }
    defer destFile.Close()
    
    // BufferedReader ve BufferedWriter oluştur
    reader := buffered.BufferedReader.New(sourceFile, 65536) // 64KB tampon
    writer := buffered.BufferedWriter.New(destFile, 65536)   // 64KB tampon
    
    // Kopyalama işlemini başlat
    startTime := time.Now()
    
    buffer := make([]byte, 8192)
    totalBytes := 0
    
    for {
        n, err := reader.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Okuma hatası:", err)
            return
        }
        
        _, err = writer.Write(buffer[:n])
        if err != nil {
            fmt.Println("Yazma hatası:", err)
            return
        }
        
        totalBytes += n
    }
    
    // Tamponu boşalt
    err = writer.Flush()
    if err != nil {
        fmt.Println("Flush hatası:", err)
        return
    }
    
    duration := time.Since(startTime)
    
    fmt.Printf("Toplam: %.2f MB, Süre: %v, Hız: %.2f MB/s\n",
        float64(totalBytes)/1024/1024,
        duration,
        float64(totalBytes)/1024/1024/duration.Seconds())
}
```

## Örnek Çıktı

```
Test dosyası oluşturuldu: test.dat (10.00 MB)

Tamponsuz kopyalama:
Toplam: 10.00 MB, Süre: 35.42ms, Hız: 282.33 MB/s

Tamponlu kopyalama:
Toplam: 10.00 MB, Süre: 12.18ms, Hız: 820.98 MB/s
```

Bu örnekler, GO-Minus'un Buffered IO paketinin nasıl kullanılacağını göstermektedir. Daha fazla bilgi için [Buffered IO Belgelendirmesi](../../stdlib/io/buffered/README.md) belgesine bakabilirsiniz.

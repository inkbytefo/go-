# GO-Minus Memory-mapped IO Paketi

Bu paket, GO-Minus programlama dili için belleğe eşlenmiş giriş/çıkış (memory-mapped I/O) işlemlerini sağlar. Memory-mapped I/O, dosyaları doğrudan belleğe eşleyerek, normal dosya I/O işlemlerine göre daha hızlı erişim sağlar.

## Özellikler

- Dosyaları belleğe eşleme
- Dosyaların belirli bölgelerini belleğe eşleme
- Okuma, yazma ve çalıştırma izinleri
- Paylaşımlı ve özel eşleme modları
- Platform bağımsız API (Windows, Unix/Linux)
- Belleğe eşlenmiş dosyaları okuma ve yazma
- Değişiklikleri diske yazma (flush)

## Kullanım

```go
import (
    "io/mmap"
    "os"
)

func main() {
    // Dosyayı aç
    file, err := os.OpenFile("data.bin", os.O_RDWR, 0)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    // Dosyayı belleğe eşle (okuma ve yazma izni ile)
    mmapFile, err := mmap.Map(file, mmap.PROT_READ|mmap.PROT_WRITE, mmap.MAP_SHARED)
    if err != nil {
        panic(err)
    }
    defer mmapFile.Close()
    
    // Belleğe eşlenmiş dosyadan oku
    data := make([]byte, 10)
    n, err := mmapFile.ReadAt(data, 0)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Okunan: %d bayt, İçerik: %v\n", n, data)
    
    // Belleğe eşlenmiş dosyaya yaz
    writeData := []byte{1, 2, 3, 4, 5}
    n, err = mmapFile.WriteAt(writeData, 0)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Yazılan: %d bayt\n", n)
    
    // Değişiklikleri diske yaz
    err = mmapFile.Flush()
    if err != nil {
        panic(err)
    }
    
    // Tüm dosya içeriğini byte dizisi olarak al
    allData := mmapFile.Bytes()
    fmt.Printf("Tüm dosya: %v\n", allData)
}
```

## Dosyanın Belirli Bir Bölgesini Eşleme

```go
import (
    "io/mmap"
    "os"
)

func main() {
    // Dosyayı aç
    file, err := os.OpenFile("large_data.bin", os.O_RDONLY, 0)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    // Dosyanın belirli bir bölgesini belleğe eşle (sadece okuma izni ile)
    offset := int64(1024) // 1KB offset
    length := int64(4096) // 4KB uzunluk
    
    mmapFile, err := mmap.MapRegion(file, length, mmap.PROT_READ, mmap.MAP_SHARED, offset)
    if err != nil {
        panic(err)
    }
    defer mmapFile.Close()
    
    // Belleğe eşlenmiş bölgeden oku
    data := make([]byte, 100)
    n, err := mmapFile.ReadAt(data, 0) // Eşlenmiş bölgenin başından itibaren
    if err != nil {
        panic(err)
    }
    fmt.Printf("Okunan: %d bayt, İçerik: %v\n", n, data)
}
```

## Performans

Memory-mapped I/O, özellikle büyük dosyalarla çalışırken, normal dosya I/O işlemlerine göre önemli performans avantajları sağlar:

1. **Daha Az Sistem Çağrısı**: Dosya belleğe eşlendikten sonra, okuma ve yazma işlemleri için sistem çağrısı gerekmez.
2. **Sayfa Önbelleği**: İşletim sistemi, belleğe eşlenmiş dosyaları sayfa önbelleğinde tutar, bu da tekrarlanan erişimleri hızlandırır.
3. **Sıfır Kopyalama**: Veriler, kullanıcı alanı ve çekirdek alanı arasında kopyalanmadan doğrudan erişilebilir.
4. **Talep Üzerine Sayfalama**: İşletim sistemi, sadece erişilen sayfaları belleğe yükler, bu da büyük dosyalarla çalışırken bellek kullanımını azaltır.

## Sınırlamalar

1. **Dosya Boyutu Değişiklikleri**: Belleğe eşlenmiş bir dosyanın boyutu değiştirilemez. Boyut değişikliği gerekiyorsa, dosya yeniden eşlenmelidir.
2. **Bellek Kullanımı**: Çok büyük dosyaları tamamen belleğe eşlemek, bellek tükenme sorunlarına neden olabilir. Bu durumda, dosyanın sadece belirli bölgelerini eşlemek daha iyidir.
3. **Hata İşleme**: Belleğe eşlenmiş bir dosyaya erişim sırasında oluşan hatalar, normal dosya I/O hatalarından farklı şekilde işlenir (örneğin, segmentation fault).

## Platform Desteği

Bu paket, aşağıdaki platformları destekler:

- Windows
- Unix/Linux
- macOS

Her platform için özel implementasyonlar sağlanmıştır, ancak API platform bağımsızdır.

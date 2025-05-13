# GO-Minus Buffered IO Paketi

Bu paket, GO-Minus programlama dili için tamponlanmış giriş/çıkış (buffered I/O) işlemlerini sağlar. Tamponlanmış I/O, performansı artırmak için veri okuma ve yazma işlemlerini gruplar halinde gerçekleştirir.

## Özellikler

- Tamponlanmış okuma işlemleri için `BufferedReader`
- Tamponlanmış yazma işlemleri için `BufferedWriter`
- Özelleştirilebilir tampon boyutu
- Satır satır okuma desteği
- Byte ve string yazma desteği

## Kullanım

```go
import (
    "io"
    "io/buffered"
)

func main() {
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
}
```

## BufferedReader

`BufferedReader`, bir `io.Reader` üzerinde tamponlanmış okuma işlemleri sağlar. Veriyi büyük bloklarda okur ve istendiğinde küçük parçalar halinde sunar.

### Metotlar

- `New(reader io.Reader, bufferSize int) *BufferedReader`: Yeni bir BufferedReader oluşturur.
- `Read(p []byte) (n int, err error)`: Tampona veri okur.
- `ReadByte() (byte, error)`: Tampondan bir byte okur.
- `ReadLine() (string, error)`: Tampondan bir satır okur.
- `Close() error`: BufferedReader'ı kapatır.

## BufferedWriter

`BufferedWriter`, bir `io.Writer` üzerinde tamponlanmış yazma işlemleri sağlar. Veriyi tamponda biriktirir ve tampon dolduğunda veya `Flush()` çağrıldığında altta yatan Writer'a yazar.

### Metotlar

- `New(writer io.Writer, bufferSize int) *BufferedWriter`: Yeni bir BufferedWriter oluşturur.
- `Write(p []byte) (n int, err error)`: Tampona veri yazar.
- `WriteByte(b byte) error`: Tampona bir byte yazar.
- `WriteString(s string) (n int, err error)`: Tampona bir string yazar.
- `Flush() error`: Tamponu boşaltır.
- `Close() error`: BufferedWriter'ı kapatır.

## Performans

Tamponlanmış I/O, aşağıdaki durumlarda performans avantajı sağlar:

- Küçük boyutlu, sık okuma/yazma işlemleri
- Satır satır okuma işlemleri
- Ağ veya disk I/O işlemleri

Tampon boyutu, performansı etkileyen önemli bir faktördür. Varsayılan tampon boyutu 4096 byte'tır, ancak kullanım senaryonuza göre özelleştirebilirsiniz.

## Sınırlamalar

- Tamponlanmış I/O, bellek kullanımını artırır.
- `Flush()` çağrılmadan önce veri kaybı riski vardır.
- Tampon boyutu çok büyük seçilirse, bellek kullanımı artar; çok küçük seçilirse, performans avantajı azalır.

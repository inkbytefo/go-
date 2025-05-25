# LLVM Toolchain Setup Guide

Bu rehber GO-Minus derleyicisi için gerekli LLVM araçlarının kurulumunu açıklar.

## Gereksinimler

GO-Minus derleyicisi executable üretebilmek için aşağıdaki LLVM araçlarına ihtiyaç duyar:

- **clang**: C/C++ derleyicisi ve linker
- **llc**: LLVM IR'dan assembly/object code üretici
- **opt**: LLVM optimizasyon aracı (opsiyonel)

## Windows Kurulumu

### Yöntem 1: LLVM Resmi İndirme (Önerilen)

1. [LLVM Releases](https://github.com/llvm/llvm-project/releases) sayfasından en son sürümü indirin
2. Windows için pre-built binary'yi seçin (örn: `LLVM-17.0.6-win64.exe`)
3. İndirilen dosyayı çalıştırın ve kurulum sihirbazını takip edin
4. **Önemli**: Kurulum sırasında "Add LLVM to the system PATH" seçeneğini işaretleyin

### Yöntem 2: Chocolatey ile Kurulum

```powershell
# Chocolatey kurulu değilse önce kurun
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

# LLVM'i kurun
choco install llvm
```

### Yöntem 3: Scoop ile Kurulum

```powershell
# Scoop kurulu değilse önce kurun
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
irm get.scoop.sh | iex

# LLVM'i kurun
scoop install llvm
```

### Yöntem 4: Winget ile Kurulum

```powershell
winget install LLVM.LLVM
```

## Linux Kurulumu

### Ubuntu/Debian

```bash
# LLVM repository'sini ekle
wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | sudo apt-key add -
sudo add-apt-repository "deb http://apt.llvm.org/focal/ llvm-toolchain-focal-17 main"

# Paketleri güncelle ve LLVM'i kur
sudo apt update
sudo apt install clang-17 llvm-17 llvm-17-dev

# Symlink'ler oluştur
sudo ln -sf /usr/bin/clang-17 /usr/bin/clang
sudo ln -sf /usr/bin/llc-17 /usr/bin/llc
```

### CentOS/RHEL/Fedora

```bash
# Fedora
sudo dnf install clang llvm llvm-devel

# CentOS/RHEL
sudo yum install clang llvm llvm-devel
```

### Arch Linux

```bash
sudo pacman -S clang llvm
```

## macOS Kurulumu

### Homebrew ile Kurulum (Önerilen)

```bash
# Homebrew kurulu değilse önce kurun
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# LLVM'i kurun
brew install llvm

# PATH'e ekleyin (gerekirse)
echo 'export PATH="/opt/homebrew/opt/llvm/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

### MacPorts ile Kurulum

```bash
sudo port install clang-17 +universal
```

## Kurulum Doğrulama

Kurulum tamamlandıktan sonra aşağıdaki komutları çalıştırarak doğrulayın:

```bash
# Clang versiyonunu kontrol et
clang --version

# LLC versiyonunu kontrol et
llc --version

# Opt versiyonunu kontrol et (opsiyonel)
opt --version
```

Başarılı bir kurulumda şu çıktıları görmelisiniz:

```
clang version 17.0.6
Target: x86_64-pc-windows-msvc
Thread model: posix
```

## GO-Minus ile Test

Kurulum tamamlandıktan sonra GO-Minus ile test edin:

```bash
# Basit bir program derleyin
echo 'package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}' > test.gom

# GO-Minus ile derleyin
gominus test.gom

# Executable oluşturulduğunu kontrol edin
ls -la test.exe  # Windows
ls -la test      # Linux/macOS
```

## Sorun Giderme

### "clang: command not found" Hatası

1. LLVM'in PATH'e eklendiğinden emin olun
2. Terminal/Command Prompt'u yeniden başlatın
3. Sistem değişkenlerini kontrol edin

### Windows'ta PATH Sorunu

1. System Properties > Environment Variables'a gidin
2. PATH değişkenine LLVM bin dizinini ekleyin (örn: `C:\Program Files\LLVM\bin`)
3. Command Prompt'u yeniden başlatın

### Linux'ta Permission Sorunu

```bash
# LLVM binary'lerinin executable olduğundan emin olun
sudo chmod +x /usr/bin/clang
sudo chmod +x /usr/bin/llc
```

## Gelişmiş Konfigürasyon

### Cross-Compilation Desteği

Farklı platformlar için derleme yapmak istiyorsanız:

```bash
# Windows'tan Linux için derleme
gominus -target=linux-x86_64 program.gom

# Linux'tan Windows için derleme  
gominus -target=windows-x86_64 program.gom
```

### Optimizasyon Seviyeleri

```bash
# Debug build (optimizasyon yok)
gominus -O0 program.gom

# Release build (tam optimizasyon)
gominus -O3 program.gom
```

## Sonraki Adımlar

LLVM kurulumu tamamlandıktan sonra:

1. [GO-Minus Tutorial](tutorial/getting-started.md) ile başlayın
2. [Example Programs](../examples/) klasöründeki örnekleri deneyin
3. [Development Guide](development.md) ile geliştirme ortamını kurun

## Destek

Kurulum sorunları için:

- [GitHub Issues](https://github.com/inkbytefo/go-minus/issues)
- [Discord Community](https://discord.gg/go-minus)
- [Documentation](https://go-minus.dev/docs)

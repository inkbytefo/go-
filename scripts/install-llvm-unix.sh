#!/bin/bash

# GO-Minus LLVM Kurulum Scripti - Linux/macOS
# Bu script Unix-like sistemlerde LLVM toolchain'ini otomatik olarak kurar

set -e  # Hata durumunda çık

# Varsayılan değerler
METHOD="auto"
VERSION="17"
FORCE=false
QUIET=false

# Renkli çıktı fonksiyonları
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

success() { echo -e "${GREEN}✅ $1${NC}"; }
error() { echo -e "${RED}❌ $1${NC}"; }
warning() { echo -e "${YELLOW}⚠️  $1${NC}"; }
info() { echo -e "${BLUE}ℹ️  $1${NC}"; }

# Yardım mesajı
show_help() {
    cat << EOF
GO-Minus LLVM Kurulum Scripti

Kullanım: $0 [SEÇENEKLER]

SEÇENEKLER:
    -m, --method METHOD     Kurulum yöntemi (auto, apt, yum, dnf, pacman, brew, port)
    -v, --version VERSION   LLVM versiyonu (varsayılan: 17)
    -f, --force            Mevcut kurulumu zorla güncelle
    -q, --quiet            Sessiz kurulum
    -h, --help             Bu yardım mesajını göster

ÖRNEKLER:
    $0                      # Otomatik kurulum
    $0 -m brew              # Homebrew ile kurulum
    $0 -v 16 -f             # LLVM 16'yı zorla kur
    $0 -q                   # Sessiz kurulum

EOF
}

# Parametreleri parse et
while [[ $# -gt 0 ]]; do
    case $1 in
        -m|--method)
            METHOD="$2"
            shift 2
            ;;
        -v|--version)
            VERSION="$2"
            shift 2
            ;;
        -f|--force)
            FORCE=true
            shift
            ;;
        -q|--quiet)
            QUIET=true
            shift
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        *)
            error "Bilinmeyen parametre: $1"
            show_help
            exit 1
            ;;
    esac
done

# İşletim sistemi tespiti
detect_os() {
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        if command -v apt-get &> /dev/null; then
            echo "ubuntu"
        elif command -v yum &> /dev/null; then
            echo "centos"
        elif command -v dnf &> /dev/null; then
            echo "fedora"
        elif command -v pacman &> /dev/null; then
            echo "arch"
        else
            echo "linux"
        fi
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        echo "macos"
    else
        echo "unknown"
    fi
}

# LLVM kurulu mu kontrol et
check_llvm_installed() {
    if command -v clang &> /dev/null && command -v llc &> /dev/null; then
        local clang_version=$(clang --version | head -n1)
        success "LLVM zaten kurulu: $clang_version"
        return 0
    fi
    return 1
}

# Sudo gerekli mi kontrol et
check_sudo() {
    if [[ $EUID -ne 0 ]] && [[ "$1" == "true" ]]; then
        if ! command -v sudo &> /dev/null; then
            error "Bu kurulum için sudo gerekli ancak sudo bulunamadı"
            exit 1
        fi
        SUDO="sudo"
    else
        SUDO=""
    fi
}

# Ubuntu/Debian kurulumu
install_llvm_apt() {
    info "APT ile LLVM kuruluyor..."
    check_sudo true
    
    # LLVM repository ekle
    if [[ $QUIET == true ]]; then
        wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | $SUDO apt-key add - &>/dev/null
        $SUDO add-apt-repository "deb http://apt.llvm.org/$(lsb_release -cs)/ llvm-toolchain-$(lsb_release -cs)-$VERSION main" -y &>/dev/null
        $SUDO apt update &>/dev/null
        $SUDO apt install -y clang-$VERSION llvm-$VERSION llvm-$VERSION-dev &>/dev/null
    else
        wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | $SUDO apt-key add -
        $SUDO add-apt-repository "deb http://apt.llvm.org/$(lsb_release -cs)/ llvm-toolchain-$(lsb_release -cs)-$VERSION main" -y
        $SUDO apt update
        $SUDO apt install -y clang-$VERSION llvm-$VERSION llvm-$VERSION-dev
    fi
    
    # Symlink'ler oluştur
    $SUDO ln -sf /usr/bin/clang-$VERSION /usr/bin/clang
    $SUDO ln -sf /usr/bin/llc-$VERSION /usr/bin/llc
    $SUDO ln -sf /usr/bin/opt-$VERSION /usr/bin/opt
    
    success "LLVM APT ile başarıyla kuruldu"
}

# CentOS/RHEL kurulumu
install_llvm_yum() {
    info "YUM ile LLVM kuruluyor..."
    check_sudo true
    
    if [[ $QUIET == true ]]; then
        $SUDO yum install -y clang llvm llvm-devel &>/dev/null
    else
        $SUDO yum install -y clang llvm llvm-devel
    fi
    
    success "LLVM YUM ile başarıyla kuruldu"
}

# Fedora kurulumu
install_llvm_dnf() {
    info "DNF ile LLVM kuruluyor..."
    check_sudo true
    
    if [[ $QUIET == true ]]; then
        $SUDO dnf install -y clang llvm llvm-devel &>/dev/null
    else
        $SUDO dnf install -y clang llvm llvm-devel
    fi
    
    success "LLVM DNF ile başarıyla kuruldu"
}

# Arch Linux kurulumu
install_llvm_pacman() {
    info "Pacman ile LLVM kuruluyor..."
    check_sudo true
    
    if [[ $QUIET == true ]]; then
        $SUDO pacman -S --noconfirm clang llvm &>/dev/null
    else
        $SUDO pacman -S --noconfirm clang llvm
    fi
    
    success "LLVM Pacman ile başarıyla kuruldu"
}

# Homebrew kurulumu (macOS)
install_llvm_brew() {
    info "Homebrew ile LLVM kuruluyor..."
    
    # Homebrew kurulu mu kontrol et
    if ! command -v brew &> /dev/null; then
        warning "Homebrew kurulu değil, kuruluyor..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    fi
    
    if [[ $QUIET == true ]]; then
        brew install llvm &>/dev/null
    else
        brew install llvm
    fi
    
    # PATH'e ekle
    local llvm_path="/opt/homebrew/opt/llvm/bin"
    if [[ -d "$llvm_path" ]]; then
        if ! grep -q "$llvm_path" ~/.zshrc 2>/dev/null; then
            echo "export PATH=\"$llvm_path:\$PATH\"" >> ~/.zshrc
            export PATH="$llvm_path:$PATH"
        fi
    fi
    
    success "LLVM Homebrew ile başarıyla kuruldu"
}

# MacPorts kurulumu (macOS)
install_llvm_port() {
    info "MacPorts ile LLVM kuruluyor..."
    check_sudo true
    
    if ! command -v port &> /dev/null; then
        error "MacPorts kurulu değil. Lütfen önce MacPorts'u kurun."
        exit 1
    fi
    
    if [[ $QUIET == true ]]; then
        $SUDO port install clang-$VERSION +universal &>/dev/null
    else
        $SUDO port install clang-$VERSION +universal
    fi
    
    success "LLVM MacPorts ile başarıyla kuruldu"
}

# Ana kurulum fonksiyonu
install_llvm() {
    info "GO-Minus LLVM Kurulum Scripti"
    info "Kurulum yöntemi: $METHOD"
    info "LLVM versiyonu: $VERSION"
    
    # Mevcut kurulum kontrolü
    if check_llvm_installed && [[ $FORCE == false ]]; then
        warning "LLVM zaten kurulu. Zorla güncellemek için -f parametresini kullanın."
        return 0
    fi
    
    local os=$(detect_os)
    info "İşletim sistemi: $os"
    
    # Kurulum yöntemini belirle
    if [[ $METHOD == "auto" ]]; then
        case $os in
            ubuntu) METHOD="apt" ;;
            centos) METHOD="yum" ;;
            fedora) METHOD="dnf" ;;
            arch) METHOD="pacman" ;;
            macos) METHOD="brew" ;;
            *) error "Desteklenmeyen işletim sistemi: $os"; exit 1 ;;
        esac
        info "Otomatik seçilen yöntem: $METHOD"
    fi
    
    # Kurulum yöntemine göre kur
    case $METHOD in
        apt) install_llvm_apt ;;
        yum) install_llvm_yum ;;
        dnf) install_llvm_dnf ;;
        pacman) install_llvm_pacman ;;
        brew) install_llvm_brew ;;
        port) install_llvm_port ;;
        *) error "Geçersiz kurulum yöntemi: $METHOD"; exit 1 ;;
    esac
    
    # Kurulumu doğrula
    info "Kurulum doğrulanıyor..."
    sleep 2
    
    if check_llvm_installed; then
        success "LLVM kurulumu başarıyla doğrulandı!"
        info "Şimdi GO-Minus ile executable üretebilirsiniz."
        return 0
    else
        warning "LLVM kuruldu ancak PATH'de bulunamıyor. Terminal'i yeniden başlatın."
        return 1
    fi
}

# GO-Minus ile test
test_llvm_with_gominus() {
    info "GO-Minus ile LLVM testi yapılıyor..."
    
    local test_file="/tmp/test_llvm.gom"
    cat > "$test_file" << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("LLVM kurulumu başarılı!")
}
EOF
    
    if command -v gominus &> /dev/null; then
        if gominus "$test_file" &>/dev/null; then
            success "GO-Minus ile LLVM testi başarılı!"
        else
            warning "GO-Minus ile derleme başarısız. Daha fazla bilgi için 'gominus --help' çalıştırın."
        fi
    else
        info "GO-Minus kurulu değil. Önce 'go build ./cmd/gominus' çalıştırın."
    fi
    
    rm -f "$test_file"
}

# Ana script
main() {
    if install_llvm; then
        echo
        success "🎉 LLVM kurulumu tamamlandı!"
        info "📖 Kullanım rehberi: docs/llvm-setup.md"
        info "🚀 GO-Minus örnekleri: examples/"
        echo
        
        test_llvm_with_gominus
    else
        error "Kurulum başarısız! Lütfen docs/llvm-setup.md dosyasını kontrol edin."
        exit 1
    fi
}

# Script'i çalıştır
main

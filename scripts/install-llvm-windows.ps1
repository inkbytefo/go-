# GO-Minus LLVM Kurulum Scripti - Windows
# Bu script Windows sistemlerde LLVM toolchain'ini otomatik olarak kurar

param(
    [string]$Method = "winget",  # winget, chocolatey, scoop, manual
    [string]$Version = "17",     # LLVM version
    [switch]$Force = $false,     # Mevcut kurulumu zorla güncelle
    [switch]$Quiet = $false      # Sessiz kurulum
)

# Renkli çıktı için fonksiyonlar
function Write-Success { param($Message) Write-Host "✅ $Message" -ForegroundColor Green }
function Write-Error { param($Message) Write-Host "❌ $Message" -ForegroundColor Red }
function Write-Warning { param($Message) Write-Host "⚠️  $Message" -ForegroundColor Yellow }
function Write-Info { param($Message) Write-Host "ℹ️  $Message" -ForegroundColor Cyan }

# Admin yetkisi kontrolü
function Test-Administrator {
    $currentUser = [Security.Principal.WindowsIdentity]::GetCurrent()
    $principal = New-Object Security.Principal.WindowsPrincipal($currentUser)
    return $principal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
}

# LLVM kurulu mu kontrol et
function Test-LLVMInstalled {
    try {
        $clangVersion = & clang --version 2>$null
        $llcVersion = & llc --version 2>$null

        if ($clangVersion -and $llcVersion) {
            Write-Success "LLVM zaten kurulu: $($clangVersion.Split("`n")[0])"
            return $true
        }
    }
    catch {
        return $false
    }
    return $false
}

# PATH'e LLVM ekle
function Add-LLVMToPath {
    param($LLVMPath)

    $currentPath = [Environment]::GetEnvironmentVariable("PATH", "Machine")
    if ($currentPath -notlike "*$LLVMPath*") {
        Write-Info "LLVM PATH'e ekleniyor: $LLVMPath"
        [Environment]::SetEnvironmentVariable("PATH", "$currentPath;$LLVMPath", "Machine")
        $env:PATH = "$env:PATH;$LLVMPath"
        Write-Success "LLVM PATH'e eklendi"
    }
}

# Winget ile kurulum
function Install-LLVMWithWinget {
    Write-Info "Winget ile LLVM kuruluyor..."

    try {
        if ($Quiet) {
            & winget install LLVM.LLVM --silent --accept-package-agreements --accept-source-agreements
        } else {
            & winget install LLVM.LLVM --accept-package-agreements --accept-source-agreements
        }

        if ($LASTEXITCODE -eq 0) {
            Write-Success "LLVM winget ile başarıyla kuruldu"
            Add-LLVMToPath "C:\Program Files\LLVM\bin"
            return $true
        }
    }
    catch {
        Write-Error "Winget kurulumu başarısız: $($_.Exception.Message)"
    }
    return $false
}

# Chocolatey ile kurulum
function Install-LLVMWithChocolatey {
    Write-Info "Chocolatey ile LLVM kuruluyor..."

    # Chocolatey kurulu mu kontrol et
    try {
        & choco --version | Out-Null
    }
    catch {
        Write-Warning "Chocolatey kurulu değil, kuruluyor..."
        Set-ExecutionPolicy Bypass -Scope Process -Force
        [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072
        iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
    }

    try {
        if ($Quiet) {
            & choco install llvm -y
        } else {
            & choco install llvm
        }

        if ($LASTEXITCODE -eq 0) {
            Write-Success "LLVM Chocolatey ile başarıyla kuruldu"
            return $true
        }
    }
    catch {
        Write-Error "Chocolatey kurulumu başarısız: $($_.Exception.Message)"
    }
    return $false
}

# Scoop ile kurulum
function Install-LLVMWithScoop {
    Write-Info "Scoop ile LLVM kuruluyor..."

    # Scoop kurulu mu kontrol et
    try {
        & scoop --version | Out-Null
    }
    catch {
        Write-Warning "Scoop kurulu değil, kuruluyor..."
        Set-ExecutionPolicy RemoteSigned -Scope CurrentUser -Force
        irm get.scoop.sh | iex
    }

    try {
        & scoop install llvm

        if ($LASTEXITCODE -eq 0) {
            Write-Success "LLVM Scoop ile başarıyla kuruldu"
            return $true
        }
    }
    catch {
        Write-Error "Scoop kurulumu başarısız: $($_.Exception.Message)"
    }
    return $false
}

# Manuel kurulum
function Install-LLVMManually {
    Write-Info "Manuel LLVM kurulumu başlatılıyor..."

    $downloadUrl = "https://github.com/llvm/llvm-project/releases/download/llvmorg-$Version.0.6/LLVM-$Version.0.6-win64.exe"
    $installerPath = "$env:TEMP\LLVM-$Version.0.6-win64.exe"

    try {
        Write-Info "LLVM indiriliyor: $downloadUrl"
        Invoke-WebRequest -Uri $downloadUrl -OutFile $installerPath -UseBasicParsing

        Write-Info "LLVM kuruluyor..."
        if ($Quiet) {
            & $installerPath /S
        } else {
            & $installerPath
        }

        # Kurulum tamamlanana kadar bekle
        Start-Sleep -Seconds 10

        # Installer dosyasını sil
        Remove-Item $installerPath -Force -ErrorAction SilentlyContinue

        Add-LLVMToPath "C:\Program Files\LLVM\bin"
        Write-Success "LLVM manuel olarak başarıyla kuruldu"
        return $true
    }
    catch {
        Write-Error "Manuel kurulum başarısız: $($_.Exception.Message)"
        Remove-Item $installerPath -Force -ErrorAction SilentlyContinue
    }
    return $false
}

# Ana kurulum fonksiyonu
function Install-LLVM {
    Write-Info "GO-Minus LLVM Kurulum Scripti"
    Write-Info "Kurulum yöntemi: $Method"
    Write-Info "LLVM versiyonu: $Version"

    # Mevcut kurulum kontrolü
    if ((Test-LLVMInstalled) -and (-not $Force)) {
        Write-Warning "LLVM zaten kurulu. Zorla güncellemek için -Force parametresini kullanın."
        return $true
    }

    # Admin yetkisi kontrolü (bazı kurulum yöntemleri için gerekli)
    if ($Method -eq "manual" -and (-not (Test-Administrator))) {
        Write-Warning "Manuel kurulum için admin yetkisi gerekli. PowerShell'i admin olarak çalıştırın."
    }

    # Kurulum yöntemine göre kur
    $success = $false
    switch ($Method.ToLower()) {
        "winget" { $success = Install-LLVMWithWinget }
        "chocolatey" { $success = Install-LLVMWithChocolatey }
        "scoop" { $success = Install-LLVMWithScoop }
        "manual" { $success = Install-LLVMManually }
        default {
            Write-Error "Geçersiz kurulum yöntemi: $Method"
            Write-Info "Geçerli yöntemler: winget, chocolatey, scoop, manual"
            return $false
        }
    }

    if ($success) {
        Write-Success "LLVM kurulumu tamamlandı!"

        # Kurulumu doğrula
        Write-Info "Kurulum doğrulanıyor..."
        Start-Sleep -Seconds 3

        if (Test-LLVMInstalled) {
            Write-Success "LLVM kurulumu başarıyla doğrulandı!"
            Write-Info "Şimdi GO-Minus ile executable üretebilirsiniz."
            return $true
        } else {
            Write-Warning "LLVM kuruldu ancak PATH'de bulunamıyor. Terminal'i yeniden başlatın."
            return $false
        }
    } else {
        Write-Error "LLVM kurulumu başarısız!"
        return $false
    }
}

# Test fonksiyonu
function Test-LLVMWithGoMinus {
    Write-Info "GO-Minus ile LLVM testi yapılıyor..."

    $testFile = "$env:TEMP\test_llvm.gom"
    $testContent = @"
package main

import "fmt"

func main() {
    fmt.Println("LLVM kurulumu başarılı!")
}
"@

    try {
        # Test dosyası oluştur
        Set-Content -Path $testFile -Value $testContent

        # GO-Minus ile derle
        & gominus $testFile

        if ($LASTEXITCODE -eq 0) {
            Write-Success "GO-Minus ile LLVM testi başarılı!"
        } else {
            Write-Warning "GO-Minus ile derleme başarısız. Daha fazla bilgi için 'gominus --help' çalıştırın."
        }

        # Test dosyasını sil
        Remove-Item $testFile -Force -ErrorAction SilentlyContinue
    }
    catch {
        Write-Error "Test sırasında hata: $($_.Exception.Message)"
    }
}

# Script'i çalıştır
if ($MyInvocation.InvocationName -ne '.') {
    $result = Install-LLVM

    if ($result) {
        Write-Info ""
        Write-Info "🎉 LLVM kurulumu tamamlandı!"
        Write-Info "📖 Kullanım rehberi: docs/llvm-setup.md"
        Write-Info "🚀 GO-Minus örnekleri: examples/"
        Write-Info ""

        # GO-Minus kurulu ise test yap
        try {
            & gominus --version | Out-Null
            Test-LLVMWithGoMinus
        }
        catch {
            Write-Info "GO-Minus kurulu değil. Önce 'go build ./cmd/gominus' çalıştırın."
        }
    } else {
        Write-Error "Kurulum basarisiz! Lutfen docs/llvm-setup.md dosyasini kontrol edin."
        exit 1
    }
}

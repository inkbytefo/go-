#!/bin/bash

# GO-Minus Derleyici Test Çalıştırıcısı

# Renk kodları
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

# Test dizini
TEST_DIR="$(dirname "$0")"
cd "$TEST_DIR" || exit 1

# GO-Minus derleyicisi
GOMINUS_COMPILER="../../bin/gominus"

# Hata ayıklama bilgisi üretimini etkinleştir
DEBUG_FLAG="--debug"

# Sonuçlar
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# Fonksiyon: Test dosyasını derle ve çalıştır
run_test() {
    local test_file="$1"
    local test_name="$(basename "$test_file" .gom)"
    local output_file="${test_name}.out"

    echo -e "${YELLOW}Test: $test_name${NC}"

    # Dosyayı derle
    echo "  Derleniyor..."
    "$GOMINUS_COMPILER" $DEBUG_FLAG "$test_file" -o "$output_file"
    local compile_result=$?

    if [ $compile_result -ne 0 ]; then
        echo -e "  ${RED}Derleme hatası!${NC}"
        FAILED_TESTS=$((FAILED_TESTS + 1))
        return
    fi

    # Derlenmiş dosyayı çalıştır
    echo "  Çalıştırılıyor..."
    ./"$output_file"
    local run_result=$?

    if [ $run_result -ne 0 ]; then
        echo -e "  ${RED}Çalıştırma hatası!${NC}"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    else
        echo -e "  ${GREEN}Başarılı!${NC}"
        PASSED_TESTS=$((PASSED_TESTS + 1))
    fi

    # Temizlik
    rm -f "$output_file"

    TOTAL_TESTS=$((TOTAL_TESTS + 1))
}

# Tüm test dosyalarını bul ve çalıştır
echo -e "${YELLOW}GO-Minus Derleyici Testleri${NC}"
echo "=============================="

for test_file in *.gom; do
    if [ -f "$test_file" ]; then
        run_test "$test_file"
        echo ""
    fi
done

# Sonuçları göster
echo "=============================="
echo -e "${YELLOW}Test Sonuçları:${NC}"
echo -e "  Toplam Test: $TOTAL_TESTS"
echo -e "  ${GREEN}Başarılı: $PASSED_TESTS${NC}"
echo -e "  ${RED}Başarısız: $FAILED_TESTS${NC}"

# Başarı durumuna göre çıkış kodu
if [ $FAILED_TESTS -eq 0 ]; then
    echo -e "${GREEN}Tüm testler başarılı!${NC}"
    exit 0
else
    echo -e "${RED}Bazı testler başarısız!${NC}"
    exit 1
fi

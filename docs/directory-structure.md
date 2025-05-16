# GO-Minus Proje Dizin Yapısı

Bu belge, GO-Minus projesinin dizin yapısını ve her bir dizinin amacını açıklar.

## Ana Dizinler

```
go-minus/
├── cmd/                      # Komut satırı uygulamaları
│   ├── gominus/              # GO-Minus derleyicisi
│   ├── gompm/                # GO-Minus paket yöneticisi
│   ├── gomfmt/               # GO-Minus kod biçimlendirme aracı
│   ├── gomdoc/               # GO-Minus belgelendirme aracı
│   ├── gomtest/              # GO-Minus test aracı
│   └── gomlsp/               # GO-Minus dil sunucusu
├── docs/                     # Belgelendirme
│   ├── tutorial/             # Öğreticiler
│   ├── reference/            # Referans belgeleri
│   ├── images/               # Belgelendirme görselleri
│   └── ...                   # Diğer belgelendirme dosyaları
├── examples/                 # Örnek kodlar
│   ├── basic/                # Temel örnekler
│   ├── advanced/             # İleri düzey örnekler
│   ├── vulkan/               # Vulkan örnekleri
│   └── memory/               # Bellek yönetimi örnekleri
├── internal/                 # Dahili paketler
│   ├── ast/                  # Soyut sözdizimi ağacı
│   ├── codegen/              # Kod üretimi
│   ├── irgen/                # IR üretimi
│   ├── lexer/                # Sözcük analizi
│   ├── optimizer/            # Optimizasyon
│   ├── parser/               # Sözdizimi analizi
│   ├── semantic/             # Semantik analiz
│   └── token/                # Token tanımları
├── pkg/                      # Dışa açık paketler
│   ├── compiler/             # Derleyici API'si
│   └── runtime/              # Çalışma zamanı API'si
├── stdlib/                   # Standart kütüphane
│   ├── async/                # Asenkron işlemler
│   ├── concurrent/           # Eşzamanlılık
│   ├── container/            # Veri yapıları
│   ├── core/                 # Çekirdek işlevler
│   ├── fmt/                  # Biçimlendirme
│   ├── io/                   # Giriş/çıkış işlemleri
│   ├── math/                 # Matematiksel işlemler
│   ├── memory/               # Bellek yönetimi
│   ├── net/                  # Ağ işlemleri
│   ├── regex/                # Düzenli ifadeler
│   ├── strings/              # Dize işlemleri
│   ├── time/                 # Zaman işlemleri
│   └── vulkan/               # Vulkan bağlayıcıları
├── tests/                    # Testler
│   ├── compiler/             # Derleyici testleri
│   ├── basic/                # Temel dil özellikleri testleri
│   └── ...                   # Diğer test dosyaları
├── website/                  # Web sitesi
│   ├── content/              # İçerik
│   ├── static/               # Statik dosyalar
│   └── templates/            # Şablonlar
├── .gitignore                # Git dışlama listesi
├── go.mod                    # Go modül tanımı
├── go.sum                    # Go bağımlılık sağlaması
├── LICENSE                   # Lisans dosyası
└── README.md                 # Ana README dosyası
```

## Dizinlerin Açıklamaları

### cmd/

Bu dizin, GO-Minus projesinin komut satırı uygulamalarını içerir. Her bir alt dizin, ayrı bir uygulamayı temsil eder.

- **gominus**: GO-Minus derleyicisi
- **gompm**: GO-Minus paket yöneticisi
- **gomfmt**: GO-Minus kod biçimlendirme aracı
- **gomdoc**: GO-Minus belgelendirme aracı
- **gomtest**: GO-Minus test aracı
- **gomlsp**: GO-Minus dil sunucusu

### docs/

Bu dizin, GO-Minus projesinin belgelendirmesini içerir.

- **tutorial/**: Öğreticiler ve başlangıç rehberleri
- **reference/**: Dil referansı ve API belgeleri
- **images/**: Belgelendirmede kullanılan görseller

### examples/

Bu dizin, GO-Minus dilinde yazılmış örnek kodları içerir.

- **basic/**: Temel dil özelliklerini gösteren basit örnekler
- **advanced/**: İleri düzey dil özelliklerini gösteren karmaşık örnekler
- **vulkan/**: Vulkan API'sini kullanan örnekler
- **memory/**: Bellek yönetimi özelliklerini gösteren örnekler

### internal/

Bu dizin, GO-Minus derleyicisinin dahili bileşenlerini içerir. Bu paketler, projenin dışından doğrudan kullanılmaz.

- **ast/**: Soyut sözdizimi ağacı tanımları ve işlemleri
- **codegen/**: Makine kodu üretimi
- **irgen/**: LLVM IR üretimi
- **lexer/**: Sözcük analizi
- **optimizer/**: Kod optimizasyonu
- **parser/**: Sözdizimi analizi
- **semantic/**: Semantik analiz
- **token/**: Token tanımları

### pkg/

Bu dizin, GO-Minus projesinin dışa açık paketlerini içerir. Bu paketler, diğer projeler tarafından kullanılabilir.

- **compiler/**: Derleyici API'si
- **runtime/**: Çalışma zamanı API'si

### stdlib/

Bu dizin, GO-Minus dilinin standart kütüphanesini içerir.

- **async/**: Asenkron işlemler için paketler
- **concurrent/**: Eşzamanlılık için paketler
- **container/**: Veri yapıları
- **core/**: Çekirdek işlevler
- **fmt/**: Biçimlendirme işlemleri
- **io/**: Giriş/çıkış işlemleri
- **math/**: Matematiksel işlemler
- **memory/**: Bellek yönetimi
- **net/**: Ağ işlemleri
- **regex/**: Düzenli ifadeler
- **strings/**: Dize işlemleri
- **time/**: Zaman işlemleri
- **vulkan/**: Vulkan API bağlayıcıları

### tests/

Bu dizin, GO-Minus projesinin test dosyalarını içerir.

- **compiler/**: Derleyici testleri
- **basic/**: Temel dil özellikleri testleri

### website/

Bu dizin, GO-Minus projesinin web sitesi dosyalarını içerir.

- **content/**: Web sitesi içeriği
- **static/**: Statik dosyalar (CSS, JS, görseller)
- **templates/**: HTML şablonları

## Dosyalar

- **.gitignore**: Git tarafından izlenmeyecek dosyaları belirtir
- **go.mod**: Go modül tanımı
- **go.sum**: Go bağımlılık sağlaması
- **LICENSE**: Lisans dosyası
- **README.md**: Ana README dosyası
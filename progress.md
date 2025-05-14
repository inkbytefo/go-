# GO-Minus Dili Geliştirme İlerleme Raporu

Bu belge, GO-Minus programlama dilinin geliştirme sürecini takip etmek için kullanılmaktadır. Tamamlanan görevler, devam eden çalışmalar ve gelecek planlar burada belgelenecektir.

## Proje Genel Bakış

GO-Minus dili, Go programlama dilinin tüm özelliklerini içeren ve C++ benzeri özelliklerle (sınıflar, şablonlar, istisna işleme vb.) genişletilmiş bir dil olarak tasarlanmıştır. GO-Minus dosyaları `.gom` uzantısını kullanır.

## Tamamlanan Görevler

### Temel Altyapı

- [x] **Token Paketi**: Token türleri ve token yapısı tanımlandı.
  - Token türleri (anahtar kelimeler, operatörler, ayırıcılar vb.) tanımlandı.
  - Token yapısı (tür, değer, satır, sütun, pozisyon) tanımlandı.
  - GO+ için özel token türleri (class, template, throw vb.) eklendi.

- [x] **Lexer Paketi**: Kaynak kodu token'lara ayıran lexer geliştirildi.
  - Temel token tanıma işlevleri eklendi.
  - Yorum, string, sayı vb. özel token türleri için işleme eklendi.
  - Satır ve sütun numaralarını takip etme eklendi.

- [x] **AST Paketi**: Soyut Sözdizimi Ağacı (AST) düğümleri tanımlandı.
  - Temel AST düğüm arayüzleri (Node, Statement, Expression) tanımlandı.
  - İfade düğümleri (Identifier, IntegerLiteral, StringLiteral vb.) tanımlandı.
  - Deyim düğümleri (VarStatement, ReturnStatement, BlockStatement vb.) tanımlandı.
  - GO+ için özel AST düğümleri (ClassStatement, TemplateExpression vb.) tanımlandı.

- [x] **Parser Paketi (Temel)**: Token dizisini AST'ye dönüştüren parser'ın temel yapısı geliştirildi.
  - Recursive descent parser yapısı oluşturuldu.
  - Temel ifade ayrıştırma (expression parsing) eklendi.
  - Operatör öncelik tablosu eklendi.
  - Basit ifadeleri ayrıştırma yeteneği eklendi.

### Testler ve Örnekler

- [x] **Minimal Örnek**: Basit bir ifadeyi ayrıştıran minimal örnek oluşturuldu.
  - `5 + 10` ifadesini ayrıştıran ve AST'yi yazdıran örnek çalıştırıldı.

## Devam Eden Çalışmalar

### Parser Paketi (Gelişmiş)

- [x] **Paket ve Import Bildirimleri**: Paket ve import bildirimlerini ayrıştırma.
  - Paket bildirimi ayrıştırma eklendi.
  - Tek ve çoklu import bildirimi ayrıştırma eklendi.

- [x] **Değişken ve Sabit Tanımlamaları**: Değişken ve sabit tanımlamalarını ayrıştırma.
  - Değişken tanımlama ayrıştırma eklendi.
  - Sabit tanımlama ayrıştırma eklendi.
  - Tip tanımlama desteği eklendi.

- [x] **Fonksiyon Tanımları**: Fonksiyon tanımlarını ayrıştırma.
  - Fonksiyon tanımlama ayrıştırma eklendi.
  - Parametre ve dönüş tipi ayrıştırma eklendi.
  - Metot tanımlama ayrıştırma eklendi.

- [x] **Sınıf ve Şablon Tanımları**: Sınıf ve şablon tanımlarını ayrıştırma.
  - Sınıf tanımlama ayrıştırma eklendi.
  - Şablon tanımlama ayrıştırma eklendi.
  - Kalıtım ve arayüz uygulamaları ayrıştırma eklendi.
  - Erişim belirleyicileri ayrıştırma eklendi.

- [x] **İstisna İşleme**: Try-catch-finally yapılarını ayrıştırma.
  - Try-catch-finally ayrıştırma eklendi.
  - Throw ifadesi ayrıştırma eklendi.

- [x] **Parser İyileştirmeleri**: Parser'ın daha karmaşık ifadeleri ayrıştırabilmesi için iyileştirmeler.
  - This ve super anahtar kelimeleri için ayrıştırma eklendi.
  - Sınıf üye erişimi için ayrıştırma iyileştirildi.
  - Hata kurtarma mekanizmaları geliştirildi.
  - Kısa değişken tanımlama (`:=`) için ayrıştırma eklendi.

## Tamamlanan Görevler

### Semantik Analiz

- [x] **Sembol Tablosu**: Kapsam (scope) yönetimi ve sembol tanımlama/çözümleme.
  - Scope yapısı oluşturuldu.
  - Sembol tanımlama ve çözümleme fonksiyonları eklendi.
  - İç içe kapsamlar için destek eklendi.

- [x] **Tip Sistemi**: Temel tipler, karmaşık tipler, şablon tipler ve tip çıkarımı.
  - Temel tipler (int, float, string, bool, char, null) eklendi.
  - Karmaşık tipler (array, map, function, class, interface) eklendi.
  - Şablon tipler için destek eklendi.

- [x] **Tip Kontrolü**: İfadelerin ve deyimlerin tip kontrolü.
  - Aritmetik operatörler için tip kontrolü eklendi.
  - Karşılaştırma operatörleri için tip kontrolü eklendi.
  - Mantıksal operatörler için tip kontrolü eklendi.
  - Atama operatörleri için tip kontrolü eklendi.

- [x] **İsim Çözümlemesi**: Değişken, fonksiyon ve tip isimlerinin çözümlenmesi.
  - Tanımlayıcıların çözümlenmesi için destek eklendi.
  - Fonksiyon ve metot çağrıları için destek eklendi.
  - Üye erişimi için destek eklendi.

- [x] **Erişim Kontrolü**: Public, private, protected erişim kontrolü.
  - Sınıf üyeleri için erişim belirleyicileri eklendi.
  - Erişim kontrolü için destek eklendi.

- [x] **Semantik Analiz İyileştirmeleri**: Semantik analiz bileşeninin daha karmaşık ifadeleri analiz edebilmesi için iyileştirmeler.
  - Hata mesajlarının iyileştirilmesi.
  - Daha karmaşık tip çıkarımı.
  - Daha iyi hata kurtarma mekanizmaları.

### Ara Kod Üretimi (IR Generation)

- [x] **IRGenerator Yapısı**: IRGenerator yapısı ve temel fonksiyonları oluşturuldu.
- [x] **LLVM Entegrasyonu**: LLVM Go bağlayıcıları (llir/llvm) projeye entegre edildi.
- [x] **Temel Tipler ve İfadeler için IR Üretimi**: Temel tipler (int, float, string, bool) ve sabit değerler için IR üretimi eklendi.
- [x] **Aritmetik ve Mantıksal İfadeler için IR Üretimi**: Aritmetik ve mantıksal ifadeler için IR üretimi eklendi.
- [x] **Değişken Tanımlamaları için IR Üretimi**: Değişken tanımlamaları için IR üretimi eklendi.
- [x] **Fonksiyon Tanımlamaları ve Çağrıları için IR Üretimi**: Fonksiyon tanımlamaları ve çağrıları için IR üretimi eklendi.
- [x] **Kontrol Akışı için IR Üretimi**: If ifadeleri ve while döngüleri için IR üretimi eklendi.
- [x] **Sınıf Tanımlamaları için IR Üretimi**: Sınıf tanımlamaları için IR üretimi eklendi.
- [x] **Optimizasyon ve Kod Üretimi**: LLVM optimizasyon geçişleri ve hedef kod üretimi.

### Optimizasyon ve Kod Üretimi

- [x] **IR Optimizasyonu Altyapısı**: LLVM optimizasyon geçişleri için altyapı oluşturuldu.
- [x] **IR Optimizasyonu**: LLVM optimizasyon geçişlerini yapılandırma.
- [x] **Hedef Kod Üretimi**: Farklı platformlar için makine kodu üretimi.
- [x] **Çalıştırılabilir Dosya Oluşturma**: Çalıştırılabilir dosya oluşturma.

### Standart Kütüphane ve Araçlar

- [x] **Standart Kütüphane**: Temel veri yapıları, I/O işlemleri, eşzamanlılık desteği.
  - Standart kütüphane dizin yapısı oluşturuldu.
  - Temel veri yapıları (list, vector) implementasyonları eklendi.
  - I/O işlemleri için temel arayüzler ve fonksiyonlar eklendi.
  - Eşzamanlılık desteği (channel, mutex, waitgroup) eklendi.
  - Dize işleme fonksiyonları eklendi.
  - Matematiksel fonksiyonlar ve sabitler eklendi.
  - Biçimlendirilmiş giriş/çıkış işlemleri için fmt paketi eklendi.
- [x] **Geliştirme Araçları**: Paket yöneticisi, derleme ve bağlama araçları, test araçları.
  - GO-Minus Paket Yöneticisi (gompm) oluşturuldu.
  - GO-Minus Test Aracı (gomtest) oluşturuldu.
  - GO-Minus Belgelendirme Aracı (gomdoc) oluşturuldu.
  - GO-Minus Kod Biçimlendirme Aracı (gomfmt) oluşturuldu.

### IDE Entegrasyonu ve Ekosistem

- [x] **IDE Desteği**: Sözdizimi vurgulama, kod tamamlama, hata ayıklama.
  - GO-Minus Dil Sunucusu (gomlsp) oluşturuldu.
  - GO-Minus Hata Ayıklama Aracı (gomdebug) oluşturuldu.
  - VS Code eklentisi geliştirildi.
  - JetBrains IDE'leri için eklenti geliştirildi.
  - Vim/Neovim eklentisi geliştirildi.
  - Emacs eklentisi geliştirildi.
  - TextMate dilbilgisi dosyaları oluşturuldu.
- [x] **Ekosistem Geliştirme**: Topluluk oluşturma, örnek projeler ve belgelendirme.
  - GO-Minus web sitesi oluşturuldu.
  - Katkı sağlama rehberi oluşturuldu.
  - Davranış kuralları oluşturuldu.
  - Örnek projeler ve şablonlar oluşturuldu.
  - Belgelendirme ve öğretici içerikler geliştirildi.

## Tamamlanan Görevler

### LLVM Entegrasyonu ve Kod Üretimi İyileştirmeleri

- [x] **LLVM IR Üretimi**: LLVM IR üretiminin tüm dil özellikleri için tamamlanması.
  - [x] Sınıf, şablon ve istisna işleme için IR üretimi
  - [x] Kalıtım ve polimorfizm için vtable implementasyonu
  - [x] Şablon örnekleme (template instantiation) mekanizması

### Hata Ayıklama Desteği

- [x] **Hata Ayıklama Bilgisi Üretimi**: Hata ayıklama bilgisi üretiminin tamamlanması.
  - [x] DWARF hata ayıklama bilgisi üretimi
  - [x] Kaynak haritalaması (source mapping)
  - [x] Değişken ve fonksiyon sembol tablosu

### Standart Kütüphane Genişletmeleri

- [x] **Container Paketi Genişletme**: Container paketinin genişletilmesi.
  - [x] Heap (öncelik kuyruğu) implementasyonu
  - [x] Deque (çift uçlu kuyruk) implementasyonu
  - [x] Trie (önek ağacı) implementasyonu

- [x] **Concurrent Paketi Genişletme**: Concurrent paketinin genişletilmesi.
  - [x] Semaphore implementasyonu
  - [x] Barrier implementasyonu
  - [x] ThreadPool implementasyonu
  - [x] Future/Promise implementasyonu

- [x] **IO Paketi Genişletme**: IO paketinin genişletilmesi.
  - [x] Buffered IO implementasyonu

- [x] **Yeni Paketler**: Yeni paketlerin eklenmesi.
  - [x] Regex paketi (düzenli ifadeler)

## Devam Eden Görevler

### Asenkron IO Implementasyonu

- [x] **Asenkron IO Arayüzleri**: Asenkron IO için temel arayüzler oluşturuldu.
  - [x] AsyncReader, AsyncWriter, AsyncCloser arayüzleri
  - [x] AsyncReadWriter, AsyncReadCloser, AsyncWriteCloser arayüzleri
  - [x] AsyncReadWriteCloser, AsyncSeeker, AsyncReadWriteSeeker arayüzleri

- [x] **Event Loop**: Asenkron IO için event loop implementasyonu.
  - [x] Temel event loop yapısı
  - [x] Olay işleme mekanizması
  - [x] Görev kuyruğu yönetimi

- [x] **Platform Bağımlı IO Multiplexing**: Farklı platformlar için IO multiplexing.
  - [x] Linux için epoll implementasyonu
  - [x] macOS/BSD için kqueue implementasyonu
  - [x] Windows için IOCP implementasyonu

- [x] **Future/Promise Pattern**: Asenkron işlemler için Future/Promise pattern.
  - [x] AsyncFuture ve AsyncPromise sınıfları
  - [x] Callback mekanizması
  - [x] Zincirlenebilir işlemler (then, catch, finally)
  - [x] Dönüştürme işlemleri (map, flatMap)

- [x] **Asenkron Dosya İşlemleri**: Asenkron dosya okuma/yazma işlemleri.
  - [x] Asenkron dosya açma/kapatma
  - [x] Asenkron okuma/yazma
  - [x] Asenkron konumlandırma

- [x] **Asenkron Soket İşlemleri**: Asenkron ağ işlemleri.
  - [x] Asenkron bağlantı kurma/dinleme
  - [x] Asenkron okuma/yazma
  - [x] Asenkron bağlantı kabul etme

- [ ] **Performans Optimizasyonu**: Asenkron IO performans optimizasyonu.
  - [x] CPU kullanımı optimizasyonu
  - [x] Lock-free veri yapıları
  - [ ] Sistem çağrıları optimizasyonu
  - [ ] Buffer havuzu implementasyonu

## Yapılacak İşler

### LLVM IR Optimizasyon Geçişleri

- [ ] **LLVM IR Optimizasyon Geçişleri**: LLVM IR optimizasyon geçişlerinin tamamlanması.
  - [ ] Fonksiyon içi optimizasyonlar (inlining, tail call optimization)
  - [ ] Döngü optimizasyonları (loop unrolling, loop fusion)
  - [ ] Vektörleştirme optimizasyonları
  - [ ] Ölü kod eliminasyonu ve sabit yayılımı
  - [ ] Fonksiyon çağrı optimizasyonları

### Hedef Kod Üretimi İyileştirmeleri

- [ ] **Hedef Kod Üretimi İyileştirmeleri**: Farklı platformlar için makine kodu üretiminin iyileştirilmesi.
  - [ ] x86_64 mimarisi için kod üretimi optimizasyonları
  - [ ] ARM64 mimarisi için kod üretimi optimizasyonları
  - [ ] WebAssembly hedefi için destek
  - [ ] Platform özel optimizasyonlar

### Çalıştırılabilir Dosya Oluşturma İyileştirmeleri

- [ ] **Çalıştırılabilir Dosya Oluşturma İyileştirmeleri**: Çalıştırılabilir dosya oluşturma sürecinin iyileştirilmesi.
  - [ ] Bağlayıcı (linker) entegrasyonu
  - [ ] Dinamik ve statik kütüphane desteği
  - [ ] Çalıştırılabilir dosya optimizasyonları
  - [ ] Çoklu platform destek iyileştirmeleri

### Standart Kütüphane Genişletme

- [ ] **Container Paketi Genişletme**: Container paketinin genişletilmesi.
  - [ ] Graph (çizge) implementasyonu

- [ ] **IO Paketi Genişletme**: IO paketinin genişletilmesi.
  - [x] Memory-mapped IO implementasyonu
  - [ ] Asenkron IO implementasyonu tamamlanması

- [ ] **Yeni Paketler**: Yeni paketlerin eklenmesi.
  - [x] Time paketi (zaman işlemleri)
  - [x] Network IO implementasyonu
  - [ ] Crypto paketi (kriptografi)
  - [ ] Encoding paketi (JSON, XML, CSV, vb.)
  - [ ] Database paketi (veritabanı işlemleri)
  - [ ] HTTP paketi (HTTP istemci ve sunucu)

### Hata Ayıklama Desteği

- [ ] **Hata Ayıklama Aracı İyileştirmeleri**: GO-Minus Hata Ayıklama Aracı'nın (gomdebug) iyileştirilmesi.
  - [ ] Kesme noktası yönetimi iyileştirmeleri
  - [ ] Değişken inceleme iyileştirmeleri
  - [ ] Yığın izi görüntüleme iyileştirmeleri
  - [ ] İfade değerlendirme iyileştirmeleri

- [ ] **IDE Entegrasyonu İyileştirmeleri**: IDE entegrasyonlarının iyileştirilmesi.
  - [ ] VS Code eklentisi iyileştirmeleri
  - [ ] JetBrains IDE'leri için eklenti iyileştirmeleri
  - [ ] Vim/Neovim eklentisi iyileştirmeleri
  - [ ] Emacs eklentisi iyileştirmeleri

### Performans Optimizasyonları

- [ ] **Derleme Süresi Optimizasyonları**: Derleme süresinin iyileştirilmesi.
  - [ ] Paralel derleme desteği
  - [ ] Artımlı derleme desteği
  - [ ] Önbellek mekanizmaları
  - [ ] Modül sistemi optimizasyonları

- [ ] **Çalışma Zamanı Performans Optimizasyonları**: Çalışma zamanı performansının artırılması.
  - [ ] Bellek yönetimi optimizasyonları
  - [ ] Garbage collector optimizasyonları
  - [ ] Fonksiyon çağrı optimizasyonları
  - [ ] Nesne düzeni optimizasyonları

- [ ] **Bellek Kullanımı Optimizasyonları**: Bellek kullanımının optimize edilmesi.
  - [ ] Veri yapıları optimizasyonları
  - [ ] Bellek havuzu (memory pool) implementasyonu
  - [ ] Bellek sızıntısı tespiti ve önleme
  - [ ] Bellek kullanımı profilleme araçları

### Belgelendirme ve Örnekler

- [ ] **Dil Referansı**: Kapsamlı dil referansı oluşturulması.
  - [ ] Sözdizimi referansı
  - [ ] Tip sistemi referansı
  - [ ] Standart kütüphane referansı
  - [ ] Operatör ve ifade referansı

- [ ] **Öğreticiler ve En İyi Uygulamalar**: Öğreticiler ve en iyi uygulamalar oluşturulması.
  - [ ] Başlangıç öğreticileri
  - [ ] İleri seviye öğreticiler
  - [ ] En iyi uygulamalar kılavuzu
  - [ ] Performans ipuçları

- [ ] **Örnek Projeler ve Şablonlar**: Örnek projeler ve şablonlar oluşturulması.
  - [ ] Konsol uygulamaları
  - [ ] Web uygulamaları
  - [ ] GUI uygulamaları
  - [ ] Sistem uygulamaları
  - [ ] Oyun geliştirme

## Notlar ve Kararlar

- GO-Minus dosyaları `.gom` uzantısını kullanacak.
- GO-Minus derleyicisi geliştirme aşamasında Go ile yazılacak.
- GO-Minus dili, Go'nun tüm özelliklerini destekleyecek ve C++ benzeri özelliklerle genişletilecek.

## Son Güncelleme

Tarih: 2024-08-01
Durum: Temel altyapı, semantik analiz, ara kod üretimi, IDE entegrasyonu ve ekosistem geliştirmesi tamamlandı. LLVM IR üretimi tüm dil özellikleri için tamamlandı (sınıf, şablon, istisna işleme, kalıtım, polimorfizm). Hata ayıklama bilgisi üretimi (DWARF) tamamlandı. Standart kütüphane genişletme çalışmaları kapsamında Container paketi (Heap, Deque, Trie), Concurrent paketi (Semaphore, Barrier, ThreadPool, Future/Promise), IO paketi (Buffered IO, Memory-mapped IO), Time paketi ve Network IO implementasyonları tamamlandı. Asenkron IO implementasyonu için temel arayüzler, event loop, platform bağımlı IO multiplexing (epoll, kqueue, IOCP), Future/Promise pattern, asenkron dosya ve soket işlemleri tamamlandı. Asenkron IO performans optimizasyonu (CPU kullanımı, lock-free veri yapıları) üzerinde çalışmalar devam ediyor. LLVM IR optimizasyon geçişleri, hedef kod üretimi iyileştirmeleri, çalıştırılabilir dosya oluşturma iyileştirmeleri, standart kütüphane genişletme, hata ayıklama desteği, performans optimizasyonları, belgelendirme ve örnekler üzerinde çalışmalar devam ediyor.

## Proje İsim Değişikliği

Tarih: 2024-05-15
Durum: Proje adı "GO+" dan "GO-Minus" olarak değiştirildi. Dosya uzantısı `.gop` yerine `.gom` olarak güncellendi. Tüm dokümantasyon ve kod tabanı bu değişikliğe göre güncellendi.

## Belgelendirme ve Örnekler Güncellemesi

Tarih: 2024-05-20
Durum: Belgelendirme yapısı genişletildi ve yeni belgeler eklendi:

1. "Neden GO-Minus" kılavuzu oluşturuldu
2. Başlangıç rehberi eklendi
3. Dil referansı belgeleri güncellendi
4. Vulkan "Hello Triangle" örneği ve belgelendirmesi eklendi
5. SSS (Sık Sorulan Sorular) belgesi oluşturuldu
6. En İyi Uygulamalar belgesi eklendi

Kısa vadeli geliştirme planı kapsamında belgelendirme çalışmaları devam ediyor. Vulkan bağlayıcıları için prototip çalışmaları başlatıldı ve manuel bellek yönetimi için araştırma ekibi oluşturuldu.

## Semantik Analiz İyileştirmeleri ve Standart Kütüphane Genişletmeleri

Tarih: 2024-06-15
Durum: GO-Minus programlama dilinin temel özelliklerini tamamlamak için aşağıdaki geliştirmeler yapıldı:

### Semantik Analiz İyileştirmeleri:
1. **Gelişmiş Hata Raporlama Sistemi**: Farklı hata seviyeleri, renkli çıktı, dosya ve konum bilgisi, ipuçları ve düzeltme önerileri, benzer tanımlayıcı önerileri eklendi.
2. **Tip Çıkarımı Modülü**: Değişken tanımlamaları, fonksiyon dönüş tipleri, karmaşık ifadeler, jenerik fonksiyonlar ve şablon sınıflar için tip çıkarımı eklendi.
3. **Hata Kurtarma Mekanizmaları**: Analiz devam etme, eksik sembol kurtarma, tip uyuşmazlığı kurtarma, eksik üye kurtarma ve sözdizimi hatası kurtarma mekanizmaları eklendi.

### Standart Kütüphane Genişletmeleri:
1. **Trie (Önek Ağacı) Implementasyonu**: Container paketine jenerik tip desteği olan, kelime ekleme, arama, silme, önek araması ve tüm kelimeleri listeleme özelliklerine sahip Trie veri yapısı eklendi.
2. **Buffered IO Implementasyonu**: IO paketine tamponlanmış okuma ve yazma işlemleri için BufferedReader ve BufferedWriter sınıfları eklendi.
3. **Regex Paketi**: Düzenli ifade deseni derleme, metin eşleştirme, tüm eşleşmeleri bulma, metin değiştirme, metin bölme, büyük/küçük harf duyarlı ve duyarsız modlar, çok satırlı mod desteği sağlayan Regex paketi eklendi.

Bu geliştirmeler, GO-Minus programlama dilinin daha güçlü, kullanıcı dostu ve yetenekli olmasını sağlamıştır. Semantik analiz iyileştirmeleri, programcıların hatalarını daha hızlı bulmasına ve düzeltmesine yardımcı olurken, standart kütüphane genişletmeleri, GO-Minus'un daha geniş bir uygulama yelpazesinde kullanılmasını sağlamaktadır.

Detaylı bilgi için [Semantik Analiz İyileştirmeleri](docs/semantic-analysis-improvements.md), [Standart Kütüphane Genişletmeleri](docs/stdlib-extensions.md) ve [Geliştirme Raporu](docs/development-report.md) belgelerine bakabilirsiniz.

## LLVM IR Üretimi Tamamlanması

Tarih: 2024-07-01
Durum: GO-Minus programlama dilinin LLVM IR üretimi tüm dil özellikleri için tamamlandı:

### Sınıf, Şablon ve İstisna İşleme için IR Üretimi:
1. **Sınıf IR Üretimi**: Sınıf tanımlamaları için LLVM IR üretimi tamamlandı. Sınıflar, LLVM struct tipleri olarak temsil edildi ve sınıf üyeleri (alanlar ve metotlar) için IR üretimi eklendi.
2. **Şablon IR Üretimi**: Şablon sınıflar ve fonksiyonlar için LLVM IR üretimi tamamlandı. Şablon örnekleme mekanizması eklendi ve şablon parametreleri için tip eşleme sistemi geliştirildi.
3. **İstisna İşleme IR Üretimi**: Try-catch-finally blokları ve throw ifadeleri için LLVM IR üretimi tamamlandı. LLVM'in istisna işleme mekanizması (landingpad, personality fonksiyonu, resume) kullanıldı.

### Kalıtım ve Polimorfizm için VTable Implementasyonu:
1. **VTable Yapısı**: Sanal metot tablosu (vtable) için LLVM struct tipi oluşturuldu ve sınıf başına bir vtable örneği oluşturuldu.
2. **Sanal Metot Çağrıları**: Sanal metot çağrıları için vtable üzerinden dinamik dispatch mekanizması eklendi.
3. **Kalıtım Desteği**: Türetilmiş sınıfların ebeveyn sınıfların vtable'larını miras alması ve gerektiğinde override etmesi için mekanizma eklendi.

### Şablon Örnekleme Mekanizması:
1. **Şablon Tanımlama**: Şablon sınıf ve fonksiyon tanımlamaları için AST düğümleri ve IR üretimi eklendi.
2. **Şablon Örnekleme**: Şablon sınıf ve fonksiyonların belirli tip argümanlarıyla örneklenmesi için mekanizma eklendi.
3. **Şablon Örneği Önbelleği**: Aynı tip argümanlarıyla örneklenmiş şablonların tekrar kullanılması için önbellek mekanizması eklendi.

Bu geliştirmeler, GO-Minus programlama dilinin nesne yönelimli programlama, jenerik programlama ve istisna işleme gibi modern dil özelliklerini tam olarak desteklemesini sağlamıştır. LLVM IR üretiminin tamamlanmasıyla, GO-Minus derleyicisi artık tüm dil özelliklerini destekleyen IR kodu üretebilmektedir.

Sonraki adımlar, LLVM IR optimizasyon geçişlerinin iyileştirilmesi ve kapsamlı test senaryoları çalıştırılarak hataların düzeltilmesi olacaktır.

## Hata Ayıklama Bilgisi Üretimi Tamamlanması

Tarih: 2024-07-05
Durum: GO-Minus programlama dilinin hata ayıklama bilgisi üretimi tamamlandı:

### DWARF Hata Ayıklama Bilgisi Üretimi:
1. **Derleme Birimi Bilgisi**: Kaynak dosya, dizin, üretici ve derleme seçenekleri gibi bilgileri içeren derleme birimi meta verisi eklendi.
2. **Fonksiyon Bilgisi**: Fonksiyon adı, kaynak konumu, dönüş tipi ve parametreler gibi bilgileri içeren fonksiyon meta verisi eklendi.
3. **Değişken Bilgisi**: Değişken adı, tipi, kapsamı ve bellek konumu gibi bilgileri içeren değişken meta verisi eklendi.
4. **Tip Bilgisi**: Temel tipler, işaretçi tipleri, dizi tipleri, yapı tipleri ve fonksiyon tipleri için meta veri eklendi.

### Kaynak Haritalaması (Source Mapping):
1. **Konum Bilgisi**: Her IR talimatı için kaynak dosya, satır ve sütun bilgisi eklendi.
2. **Sözcüksel Blok Bilgisi**: Blok deyimleri için sözcüksel blok meta verisi eklendi.
3. **Kapsam Bilgisi**: Değişkenlerin kapsamını belirten meta veri eklendi.

### Değişken ve Fonksiyon Sembol Tablosu:
1. **Değişken Bildirimleri**: Yerel değişkenler için bildirim meta verisi eklendi.
2. **Parametre Bildirimleri**: Fonksiyon parametreleri için bildirim meta verisi eklendi.
3. **Global Değişken Bildirimleri**: Global değişkenler için bildirim meta verisi eklendi.
4. **Fonksiyon Bildirimleri**: Fonksiyonlar için bildirim meta verisi eklendi.

Bu geliştirmeler, GO-Minus programlama dilinin hata ayıklama araçları ile entegrasyonunu sağlamıştır. Hata ayıklama bilgisi üretiminin tamamlanmasıyla, GO-Minus derleyicisi artık kaynak düzeyinde hata ayıklama yapılabilecek şekilde DWARF hata ayıklama bilgisi üretebilmektedir.

Sonraki adımlar, LLVM IR optimizasyon geçişlerinin iyileştirilmesi olacaktır.

## Kapsamlı Test Senaryoları Tamamlanması

Tarih: 2024-07-10
Durum: GO-Minus programlama dili için kapsamlı test senaryoları tamamlandı:

### Temel Dil Özellikleri için Test Senaryoları:
1. **Sözdizimi Testleri**: Değişken tanımlamaları, fonksiyon tanımlamaları, kontrol yapıları ve döngüler için sözdizimi testleri eklendi.
2. **Tip Sistemi Testleri**: Temel tipler, karmaşık tipler, tip dönüşümleri ve tip çıkarımı için testler eklendi.
3. **Operatör Testleri**: Aritmetik, mantıksal, karşılaştırma ve atama operatörleri için testler eklendi.
4. **Kapsam Testleri**: Yerel, global ve blok kapsamları için testler eklendi.

### Gelişmiş Dil Özellikleri için Test Senaryoları:
1. **Sınıf Testleri**: Sınıf tanımlamaları, kalıtım, arayüzler ve polimorfizm için testler eklendi.
2. **Şablon Testleri**: Şablon sınıflar, şablon fonksiyonlar ve şablon özelleştirme için testler eklendi.
3. **İstisna İşleme Testleri**: Try-catch-finally blokları ve throw ifadeleri için testler eklendi.
4. **Standart Kütüphane Testleri**: Container, IO, Regex ve Concurrent paketleri için testler eklendi.

### Test Çalıştırıcı Script:
1. **Otomatik Test Çalıştırma**: Tüm test dosyalarını otomatik olarak derleyen ve çalıştıran bir script eklendi.
2. **Hata Ayıklama Desteği**: Hata ayıklama bilgisi üretimini etkinleştiren bir seçenek eklendi.
3. **Test Sonuçları Raporlama**: Başarılı ve başarısız testleri raporlayan bir mekanizma eklendi.

Bu geliştirmeler, GO-Minus programlama dilinin doğru çalıştığından emin olmak için kapsamlı bir test altyapısı sağlamıştır. Kapsamlı test senaryoları, dilin tüm özelliklerini test etmekte ve hataları erken aşamada tespit etmektedir.

Sonraki adımlar, LLVM IR optimizasyon geçişlerinin iyileştirilmesi olacaktır.
# GO+ Projesine Katkıda Bulunma Rehberi

GO+ projesine katkıda bulunmak istediğiniz için teşekkür ederiz! Bu belge, katkıda bulunma sürecini anlamanıza yardımcı olacaktır.

## Katkıda Bulunma Yolları

GO+ projesine çeşitli şekillerde katkıda bulunabilirsiniz:

1. **Kod Katkıları**: Yeni özellikler ekleyebilir, hataları düzeltebilir veya performans iyileştirmeleri yapabilirsiniz.
2. **Belgelendirme**: Belgelendirmeyi iyileştirebilir, örnekler ekleyebilir veya öğreticiler yazabilirsiniz.
3. **Hata Raporları**: Bulduğunuz hataları bildirebilirsiniz.
4. **Özellik İstekleri**: Yeni özellikler önerebilirsiniz.
5. **Testler**: Test kapsamını artırabilir veya mevcut testleri iyileştirebilirsiniz.
6. **Topluluk Desteği**: Forumlarda veya Discord'da diğer kullanıcılara yardımcı olabilirsiniz.

## Geliştirme Ortamı Kurulumu

GO+ projesini geliştirmek için aşağıdaki adımları izleyin:

1. Depoyu klonlayın:
   ```bash
   git clone https://github.com/goplus/goplus.git
   cd goplus
   ```

2. Bağımlılıkları yükleyin:
   ```bash
   go mod download
   ```

3. Derleyiciyi derleyin:
   ```bash
   go build -o goplus ./cmd/goplus
   ```

4. Testleri çalıştırın:
   ```bash
   go test ./...
   ```

## Kod Katkıları

### Dallanma Modeli

GO+ projesi, [GitHub Flow](https://guides.github.com/introduction/flow/) dallanma modelini kullanır:

1. Ana depoyu forklayın
2. Özellik dalı oluşturun (`git checkout -b feature/amazing-feature`)
3. Değişikliklerinizi commit edin (`git commit -m 'Add amazing feature'`)
4. Dalınızı uzak depoya itin (`git push origin feature/amazing-feature`)
5. Bir Pull Request açın

### Kod Stili

GO+ projesi, aşağıdaki kod stili kurallarını takip eder:

- Go için [Effective Go](https://golang.org/doc/effective_go.html) ve [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) belgelerindeki kurallar geçerlidir.
- GO+ için ek olarak:
  - Sınıf isimleri PascalCase kullanır
  - Metot isimleri camelCase kullanır
  - Özel üyeler için `private:` erişim belirleyicisi kullanılır
  - Genel üyeler için `public:` erişim belirleyicisi kullanılır
  - Korumalı üyeler için `protected:` erişim belirleyicisi kullanılır

### Commit Mesajları

Commit mesajları aşağıdaki formatta olmalıdır:

```
<type>: <subject>

<body>

<footer>
```

Örnek:

```
feat: Add class inheritance support

Implement C++-like class inheritance with single and multiple inheritance.
Support for virtual methods and abstract classes.

Closes #123
```

Tip alanı aşağıdakilerden biri olabilir:
- `feat`: Yeni bir özellik
- `fix`: Bir hata düzeltmesi
- `docs`: Sadece belgelendirme değişiklikleri
- `style`: Kod davranışını etkilemeyen değişiklikler (boşluk, biçimlendirme, vb.)
- `refactor`: Hata düzeltmesi veya özellik eklemeyen kod değişiklikleri
- `perf`: Performansı artıran değişiklikler
- `test`: Test ekleme veya düzeltme
- `chore`: Derleme süreci veya yardımcı araçlardaki değişiklikler

### Pull Request Süreci

1. Pull Request açmadan önce, kodunuzun tüm testleri geçtiğinden emin olun.
2. Pull Request'inizde, değişikliklerinizi açıklayan ayrıntılı bir açıklama sağlayın.
3. Eğer Pull Request'iniz bir sorunu çözüyorsa, açıklamada "Closes #123" gibi bir referans ekleyin.
4. Pull Request'iniz gözden geçirilecek ve geri bildirim alacaksınız.
5. Gerekirse, geri bildirimlere göre değişiklikler yapın.
6. Pull Request'iniz onaylandığında, ana dala birleştirilecektir.

## Hata Raporları

Bir hata raporu gönderirken, lütfen aşağıdaki bilgileri sağlayın:

1. GO+ sürümü
2. İşletim sistemi ve sürümü
3. Hatayı yeniden oluşturmak için adımlar
4. Beklenen davranış
5. Gerçek davranış
6. Varsa, ilgili log çıktıları veya ekran görüntüleri

## Özellik İstekleri

Bir özellik isteği gönderirken, lütfen aşağıdaki bilgileri sağlayın:

1. Özelliğin ayrıntılı bir açıklaması
2. Özelliğin neden faydalı olacağına dair bir gerekçe
3. Varsa, diğer dillerdeki benzer özelliklere referanslar
4. Varsa, özelliğin nasıl uygulanabileceğine dair fikirler

## Belgelendirme Katkıları

Belgelendirme katkıları, GO+ projesinin kullanıcılar tarafından daha iyi anlaşılmasına yardımcı olur. Belgelendirme katkıları şunları içerebilir:

1. API belgelendirmesi
2. Öğreticiler ve kılavuzlar
3. Kod örnekleri
4. Dil referansı
5. Hata ayıklama ve sorun giderme belgeleri

## Lisans

GO+ projesi, [MIT Lisansı](LICENSE) altında lisanslanmıştır. Katkıda bulunarak, katkılarınızın aynı lisans altında yayınlanmasını kabul etmiş olursunuz.

## İletişim

Sorularınız veya geri bildirimleriniz için:

- GitHub Issues: [https://github.com/goplus/goplus/issues](https://github.com/goplus/goplus/issues)
- Discord: [GO+ Discord Sunucusu](https://discord.gg/goplus)
- Forum: [GO+ Forum](https://forum.goplus.org)

## Teşekkürler

GO+ projesine katkıda bulunduğunuz için tekrar teşekkür ederiz! Katkılarınız, GO+ dilinin gelişmesine ve büyümesine yardımcı olacaktır.
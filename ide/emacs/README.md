# GO-Minus Emacs Eklentisi

GO-Minus Emacs Eklentisi, GO-Minus programlama dili için Emacs desteği sağlar. Bu eklenti, sözdizimi vurgulama, kod tamamlama, hata işaretleme, tanıma gitme gibi özellikler sunar.

## Özellikler

- Sözdizimi vurgulama
- Kod tamamlama (lsp-mode ile)
- Hata işaretleme
- Tanıma gitme
- Referansları bulma
- Kod biçimlendirme
- Testleri çalıştırma
- Kod kapsama analizi

## Kurulum

### Ön Koşullar

- GO-Minus derleyicisi ve araçları yüklü olmalıdır.
- `gomlsp` (GO-Minus Dil Sunucusu) PATH ortam değişkeninizde bulunmalıdır.
- Emacs 25.1+
- (İsteğe bağlı) lsp-mode ve company-mode

### MELPA ile Kurulum

```elisp
;; .emacs veya init.el dosyanıza ekleyin
(require 'package)
(add-to-list 'package-archives '("melpa" . "https://melpa.org/packages/") t)
(package-initialize)

(unless package-archive-contents
  (package-refresh-contents))

(dolist (package '(goplus-mode lsp-mode company-mode))
  (unless (package-installed-p package)
    (package-install package)))
```

### Manuel Kurulum

1. Eklentiyi indirin
2. Dosyaları Emacs yükleme dizinine kopyalayın
3. `.emacs` veya `init.el` dosyanıza aşağıdaki kodu ekleyin:

```elisp
(add-to-list 'load-path "/path/to/goplus-mode")
(require 'goplus-mode)
```

## Yapılandırma

### Temel Yapılandırma

```elisp
;; .emacs veya init.el dosyanıza ekleyin
(require 'goplus-mode)
(add-to-list 'auto-mode-alist '("\\.gop\\'" . goplus-mode))

;; Girinti ayarları
(setq goplus-tab-width 4)
(setq goplus-indent-offset 4)

;; Biçimlendirme ayarları
(setq goplus-format-on-save t)
```

### LSP Yapılandırması

```elisp
;; .emacs veya init.el dosyanıza ekleyin
(require 'lsp-mode)
(add-hook 'goplus-mode-hook #'lsp-deferred)

;; LSP UI
(require 'lsp-ui)
(setq lsp-ui-doc-enable t
      lsp-ui-doc-use-childframe t
      lsp-ui-doc-position 'top
      lsp-ui-doc-include-signature t
      lsp-ui-sideline-enable t
      lsp-ui-flycheck-enable t
      lsp-ui-flycheck-list-position 'right
      lsp-ui-flycheck-live-reporting t
      lsp-ui-peek-enable t
      lsp-ui-peek-list-width 60
      lsp-ui-peek-peek-height 25)

;; LSP için GO+ dil sunucusu yapılandırması
(lsp-register-client
 (make-lsp-client :new-connection (lsp-stdio-connection "goplsp")
                  :major-modes '(goplus-mode)
                  :server-id 'goplsp
                  :initialized-fn (lambda (workspace)
                                    (with-lsp-workspace workspace
                                      (lsp--set-configuration
                                       (lsp-configuration-section "goplus"))))))
```

### Kod Tamamlama

```elisp
;; .emacs veya init.el dosyanıza ekleyin
(require 'company)
(add-hook 'goplus-mode-hook #'company-mode)
(setq company-idle-delay 0.1
      company-minimum-prefix-length 1)
```

### Anahtar Bağlamaları

```elisp
;; .emacs veya init.el dosyanıza ekleyin
(define-key goplus-mode-map (kbd "C-c C-f") 'goplus-format-buffer)
(define-key goplus-mode-map (kbd "C-c C-t") 'goplus-test-current-file)
(define-key goplus-mode-map (kbd "C-c C-r") 'goplus-run-current-file)
(define-key goplus-mode-map (kbd "C-c C-d") 'goplus-describe)
(define-key goplus-mode-map (kbd "C-c C-j") 'goplus-jump-to-definition)
(define-key goplus-mode-map (kbd "C-c C-k") 'goplus-find-references)
```

## Kullanım

### Sözdizimi Vurgulama

GO+ dosyaları (`.gop` uzantılı) otomatik olarak sözdizimi vurgulaması ile açılır.

### Kod Tamamlama

Kod yazarken, otomatik tamamlama önerileri görünecektir (company-mode ve lsp-mode yapılandırılmışsa).

### Tanıma Gitme

Bir sembolün tanımına gitmek için, sembolün üzerindeyken `C-c C-j` tuşlarına basın.

### Kod Biçimlendirme

Bir belgeyi biçimlendirmek için, `C-c C-f` tuşlarına basın.

### Testleri Çalıştırma

Testleri çalıştırmak için, `C-c C-t` tuşlarına basın.

### Mevcut Dosyayı Çalıştırma

Mevcut dosyayı çalıştırmak için, `C-c C-r` tuşlarına basın.

## Komutlar

- `goplus-format-buffer`: Mevcut tamponu biçimlendirir
- `goplus-test-current-file`: Mevcut dosyanın testlerini çalıştırır
- `goplus-run-current-file`: Mevcut dosyayı çalıştırır
- `goplus-describe`: Sembol hakkında bilgi gösterir
- `goplus-jump-to-definition`: Sembolün tanımına gider
- `goplus-find-references`: Sembolün referanslarını bulur
- `goplus-list-packages`: Paketleri listeler
- `goplus-import-add`: İçe aktarma ekler
- `goplus-import-remove`: İçe aktarmayı kaldırır

## Sorun Giderme

### Sözdizimi Vurgulama Çalışmıyor

1. Mod'un doğru yüklendiğini kontrol edin (`M-x goplus-mode`)
2. Dosya uzantısının doğru ilişkilendirildiğini kontrol edin

### LSP Çalışmıyor

1. `goplsp` komutunun PATH ortam değişkeninizde olduğunu kontrol edin
2. LSP durumunu kontrol edin (`M-x lsp-describe-session`)
3. LSP günlüklerini kontrol edin (`M-x lsp-workspace-show-log`)

## Katkıda Bulunma

GO+ Emacs Eklentisi, açık kaynaklı bir projedir. Katkıda bulunmak için, lütfen [katkı sağlama rehberini](../../CONTRIBUTING.md) okuyun.

## Lisans

GO+ Emacs Eklentisi, GO+ projesi ile aynı lisans altında dağıtılmaktadır.
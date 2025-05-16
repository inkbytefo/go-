# GO-Minus Vim/Neovim Eklentisi

GO-Minus Vim/Neovim Eklentisi, GO-Minus programlama dili için Vim ve Neovim desteği sağlar. Bu eklenti, sözdizimi vurgulama, kod tamamlama, hata işaretleme, tanıma gitme gibi özellikler sunar.

## Özellikler

- Sözdizimi vurgulama
- Kod tamamlama (coc.nvim veya vim-lsp ile)
- Hata işaretleme
- Tanıma gitme
- Referansları bulma
- Kod biçimlendirme
- Testleri çalıştırma

## Kurulum

### Ön Koşullar

- GO-Minus derleyicisi ve araçları yüklü olmalıdır.
- `gomlsp` (GO-Minus Dil Sunucusu) PATH ortam değişkeninizde bulunmalıdır.
- Vim 8.0+ veya Neovim 0.4.0+
- (İsteğe bağlı) coc.nvim veya vim-lsp gibi bir LSP istemcisi

### vim-plug ile Kurulum

```vim
" vimrc veya init.vim dosyanıza ekleyin
Plug 'gominus/vim-gominus'

" coc.nvim kullanıyorsanız
Plug 'neoclide/coc.nvim', {'branch': 'release'}

" vim-lsp kullanıyorsanız
Plug 'prabirshrestha/vim-lsp'
Plug 'mattn/vim-lsp-settings'
```

Vim'i açın ve `:PlugInstall` komutunu çalıştırın.

### Manuel Kurulum

1. Eklentiyi indirin
2. Dosyaları Vim eklenti dizinine kopyalayın:
   - Vim: `~/.vim/`
   - Neovim: `~/.config/nvim/`

## Yapılandırma

### Sözdizimi Vurgulama

Sözdizimi vurgulama otomatik olarak etkinleştirilir. Özel renk şeması için:

```vim
" vimrc veya init.vim dosyanıza ekleyin
let g:gominus_highlight_types = 1
let g:gominus_highlight_fields = 1
let g:gominus_highlight_functions = 1
let g:gominus_highlight_function_calls = 1
let g:gominus_highlight_operators = 1
let g:gominus_highlight_build_constraints = 1
```

### coc.nvim ile LSP Yapılandırması

```json
// coc-settings.json dosyanıza ekleyin
{
  "languageserver": {
    "gominus": {
      "command": "gomlsp",
      "filetypes": ["gominus", "gom"],
      "rootPatterns": ["gom.mod", ".git/"],
      "initializationOptions": {
        "goplsEnv": {
          "GOPATH": "/path/to/gopath"
        }
      }
    }
  }
}
```

### vim-lsp ile LSP Yapılandırması

```vim
" vimrc veya init.vim dosyanıza ekleyin
if executable('gomlsp')
  au User lsp_setup call lsp#register_server({
    \ 'name': 'gominus',
    \ 'cmd': {server_info->['gomlsp']},
    \ 'whitelist': ['gominus', 'gom'],
    \ 'workspace_config': {
    \   'gopls': {
    \     'staticcheck': v:true,
    \     'completeUnimported': v:true,
    \     'usePlaceholders': v:true,
    \     'completionDocumentation': v:true,
    \     'watchFileChanges': v:true,
    \     'hoverKind': 'FullDocumentation',
    \   }
    \ }
    \ })
endif
```

### Dosya Tipi Tanımlama

```vim
" vimrc veya init.vim dosyanıza ekleyin
au BufRead,BufNewFile *.gom set filetype=gominus
```

### Anahtar Eşlemeleri

```vim
" vimrc veya init.vim dosyanıza ekleyin
" Tanıma gitme
autocmd FileType gominus nmap <buffer> gd <plug>(lsp-definition)
" Referansları bulma
autocmd FileType gominus nmap <buffer> gr <plug>(lsp-references)
" Belgeyi biçimlendir
autocmd FileType gominus nmap <buffer> gf <plug>(lsp-document-format)
" Testleri çalıştır
autocmd FileType gominus nmap <buffer> <leader>t :GominusTest<CR>
" Mevcut dosyayı çalıştır
autocmd FileType gominus nmap <buffer> <leader>r :GominusRun<CR>
```

## Kullanım

### Sözdizimi Vurgulama

GO-Minus dosyaları (`.gom` uzantılı) otomatik olarak sözdizimi vurgulaması ile açılır.

### Kod Tamamlama

Kod yazarken, otomatik tamamlama önerileri görünecektir (LSP istemcisi yapılandırılmışsa).

### Tanıma Gitme

Bir sembolün tanımına gitmek için, sembolün üzerindeyken `gd` tuşlarına basın (LSP istemcisi yapılandırılmışsa).

### Kod Biçimlendirme

Bir belgeyi biçimlendirmek için, `gf` tuşlarına basın (LSP istemcisi yapılandırılmışsa).

### Testleri Çalıştırma

Testleri çalıştırmak için, `<leader>t` tuşlarına basın.

### Mevcut Dosyayı Çalıştırma

Mevcut dosyayı çalıştırmak için, `<leader>r` tuşlarına basın.

## Komutlar

- `:GominusRun`: Mevcut dosyayı çalıştırır
- `:GominusTest`: Testleri çalıştırır
- `:GominusLint`: Kodu denetler
- `:GominusFmt`: Kodu biçimlendirir
- `:GominusImports`: İçe aktarmaları düzenler
- `:GominusInfo`: GO-Minus ortam bilgilerini gösterir

## Sorun Giderme

### Sözdizimi Vurgulama Çalışmıyor

1. Dosya tipinin doğru tanımlandığını kontrol edin (`:set filetype?` komutunu çalıştırın)
2. Sözdizimi vurgulamanın etkin olduğunu kontrol edin (`:syntax on` komutunu çalıştırın)

### LSP Çalışmıyor

1. `gomlsp` komutunun PATH ortam değişkeninizde olduğunu kontrol edin
2. LSP istemcisinin doğru yapılandırıldığını kontrol edin
3. LSP durumunu kontrol edin (coc.nvim için `:CocInfo`, vim-lsp için `:LspStatus`)

## Katkıda Bulunma

GO-Minus Vim/Neovim Eklentisi, açık kaynaklı bir projedir. Katkıda bulunmak için, lütfen [katkı sağlama rehberini](../../CONTRIBUTING.md) okuyun.

## Lisans

GO-Minus Vim/Neovim Eklentisi, GO-Minus projesi ile aynı lisans altında dağıtılmaktadır.
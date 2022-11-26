# vim

## basic settings

限制最大列数
-----------
```vimrc
:set colorcolumn=80
```

python tabs
-----------

```vimrc
set expandtab           " enter spaces when tab is pressed
set textwidth=120       " break lines when line length increases
set tabstop=4           " use 4 spaces to represent tab
set softtabstop=4
set shiftwidth=4        " number of spaces to use for auto indent
set autoindent          " copy indent from current line when starting a new line
" make backspaces more powerfull
set backspace=indent,eol,start
```

## vim-plug

### install

```
curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
```

### config


```
call plug#begin('~/.vim/plugged')

Plug 'scrooloose/nerdtree', { 'on': 'NERDTreeToggle' }

call plug#end()
```

then `:source %`  and `:PlugInstall`


### coc.nvim

install node

```sh
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | bash
nvm install 11
```

edit `~/.vimrc`

```
Plug 'neoclide/coc.nvim', {'branch': 'release'}
```

`:PlugInstall`

edit `~/.vimrc`

```
" GoTo code navigation.
nmap <silent> gd <Plug>(coc-definition)
nmap <silent> gy <Plug>(coc-type-definition)
nmap <silent> gi <Plug>(coc-implementation)
nmap <silent> gr <Plug>(coc-references)

" Use K to show documentation in preview window.
nnoremap <silent> K :call <SID>show_documentation()<CR>

function! s:show_documentation()
  if (index(['vim','help'], &filetype) >= 0)
    execute 'h '.expand('<cword>')
  elseif (coc#rpc#ready())
    call CocActionAsync('doHover')
  else
    execute '!' . &keywordprg . " " . expand('<cword>')
  endif
endfunction
```


### C/C++

```sh
apt install ccls
```

edit config file via `:CocConfig`

```json
{
  "languageserver": {
    "ccls": {
      "command": "ccls",
      "filetypes": ["c", "cpp", "cuda", "objc", "objcpp"],
      "rootPatterns": [".ccls-root", "compile_commands.json"],
      "initializationOptions": {
        "cache": {
          "directory": ".ccls-cache"
        },
        "client": {
          "snippetSupport": true
        }
      }
    }
  }
}
```



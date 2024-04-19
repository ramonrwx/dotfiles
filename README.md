# dotfiles

my dotfiles for termux app

## install

### required packages

```
pkg install -y aerc atuin binutils build-essential curl
    direnv fd git gnupg golang
    gum iconv jq luarocks lazygit \
    lsd neovim nodejs openssh openssl
    openssl-tool python ripgrep ruby rust \
    sqlite stylua zsh
```

### oh-my-zsh

```
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

### required python packages

```
python -m pip install --user \
    httpie \
    pre-commit \
    requests \
    virtualenv
```

### lua packages

```
luarocks install luacheck
```

### install dotfiles

```
git clone --recurse-submodules https://git.sr.ht/~ramonrw/dotfiles ~/.dotfiles
```

```
cd ~/.dotfiles
./install
```

# dotfiles

my dotfiles for termux app

## install

### required packages

```
pkg install -y aerc atuin binutils build-essential curl \
    direnv fd git gnupg golang \
    gum hut iconv isync jq \
    luarocks lazygit lsd man neovim \
    nodejs openssh openssl openssl-tool python \
    pass ripgrep ruby rust sqlite \
    stylua ugrep w3m zsh
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

### required go packages

```
go install github.com/ferdinandyb/maildir-rank-addr@latest
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

### install better-sqlite3

```
mkdir ~/.gyp
cat << EOF > ~/.gyp/include.gypi
{
	'variables': {
		'android_ndk_path': ''
	}
}
EOF
```

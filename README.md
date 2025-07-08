# dotfiles

my dotfiles for termux app

## install

### required packages

```
pkg install -y aerc atuin binutils build-essential curl \
    direnv fd git git-lfs gnupg \
    golang gum hut iconv isync \
    jq just luarocks lazygit lsd \
    neovim nodejs man msmtp openssh \
    openssl openssl-tool python pass ripgrep \
    ruby rust scdoc sqlite shfmt \
    shellcheck stylua ugrep uv w3m \
    wget zsh
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
    virtualenv \
    virtualenvwrapper
```

### required go packages

*mjmap*
a sendmail-compatible JMAP client.

```
git clone https://git.sr.ht/~rockorager/mjmap
gmake -C mjmap PREFIX=~/.local
cd mjmap
mv mjmap ~/.go/bin
```

### lua packages

```
luarocks install luacheck
```

### install dotfiles

```
git clone --recurse-submodules https://git.sr.ht/~rmon/dotfiles ~/.dotfiles
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

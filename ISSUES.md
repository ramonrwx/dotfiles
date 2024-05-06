# Python

when trying to install pre-commit-hooks there is a [bug](https://github.com/termux/termux-packages/issues/16746)
in the ruamel.yaml.clib library in which it cannot
compile.

you can omit the compilation failure by setting
`-Wno-incompatible-function-pointer-types` in CFLAGS

```
export CFLAGS="-Wno-incompatible-function-pointer-types"
python -m pip install ruamel.yaml.clib
```

# Node

[Error: EACCES: permission denied](https://stackoverflow.com/a/49714908)

- make a directory for global installations:
```
mkdir ~/.npm-global
```
- configure npm to use the new directory path:
```
npm config set prefix ~/.npm-global
```
- open or create a ~/.profile file and add this line:
```
export PATH=~/.npm-global/bin:$PATH
```
- back on the command line, update your system variables:
```
exec zsh
```

# C
[Unable To Install CS50 Library for C on Android Termux](https://cs50.stackexchange.com/a/38316)

## Build

```zsh
  go build -o pulumiGo
```

##

xattr pulumiGo

```zsh
com.apple.quarantine
```

- remove quarantine

```zsh
xattr -d com.apple.quarantine pulumiGo
```

```zsh
mkdir -p ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/pulumiGo
pulumiGo completion zsh > ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/pulumiGo/_pulumiGo

# Eidt ~/.zshrc
# plugins=(git ... pulumiGo)

# Reload
source ~/.zshrc

```

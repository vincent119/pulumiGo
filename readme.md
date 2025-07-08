## Build

- go build -o pulumiGo

##

xattr pulumiGo

```
com.apple.quarantine
```

- remove quarantine

```
xattr -d com.apple.quarantine pulumiGo
```

```

mkdir -p ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/pulumiGo
pulumiGo completion zsh > ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/pulumiGo/_pulumiGo

# Eidt ~/.zshrc
# plugins=(git ... pulumiGo)

# Reload
source ~/.zshrc

```

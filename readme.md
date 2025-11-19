## Build for Mac

```zsh
  go build -o pulumiGo
```

## Build for Windows

```zsh
GOOS=windows GOARCH=amd64 go build -o pulumiGo.exe
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

# Edit ~/.zshrc
# plugins=(git ... pulumiGo)

# Reload
source ~/.zshrc

```

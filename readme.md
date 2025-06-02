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

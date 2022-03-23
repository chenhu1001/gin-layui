## 使用 GoReleaser 发布你的应用
初始化
```
goreleaser init
```

执行自动发布流程
Mac
```
brew install goreleaser
goreleaser --snapshot --skip-publish --rm-dist
```
Linux
```
./goreleaser --snapshot --skip-publish --rm-dist
```
Windows
```
goreleaser.exe --snapshot --skip-publish --rm-dist
```
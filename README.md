## 同じリポジトリの別階層のモジュール/packageをimportしたいとき

例として、このリポジトリの構造を利用する。

```
play-go-module/
├── README.md
├── cmd
│   └── sub
│       ├── main.go
│       └── sub.exe
├── go.mod
└── model.go
```

一番上位のディレクトリ`play-go-module`直下の`go.mod`でモジュール定義(play-go-module)し、`model.go`はそのモジュールに属するファイルである。`model.go`では`User`構造体を定義しておき、外部から利用できるようにしている。

`cmd/sub`ディレクトリでは、`play-go-module`を`pgm`としてimportしている。

`cmd/sub`には`*.mod`ファイルをおいていないので、`go mod *`のコマンドは`play-go-module`直下で行うことになる。

module名の定義ではハイフンを使用できるが、コードの中では使用できないため、あらかじめハイフンなしでmodule名を定義したほうがよさげ。

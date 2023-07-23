# main.go

```go
// パッケージをmainに指定する
// パッケージの宣言は1つだけにする
package main

// formatパッケージをimportする
import "fmt"

// entrypoint
func main() {
    // Hello World
	fmt.Println("Hello World")
}

```

# プログラムの実行

```
go run main.go
```

上記のように go run ファイル名<br>
と指定することで実行できる

# 変数

```go
package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println(i)
    // 100
    var s string = "Hello GO"
    fmt.Println(s)
    // Hello GO
	var t,f bool = true, false
	fmt.Println(t,f)
    // true false
	var (
		num = 200
		str = "Golang"
	)
	fmt.Println(num, str)
    // 200 Golang

    // 暗黙的な定義
    // 変数 := 値
    i4 := 400
    fmt.Println(i4)
    i4 = 450
    fmt.Println(i4)
    // 一度定義したら再定義できない
    // また、暗黙的な定義に関しては関数外で定義できない
    // ただし、明示的な定義は関数外で定義できる
}
```

## 明示的な定義と暗黙的な定義の使い分け

- 関数内か関数外かで使い分ける
- 明示的な定義を推奨

## 型

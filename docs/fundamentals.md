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

### Integer

```go
func main() {
	var i int = 100
	fmt.Println(i)
    fmt.Printf("%T\n",i)
}
//100
//int
```

- int8
- int16
- int32
- int64

がある。同じ Int 型でも種類が違うと計算できない
また、int 型は環境に依存する変数なので PC が 64bit だと int64 になる

### Float

```go
func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)
}
// 2.4
```

### Bool

```go
func main() {
	var t,f bool = true, false
	fmt.Println(t,f)
}
//true false
```

### String

```go
func main() {
    var s string = "Hello GO"
    fmt.Println(s)
}
//Hello GO
```

### Array

GO の場合は後から配列の要素数を変更できない

```go
func main() {
    var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n",arr1)
	var arr2 [3]string = [3]string{"A","B","C"}
	fmt.Println(arr2)
	arr3 := [3]int{1,2,3}
	fmt.Println(arr3)
    // 要素数を省略できる
	arr4 := [...]string{"D","E"}
	fmt.Println(arr4)
    // arr2のindex番号0をprint
    fmt.Println(arr2[0])
    // 要素数-1の値をprint
    fmt.Println(arr2[2-1])
    // 値を更新する
	arr2[2] = "F"
	fmt.Println(arr2)
    // 要素数をprint
	fmt.Println(len(arr2))
}
// [0 0 0]
// [3]int
// [A B C]
// [1 2 3]
// [D E]
// A
// B
// [A B F]
// 3
```

### Interface

あらゆる型と互換性のある型で初期値は Nil。ただし、演算ができない

```go
func main() {
	var x interface{}
	fmt.Println(x)
    x = 1
    fmt.Println(x)
}
// <nil>
// 1
```

### 型変換

```go
func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)
    // 2つ目の値(_)を破棄する
	j, _ := strconv.Atoi(s)
	fmt.Printf("i = %T\n", j)
	var i2 int = 100
	s2 := strconv.Itoa(i2)
	fmt.Printf("s2 = %T\n", s2)
}
// i = int
// fl64 = float64
// s = string
// i = int
// s2 = string
```

## 定数

後から上書きできない

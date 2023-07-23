package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Testing")
	fmt.Println(time.Now())
	var i int = 100
	fmt.Println(i)
    var s string = "Hello GO"
    fmt.Println(s)
	var t,f bool = true, false
	fmt.Println(t,f)
	var (
		num = 200
		str = "Golang"
	)
	fmt.Println(num, str)
    i4 := 400
    fmt.Println(i4)
    i4 = 450
    fmt.Println(i4)
}


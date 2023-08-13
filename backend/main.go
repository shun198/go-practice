package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)
	var s string = "100"
	fmt.Printf("s = %T\n", s)
	j, _ := strconv.Atoi(s)
	fmt.Printf("i = %T\n", j)
	var i2 int = 100
	s2 := strconv.Itoa(i2)
	fmt.Printf("s2 = %T\n", s2)
}


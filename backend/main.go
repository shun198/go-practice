package main

import (
	"fmt"
)

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func main() {
	TestDefer()
}


package gomodule

import "fmt"

func Hello() {
	fmt.Println("==> You called gomodule hello world")
}

func NewHello(value string) {
	fmt.Printf("==> This is value: %v\n", value)
}

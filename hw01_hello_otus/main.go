package main

import (
	"fmt"
)

func main() {
	text := "Hello, OTUS!"
	res := ""
	for i := len(text) - 1; i >= 0; i-- {
		res += string(text[i])
	}
	fmt.Println(res)
}

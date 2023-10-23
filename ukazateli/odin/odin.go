package main

import "fmt"

func main() {
	a := 7
	var b *int = &a
	*b += *b
	fmt.Println(b)
}

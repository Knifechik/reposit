package main

import "fmt"

func main() {
	b := 7
	plus(&b)
	fmt.Println(b)
}

func plus(a *int) {
	*a += *a
}

package main

import "fmt"

func main() {
	var b = tri()
	fmt.Println(*b)

}

func tri() *int {
	s := 7
	var e *int
	e = &s
	return e

}

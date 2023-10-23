package main

import "fmt"

type kekw struct {
	test int
}

func main() {
	var kek kekw
	kek.test = 45
	var p = &kek
	p.test = 34
	fmt.Println(kek)
}

package main

import "fmt"

func main() {
	kek := new(int)
	*kek = 5
	lol := new(string)
	*lol = "ponyal"
	tt := new(bool)
	*tt = true
	fmt.Println(*kek, *lol, *tt)
	*kek = *kek + 7
	*lol = *lol + " prinyal"
	*tt = !*tt
	fmt.Println(*kek, *lol, *tt)
}

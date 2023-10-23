package main

import "fmt"

func main() {
	s := 5
	e := &s
	a := 4
	v := &a
	var slay = []int{*e, *v}
	for i, elem := range slay {
		slay[i] = elem + elem
	}
	fmt.Println(slay)
}

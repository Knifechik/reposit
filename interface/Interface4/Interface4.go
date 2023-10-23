package main

import (
	"fmt"
)

type Seconds int

type ConvertToMinutes interface {
	convert() Seconds
}

func (s Seconds) convert() Seconds {
	s = s * 60
	return s
}

func main() {
	var minutes int
	fmt.Scan(&minutes)
	var inter ConvertToMinutes = Seconds(minutes)
	fmt.Println(inter.convert())
}

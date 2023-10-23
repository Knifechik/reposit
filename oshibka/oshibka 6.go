package main

import (
	"errors"
	"fmt"
)

type YourError string

func (hem *YourError) Error() string {
	return fmt.Sprint("Одна ошибка и  ты ошибся")
}

var meh YourError

func main() {
	a := 5
	b := 0
	c, err := hm(a, b)
}

func hm(a, b int) (int, error) {
	if a != 0 || b != 0 {
		return 0, meh
	}
}

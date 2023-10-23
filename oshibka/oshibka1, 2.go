package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	a, err := oshibka(5, 0) // таков путь
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(a)
}

func oshibka(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("что-то из этого ноль, ошибочка")
	}
	return a * b, nil
}

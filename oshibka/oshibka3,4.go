package main

import (
	"fmt"
	"os"
)

func main() {
	a, err := oshibk(5, 0) // таков путь
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(a)
}

func oshibk(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("что-то из этого ноль, ошибочка: a = %d, b = %d", a, b)
	}
	return a / b, nil
}

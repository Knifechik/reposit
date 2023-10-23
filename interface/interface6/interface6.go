package main

import "fmt"

func main() {
	q := "kekw"
	inter(q)
	p := 3.21
	inter(p)
	e := true
	inter(e)
	a := 6
	inter(a)
}

func inter(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("это int")
	case string:
		fmt.Println("эт string")
	case bool:
		fmt.Println("эт bool")
	case float64:
		fmt.Println("эт float64")
	}
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("\tПервое Задание\n")
	odin()
	fmt.Printf("\tВторое Задание\n")
	dva()
	fmt.Printf("\tТретье Задание\n")
	tri()
	fmt.Printf("\tЧетвёртое Задание\n")
	chetire()
	fmt.Printf("\tПятое Задание\n")
	pyat()
	fmt.Printf("\tШестое Задание\n")
	shest()
	fmt.Printf("\tСедьмое Задание\n")
	sem()
	fmt.Printf("\tВосьмое Задание\n")
	vosem()
	fmt.Printf("\tДевятое Задание\n")
	devyat()
	fmt.Printf("\tДесятое Задание\n")
	desyat()
}

func odin() {
	i := 5
	f := 5.5
	s := "5"
	b := true
	fmt.Printf("%v\n%v\n%v\n%v\n", i, f, s, b)
}

func dva() {
	raz := 5
	dva := 5
	fmt.Printf("%v\n%v\n%v\n%v\n", raz+dva, raz-dva, raz/dva, raz*dva)
}

func tri() {
	t := true
	f := false
	fmt.Printf("%v\n%v\n", t && f, t || f)
}

func chetire() {
	s := "5"
	ss, _ := strconv.Atoi(s)
	fmt.Println(ss + ss)
}

func pyat() {
	var i = 5
	ii := 5
	fmt.Printf("%v\n%v\n", i, ii)
}

func shest() {
	f := 5.5
	ff := 5.5
	fmt.Printf("%f\n%f\n%f\n%f\n", f+ff, f*ff, f/ff, f-ff)
}

func sem() {
	t := true
	f := false
	if t != f {
		fmt.Printf("%v != %v\n", t, f)
	}
}
func vosem() {
	f := 5.5
	fmt.Println(f)
	i := int(f)
	fmt.Println(i)
	f = float64(i)
	fmt.Println(f)
}

func devyat() {
	s := "pyat"
	fmt.Println(s)
	s = "5"
	fmt.Println(s)
}

func desyat() {
	fmt.Println("Введите два числа через пробелы")
	var s string
	fmt.Scan(&s)
	var i int
	fmt.Scan(&i)
	ii, _ := strconv.Atoi(s)
	fmt.Println(i + ii)
}

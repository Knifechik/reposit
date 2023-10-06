package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() int
}

type Rectangle struct {
	Width  int
	Height int
}

type ReTriangle struct {
	odin float64
	dva  float64
	tri  float64
}

func (rekt Rectangle) Area() int {
	return rekt.Width * rekt.Height
}

func (tri ReTriangle) Area() int {
	p := (tri.odin + tri.dva + tri.tri) / 2
	pp := math.Sqrt(p * (p - tri.odin) * (p - tri.dva) * (p - tri.tri))
	return int(pp)
}

func main() {
	var square Rectangle
	fmt.Scan(&square.Width, &square.Height)
	var hz Shape = square
	fmt.Println(hz.Area())
	var triangle ReTriangle
	fmt.Scan(&triangle.odin, &triangle.dva, &triangle.tri)
	hz = triangle
	fmt.Println(hz.Area())
	Zadanie2()
	//hz = newRentanglefn(square.Width, square.Height)
	//fmt.Println(hz.Area())
}

//type Rectanglefn func() int
//
//func newRentanglefn(width, height int) Rectanglefn {
//	return Rectanglefn(func() int {
//		return width + height
//	})
//}
//
//func (receiver Rectanglefn) Area() int {
//	return receiver()
//}

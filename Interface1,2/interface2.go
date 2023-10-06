package main

import (
	"fmt"
	"math/rand"
)

// Интерфейс Shape и Структура Rectangle из первой задачки

type Resizable interface {
	Resize()
	DoubleSize()
}

func (rekt Rectangle) Resize() {
	random := rand.Intn(10)
	rekt.Height += random
	rekt.Width += random
}

func (rekt Rectangle) DoubleSize() {
	rekt.Height *= rekt.Height
	rekt.Width *= rekt.Width
}

func Zadanie2() {
	var square Rectangle
	fmt.Scan(&square.Width, &square.Height)
	var size Resizable = square
	fmt.Println(size.DoubleSize, size.Resize)
}

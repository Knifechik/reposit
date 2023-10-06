package main

import (
	"fmt"
	"math/rand"
)

// Интерфейс Shape и Структура Rectangle из первой задачки

type Resizable interface {
	Resize() (int, int)
	DoubleSize() (int, int)
}

func (rekt Rectangle) Resize() (int, int) {
	random := rand.Intn(10)
	rekt.Height += random
	rekt.Width += random
	return rekt.Height, rekt.Width
}

func (rekt Rectangle) DoubleSize() (int, int) {
	rekt.Height += rekt.Height
	rekt.Width += rekt.Width
	return rekt.Height, rekt.Width
}

func Zadanie2() {
	fmt.Printf("\tВторое задание\nВведите значение высоты и ширины прямоугольника:\n")
	var square Rectangle
	fmt.Scan(&square.Width, &square.Height)
	var size Resizable = square
	var ploshad Shape = square
	fmt.Printf("Площадь: %v\n", fmt.Sprint(ploshad.Area()))
	fmt.Printf("Изменение величин на случайное число: %v\n", fmt.Sprint(size.Resize()))
	fmt.Printf("Увеличение размеров вдвое: %v\n", fmt.Sprint(size.DoubleSize()))
}

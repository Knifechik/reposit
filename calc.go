package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Умножительные операции
func multi(flot []float64, st string) ([]float64, string) {
	i := strings.Index(st, "*")
	st = strings.Replace(st, "*", "", 1)
	flot[i] = flot[i] * flot[i+1]
	if len(flot) == 2 {
		return flot[:1], st
	} else {
		flot = append(flot[:i+1], flot[i+2:]...)
		return flot, st
	}
}

// Делительные операции
func div(flot []float64, st string) ([]float64, string) {
	i := strings.Index(st, "/")
	st = strings.Replace(st, "/", "", 1)
	if flot[i+1] == 0 {
		p := errors.New("divizion by zero")
		fmt.Println(p)
		os.Exit(0)
	}
	flot[i] = flot[i] / flot[i+1]
	if len(flot) == 2 {
		return flot[:1], st
	} else {
		flot = append(flot[:i+1], flot[i+2:]...)
		return flot, st
	}
}

// Плюсовые операции
func plus(flot []float64, st string) ([]float64, string) {
	i := strings.Index(st, "+")
	st = strings.Replace(st, "+", "", 1)
	flot[i] = flot[i] + flot[i+1]
	if len(flot) == 2 {
		return flot[:1], st
	} else {
		flot = append(flot[:i+1], flot[i+2:]...)
		return flot, st
	}
}

// Минусовые операции
func minus(flot []float64, st string) ([]float64, string) {
	i := strings.Index(st, "-")
	st = strings.Replace(st, "-", "", 1)
	flot[i] = flot[i] - flot[i+1]
	if len(flot) == 2 {
		return flot[:1], st
	} else {
		flot = append(flot[:i+1], flot[i+2:]...)
		return flot, st
	}
}

// Не, ну это перебор
func per(fl []float64, op string) float64 {
	for len(op) != 0 {
		if strings.Contains(op, "*") || strings.Contains(op, "/") {
			if !strings.Contains(op, "/") || strings.Contains(op, "*") && strings.Index(op, "/") > strings.Index(op, "*") {
				fl, op = multi(fl, op)
				continue
			} else if !strings.Contains(op, "*") || strings.Contains(op, "/") && strings.Index(op, "*") > strings.Index(op, "/") {
				fl, op = div(fl, op)
				continue
			}
		}
		if !strings.Contains(op, "*") && !strings.Contains(op, "/") && (strings.Contains(op, "+") || strings.Contains(op, "-")) {
			if !strings.Contains(op, "-") || strings.Contains(op, "+") && strings.Index(op, "-") > strings.Index(op, "+") {
				fl, op = plus(fl, op)
				continue
			} else if !strings.Contains(op, "+") || strings.Contains(op, "-") && strings.Index(op, "+") > strings.Index(op, "-") {
				fl, op = minus(fl, op)
				continue
			}
		}
	}
	return fl[0]
}

func main() {
	chtoto := read()
	fl, op := conv(chtoto)
	result := per(fl, op)
	fmt.Println(result)
}

// Чтение строки
func read() string {
	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSuffix(expression, "\r\n")
	expression = strings.ReplaceAll(expression, " ", "")
	return expression
}

// Конвертация строки в срез чисел и строку операторов
func conv(stroka string) ([]float64, string) {
	var cifra string
	var oper string
	var flot float64
	norm := make([]float64, 0, 100)
	for _, elem := range stroka {
		if unicode.IsDigit(elem) || elem == '.' {
			cifra = cifra + string(elem)
		} else {
			flot, _ = strconv.ParseFloat(cifra, 64)
			norm = append(norm, flot)
			oper = oper + string(elem)
			cifra = ""
		}
	}
	flot, _ = strconv.ParseFloat(cifra, 64)
	norm = append(norm, flot)
	return norm, oper
}

//Наговнокодил знатно, можно сделать гораздо лаконичнее и читабильнее, пофиксить действия с отрицательными числами и засунуть в цикл до знака "=", пока впадлу

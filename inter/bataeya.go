package main

import "fmt"

type batteryForTest struct {
	energy string
}

func (bat batteryForTest) String() string {
	nach := "["
	kon := "]"
	for _, elem := range bat.energy {
		if elem == '0' {
			nach = nach + " "
		} else {
			kon = "X" + kon
		}
	}
	return nach + kon
}

func main() {
	var battery batteryForTest
	fmt.Scan(&battery.energy)
	fmt.Println(battery)
}

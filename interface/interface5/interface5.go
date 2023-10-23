package main

import (
	"fmt"
	"sort"
)

type FloatSlice []float64

var _ sort.Interface = FloatSlice{}

func (f FloatSlice) Len() int {
	return len(f)
}

func (f FloatSlice) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f FloatSlice) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func main() {
	v := FloatSlice{1, 5, 2, 5, 8, 15, 7, 3}
	sort.Sort(v)
	fmt.Println(v)
}

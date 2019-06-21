package main

import (
	"fmt"
	"math"
)

var a1, a2 int
var b1, b2 = 1, 2
var (
	e float32
	f bool
)

const (
	a    = iota
	b, c = iota, iota
	d    = iota
)

const j = iota
const k, y = iota, iota

func basic() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
	nan := math.NaN()
	fmt.Println(math.IsNaN(z / z))                // true
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false" NaN和任何数都是不相等
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(f)
	fmt.Println(c)
	fmt.Println(d)
}

func complex() {
	var t = 2.1 + 0.1i
	fmt.Println(real(t))
	fmt.Println(imag(t))
}
func main() {
	complex()

}

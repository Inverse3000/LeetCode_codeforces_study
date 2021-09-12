package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var ss = "kkk"
var bb = true

func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("%d, %q\n", a, b)
}
func variableInitialValue() {
	var a, s int = 3, 5
	var b string = "abc"
	fmt.Printf("%d, %q\n", a, b)
	fmt.Printf("%d", s)
}
func variableTypeDeduction() {
	var a, b, c, d = 1, 2, 3, 4
	fmt.Println(a, b, c, d)
}
func variableTypeDeduction2() {
	a := 1
	a = 3
	fmt.Println(a)
}
func euler() {
	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1)
	//	fmt.Println(cmplx.Pow(math.E,1i * math.Pi)+1)
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
func enums() {
	const (
		cpp    = 1
		java   = 0
		golang = 2
		python = 3
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)
	fmt.Println(b, kb, mb, gb)
	fmt.Println(cpp, python, java, golang)
}
func main() {
	enums()
}

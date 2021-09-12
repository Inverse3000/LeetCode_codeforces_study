package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	var res int
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		panic("unsupported op: " + op)
	}
	return res
}
func plus(a, b int) int {
	return a + b
}
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(plus).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println(opName)
	return op(a, b)
}
func div(a, b int) (q, r int) {
	return a / b, a % b
}
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
func swap(a, b int) (int, int) {
	return b, a // wont work
}
func main() {
	fmt.Println(apply(plus, 3, 5))
	fmt.Println(sum(1, 2, 3, 4))
	a := 1
	b := 2
	a, b = swap(a, b)
	fmt.Println(a, b)
}

package main

import (
	"fmt"
)

func grade(score int) string {
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		return "FFF"
	case score < 80:
		return "C"
	case score < 90:
		return "B"
	case score <= 100:
		return "A"
	}
	return "why"
}
func main() {
	fmt.Println(grade(99))
	fmt.Println(grade(-9))
}

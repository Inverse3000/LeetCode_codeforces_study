package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func main() {
	var arr [5]int
	arr2 := [3]int{1, 3, 4}
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr, arr2, arr3)
	var cool [25][25]bool
	fmt.Println(len(cool), len(cool[0]))
	for i, v := range arr3 {
		fmt.Printf("%d, %d\n", i, v)
	}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
}

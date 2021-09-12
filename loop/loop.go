package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
func printFile() {
	file, err := os.Open("abc.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(1)
		fmt.Println(scanner.Text())
	}
}
func forever() {
	for true {
		fmt.Println("haha")
	}
}
func main() {
	b := "abcd"
	fmt.Println(len(b))
	c := []rune(b)
	for i, v := range c {
		fmt.Printf("%d, %c\n", i, rune(v))
	}
	//	printFile()
	//	fmt.Printf("%s\n",convertToBin(105))
}

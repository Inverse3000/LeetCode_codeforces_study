package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var t int
	fmt.Fscan(in, &t)
	for c := 1; c <= t; c++ {
		var n int
		fmt.Fscan(in, &n)
		var s string
		fmt.Fscan(in, &s)
		now := []byte(s)
		flag := false
		for i := 1; i <= len(s); i++ {
			if now[i-1] == '0' {
				flag = true
				if i <= len(s)/2 {
					fmt.Fprintln(out, i, len(s), i+1, len(s))
				}
				if i > len(s)/2 {
					fmt.Fprintln(out, 1, i, 1, i-1)
				}
				break
			}
		}
		if !flag {
			fmt.Fprintln(out, 1, len(s)-1, 2, len(s))
		}
	}
}

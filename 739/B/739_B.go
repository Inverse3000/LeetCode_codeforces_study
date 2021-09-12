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
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		var numa, numb int
		if a < b {
			numa = a
			numb = b
		} else {
			numa = b
			numb = a
		}
		ans := 0
		//	fmt.Println(numb+1-numa,numa)
		if numb+1-numa <= numa {
			ans = -1
		}
		mid := numb - numa + 1
		if c > mid+(mid-2) {
			ans = -1
		} else if c < mid && ans != -1 {
			ans = mid + (c - 1)
		} else if c == mid && ans != -1 {
			ans = 1
		} else if c > mid && ans != -1 {
			ans = 1 + (c - mid)
		}
		fmt.Fprintln(out, ans)
	}
}

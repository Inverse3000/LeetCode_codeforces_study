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
		var n, m, k int
		fmt.Fscan(in, &n, &m, &k)
		flag := true
		tk := k
		for i := 1; i <= n; {
			if (n-i)%2 != 0 {
				if tk < (m/2)*2 {
					if tk%2 != 0 {
						flag = false
					} else {
						tk = 0
					}
					break
				} else {
					tk -= (m / 2) * 2
				}
				i += 2
			} else {
				if tk < (m / 2) {
					flag = false
					break
				}
				tk -= (m / 2)
				i++
			}
		}
		if tk != 0 {
			flag = false
		}
		if flag {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

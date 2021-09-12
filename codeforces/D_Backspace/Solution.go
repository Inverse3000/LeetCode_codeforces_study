package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var s []byte
	var t []byte
	var n int
	fmt.Fscan(in, &n)
	for u := 1; u <= n; u++ {
		fmt.Fscan(in, &s)
		fmt.Fscan(in, &t)
		if len(t) > len(s) {
			fmt.Fprintln(out, "NO")
			continue
		}
		flag := false
		for i, j := len(s)-1, len(t)-1; i >= 0 && j >= 0; {
			if s[i] == t[j] {
				i--
				j--
				if j+1 == 0 {
					fmt.Fprintln(out, "YES")
					flag = true
					break
				}
			} else {
				i -= 2
			}
		}
		if !flag {
			fmt.Fprintln(out, "NO")
		}
	}
}

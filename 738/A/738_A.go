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
		var a [105]int
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
		}
		ans := 0
		for i := 1; i <= 31; i++ {
			flag := true
			for j := 1; j <= n; j++ {
				if a[j]&(1<<(i-1)) == 0 {
					flag = false
				}
			}
			//	fmt.Println(i,flag,1<<(i-1))
			if flag {
				ans += 1 << (i - 1)
			}
		}
		fmt.Fprintln(out, ans)
	}
}

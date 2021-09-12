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
		var a [30005]int
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
		}
		start := -1
		for i := 1; i <= n; i++ {
			if i == n {
				if a[i] == 0 && a[1] == 1 {
					start = i
				}
			} else {
				if a[i] == 0 && a[i+1] == 1 {
					start = i
				}
			}
		}
		ans := []int{}
		if start != -1 {
			now := start
			for i := 1; i < now; i++ {
				ans = append(ans, i)
			}
			ans = append(ans, now)
			ans = append(ans, n+1)
			for i := now + 1; i <= n; i++ {
				ans = append(ans, i)
			}
		} else if a[1] == 1 {
			ans = append(ans, n+1)
			for i := 1; i <= n; i++ {
				ans = append(ans, i)
			}
		} else if a[n] == 0 {
			for i := 1; i <= n+1; i++ {
				ans = append(ans, i)
			}
		} else if a[1] == 0 && n == 1 {
			ans = append(ans, 1)
			ans = append(ans, 2)
		}
		if len(ans) == 0 {
			fmt.Fprintln(out, -1)
		} else {
			for i := 0; i < len(ans); i++ {
				fmt.Fprintf(out, "%d ", ans[i])
			}
		}
	}
}

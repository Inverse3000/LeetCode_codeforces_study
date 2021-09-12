package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func d(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for c := 1; c <= t; c++ {
		var n int
		fmt.Fscan(in, &n)
		var a [200005]int
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
		}
		var ans int
		ans = n + n - 1 // 1s and 2s
		for i := 1; i <= n-2; i++ {
			flag := false
			if d(a[i], i, a[i+2], i+2) == d(a[i+1], i+1, a[i], i)+d(a[i+1], i+1, a[i+2], i+2) {
				if i == 1 {
				}
				flag = true
			}

			if flag == false {
				ans++
			}
		}
		for i := 1; i <= n-3; i++ {
			flag := false
			for j := i; j <= i+3; j++ {
				for k := j + 1; k <= i+3; k++ {
					for v := k + 1; v <= i+3; v++ {
						if k != j && v != k && j != v {
							if d(a[j], j, a[v], v) == d(a[k], k, a[j], j)+d(a[k], k, a[v], v) {
								flag = true
							}
						}
					}
				}
			}
			if flag == false {
				ans++
			}
		}
		fmt.Fprintln(out, ans)
	}
}

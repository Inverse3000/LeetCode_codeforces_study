package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type ll int64

var spt [200005][19]ll

func gcd(x, y ll) ll {
	for y != 0 {
		temp := x
		x = y
		y = temp % y
	}
	return x
}
func abs(x ll) ll {
	if x < 0 {
		return -x
	}
	return x
}
func rq(l, r ll) ll {
	j := ll(math.Log2(float64(r - l + 1)))
	return gcd(spt[l][j], spt[r-(1<<j)+1][j])
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t ll
	fmt.Fscan(in, &t)
	var c ll
	for c = 1; c <= t; c++ {
		var n ll
		fmt.Fscan(in, &n)
		a := make([]ll, n+1)
		var i, j ll
		for i = 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
		}
		if n < 2 {
			fmt.Fprintln(out, 1)
			continue
		}
		d := make([]ll, n+1)
		for i = 1; i < n; i++ {
			d[i] = abs(a[i+1] - a[i])
		}
		for i = 1; i <= n; i++ {
			spt[i][0] = d[i]
		}
		for j = 1; j <= 18; j++ {
			for i = 1; i < n; i++ {
				if i+1<<(j-1) <= 200000 {
					spt[i][j] = gcd(spt[i][j-1], spt[i+1<<(j-1)][j-1])
				}
			}
		}
		var res ll
		res = 1
		for i = 1; i < n; i++ {
			r := n - 1
			l := i
			ans := ll(0)
			for l <= r {
				mid := (l + r) / 2
				if rq(i, mid) > 1 {
					ans = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			if res < ans-i+2 {
				res = r - i + 2
			}
		}
		fmt.Fprintln(out, res)
	}
}

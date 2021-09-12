package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(x, y ll) ll {
	if y != 0 {
		return gcd(y, x%y)
	} else {
		return x
	}
}
func LCM(x, y ll) ll {
	return x * y / gcd(x, y)
}

type ll int64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var t ll
	fmt.Fscan(in, &t)
	var c ll
	for c = 1; c <= t; c++ {
		var n ll
		fmt.Fscan(in, &n)
		var i ll
		var G ll
		G = 1
		var ans ll
		ans = 0
		for i = 1; G <= n; i++ {
			G = LCM(G, i)
			if G > n {
				break
			}
			ans += n / G
		}
		fmt.Fprintln(out, (ans+n)%1000000007)
	}
}

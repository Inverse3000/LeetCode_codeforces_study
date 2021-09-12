package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout
	var t int
	fmt.Fscan(in, &t)
	for u := 0; u < t; u++ {
		a := []int{}
		var temp int
		var n int
		fmt.Fscan(in, &n)
		var sum int64
		sum = 0
		mx := -math.MaxInt32
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &temp)
			a = append(a, temp)
			sum += int64(temp)
			if mx < temp {
				mx = temp
			}
		}
		ans := float64(sum-int64(mx))/float64(n-1) + float64(mx)
		fmt.Fprintln(out, ans)
	}
}

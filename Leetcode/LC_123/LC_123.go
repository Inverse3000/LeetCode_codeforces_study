package main

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxProfit(prices []int) int {
	var f [300005]int
	mi := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if f[i] < f[i]-mi {
			f[i] = f[i] - mi
		}
		if mi > f[i] {
			mi = f[i]
		}
	}
	ma := -math.MaxInt32
	var d [300005]int
	for i := len(prices) - 1; i >= 0; i-- {
		if d[i] < ma-d[i] {
			d[i] = ma - d[i]
		}
		if d[i] > ma {
			d[i] = ma
		}
	}
	ma = -math.MaxInt32
	ans := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if ans < f[i]+max(0, ma) {
			ans = f[i] + max(0, ma)
		}
		if d[i] > ma {
			ma = d[i]
		}
	}
	return ans
}
func main() {

}

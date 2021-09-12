package main

var f [50005]int
var dp [50005]int

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func numSquares(n int) int {
	f = [50005]int{}
	dp = [50005]int{}
	for i := 1; i*i <= n; i++ {
		f[i] = i * i
		dp[i*i] = 1
	}
	for i := 1; i <= n; i++ {
		if dp[i] == 0 {
			continue
		}
		for j := 1; j*j <= n; j++ {
			if dp[i+j*j] == 0 {
				dp[i+j*j] = dp[i] + 1
			} else {
				dp[i+j*j] = min(dp[i+j*j], dp[i]+1)
			}
		}
	}
	return dp[n]
}
func main() {

}

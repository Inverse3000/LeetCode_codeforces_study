package main

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func maxProfit(k int, prices []int) int {
	var f [1005][2][105]int
	f[0][0][0] = 0
	f[0][1][0] = -prices[0]
	for i := 1; i <= k; i++ {
		f[0][0][i] = -999999
		f[0][1][i] = -999999
	}
	ans := 0
	for i := 0; i < len(prices); i++ {
		for j := 0; j <= k; j++ {
			f[i][0][j] = f[i-1][0][j]
			if j >= 1 {
				f[i][0][j] = max(f[i-1][1][j-1]+prices[i], f[i][0][j])
			}
			f[i][1][j] = max(f[i-1][0][j]-prices[i], f[i-1][1][j])
			ans = max(ans, max(f[i][0][j], f[i][1][j]))
		}
	}
	return ans
}
func main() {

}

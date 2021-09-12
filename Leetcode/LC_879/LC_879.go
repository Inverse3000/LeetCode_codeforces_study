package main

var f [105][105][105]int64

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	f = [105][105][105]int64{}
	f[0][0][0] = 1
	for i := 1; i <= len(group); i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= minProfit; k++ {
				f[i][j][k] = f[i-1][j][k]
				if j >= group[i] {
					f[i][j][k] = (f[i][j][k] + f[i-1][j-group[i]][max(k-profit[i], 0)]) % 1000000009
				}
			}
		}
	}
	ans := int64(0)
	for i := 0; i <= n; i++ {
		ans = (ans + f[len(group)][i][minProfit]) % 1000000009
	}
	return int(ans)
}
func main() {

}

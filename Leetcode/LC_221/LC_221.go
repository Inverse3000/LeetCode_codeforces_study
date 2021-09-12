package main

var f [305][305]int

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
func maximalSquare(matrix [][]byte) int {
	ans := 0
	for i := 0; i < len(matrix); i++ {
		f[i][0] = int(matrix[i][0])
	}
	for i := 0; i < len(matrix[0]); i++ {
		f[0][i] = int(matrix[0][i])
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i-1 >= 0 && j-1 >= 0 {
				f[i][j] = min(f[i-1][j], min(f[i-1][j-1], f[i][j-1])) + 1
			}
			ans = max(ans, f[i][j]*f[i][j])
		}
	}
	return ans
}
func main() {

}

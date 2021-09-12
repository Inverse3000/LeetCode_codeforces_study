package main

func findTargetSumWays(nums []int, target int) int {
	f := [2005][2005]int{}
	f[0][1000] = 1
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= 2000; j++ {
			f[i][j] = f[i][j] + f[i-1][j-nums[i-1]]
			f[i][j] = f[i][j] + f[i-1][j+nums[i-1]]
		}
	}
	return f[len(nums)][target+1000]
}
func main() {

}

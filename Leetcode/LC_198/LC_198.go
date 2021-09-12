package main

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// dp[i][0] = max(dp[i-1][1],dp[i-1][0])
// dp[i][1] = max(dp[i-1][0]+num
func rob(nums []int) int {
	f := [100005][2]int{}
	f[0][0] = 0
	f[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		f[i][1] = f[i-1][0] + nums[i]
		f[i][0] = max(f[i-1][0], f[i-1][1])
	}
	return max(f[len(nums)-1][0], f[len(nums)-1][1])
}
func main() {
}

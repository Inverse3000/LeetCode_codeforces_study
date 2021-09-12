package main

func numberOfArithmeticsSlices(nums []int) int {
	var f [5005][5005]int
	for i := 0; i <= 5000; i++ {
		for j := 0; j <= 5000; j++ {
			f[i][j] = -999999
		}
	}
	ans := 0
	for i := 0; i+2 <= len(nums)-1; i++ {
		if nums[i+1]-nums[i] == nums[i+2]-nums[i] {
			f[i][i+2] = nums[i+1] - nums[i]
			ans++
		}
	}
	for k := 4; k <= len(nums); k++ {
		for i := 0; i+k-1 <= len(nums)-1; i++ {
			if nums[i]-nums[i+1] == f[i+1][i+k-1] {
				ans++
			}
		}
	}
	return ans
}
func main() {

}
